package transport

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/algolia/algoliasearch-client-go/v4/algolia/call"
	"github.com/algolia/algoliasearch-client-go/v4/algolia/compression"
	"github.com/algolia/algoliasearch-client-go/v4/algolia/debug"
	"github.com/algolia/algoliasearch-client-go/v4/algolia/errs"
)

type Transport struct {
	requester                       Requester
	retryStrategy                   *RetryStrategy
	compression                     compression.Compression
	connectTimeout                  time.Duration
	exposeIntermediateNetworkErrors bool
}

func New(cfg Configuration) *Transport {
	transport := &Transport{
		requester:                       cfg.Requester,
		retryStrategy:                   newRetryStrategy(cfg.Hosts, cfg.ReadTimeout, cfg.WriteTimeout),
		connectTimeout:                  cfg.ConnectTimeout,
		compression:                     cfg.Compression,
		exposeIntermediateNetworkErrors: cfg.ExposeIntermediateNetworkErrors,
	}

	if transport.connectTimeout == 0 {
		transport.connectTimeout = DefaultConnectTimeout
	}

	if transport.requester == nil {
		transport.requester = NewDefaultRequester(&transport.connectTimeout)
	}

	return transport
}

func prepareRetryableRequest(req *http.Request) (*http.Request, error) {
	// Read the original body
	if req.Body == nil {
		return req, nil // Nothing to do if there's no body
	}

	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read body: %w", err)
	}

	_ = req.Body.Close() // close the original body

	// Set up GetBody to recreate the body for retries
	req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	req.GetBody = func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewReader(bodyBytes)), nil
	}

	return req, nil
}

func (t *Transport) Request(ctx context.Context, req *http.Request, k call.Kind, c RequestConfiguration) (*http.Response, []byte, error) {
	var intermediateNetworkErrors []error

	// Add Content-Encoding header, if needed
	if t.compression == compression.GZIP && shouldCompress(t.compression, req.Method, req.Body) {
		req.Header.Add("Content-Encoding", "gzip")
	}

	// Prepare the request to be retryable.
	req, err := prepareRetryableRequest(req)
	if err != nil {
		return nil, nil, err
	}

	for i, h := range t.retryStrategy.GetTryableHosts(k) {
		// Handle per-request timeout by using a context with timeout.
		// Note that because we are in a loop, the cancel() callback cannot be
		// deferred. Instead, we call it precisely after the end of each loop or
		// before the early returns, but when we do so, we do it **after**
		// reading the body content of the response. Otherwise, a `context
		// cancelled` error may happen when the body is read.
		var (
			ctxTimeout     time.Duration
			connectTimeout time.Duration
			err            error
		)

		// Reassign a fresh body for the retry
		if i > 0 && req.GetBody != nil {
			req.Body, err = req.GetBody()
			if err != nil {
				break
			}
		}

		switch {
		case k == call.Read && c.ReadTimeout != nil:
			ctxTimeout = *c.ReadTimeout
		case k == call.Write && c.WriteTimeout != nil:
			ctxTimeout = *c.WriteTimeout
		default:
			ctxTimeout = h.timeout
		}

		if c.ConnectTimeout != nil {
			connectTimeout = *c.ConnectTimeout
		} else {
			connectTimeout = t.connectTimeout
		}

		perRequestCtx, cancel := context.WithTimeout(ctx, ctxTimeout)
		req = req.WithContext(perRequestCtx)
		res, err := t.request(req, h, ctxTimeout, connectTimeout)

		code := 0
		if res != nil {
			code = res.StatusCode
		}

		// Context error only returns a non-nil error upon context
		// cancellation, which is a signal we interpret as an early return.
		// Indeed, we do not want to retry on other hosts if the context is
		// already cancelled.
		if ctx.Err() != nil {
			cancel()

			return res, nil, err
		}

		switch t.retryStrategy.Decide(h, code, err) {
		case Success, Failure:
			body, errBody := io.ReadAll(res.Body)
			errClose := res.Body.Close()

			cancel()

			res.Body = io.NopCloser(bytes.NewBuffer(body))
			if errBody != nil {
				return res, nil, fmt.Errorf("cannot read body: %w", errBody)
			}

			if errClose != nil {
				return res, nil, fmt.Errorf("cannot close response's body: %w", errClose)
			}

			return res, body, err
		default:
			if err != nil {
				intermediateNetworkErrors = append(intermediateNetworkErrors, err)
			} else if res != nil {
				msg := fmt.Sprintf("cannot perform request:\n\tStatusCode=%d\n\tmethod=%s\n\turl=%s\n\t", res.StatusCode, req.Method, req.URL)
				intermediateNetworkErrors = append(intermediateNetworkErrors, errors.New(msg))
			}

			if res != nil && res.Body != nil {
				err = res.Body.Close()
				if err != nil {
					cancel()

					return res, nil, fmt.Errorf("cannot close response's body before retry: %w", err)
				}
			}
		}

		cancel()
	}

	if t.exposeIntermediateNetworkErrors {
		return nil, nil, errs.NewNoMoreHostToTryError(intermediateNetworkErrors...)
	}

	return nil, nil, errs.ErrNoMoreHostToTry
}

func (t *Transport) request(req *http.Request, host Host, timeout time.Duration, connectTimeout time.Duration) (*http.Response, error) {
	req.URL.Scheme = host.scheme
	req.URL.Host = host.host

	debug.Display(req)
	res, err := t.requester.Request(req, timeout, connectTimeout)
	debug.Display(res)

	if err != nil {
		msg := fmt.Sprintf("cannot perform request:\n\terror=%v\n\tmethod=%s\n\turl=%s", err, req.Method, req.URL)

		var nerr net.Error
		if errors.As(err, &nerr) {
			// Because net.Error and error have different meanings for the
			// retry strategy, we cannot simply return a new error, which
			// would make all net.Error simple errors instead. To keep this
			// behaviour, we wrap the message into a custom netError that
			// implements the net.Error interface if the original error was
			// already a net.Error.
			err = errs.NetError(nerr, msg)
		} else {
			err = errors.New(msg)
		}

		return nil, err
	}

	return res, nil
}

func shouldCompress(c compression.Compression, method string, body any) bool {
	isValidMethod := method == http.MethodPut || method == http.MethodPost
	isCompressionEnabled := c != compression.NONE
	isBodyNonEmpty := body != nil

	return isCompressionEnabled && isValidMethod && isBodyNonEmpty
}
