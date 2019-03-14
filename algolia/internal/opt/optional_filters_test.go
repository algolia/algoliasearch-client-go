package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestOptionalFilters(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.OptionalFiltersOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.OptionalFilterAnd(),
		},
		{
			opts: []interface{}{opt.OptionalFilter("filter1:value1")},
			expected: opt.OptionalFilterAnd(
				[]string{"filter1:value1"},
			),
		},
		{
			opts: []interface{}{opt.OptionalFilterOr(
				"filter1:value1",
				"filter2:value2",
			)},
			expected: opt.OptionalFilterAnd(
				[]string{"filter1:value1", "filter2:value2"},
			),
		},
		{
			opts: []interface{}{opt.OptionalFilterOr(
				"filter1:value1",
				opt.OptionalFilter("filter2:value2"),
			)},
			expected: opt.OptionalFilterAnd(
				[]string{"filter1:value1", "filter2:value2"},
			),
		},
		{
			opts: []interface{}{opt.OptionalFilterAnd(
				[]string{"filter1:value1"},
				[]string{"filter2:value2"},
			)},
			expected: opt.OptionalFilterAnd(
				[]string{"filter1:value1"},
				[]string{"filter2:value2"},
			),
		},
		{
			opts: []interface{}{opt.OptionalFilterAnd(
				"filter1:value1",
				[]string{"filter2:value2"},
			)},
			expected: opt.OptionalFilterAnd(
				[]string{"filter1:value1"},
				[]string{"filter2:value2"},
			),
		},
		{
			opts: []interface{}{opt.OptionalFilterAnd(
				"filter1:value1",
				[]string{"filter2:value2", "filter3:value3"},
			)},
			expected: opt.OptionalFilterAnd(
				[]string{"filter1:value1"},
				[]string{"filter2:value2", "filter3:value3"},
			),
		},
		{
			opts: []interface{}{opt.OptionalFilterAnd(
				"filter1:value1",
				opt.OptionalFilterOr("filter2:value2", "filter3:value3"),
			)},
			expected: opt.OptionalFilterAnd(
				[]string{"filter1:value1"},
				[]string{"filter2:value2", "filter3:value3"},
			),
		},
	} {
		var (
			in  = ExtractOptionalFilters(c.opts...)
			out opt.OptionalFiltersOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, *c.expected, out)
	}
}
