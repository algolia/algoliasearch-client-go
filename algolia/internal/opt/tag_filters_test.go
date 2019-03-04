package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestTagFilters(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.TagFiltersOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.TagFilterAnd(),
		},
		{
			opts: []interface{}{opt.TagFilter("filter1:value1")},
			expected: opt.TagFilterAnd(
				[]string{"filter1:value1"},
			),
		},
		{
			opts: []interface{}{opt.TagFilterOr(
				"filter1:value1",
				"filter2:value2",
			)},
			expected: opt.TagFilterAnd(
				[]string{"filter1:value1", "filter2:value2"},
			),
		},
		{
			opts: []interface{}{opt.TagFilterOr(
				"filter1:value1",
				opt.TagFilter("filter2:value2"),
			)},
			expected: opt.TagFilterAnd(
				[]string{"filter1:value1", "filter2:value2"},
			),
		},
		{
			opts: []interface{}{opt.TagFilterAnd(
				[]string{"filter1:value1"},
				[]string{"filter2:value2"},
			)},
			expected: opt.TagFilterAnd(
				[]string{"filter1:value1"},
				[]string{"filter2:value2"},
			),
		},
		{
			opts: []interface{}{opt.TagFilterAnd(
				"filter1:value1",
				[]string{"filter2:value2"},
			)},
			expected: opt.TagFilterAnd(
				[]string{"filter1:value1"},
				[]string{"filter2:value2"},
			),
		},
		{
			opts: []interface{}{opt.TagFilterAnd(
				"filter1:value1",
				[]string{"filter2:value2", "filter3:value3"},
			)},
			expected: opt.TagFilterAnd(
				[]string{"filter1:value1"},
				[]string{"filter2:value2", "filter3:value3"},
			),
		},
		{
			opts: []interface{}{opt.TagFilterAnd(
				"filter1:value1",
				opt.TagFilterOr("filter2:value2", "filter3:value3"),
			)},
			expected: opt.TagFilterAnd(
				[]string{"filter1:value1"},
				[]string{"filter2:value2", "filter3:value3"},
			),
		},
	} {
		var (
			in  = ExtractTagFilters(c.opts...)
			out opt.TagFiltersOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, *c.expected, out)
	}
}
