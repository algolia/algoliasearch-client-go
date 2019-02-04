package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestSortFacetValuesBy(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected opt.SortFacetValuesByOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.SortFacetValuesBy("count"),
		},
		{
			opts:     []interface{}{opt.SortFacetValuesBy("count")},
			expected: opt.SortFacetValuesBy("count"),
		},
		{
			opts:     []interface{}{opt.SortFacetValuesBy("alpha")},
			expected: opt.SortFacetValuesBy("alpha"),
		},
	} {
		var (
			in  = ExtractSortFacetValuesBy(c.opts...)
			out opt.SortFacetValuesByOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, c.expected, out)
	}
}
