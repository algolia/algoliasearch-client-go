package algoliasearch

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransport_BuildRequest(t *testing.T) {
	t.Log("TestTransport_BuildRequest: Initialize the transport layer")
	appid := "appid"
	apikey := "apikey"
	transport := NewTransport(appid, apikey)

	t.Log("TestTransport_BuildRequest: Set up all request's parameters")
	method := "POST"
	host := "APPID.algolia.net"
	path := "/1/some/path"
	body := map[string]interface{}{
		"one": 1.0,
		"two": "two",
	}
	opts := &RequestOptions{
		ForwardedFor: "127.0.0.1",
		ExtraHeaders: map[string]string{
			"extra-header-1": "header-value-1",
			"extra-header-2": "header-value-2",
		},
		ExtraUrlParams: map[string]string{
			"url-param-key-1": "url-param-value-1",
			"url-param-key-2": "url-param-value-2",
		},
	}

	req, err := transport.buildRequest(method, host, path, body, opts)
	require.Nil(t, err, "should build a new request without error")

	t.Log("TestTransport_BuildRequest: Check URL")
	require.NotNil(t, req.URL, "should populate the URL field")
	require.Equal(t, host, req.URL.Host, "should hold the correct URL host")
	require.Equal(t, path, req.URL.Path, "should hold the correct URL path")
	require.Equal(t, "url-param-key-1=url-param-value-1&url-param-key-2=url-param-value-2", req.URL.RawQuery, "should hold the correct URL path")

	t.Log("TestTransport_BuildRequest: Check headers")
	require.Contains(t, req.Header, "User-Agent", "missing header")
	checkHeader(t, "X-Algolia-Application-Id", appid, req.Header)
	checkHeader(t, "X-Algolia-Api-Key", apikey, req.Header)
	checkHeader(t, "X-Forwarded-For", opts.ForwardedFor, req.Header)

	for header, value := range opts.ExtraHeaders {
		checkHeader(t, header, value, req.Header)
	}

	t.Log("TestTransport_BuildRequest: Check HTTP method")
	require.Equal(t, method, req.Method, "should use the expected HTTP method")

	t.Log("TestTransport_BuildRequest: Check body")

	var readBody map[string]interface{}
	err = json.NewDecoder(req.Body).Decode(&readBody)
	require.Nil(t, err, "should decode the body without any error")
	require.Equal(t, len(body), len(readBody), "body should contain the correct number of elements")
	for key, value := range body {
		require.Contains(t, readBody, key)
		require.Equal(t, value, readBody[key], "body should contain the expected key/value pairs")
	}
}

func checkHeader(t *testing.T, header, value string, headers http.Header) {
	header = strings.Title(header)
	require.Contains(t, headers, header)
	require.Equal(t, 1, len(headers[header]), "header value slice should only contain one element")
	require.Equal(t, value, headers[header][0], "header should have the correct value")
}
