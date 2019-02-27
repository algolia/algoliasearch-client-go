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
	"net/url"
	"runtime"
	"strings"
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
		headers   = mergeHeaders(t.headers, iopt.ExtractExtraHeaders(opts...).Get())
		urlParams = iopt.ExtractExtraURLParams(opts...).Get()
	)

	for _, h := range t.retryStrategy.GetTryableHosts(k) {
		req, err := buildRequest(method, h.host, path, body, headers, urlParams)
		if err != nil {
			return err
		}

		// Handle per-request timeout by using a context with timeout
		perRequestCtx, cancel := context.WithTimeout(ctx, h.timeout)
		req = req.WithContext(perRequestCtx)
		body, code, err := t.request(req)
		cancel()

		switch t.retryStrategy.Decide(h, code, err) {
		case Success:
			return unmarshalTo(body, &res)
		case Failure:
			return unmarshalToError(body)
		default:
			if body != nil {
				if err = body.Close(); err != nil {
					return fmt.Errorf("cannot close response's body before retry: %v", err)
				}
			}
		}
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
		data, err := json.Marshal(body)
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
		values.Set(k, fmt.Sprintf("%v", v))
	}
	req.URL.RawQuery = values.Encode()

	// TODO: Understand if this is still useful
	if strings.Contains(path, "/*/") {
		req.URL = &url.URL{
			Scheme: "https",
			Host:   host,
			Opaque: "//" + host + path, // Remove url encoding
		}
	}

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
	var msg string
	err := unmarshalTo(r, &msg)
	if err != nil {
		return err
	}
	return errors.New(msg)
}
