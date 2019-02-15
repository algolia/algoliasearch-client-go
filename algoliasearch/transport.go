package algoliasearch

import (
	"bytes"
	"context"
	_ "crypto/sha512" // Fix certificates
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/algolia/algoliasearch-client-go/algoliasearch/call"
)

const (
	version = "2.27.0"
)

// Define the constants used to specify the type of request.
const (
	search = 1 << iota
	write
	read
	analyticsCall
	insightsCall
)

// Transport is responsible for the connection and the retry strategy to
// Algolia servers.
type Transport struct {
	headers       map[string]string
	httpClient    *http.Client
	retryStrategy RetryStrategy
}

const (
	DefaultConnectTimeout      = 2 * time.Second
	DefaultKeepAliveDuration   = 5 * time.Minute
	DefaultMaxIdleConnsPerHost = 64
)

// NewTransport instantiates a new Transport with the default Algolia hosts to
// connect to.
func NewTransport(appID, apiKey string) *Transport {
	return NewTransportWithHosts(appID, apiKey, nil)
}

// NewTransport instantiates a new Transport with the specificed hosts as main
// servers to connect to.
func NewTransportWithHosts(appID, apiKey string, hosts []string) *Transport {
	return &Transport{
		headers: map[string]string{
			"Connection":               "keep-alive",
			"User-Agent":               fmt.Sprintf("Algolia for Go (%s); Go (%s); ", version, runtime.Version()),
			"X-Algolia-Application-Id": appID,
			"X-Algolia-API-Key":        apiKey,
		},
		httpClient: &http.Client{
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					KeepAlive: DefaultKeepAliveDuration,
					Timeout:   DefaultConnectTimeout,
				}).Dial,
				DisableKeepAlives:   false,
				MaxIdleConnsPerHost: DefaultMaxIdleConnsPerHost,
				Proxy:               http.ProxyFromEnvironment,
				TLSHandshakeTimeout: DefaultConnectTimeout,
			},
		},
		retryStrategy: NewRetryStrategy(appID, hosts),
	}
}

// request is the method used by the `Client` to perform the request against
// the Algolia servers (or to the list of specified hosts).
func (t *Transport) request(method, path string, body interface{}, typeCall int, opts *RequestOptions) ([]byte, error) {
	var k call.Kind
	switch typeCall {
	case search, read:
		k = call.Read
	case write:
		k = call.Write
	case analyticsCall:
		k = call.Analytics
	case insightsCall:
		k = call.Insights
	default:
		return nil, fmt.Errorf("unsupported call type %d", typeCall)
	}

	for _, h := range t.retryStrategy.GetTryableHosts(k) {
		req, err := t.buildRequest(method, h.Host(), path, body, opts)
		if err != nil {
			return nil, err
		}

		debug("* REQUEST [%s] url=%s", method, req.URL)
		bodyRes, code, err := t.do(req, h.Timeout())
		debug("* RESPONSE [%d] err=%v body=%s", code, err, bodyRes)

		switch t.retryStrategy.Decide(h, code, err) {
		case Success:
			return bodyRes, err
		case Failure:
			return nil, errors.New(string(bodyRes))
		}
	}

	return nil, ExhaustionOfTryableHostsErr
}

func (t *Transport) buildRequest(method, host, path string, body interface{}, opts *RequestOptions) (*http.Request, error) {
	var req *http.Request
	var err error

	urlStr := "https://" + host + path

	if body == nil {
		// As the body is nil, an empty body request is instantiated
		req, err = buildRequestWithEmptyBody(method, urlStr)
	} else {
		// If the body is non-nil and the HTTP method is GET, the body request
		// is translated into HTTP query parameters (needed for the
		// `Client.GetLogs` for instance.
		if method == "GET" {
			req, err = buildRequestWithURLParameters(method, urlStr, body)
		} else {
			req, err = buildRequestWithBodyParameters(method, urlStr, body)
		}
	}

	if err != nil {
		return nil, err
	}

	if strings.Contains(path, "/*/") {
		req.URL = &url.URL{
			Scheme: "https",
			Host:   host,
			Opaque: "//" + host + path, // Remove url encoding
		}
	}

	addHeaders(req, t.headers)

	if opts != nil {
		addHeaders(req, opts.ExtraHeaders)
		addHeaders(req, map[string]string{"X-Forwarded-For": opts.ForwardedFor})
		addUrlParameters(req, opts.ExtraUrlParams)
	}

	return req, nil
}

