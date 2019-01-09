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
	"github.com/algolia/algoliasearch-client-go/algolia/errs"
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
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
		ctx            = opt.ExtractContext(opts...)
		requestOptions = opt.ExtractRequestOptions(opts...)
		headers        = mergeHeaders(t.headers, requestOptions.ExtraHeaders)
		urlParams      = requestOptions.ExtraURLParams
	)

	for _, h := range t.retryStrategy.GetTryableHosts(k) {
		req, err := buildRequest(ctx, method, h.host, path, body, headers, urlParams, h.timeout)
		if err != nil {
			return err
		}

		body, code, err := t.request(req)

		switch t.retryStrategy.Decide(h, code, err) {
		case Success:
			return unmarshalTo(body, &res)
		case Failure:
			return unmarshalToError(body)
		default:
			_ = body.Close()
		}
	}

	return errs.NoMoreHostToTry
}

func (t *Transport) request(req *http.Request) (io.ReadCloser, int, error) {
	res, err := t.requester.Request(req)

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
			return nil, 0, errs.NetError(nerr, msg)
		} else {
			return nil, 0, errors.New(msg)
		}
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
	ctx context.Context,
	method string,
	host string, path string,
	body interface{},
	headers map[string]string,
	urlParams map[string]string,
	timeout time.Duration,
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

	// Handle per-request timeout by using a context with timeout
	ctx, _ = context.WithTimeout(ctx, timeout)
	req = req.WithContext(ctx)

	return req, nil
}

func unmarshalTo(r io.ReadCloser, v interface{}) error {
	defer r.Close()
	err := json.NewDecoder(r).Decode(&v)
	if err != nil {
		return fmt.Errorf("cannot deserialize response: %v", err)
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