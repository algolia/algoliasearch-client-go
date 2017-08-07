package algoliasearch

import (
	"bytes"
	_ "crypto/sha512" // Fix certificates
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	version = "2.17.0"
)

// Define the constants used to specify the type of request.
const (
	search = 1 << iota
	write
	read
)

// Seed the RNG used to shuffle the hosts slice (see `defaultHosts` function).
func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

// Transport is responsible for the connection and the retry strategy to
// Algolia servers.
type Transport struct {
	activeReadHost    string
	activeReadSince   time.Time
	activeWriteHost   string
	activeWriteSince  time.Time
	apiKey            string
	appId             string
	dialTimeout       time.Duration
	headers           map[string]string
	httpClient        *http.Client
	keepAliveDuration time.Duration
	providedHosts     []string
}

// NewTransport instantiates a new Transport with the default Algolia hosts to
// connect to.
func NewTransport(appId, apiKey string) *Transport {
	return &Transport{
		activeReadHost:    "",
		activeWriteHost:   "",
		apiKey:            apiKey,
		appId:             appId,
		dialTimeout:       1 * time.Second,
		headers:           defaultHeaders(appId, apiKey),
		httpClient:        defaultHttpClient(),
		keepAliveDuration: 5 * time.Minute,
		providedHosts:     nil,
	}
}

// NewTransport instantiates a new Transport with the specificed hosts as main
// servers to connect to.
func NewTransportWithHosts(appId, apiKey string, hosts []string) *Transport {
	return &Transport{
		activeReadHost:    "",
		activeWriteHost:   "",
		apiKey:            apiKey,
		appId:             appId,
		dialTimeout:       1 * time.Second,
		headers:           defaultHeaders(appId, apiKey),
		httpClient:        defaultHttpClient(),
		keepAliveDuration: 5 * 60 * time.Second,
		providedHosts:     hosts,
	}
}

// defaultHeaders is used to set the default HTTP headers to use with each
// requests.
func defaultHeaders(appId, apiKey string) map[string]string {
	return map[string]string{
		"Connection":               "keep-alive",
		"User-Agent":               "Algolia for Go (" + version + ")",
		"X-Algolia-API-Key":        apiKey,
		"X-Algolia-Application-Id": appId,
	}
}

// defaultHosts returns the list of the default Algolia hosts to use. The
// entries are shuffled.
func (t *Transport) defaultHosts() []string {
	hosts := []string{
		t.appId + "-1.algolianet.com",
		t.appId + "-2.algolianet.com",
		t.appId + "-3.algolianet.com",
	}

	shuffled := make([]string, len(hosts))
	for i, v := range rand.Perm(len(hosts)) {
		shuffled[i] = hosts[v]
	}

	return shuffled
}

// defaultHttpClient returns the `*http.Client` which will perform all the
// requests. All the timeout settings are explicitely defined here.
func defaultHttpClient() *http.Client {
	return &http.Client{
		Timeout:   time.Second * 30,
		Transport: defaultTransport(1 * time.Second),
	}
}

// defaultTransport returns the `*http.Transport` which starts and maintain the
// connection with the server. The `dialTimeout` is used to specify the timeout
// beyond which the connection is considered as failed (used to control DNS
// lookup timeouts).
func defaultTransport(dialTimeout time.Duration) *http.Transport {
	return &http.Transport{
		Proxy:               http.ProxyFromEnvironment,
		Dial:                defaultDial(dialTimeout).Dial,
		DisableKeepAlives:   false,
		MaxIdleConnsPerHost: 2,
		TLSHandshakeTimeout: 2 * time.Second,
	}
}

