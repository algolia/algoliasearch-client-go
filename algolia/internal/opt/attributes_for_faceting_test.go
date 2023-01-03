// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestAttributesForFaceting(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.AttributesForFacetingOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.AttributesForFaceting([]string{}...),
		},
		{
			opts:     []interface{}{opt.AttributesForFaceting("value1")},
			expected: opt.AttributesForFaceting("value1"),
		},
		{
			opts:     []interface{}{opt.AttributesForFaceting("value1", "value2", "value3")},
			expected: opt.AttributesForFaceting("value1", "value2", "value3"),
		},
	} {
		var (
			in  = ExtractAttributesForFaceting(c.opts...)
			out opt.AttributesForFacetingOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), out.Get())
	}
}

func TestAttributesForFaceting_CommaSeparatedString(t *testing.T) {
	for _, c := range []struct {
		payload  string
		expected *opt.AttributesForFacetingOption
	}{
		{
			payload:  `""`,
			expected: opt.AttributesForFaceting([]string{}...),
		},
		{
			payload:  `"value1"`,
			expected: opt.AttributesForFaceting("value1"),
		},
		{
			payload:  `"value1,value2,value3"`,
			expected: opt.AttributesForFaceting("value1", "value2", "value3"),
		},
	} {
		var got opt.AttributesForFacetingOption
		err := json.Unmarshal([]byte(c.payload), &got)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), got.Get())
	}
}