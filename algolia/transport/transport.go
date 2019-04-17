package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"runtime"
	"time"

	"github.com/algolia/algoliasearch-client-go/algolia/call"
	"github.com/algolia/algoliasearch-client-go/algolia/debug"
	"github.com/algolia/algoliasearch-client-go/algolia/errs"
	iopt "github.com/algolia/algoliasearch-client-go/algolia/internal/opt"
)

const version = "3.0.0"

type Transport struct {
	requester     Requester
	retryStrategy *RetryStrategy
	headers       map[string]string
}

func New(
	hosts []*StatefulHost,
	requester Requester,
	appID string,
	apiKey string,
	readTimeout time.Duration,
	writeTimeout time.Duration,
	defaultHeaders map[string]string,
) *Transport {

	if requester == nil {
		requester = newDefaultRequester()
	}

	headers := map[string]string{
		"Connection":               "Keep-Alive",
		"Content-Type":             "application/json; charset=utf-8",
		"User-Agent":               fmt.Sprintf("Algolia for Go (%s); Go (%s); ", version, runtime.Version()),
		"X-Algolia-Application-Id": appID,
		"X-Algolia-API-Key":        apiKey,
	}

	for k, v := range defaultHeaders {
		headers[k] = v
	}

	return &Transport{
		requester:     requester,
		retryStrategy: newRetryStrategy(hosts, readTimeout, writeTimeout),
		headers:       headers,
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
		headers   = t.headers
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

	for _, h := range t.retryStrategy.GetTryableHosts(k) {
		req, err := buildRequest(method, h.host, path, body, headers, urlParams)
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
		body, code, err := t.request(req)

		switch t.retryStrategy.Decide(h, code, err) {
		case Success:
			err = unmarshalTo(body, &res)
			cancel()
			return err
		case Failure:
			err = unmarshalToError(body)
			cancel()
			return err
		default:
			if body != nil {
				if err = body.Close(); err != nil {
					cancel()
					return fmt.Errorf("cannot close response's body before retry: %v", err)
				}
			}
		}

		cancel()
	}

	return errs.NoMoreHostToTry
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

func mergeHeaders(defaultHeaders, extraHeaders map[string]string) map[string]string {
	headers := make(map[string]string)

	for key, value := range defaultHeaders {
		headers[key] = value
	}
	for key, value := range extraHeaders {
		headers[key] = value
	}

	return headers
}

func buildRequest(
	method string,
	host string, path string,
	body interface{},
	headers map[string]string,
	urlParams map[string]string,
) (req *http.Request, err error) {

	urlStr := "https://" + host + path

	// Build the body payload if the body is not empty
	if body == nil {
		req, err = http.NewRequest(method, urlStr, nil)
	} else {
		var data []byte
		data, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("cannot serialize request body:\n\terr=%v\n\tmethod=%s\n\turl=%s\n\tbody=%#v\n", err, method, urlStr, body)
		}
		req, err = http.NewRequest(method, urlStr, bytes.NewReader(data))
	}

	if err != nil {
		return nil, fmt.Errorf("cannot instantiate request:\n\terr=%v\n\tmethod=%s\n\turl=%s\n\tbody=%#v\n", err, method, urlStr, body)
	}

	// Add headers
	for k, v := range headers {
		req.Header.Add(k, v)
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
	err := json.NewDecoder(r).Decode(&v)
	errClose := r.Close()
	if err != nil {
		return fmt.Errorf("cannot deserialize response's body: %v", err)
	}
	if errClose != nil {
		return fmt.Errorf("cannot close response's body: %v", errClose)
	}
	return nil
}

func unmarshalToError(r io.ReadCloser) error {
	var algoliaErr errs.EngineError
	err := unmarshalTo(r, &algoliaErr)
	if err != nil {
		return err
	}
	return algoliaErr
}