func (t *Transport) do(req *http.Request, timeout time.Duration) ([]byte, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req = req.WithContext(ctx)

	res, err := t.httpClient.Do(req)
	if err != nil {
		msg := fmt.Sprintf("cannot perform request %s %s: %s", req.Method, req.URL, err)
		nerr, ok := err.(net.Error)
		if ok {
			// Because net.Error and error have different meanings for the
			// retry strategy, we cannot simply return a new fmt.Errorf, which
			// would make all net.Error simple errors instead. To keep this
			// behavior, we wrap the message into a custom NetError that
			// implements the net.Error interface if the original error was
			// already a net.Error.
			return nil, 0, NewNetError(nerr, msg)
		} else {
			return nil, 0, errors.New(msg)
		}
	}
	defer res.Body.Close()

	bodyRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, 0, fmt.Errorf("cannot read response: %s", err)
	}

	return bodyRes, res.StatusCode, nil
}

// setExtraHeader lets the user (through the exported `Client.SetExtraHeader`)
// add custom headers to the requests.
func (t *Transport) setExtraHeader(key, value string) {
	t.headers[key] = value
}

func (t *Transport) setTimeouts(read, write, analytics, insights time.Duration) {
	t.retryStrategy.SetTimeouts(read, write, analytics, insights)
}

// setMaxIdleConnsPerHost sets the `MaxIdleConnsPerHost` via the given
// `perHosts` value of the underlying RoundTripper of the HTTP client if it is
// an instance of `http.Transport`.
func (t *Transport) setMaxIdleConnsPerHost(maxIdleConnsPerHost int) {
	switch transport := t.httpClient.Transport.(type) {
	case (*http.Transport):
		transport.MaxIdleConnsPerHost = maxIdleConnsPerHost
		t.httpClient.Transport = transport
	}
}

// addHeaders adds the key/value pairs from `headers` to the header list of the
// `req` request.
func addHeaders(req *http.Request, headers map[string]string) {
	for k, v := range headers {
		req.Header.Add(k, v)
	}
}

// addUrlParameters adds the key/value pairs from `params` to the URL query
// parameter list of the `req` request.
func addUrlParameters(req *http.Request, params map[string]string) {
	q := req.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
}

// buildRequestWithEmptyBody returns a new `http.Request` for the given
// HTTP method and url whose body is empty. If the request could not have been
// instantiated correctly, a non-nil error is returned.
func buildRequestWithEmptyBody(method, url string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("Cannot instantiate request: [%s] %s", method, url)
	}
	return req, nil
}

// buildRequestWithURLParameters returns a new `http.Request` for the given
// HTTP method and url whose body is empty but the URL parameters are filled
// with the values from the given body (which must be an `algoliasearch.Map`).
// If the request could not have been instantiated correctly, a non-nil error
// is returned.
func buildRequestWithURLParameters(method, url string, body interface{}) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("Cannot instantiate request: [%s] %s", method, url)
	}

	params, ok := body.(Map)
	if !ok {
		return nil, fmt.Errorf("Cannot instantiate request: GET request has non-Map body")
	}

	values := req.URL.Query()
	var value string
	for k, v := range params {
		value = fmt.Sprintf("%v", v)
		values.Set(k, value)
	}
	req.URL.RawQuery = values.Encode()

	return req, nil
}

// buildRequestWithBodyParameters returns a new `http.Request` for the given
// HTTP method and url whose body is filled with the given body `interface{}`.
// If the request could not have been instantiated correctly, a non-nil error
// is returned.
func buildRequestWithBodyParameters(method, url string, body interface{}) (*http.Request, error) {
	// As the body is non-nil, the content is read
	data, err := json.Marshal(body)
	if err != nil {
		return nil, errors.New("Invalid JSON in the query")
	}
	reader := bytes.NewReader(data)

	// The request is then instantiated with the body content
	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		return nil, fmt.Errorf("Cannot instantiate request: [%s] %s", method, url)
	}

	// Add content specific headers
	req.Header.Add("Content-Length", strconv.Itoa(len(string(data))))
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	return req, nil
}
