package transport

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/call"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/compression"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/debug"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/errs"
	iopt "github.com/algolia/algoliasearch-client-go/v3/algolia/internal/opt"
)

const version = "3.31.2"

type Transport struct {
	requester     Requester
	retryStrategy *RetryStrategy
	headers       http.Header
	compression   compression.Compression
}

func New(
	hosts []*StatefulHost,
	requester Requester,
	appID string,
	apiKey string,
	readTimeout time.Duration,
	writeTimeout time.Duration,
	defaultHeaders map[string]string,
	extraUserAgent string,
	compression compression.Compression,
) *Transport {
	if requester == nil {
		requester = newDefaultRequester()
	}

	userAgents := []string{
		fmt.Sprintf("Algolia for Go (%s)", version),
		fmt.Sprintf("Go (%s)", runtime.Version()),
	}

	if extraUserAgent != "" {
		userAgents = append([]string{extraUserAgent}, userAgents...)
	}

	headers := http.Header{}
	for k, v := range map[string]string{
		"Connection":               "Keep-Alive",
		"Content-Type":             "application/json; charset=utf-8",
		"User-Agent":               strings.Join(userAgents, ";"),
		"X-Algolia-Application-Id": appID,
		"X-Algolia-API-Key":        apiKey,
	} {
		headers.Set(k, v)
	}

	for k, v := range defaultHeaders {
		headers.Set(k, v)
	}

	return &Transport{
		requester:     requester,
		retryStrategy: newRetryStrategy(hosts, readTimeout, writeTimeout),
		headers:       headers,
		compression:   compression,
	}
}

func (t *Transport) Request(
	res interface{},
	method string,
	path string,
	body interface{},
	k call.Kind,
	opts ...interface{},
) error {
	var (
		ctx       = iopt.ExtractContext(opts...)
		headers   = t.headers.Clone()
		urlParams = make(map[string]string)
	)

	if extraHeaders := iopt.ExtractExtraHeaders(opts...); extraHeaders != nil {
		headers = mergeHeaders(headers, extraHeaders.Get())
	}

	if extraURLParams := iopt.ExtractExtraURLParams(opts...); extraURLParams != nil {
		urlParams = extraURLParams.Get()
	}

	if forwardToReplicas := iopt.ExtractForwardToReplicas(opts...); forwardToReplicas != nil {
		urlParams["forwardToReplicas"] = fmt.Sprintf("%t", forwardToReplicas.Get())
	}

	exposeIntermediateNetworkErrors := iopt.ExtractExposeIntermediateNetworkErrors(opts...).Get()
	var intermediateNetworkErrors []error

	for _, h := range t.retryStrategy.GetTryableHosts(k) {
		req, err := buildRequest(t.compression, method, h.host, path, body, headers, urlParams)
		if err != nil {
			return err
		}

		// Handle per-request timeout by using a context with timeout.
		// Note that because we are in a loop, the cancel() callback cannot be
		// deferred. Instead, we call it precisely after the end of each loop or
		// before the early returns, but when we do so, we do it **after**
		// reading the body content of the response. Otherwise, a `context
		// cancelled` error may happen when the body is read.
		perRequestCtx, cancel := context.WithTimeout(ctx, h.timeout)
		req = req.WithContext(perRequestCtx)
		bodyRes, code, err := t.request(req)

		// Context error only returns a non-nil error upon context
		// cancellation, which is a signal we interpret as an early return.
		// Indeed, we do not want to retry on other hosts if the context is
		// already cancelled.
		if ctx.Err() != nil {
			cancel()
			return err
		}

		switch t.retryStrategy.Decide(h, code, err) {
		case Success:
			err = unmarshalTo(bodyRes, &res)
			cancel()
			return err
		case Failure:
			if bodyRes != nil {
				err = unmarshalToError(bodyRes)
			} else if err == nil {
				err = fmt.Errorf("undefined network error with code: %v", code)
			}
			cancel()
			return err
		default:
			if err != nil {
				intermediateNetworkErrors = append(intermediateNetworkErrors, err)
			} else {
				responseErr := unmarshalToError(bodyRes)
				intermediateNetworkErrors = append(intermediateNetworkErrors, responseErr)
			}
			if bodyRes != nil {
				if err = bodyRes.Close(); err != nil {
					cancel()
					return fmt.Errorf("cannot close response's body before retry: %v", err)
				}
			}
		}

		cancel()
	}

	if exposeIntermediateNetworkErrors {
		return errs.NewNoMoreHostToTryError(intermediateNetworkErrors...)
	}

	return errs.ErrNoMoreHostToTry
}

