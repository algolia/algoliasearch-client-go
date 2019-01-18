package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestFilters(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected opt.FiltersOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.FiltersOption{},
		},
		{
			opts:     []interface{}{opt.Filters("")},
			expected: opt.Filters(""),
		},
		{
			opts:     []interface{}{opt.Filters("price < 10 AND (category:Book OR NOT category:Ebook)")},
			expected: opt.Filters("price < 10 AND (category:Book OR NOT category:Ebook)"),
		},
	} {
		var (
			in  = ExtractFilters(c.opts...)
			out opt.FiltersOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, c.expected, out)
	}
}
