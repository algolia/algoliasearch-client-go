package algoliasearch

import (
	"bytes"
	_ "crypto/sha512" // Fix certificates
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
	version = "2.0.0"
)

const (
	search = 1 << iota
	write
	read
)

// Transport defines low level functions to trade with Algolia servers.
type Transport struct {
	apiKey        string
	appID         string
	headers       map[string]string
	hosts         []string
	hostsProvided bool
	httpClient    *http.Client
}

// newHTTPClient creates and initializes an http.Client with sane defaults.
func newHTTPClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 30,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   2 * time.Second,
				KeepAlive: 180 * time.Second,
			}).Dial,
			DisableKeepAlives:   false,
			MaxIdleConnsPerHost: 2,
			TLSHandshakeTimeout: 2 * time.Second,
		},
	}
}

// NewTransport creates and initializes a new Transport targeting the Algolia
// application `appID` using the API key `apiKey`. The hosts are deduced from
// `appID`.
func NewTransport(appID, apiKey string) *Transport {
	hosts := []string{
		appID + "-1.algolianet.com",
		appID + "-2.algolianet.com",
		appID + "-3.algolianet.com",
	}

	// Randomize the hosts order
	randHosts := make([]string, len(hosts))
	for i, v := range rand.Perm(len(hosts)) {
		randHosts[i] = hosts[v]
	}

	return &Transport{
		apiKey:        apiKey,
		appID:         appID,
		headers:       make(map[string]string),
		hosts:         randHosts,
		hostsProvided: false,
		httpClient:    newHTTPClient(),
	}
}

// NewTransportWithHosts creates and initializes a new Transport targeting the
// Algolia application `appID` using the API key `apiKey` via the specified
// `hosts`.
func NewTransportWithHosts(appID, apiKey string, hosts []string) *Transport {
	return &Transport{
		apiKey:        apiKey,
		appID:         appID,
		headers:       make(map[string]string),
		hosts:         hosts,
		hostsProvided: true,
		httpClient:    newHTTPClient(),
	}
}

// setTimeout changes the timeouts used by the underlying HTTP client.
func (t *Transport) setTimeout(connectTimeout, readTimeout time.Duration) {
	switch transport := t.httpClient.Transport.(type) {
	case *http.Transport:
		transport.TLSHandshakeTimeout = connectTimeout
		transport.ResponseHeaderTimeout = readTimeout
	default:
		fmt.Fprintln(os.Stderr, "Timeouts not set for nonstandard underlying Transport")
	}
}

// setExtraHeader adds a custom header to be used when exchanging with Algolia
// servers.
func (t *Transport) setExtraHeader(key, value string) {
	t.headers[key] = value
}

// request performs a `method` HTTP request at `path` sending `body`.
// `typeCall` represents the operation intended on the Algolia servers and can
// be one of the following constants: `search`, `write` or `read`.
func (t *Transport) request(method, path string, body interface{}, typeCall int) ([]byte, error) {
	var host string
	errorMsg := ""
	if typeCall == write {
		host = t.appID + ".algolia.net"
	} else {
		host = t.appID + "-dsn.algolia.net"
	}

	if !t.hostsProvided {
		req, err := t.buildRequest(method, host, path, body)
		if err != nil {
			return nil, err
		}

		req = t.addHeaders(req)
		resp, err := t.httpClient.Do(req)

		if err != nil {
			if len(errorMsg) > 0 {
				errorMsg = fmt.Sprintf("%s, %s:%s", errorMsg, host, err)
			} else {
				errorMsg = fmt.Sprintf("%s:%s", host, err)
			}
		} else if (resp.StatusCode/100) == 2 || (resp.StatusCode/100) == 4 { // Bad request, not found, forbidden
			return t.handleResponse(resp)
		} else {
			io.Copy(ioutil.Discard, resp.Body)
			resp.Body.Close()
		}
	}

	for _, host := range t.hosts {
		req, err := t.buildRequest(method, host, path, body)
		if err != nil {
			return nil, err
		}

		req = t.addHeaders(req)
		resp, err := t.httpClient.Do(req)

		if err != nil {
			if len(errorMsg) > 0 {
				errorMsg = fmt.Sprintf("%s, %s:%s", errorMsg, host, err)
			} else {
				errorMsg = fmt.Sprintf("%s:%s", host, err)
			}
			continue
		}

		if (resp.StatusCode/100) == 2 || (resp.StatusCode/100) == 4 { // Bad request, not found, forbidden
			return t.handleResponse(resp)
		}
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}

	return nil, fmt.Errorf("Cannot reach any host. (%s)", errorMsg)
}

// buildRequest builds an http.Request object. The built request uses `method`,
// tries to reach `path` at `host`, sending `body`.
func (t *Transport) buildRequest(method, host, path string, body interface{}) (*http.Request, error) {
	var req *http.Request
	var err error

	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return nil, errors.New("Invalid JSON in the query")
		}

		reader := bytes.NewReader(bodyBytes)
		req, err = http.NewRequest(method, "https://"+host+path, reader)
		req.Header.Add("Content-Length", strconv.Itoa(len(string(bodyBytes))))
		req.Header.Add("Content-Type", "application/json; charset=utf-8")
	} else {
		req, err = http.NewRequest(method, "https://"+host+path, nil)
	}

	if strings.Contains(path, "/*/") {
		req.URL = &url.URL{
			Scheme: "https",
			Host:   host,
			Opaque: "//" + host + path, //Remove url encoding
		}
	}

	return req, err
}

// addHeaders adds the mandatory Algolia headers and the custom headers of `t`
// to `req`.
func (t *Transport) addHeaders(req *http.Request) *http.Request {
	req.Header.Add("X-Algolia-API-Key", t.apiKey)
	req.Header.Add("X-Algolia-Application-Id", t.appID)
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("User-Agent", "Algolia for go "+version)

	for k, v := range t.headers {
		req.Header.Add(k, v)
	}

	return req
}

// handleResponse takes care of reading a response as JSON, and returns the
// parsed object. If the status code of the response indicates a failed
// request, or if the body of the response is not a valid JSON object, an error
// is returned.
func (t *Transport) handleResponse(resp *http.Response) ([]byte, error) {
	res, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var jsonResp interface{}
	if err = json.Unmarshal(res, &jsonResp); err != nil {
		return nil, fmt.Errorf("Invalid JSON response: %s", err)
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return res, nil
	} else {
		return nil, errors.New(string(res))
	}
}
