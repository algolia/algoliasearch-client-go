package opt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFilters(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected string
	}{
		{
			opts:     []interface{}{nil},
			expected: "",
		},
		{
			opts:     []interface{}{Filters("")},
			expected: "",
		},
		{
			opts:     []interface{}{Filters("price < 10 AND (category:Book OR NOT category:Ebook)")},
			expected: "price < 10 AND (category:Book OR NOT category:Ebook)",
		},
	} {
		res := ExtractFilters(c.opts...)
		require.Equal(t, c.expected, res)
	}
}
