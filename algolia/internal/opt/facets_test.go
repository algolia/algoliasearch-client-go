// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestFacets(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.FacetsOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.Facets([]string{}...),
		},
		{
			opts:     []interface{}{opt.Facets("value1")},
			expected: opt.Facets("value1"),
		},
		{
			opts:     []interface{}{opt.Facets("value1", "value2", "value3")},
			expected: opt.Facets("value1", "value2", "value3"),
		},
	} {
		var (
			in  = ExtractFacets(c.opts...)
			out opt.FacetsOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), out.Get())
	}
}