// defaultDial returns the `*net.Dialer` which will connect to the hosts
// according to the given timeout.
func defaultDial(dialTimeout time.Duration) *net.Dialer {
	return &net.Dialer{
		KeepAlive: 180 * time.Second,
		Timeout:   dialTimeout,
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

// setExtraHeader lets the user (through the exported `Client.SetExtraHeader`)
// add custom headers to the requests.
func (t *Transport) setExtraHeader(key, value string) {
	t.headers[key] = value
}

// setTimeout lets the user (through the exported `Client.SetTimeout`) replace
// the default values of `TLSHandshakeTimeout` (via `connectTimeout`) and
// `ResponseHeaderTimeout` (via `readTimeout`).
func (t *Transport) setTimeout(connectTimeout, readTimeout time.Duration) {
	switch transport := t.httpClient.Transport.(type) {
	case *http.Transport:
		transport.TLSHandshakeTimeout = connectTimeout
		transport.ResponseHeaderTimeout = readTimeout
	default:
		fmt.Fprintln(os.Stderr, "Timeouts not set for nonstandard underlying Transport")
	}
}

// request is the method used by the `Client` to perform the request against
// the Algolia servers (or to the list of specified hosts).
func (t *Transport) request(method, path string, body interface{}, typeCall int, opts *RequestOptions) ([]byte, error) {
	var res []byte
	var err error

	for _, host := range t.hostsToTry(typeCall) {
		res, err = t.tryRequest(method, host, path, body, opts)
		if err == nil {
			t.resetDialTimeout()
			if typeCall == write {
				t.activeWriteSince = time.Now()
				t.activeWriteHost = host
			} else {
				t.activeReadSince = time.Now()
				t.activeReadHost = host
			}
			return res, nil
		}
		t.increaseDialTimeout()
	}

	if typeCall == write {
		t.activeWriteHost = ""
	} else {
		t.activeReadHost = ""
	}

	return nil, err
}

// hostsToTry returns the list of hosts to try ordered by priority according to
// the type of request (write vs. read/search) and if a previous host was
// marked as active.
func (t *Transport) hostsToTry(typeCall int) []string {
	var hosts []string

	// Step 1:
	//
	// We set the first host to try to the last active one if any and
	// if it was active recently.

	if typeCall == write {
		// In case the request is a write query, we put the last active write
		// host first in the list of hosts to try if it was used in the last
		// `keepAliveDuration` seconds. We then put the main algolia.net host.
		if t.activeWriteHost != "" &&
			time.Now().Sub(t.activeWriteSince) <= t.keepAliveDuration {
			hosts = []string{t.activeWriteHost}
		}
	} else {
		// In case the request is not a write query, we put the last active
		// read host first in the list of hosts to try if it was used in the
		// last `keepAliveDuration` seconds. We then put the DSN host.
		if t.activeReadHost != "" &&
			time.Now().Sub(t.activeReadSince) <= t.keepAliveDuration {
			hosts = []string{t.activeReadHost}
		}
	}

	// Step 2:
	//
	// If the hosts were provided we use them first to make sure they are tried
	// first. Otherwise, we use put the default ones after the ones already
	// generated.

	if len(t.providedHosts) > 0 {
		hosts = append(hosts, t.providedHosts...)
	}

	// Step 3:
	//
	// The main host is added to the list, along with the default ones.

	if typeCall == write {
		hosts = append(hosts, t.appId+".algolia.net")
	} else {
		hosts = append(hosts, t.appId+"-dsn.algolia.net")
	}
	hosts = append(hosts, t.defaultHosts()...)

	return hosts
}

// tryRequest is the underlying method which actually performs the request. It
// returns the response as a byte slice or a non-nil error if anything went
// wrong.
func (t *Transport) tryRequest(method, host, path string, body interface{}, opts *RequestOptions) ([]byte, error) {
	// Build the request
	req, err := t.buildRequest(method, host, path, body, opts)
	if err != nil {
		return nil, err
	}

	// Perform the request
	res, err := t.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Cannot perform request [%s] %s (%s): %s", method, path, host, err)
	}
	defer res.Body.Close()

	// Read response's body
	bodyRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Cannot read response body: %s", err)
	}

	// Return the body as an error if the status code is not 2XX
	code := res.StatusCode
	if !(200 <= code && code < 300) {
		return nil, errors.New(string(bodyRes))
	}

	return bodyRes, nil
}

// buildRequest returns a valid `http.Request` with the headers and body (if
// any) correctly set. The return error is non-nil if the request is invalid or
// if the body, if non-nil, is not a valid JSON.
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

	// Add default and Algolia specific headers
	addHeaders(req, t.headers)

	if strings.Contains(path, "/*/") {
		req.URL = &url.URL{
			Scheme: "https",
			Host:   host,
			Opaque: "//" + host + path, // Remove url encoding
		}
	}

	// Add extra headers and URL parameters if a `RequestOptions` is provided
	if opts != nil {
		addHeaders(req, opts.ExtraHeaders)
		addHeaders(req, map[string]string{"X-Forwarded-For": opts.ForwardedFor})
		addUrlParameters(req, opts.ExtraUrlParams)
	}

	return req, nil
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

// resetDialTimeout increases the `Timeout` value of the underlying dialer by 1
// second if the underyling RoundTripper of the HTTP client is an instance of
// http.Transport.
func (t *Transport) increaseDialTimeout() {
	t.dialTimeout = t.dialTimeout + time.Second
	t.setDialTimeout(t.dialTimeout)
}

// resetDialTimeout resets the `Timeout` value of the underlying dialer to 1
// second if the underyling RoundTripper of the HTTP client is an instance of
// http.Transport.
func (t *Transport) resetDialTimeout() {
	t.dialTimeout = 1 * time.Second
	t.setDialTimeout(t.dialTimeout)
}

// setDialTimeout sets the `Timeout` value of the underyling dialer to the
// given value if the underlying RoundTripper of the HTTP client is an instance
// of http.Transport.
func (t *Transport) setDialTimeout(dialTimeout time.Duration) {
	switch transport := t.httpClient.Transport.(type) {
	case (*http.Transport):
		transport.Dial = defaultDial(dialTimeout).Dial
		t.httpClient.Transport = transport
	default:
		// Do nothing if the HTTP client was overriden and the RoundTripper is
		// not an instance of http.Transport.
	}
}
