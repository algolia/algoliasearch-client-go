package transport

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/algolia/algoliasearch-client-go/algolia/compression"
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
