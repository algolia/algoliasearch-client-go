package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
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

func TestOptionalFilters_LegacyDeserialization(t *testing.T) {
	for _, c := range []struct {
		payload  string
		expected *opt.OptionalFiltersOption
	}{
		{
			`"filter1:value1"`,
			opt.OptionalFilter("filter1:value1"),
		},
		{
			`" filter1:value1 "`,
			opt.OptionalFilter("filter1:value1"),
		},
		{
			`["filter1:value1"]`,
			opt.OptionalFilter("filter1:value1"),
		},
		{
			`[" filter1:value1 "]`,
			opt.OptionalFilter("filter1:value1"),
		},
		{
			`"filter1:value1,filter2:value2"`,
			opt.OptionalFilterAnd("filter1:value1", "filter2:value2"),
		},
		{
			`" filter1:value1 , filter2:value2 "`,
			opt.OptionalFilterAnd("filter1:value1", "filter2:value2"),
		},
		{
			`"(filter1:value1,filter2:value2)"`,
			opt.OptionalFilterOr("filter1:value1", "filter2:value2"),
		},
		{
			`["filter1:value1","filter2:value2"]`,
			opt.OptionalFilterAnd("filter1:value1", "filter2:value2"),
		},
		{
			`[" filter1:value1 "," filter2:value2 "]`,
			opt.OptionalFilterAnd("filter1:value1", "filter2:value2"),
		},
		{
			`[["filter1:value1","filter2:value2"]]`,
			opt.OptionalFilterOr("filter1:value1", "filter2:value2"),
		},
		{
			`[["filter1:value1","filter2:value2"], "filter3:value3"]`,
			opt.OptionalFilterAnd(opt.OptionalFilterOr("filter1:value1", "filter2:value2"), "filter3:value3"),
		},
		{
			`["filter1:value1,filter2:value2","filter3:value3"]`,
			opt.OptionalFilterAnd(opt.OptionalFilterOr("filter1:value1", "filter2:value2"), "filter3:value3"),
		},
		{
			`"(filter1:value1,filter2:value2),filter3:value3"`,
			opt.OptionalFilterAnd(opt.OptionalFilterOr("filter1:value1", "filter2:value2"), "filter3:value3"),
		},
	} {
		var got opt.OptionalFiltersOption
		err := json.Unmarshal([]byte(c.payload), &got)
		require.NoError(t, err, "cannot unmarshal legacy payload %q for opt.OptionalFiltersOption", c.payload)

		require.True(t, got.Equal(c.expected), "legacy payload %q should deserialize to %#v but got %#v instead", c.payload, c.expected, got)
		require.True(t, c.expected.Equal(&got), "legacy payload %q should deserialize to %#v but got %#v instead", c.payload, c.expected, got)
	}
}
