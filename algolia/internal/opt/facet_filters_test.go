package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestFacetFilters(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected opt.FacetFiltersOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.FacetFiltersOption{},
		},
		{
			opts: []interface{}{opt.FacetFilter("filter1:value1")},
			expected: opt.FacetFilterAnd(
				[]string{"filter1:value1"},
			),
		},
		{
			opts: []interface{}{opt.FacetFilterOr(
				"filter1:value1",
				"filter2:value2",
			)},
			expected: opt.FacetFilterAnd(
				[]string{"filter1:value1", "filter2:value2"},
			),
		},
		{
			opts: []interface{}{opt.FacetFilterOr(
				"filter1:value1",
				opt.FacetFilter("filter2:value2"),
			)},
			expected: opt.FacetFilterAnd(
				[]string{"filter1:value1", "filter2:value2"},
			),
		},
		{
			opts: []interface{}{opt.FacetFilterAnd(
				[]string{"filter1:value1"},
				[]string{"filter2:value2"},
			)},
			expected: opt.FacetFilterAnd(
				[]string{"filter1:value1"},
				[]string{"filter2:value2"},
			),
		},
		{
			opts: []interface{}{opt.FacetFilterAnd(
				"filter1:value1",
				[]string{"filter2:value2"},
			)},
			expected: opt.FacetFilterAnd(
				[]string{"filter1:value1"},
				[]string{"filter2:value2"},
			),
		},
		{
			opts: []interface{}{opt.FacetFilterAnd(
				"filter1:value1",
				[]string{"filter2:value2", "filter3:value3"},
			)},
			expected: opt.FacetFilterAnd(
				[]string{"filter1:value1"},
				[]string{"filter2:value2", "filter3:value3"},
			),
		},
		{
			opts: []interface{}{opt.FacetFilterAnd(
				"filter1:value1",
				opt.FacetFilterOr("filter2:value2", "filter3:value3"),
			)},
			expected: opt.FacetFilterAnd(
				[]string{"filter1:value1"},
				[]string{"filter2:value2", "filter3:value3"},
			),
		},
	} {
		var (
			in  = ExtractFacetFilters(c.opts...)
			out opt.FacetFiltersOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, c.expected, out)
	}
}
