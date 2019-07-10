package transport

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/algolia/algoliasearch-client-go/algolia/compression"
)

func TestShouldCompress(t *testing.T) {
	for _, c := range []struct {
		compression compression.Compression
		method      string
		expected    bool
	}{
		{0, "", false},
		{compression.None, "", false},
		{compression.None, http.MethodPost, false},
		{compression.None, http.MethodPut, false},
		{compression.None, http.MethodDelete, false},
		{compression.None, http.MethodGet, false},
		{compression.GZIP, http.MethodPost, true},
		{compression.GZIP, http.MethodPut, true},
		{compression.GZIP, http.MethodDelete, false},
		{compression.GZIP, http.MethodGet, false},
	} {
		got := shouldCompress(c.compression, c.method)
		require.Equal(t, c.expected, got, "compression=%d method=%q", c.compression, c.method)
	}
}
