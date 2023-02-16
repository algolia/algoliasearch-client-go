package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
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

func TestNumericFilters_LegacyDeserialization(t *testing.T) {
	for _, c := range []struct {
		payload  string
		expected *opt.NumericFiltersOption
	}{
		{
			`"filter1:value1"`,
			opt.NumericFilter("filter1:value1"),
		},
		{
			`" filter1:value1 "`,
			opt.NumericFilter("filter1:value1"),
		},
		{
			`["filter1:value1"]`,
			opt.NumericFilter("filter1:value1"),
		},
		{
			`[" filter1:value1 "]`,
			opt.NumericFilter("filter1:value1"),
		},
		{
			`"filter1:value1,filter2:value2"`,
			opt.NumericFilterOr("filter1:value1", "filter2:value2"),
		},
		{
			`" filter1:value1 , filter2:value2 "`,
			opt.NumericFilterOr("filter1:value1", "filter2:value2"),
		},
		{
			`["filter1:value1","filter2:value2"]`,
			opt.NumericFilterAnd("filter1:value1", "filter2:value2"),
		},
		{
			`[" filter1:value1 "," filter2:value2 "]`,
			opt.NumericFilterAnd("filter1:value1", "filter2:value2"),
		},
		{
			`[["filter1:value1","filter2:value2"]]`,
			opt.NumericFilterOr("filter1:value1", "filter2:value2"),
		},
		{
			`[["filter1:value1","filter2:value2"], "filter3:value3"]`,
			opt.NumericFilterAnd(opt.NumericFilterOr("filter1:value1", "filter2:value2"), "filter3:value3"),
		},
	} {
		var got opt.NumericFiltersOption
		err := json.Unmarshal([]byte(c.payload), &got)
		require.NoError(t, err, "cannot unmarshal legacy payload %q for opt.NumericFiltersOption", c.payload)

		require.True(t, got.Equal(c.expected), "legacy payload %q should deserialize to %#v but got %#v instead", c.payload, c.expected, got)
		require.True(t, c.expected.Equal(&got), "legacy payload %q should deserialize to %#v but got %#v instead", c.payload, c.expected, got)
	}
}
