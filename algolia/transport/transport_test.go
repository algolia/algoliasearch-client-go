package transport

import (
	"net/http"
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