func (t *Transport) request(req *http.Request) (io.ReadCloser, int, error) {
	debug.Display(req)
	res, err := t.requester.Request(req)
	debug.Display(res)

	if err != nil {
		msg := fmt.Sprintf("cannot perform request:\n\terror=%v\n\tmethod=%s\n\turl=%s", err, req.Method, req.URL)
		nerr, ok := err.(net.Error)
		if ok {
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
		return nil, 0, err
	}

	return res.Body, res.StatusCode, nil
}

func mergeHeaders(defaultHeaders http.Header, extraHeaders map[string]string) http.Header {
	headers := defaultHeaders.Clone()

	for key, value := range extraHeaders {
		headers.Set(key, value)
	}

	return headers
}

func wrapJSONEncoder(in interface{}) io.ReadCloser {
	pr, pw := io.Pipe()
	go func() {
		errEncode := json.NewEncoder(pw).Encode(in)
		errClose := pw.Close()
		if errEncode != nil {
			debug.Printf("cannot JSON encode request body: %v", errEncode)
		}
		if errClose != nil {
			debug.Printf("cannot close JSON encoder writer for request body: %v", errClose)
		}
	}()
	return pr
}

func wrapGZIPEncoder(in io.ReadCloser) io.ReadCloser {
	pr, pw := io.Pipe()
	go func() {
		gw := gzip.NewWriter(pw)
		_, errCopy := io.Copy(gw, in)
		errCloseGw := gw.Close()
		errClosePw := pw.Close()
		if errCopy != nil {
			debug.Printf("cannot GZIP request body: %v", errCopy)
		}
		if errCloseGw != nil {
			debug.Printf("cannot close gzip.Writer of request body: %v", errCloseGw)
		}
		if errClosePw != nil {
			debug.Printf("cannot close io.PipeWriter of request body: %v", errClosePw)
		}
	}()
	return pr
}

func buildRequestWithoutBody(method, url string) (*http.Request, error) {
	return http.NewRequest(method, url, nil)
}

func buildRequestWithBody(method, url string, body interface{}, c compression.Compression) (*http.Request, error) {
	var r io.ReadCloser
	jsonEncoder := wrapJSONEncoder(body)
	switch c {
	case compression.GZIP:
		r = wrapGZIPEncoder(jsonEncoder)
	case compression.None:
		r = jsonEncoder
	default:
		r = jsonEncoder
	}
	return http.NewRequest(method, url, r)
}

func buildRequest(
	c compression.Compression,
	method string,
	host string,
	path string,
	body interface{},
	headers http.Header,
	urlParams map[string]string,
) (req *http.Request, err error) {
	urlStr := "https://" + host + path
	isCompressionEnabled := shouldCompress(c, method, body)

	if body == nil {
		req, err = buildRequestWithoutBody(method, urlStr)
	} else {
		req, err = buildRequestWithBody(method, urlStr, body, c)
	}

	if err != nil {
		return nil, fmt.Errorf("cannot instantiate request:\n\terr=%v\n\tmethod=%s\n\turl=%s\n\tbody=%#v", err, method, urlStr, body)
	}

	// Add headers
	req.Header = headers

	// Add Content-Encoding header, if needed
	if isCompressionEnabled {
		switch c {
		case compression.GZIP:
			req.Header.Set("Content-Encoding", "gzip")
		default:
			// Do nothing
		}
	}

	// Add URL params
	values := req.URL.Query()
	for k, v := range urlParams {
		values.Set(k, v)
	}
	req.URL.RawQuery = values.Encode()

	return req, nil
}

func unmarshalTo(r io.ReadCloser, v interface{}) error {
	body, err := ioutil.ReadAll(r)
	errClose := r.Close()
	if err != nil {
		return fmt.Errorf("cannot read body: %v", err)
	}
	err = json.Unmarshal(body, &v)
	if err != nil {
		return fmt.Errorf("cannot deserialize response's body: %v: %s", err, string(body))
	}
	if errClose != nil {
		return fmt.Errorf("cannot close response's body: %v", errClose)
	}
	return nil
}

func unmarshalToError(r io.ReadCloser) error {
	var algoliaErr errs.AlgoliaErr
	err := unmarshalTo(r, &algoliaErr)
	if err != nil {
		return err
	}
	return &algoliaErr
}

func shouldCompress(c compression.Compression, method string, body interface{}) bool {
	isValidMethod := method == http.MethodPut || method == http.MethodPost
	isCompressionEnabled := c != compression.None
	isBodyNonEmpty := body != nil
	return isCompressionEnabled && isValidMethod && isBodyNonEmpty
}
