// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestCamelCaseAttributes(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.CamelCaseAttributesOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.CamelCaseAttributes([]string{}...),
		},
		{
			opts:     []interface{}{opt.CamelCaseAttributes("value1")},
			expected: opt.CamelCaseAttributes("value1"),
		},
		{
			opts:     []interface{}{opt.CamelCaseAttributes("value1", "value2", "value3")},
			expected: opt.CamelCaseAttributes("value1", "value2", "value3"),
		},
	} {
		var (
			in  = ExtractCamelCaseAttributes(c.opts...)
			out opt.CamelCaseAttributesOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), out.Get())
	}
}

func TestCamelCaseAttributes_CommaSeparatedString(t *testing.T) {
	for _, c := range []struct {
		payload  string
		expected *opt.CamelCaseAttributesOption
	}{
		{
			payload:  `""`,
			expected: opt.CamelCaseAttributes([]string{}...),
		},
		{
			payload:  `"value1"`,
			expected: opt.CamelCaseAttributes("value1"),
		},
		{
			payload:  `"value1,value2,value3"`,
			expected: opt.CamelCaseAttributes("value1", "value2", "value3"),
		},
	} {
		var got opt.CamelCaseAttributesOption
		err := json.Unmarshal([]byte(c.payload), &got)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), got.Get())
	}
}