package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
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

func TestTagFilters_LegacyDeserialization(t *testing.T) {
	for _, c := range []struct {
		payload  string
		expected *opt.TagFiltersOption
	}{
		{
			`"filter1:value1"`,
			opt.TagFilter("filter1:value1"),
		},
		{
			`" filter1:value1 "`,
			opt.TagFilter("filter1:value1"),
		},
		{
			`["filter1:value1"]`,
			opt.TagFilter("filter1:value1"),
		},
		{
			`[" filter1:value1 "]`,
			opt.TagFilter("filter1:value1"),
		},
		{
			`"filter1:value1,filter2:value2"`,
			opt.TagFilterOr("filter1:value1", "filter2:value2"),
		},
		{
			`" filter1:value1 , filter2:value2 "`,
			opt.TagFilterOr("filter1:value1", "filter2:value2"),
		},
		{
			`["filter1:value1","filter2:value2"]`,
			opt.TagFilterAnd("filter1:value1", "filter2:value2"),
		},
		{
			`[" filter1:value1 "," filter2:value2 "]`,
			opt.TagFilterAnd("filter1:value1", "filter2:value2"),
		},
		{
			`[["filter1:value1","filter2:value2"]]`,
			opt.TagFilterOr("filter1:value1", "filter2:value2"),
		},
		{
			`[["filter1:value1","filter2:value2"], "filter3:value3"]`,
			opt.TagFilterAnd(opt.TagFilterOr("filter1:value1", "filter2:value2"), "filter3:value3"),
		},
	} {
		var got opt.TagFiltersOption
		err := json.Unmarshal([]byte(c.payload), &got)
		require.NoError(t, err, "cannot unmarshal legacy payload %q for opt.TagFiltersOption", c.payload)

		require.True(t, got.Equal(c.expected), "legacy payload %q should deserialize to %#v but got %#v instead", c.payload, c.expected, got)
		require.True(t, c.expected.Equal(&got), "legacy payload %q should deserialize to %#v but got %#v instead", c.payload, c.expected, got)
	}
}
