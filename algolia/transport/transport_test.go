package transport

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/call"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/compression"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/errs"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
)

func TestShouldCompress(t *testing.T) {
	fakeBody := map[string]interface{}{}
	for _, c := range []struct {
		compression compression.Compression
		method      string
		body        interface{}
		expected    bool
	}{
		{0, "", nil, false},
		{compression.None, "", nil, false},
		{compression.None, http.MethodPost, fakeBody, false},
		{compression.None, http.MethodPut, fakeBody, false},
		{compression.None, http.MethodDelete, fakeBody, false},
		{compression.None, http.MethodGet, fakeBody, false},
		{compression.GZIP, http.MethodPost, fakeBody, true},
		{compression.GZIP, http.MethodPut, fakeBody, true},
		{compression.GZIP, http.MethodDelete, fakeBody, false},
		{compression.GZIP, http.MethodGet, fakeBody, false},
		{compression.GZIP, http.MethodPost, nil, false},
		{compression.GZIP, http.MethodPut, nil, false},
	} {
		got := shouldCompress(c.compression, c.method, c.body)
		require.Equal(t, c.expected, got, "compression=%d method=%q", c.compression, c.method)
	}
}

func TestShouldNotExposeIntermediateNetworkErrors(t *testing.T) {
	hosts := []*StatefulHost{NewStatefulHost("", call.IsReadWrite)}
	requester := newDefaultRequester()
	transporter := New(
		hosts,
		requester,
		"appID",
		"apiKey",
		time.Second,
		time.Second,
		nil,
		"",
		compression.None,
	)
	var res string
	err := transporter.Request(&res, http.MethodGet, "", nil, call.Read)
	require.Equal(t, errs.ErrNoMoreHostToTry, err)
}

func TestShouldExposeIntermediateNetworkErrors(t *testing.T) {
	hosts := []*StatefulHost{NewStatefulHost("", call.IsReadWrite)}
	requester := newDefaultRequester()
	transporter := New(
		hosts,
		requester,
		"appID",
		"apiKey",
		time.Second,
		time.Second,
		nil,
		"",
		compression.None,
	)
	opts := []interface{}{opt.ExposeIntermediateNetworkErrors(true)}
	var res string
	err := transporter.Request(&res, http.MethodGet, "", nil, call.Read, opts...)
	noMoreHostToTryErr, ok := err.(*errs.NoMoreHostToTryErr)
	require.True(t, ok)
	require.Len(t, noMoreHostToTryErr.IntermediateNetworkErrors(), 1)
}

func TestUnmarshallTo(t *testing.T) {
	type fakeStruct struct {
		Attr string `json:"attr"`
	}
	for _, v := range []struct {
		bodyRaw       string
		expectedError error
		expectedBody  fakeStruct
	}{
		{"<html>Non json answer</html>", fmt.Errorf("cannot deserialize response's body: invalid character '<' looking for beginning of value: <html>Non json answer</html>"), fakeStruct{}},
		{`{"attr":"value"}`, nil, fakeStruct{"value"}},
	} {
		bodyDeserialized := fakeStruct{}
		err := unmarshalTo(ioutil.NopCloser(bytes.NewReader([]byte(v.bodyRaw))), &bodyDeserialized)
		require.Equal(t, v.expectedError, err)
		require.Equal(t, v.expectedBody, bodyDeserialized)
	}
}

type FakeNetworkError struct {
	Description string
}

func (e *FakeNetworkError) Error() string {
	return e.Description
}

func (e *FakeNetworkError) Timeout() bool {
	return false
}

func (e *FakeNetworkError) Temporary() bool {
	return false
}

type FakeRequester struct {
	Error *FakeNetworkError
}

func (f *FakeRequester) Request(req *http.Request) (*http.Response, error) {
	if f.Error == nil {
		return http.DefaultClient.Do(req)
	}

	return nil, f.Error
}

func TestOnNetworkErrorWithNilBody(t *testing.T) {
	hosts := []*StatefulHost{NewStatefulHost("", call.IsReadWrite)}
	requester := &FakeRequester{Error: &FakeNetworkError{Description: "oh no"}}
	transporter := New(
		hosts,
		requester,
		"appID",
		"apiKey",
		time.Second,
		time.Second,
		nil,
		"",
		compression.None,
	)

	opts := []interface{}{opt.ExposeIntermediateNetworkErrors(true)}
	var res string
	err := transporter.Request(&res, http.MethodGet, "", nil, call.Read, opts...)
	noMoreHostToTryErr, ok := err.(*errs.NoMoreHostToTryErr)
	require.True(t, ok)
	require.Len(t, noMoreHostToTryErr.IntermediateNetworkErrors(), 1)
	require.Equal(t, noMoreHostToTryErr.IntermediateNetworkErrors()[0].Error(), "cannot perform request:\n\terror=oh no\n\tmethod=GET\n\turl=https:")
}

type MockRequester struct {
	validator func(req *http.Request)
}

func (m *MockRequester) Request(req *http.Request) (*http.Response, error) {
	m.validator(req)
	return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader("\"\""))}, nil
}

func TestShouldSupportCaseSensitiveHeader(t *testing.T) {
	hosts := []*StatefulHost{NewStatefulHost("", call.IsReadWrite)}
	requester := &MockRequester{validator: func(req *http.Request) {
		require.Equal(t, "/1/test", req.URL.Path)
		require.Equal(t, http.Header{
			"Connection":               []string{"Keep-Alive"},
			"Content-Type":             []string{"application/json; charset=utf-8"},
			"User-Agent":               []string{"Algolia for Go (3.30.0);Go (go1.20.5)"},
			"X-Algolia-Api-Key":        []string{"newKey"},
			"X-Algolia-Application-Id": []string{"appID"},
		}, req.Header)
	}}

	transporter := New(
		hosts,
		requester,
		"appID",
		"wrong key",
		time.Second,
		time.Second,
		map[string]string{"x-algolia-api-keY": "still wrong key"},
		"",
		compression.None,
	)
	opts := []interface{}{opt.ExtraHeaders(map[string]string{"x-algolia-api-key": "newKey"})}
	var res string
	err := transporter.Request(&res, http.MethodGet, "/1/test", nil, call.Read, opts...)
	require.NoError(t, err)
}
