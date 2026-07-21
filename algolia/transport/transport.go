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
		var err error

		// Reassign a fresh body for the retry
		if i > 0 && req.GetBody != nil {
			req.Body, err = req.GetBody()
			if err != nil {
				break
			}
		}

		ctxTimeout, connectTimeout := t.resolveTimeouts(k, c, h)

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

// maxErrorBodySize bounds the error body read of a failed streaming request,
// so that a server streaming an endless error body cannot stall the caller
// forever.
const maxErrorBodySize = 1 << 20

// RequestStream performs the given request and returns the raw response
// without reading its body, so that the caller can consume it as a stream.
// Unlike Request, it does not retry: the request is only sent to the first
// available host, and no read deadline is applied to the response body, as it
// would abort the stream while it is being consumed. Cancellation is
// controlled by the caller through ctx. The outcome does not update the host
// health state used by the retry strategy, consistent with the JavaScript and
// Python clients. The Accept header is always overwritten with
// text/event-stream.
//
// The RequestConfiguration timeouts are forwarded to the [Requester], but the
// default requester ignores them and no context deadline is applied here:
// with the default requester, the time to the response headers is bounded
// only by ctx. Use a custom [Requester] to enforce a time-to-first-byte
// limit.
//
// A response with a non-2xx status code is consumed and returned as an
// [errs.HTTPStatusError] carrying the status code and the error body,
// consistent with the JavaScript and Python clients.
//
// The caller is responsible for closing the response body.
func (t *Transport) RequestStream(ctx context.Context, req *http.Request, k call.Kind, c RequestConfiguration) (*http.Response, error) {
	// Add Content-Encoding header, if needed
	if t.compression == compression.GZIP && shouldCompress(t.compression, req.Method, req.Body) {
		req.Header.Add("Content-Encoding", "gzip")
	}

	req.Header.Set("Accept", "text/event-stream")

	hosts := t.retryStrategy.GetTryableHosts(k)
	if len(hosts) == 0 {
		return nil, errs.ErrNoMoreHostToTry
	}

	host := hosts[0]

	timeout, connectTimeout := t.resolveTimeouts(k, c, host)

	req = req.WithContext(ctx)
	req.URL.Scheme = host.scheme
	req.URL.Host = host.host

	debug.Display(req)

	// Unlike in request, the response is voluntarily not passed to
	// debug.Display: displaying it would buffer the whole body in memory,
	// defeating streaming.
	res, err := t.requester.Request(req, timeout, connectTimeout)
	if err != nil {
		return nil, wrapRequestError(req, err)
	}

	if !is2xx(res.StatusCode) {
		body, errBody := io.ReadAll(io.LimitReader(res.Body, maxErrorBodySize))
		errClose := res.Body.Close()

		if errBody != nil {
			return nil, fmt.Errorf("cannot read error response body: %w: %w", errBody, errs.NewHTTPStatusError(res.StatusCode, nil))
		}

		if errClose != nil {
			return nil, fmt.Errorf("cannot close error response body: %w: %w", errClose, errs.NewHTTPStatusError(res.StatusCode, body))
		}

		return nil, errs.NewHTTPStatusError(res.StatusCode, body)
	}

	return res, nil
}

// resolveTimeouts returns the request and connect timeouts applying to a call
// of kind k against host h, honoring the overrides of c.
func (t *Transport) resolveTimeouts(k call.Kind, c RequestConfiguration, h Host) (time.Duration, time.Duration) {
	var timeout time.Duration

	switch {
	case k == call.Read && c.ReadTimeout != nil:
		timeout = *c.ReadTimeout
	case k == call.Write && c.WriteTimeout != nil:
		timeout = *c.WriteTimeout
	default:
		timeout = h.timeout
	}

	var connectTimeout time.Duration
	if c.ConnectTimeout != nil {
		connectTimeout = *c.ConnectTimeout
	} else {
		connectTimeout = t.connectTimeout
	}

	return timeout, connectTimeout
}

func (t *Transport) request(req *http.Request, host Host, timeout time.Duration, connectTimeout time.Duration) (*http.Response, error) {
	req.URL.Scheme = host.scheme
	req.URL.Host = host.host

	debug.Display(req)
	res, err := t.requester.Request(req, timeout, connectTimeout)
	debug.Display(res)

	if err != nil {
		return nil, wrapRequestError(req, err)
	}

	return res, nil
}

func wrapRequestError(req *http.Request, err error) error {
	msg := fmt.Sprintf("cannot perform request:\n\terror=%v\n\tmethod=%s\n\turl=%s", err, req.Method, req.URL)

	var nerr net.Error
	if errors.As(err, &nerr) {
		// Because net.Error and error have different meanings for the
		// retry strategy, we cannot simply return a new error, which
		// would make all net.Error simple errors instead. To keep this
		// behaviour, we wrap the message into a custom netError that
		// implements the net.Error interface if the original error was
		// already a net.Error.
		return errs.NetError(nerr, msg)
	}

	return errors.New(msg)
}

func shouldCompress(c compression.Compression, method string, body any) bool {
	isValidMethod := method == http.MethodPut || method == http.MethodPost
	isCompressionEnabled := c != compression.NONE
	isBodyNonEmpty := body != nil

	return isCompressionEnabled && isValidMethod && isBodyNonEmpty
}
