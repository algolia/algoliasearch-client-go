package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestNumericFilters(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.NumericFiltersOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.NumericFilterAnd(),
		},
		{
			opts: []interface{}{opt.NumericFilter("filter1:value1")},
			expected: opt.NumericFilterAnd(
				[]string{"filter1:value1"},
			),
		},
		{
			opts: []interface{}{opt.NumericFilterOr(
				"filter1:value1",
				"filter2:value2",
			)},
			expected: opt.NumericFilterAnd(
				[]string{"filter1:value1", "filter2:value2"},
			),
		},
		{
			opts: []interface{}{opt.NumericFilterOr(
				"filter1:value1",
				opt.NumericFilter("filter2:value2"),
			)},
			expected: opt.NumericFilterAnd(
				[]string{"filter1:value1", "filter2:value2"},
			),
		},
		{
			opts: []interface{}{opt.NumericFilterAnd(
				[]string{"filter1:value1"},
				[]string{"filter2:value2"},
			)},
			expected: opt.NumericFilterAnd(
				[]string{"filter1:value1"},
				[]string{"filter2:value2"},
			),
		},
		{
			opts: []interface{}{opt.NumericFilterAnd(
				"filter1:value1",
				[]string{"filter2:value2"},
			)},
			expected: opt.NumericFilterAnd(
				[]string{"filter1:value1"},
				[]string{"filter2:value2"},
			),
		},
		{
			opts: []interface{}{opt.NumericFilterAnd(
				"filter1:value1",
				[]string{"filter2:value2", "filter3:value3"},
			)},
			expected: opt.NumericFilterAnd(
				[]string{"filter1:value1"},
				[]string{"filter2:value2", "filter3:value3"},
			),
		},
		{
			opts: []interface{}{opt.NumericFilterAnd(
				"filter1:value1",
				opt.NumericFilterOr("filter2:value2", "filter3:value3"),
			)},
			expected: opt.NumericFilterAnd(
				[]string{"filter1:value1"},
				[]string{"filter2:value2", "filter3:value3"},
			),
		},
	} {
		var (
			in  = ExtractNumericFilters(c.opts...)
			out opt.NumericFiltersOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, *c.expected, out)
	}
}
