package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestFacetFilters(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.FacetFiltersOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.FacetFilterAnd(),
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
		require.Equal(t, *c.expected, out)
	}
}

func TestFacetFilters_LegacyDeserialization(t *testing.T) {
	for _, c := range []struct {
		payload  string
		expected *opt.FacetFiltersOption
	}{
		{
			`"filter1:value1"`,
			opt.FacetFilter("filter1:value1"),
		},
		{
			`" filter1:value1 "`,
			opt.FacetFilter("filter1:value1"),
		},
		{
			`["filter1:value1"]`,
			opt.FacetFilter("filter1:value1"),
		},
		{
			`[" filter1:value1 "]`,
			opt.FacetFilter("filter1:value1"),
		},
		{
			`"filter1:value1,filter2:value2"`,
			opt.FacetFilterAnd("filter1:value1", "filter2:value2"),
		},
		{
			`" filter1:value1 , filter2:value2 "`,
			opt.FacetFilterAnd("filter1:value1", "filter2:value2"),
		},
		{
			`"(filter1:value1,filter2:value2)"`,
			opt.FacetFilterOr("filter1:value1", "filter2:value2"),
		},
		{
			`["filter1:value1","filter2:value2"]`,
			opt.FacetFilterAnd("filter1:value1", "filter2:value2"),
		},
		{
			`[" filter1:value1 "," filter2:value2 "]`,
			opt.FacetFilterAnd("filter1:value1", "filter2:value2"),
		},
		{
			`[["filter1:value1","filter2:value2"]]`,
			opt.FacetFilterOr("filter1:value1", "filter2:value2"),
		},
		{
			`[["filter1:value1","filter2:value2"], "filter3:value3"]`,
			opt.FacetFilterAnd(opt.FacetFilterOr("filter1:value1", "filter2:value2"), "filter3:value3"),
		},
		{
			`["filter1:value1,filter2:value2","filter3:value3"]`,
			opt.FacetFilterAnd(opt.FacetFilterOr("filter1:value1", "filter2:value2"), "filter3:value3"),
		},
		{
			`"(filter1:value1,filter2:value2),filter3:value3"`,
			opt.FacetFilterAnd(opt.FacetFilterOr("filter1:value1", "filter2:value2"), "filter3:value3"),
		},
	} {
		var got opt.FacetFiltersOption
		err := json.Unmarshal([]byte(c.payload), &got)
		require.NoError(t, err, "cannot unmarshal legacy payload %q for opt.FacetFiltersOption", c.payload)

		require.True(t, got.Equal(c.expected), "legacy payload %q should deserialize to %#v but got %#v instead", c.payload, c.expected, got)
		require.True(t, c.expected.Equal(&got), "legacy payload %q should deserialize to %#v but got %#v instead", c.payload, c.expected, got)
	}
}
