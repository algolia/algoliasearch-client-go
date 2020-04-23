// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestDisableTypoToleranceOnWords(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.DisableTypoToleranceOnWordsOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.DisableTypoToleranceOnWords([]string{}...),
		},
		{
			opts:     []interface{}{opt.DisableTypoToleranceOnWords("value1")},
			expected: opt.DisableTypoToleranceOnWords("value1"),
		},
		{
			opts:     []interface{}{opt.DisableTypoToleranceOnWords("value1", "value2", "value3")},
			expected: opt.DisableTypoToleranceOnWords("value1", "value2", "value3"),
		},
	} {
		var (
			in  = ExtractDisableTypoToleranceOnWords(c.opts...)
			out opt.DisableTypoToleranceOnWordsOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), out.Get())
	}
}

func TestDisableTypoToleranceOnWords_CommaSeparatedString(t *testing.T) {
	for _, c := range []struct {
		payload  string
		expected *opt.DisableTypoToleranceOnWordsOption
	}{
		{
			payload:  `""`,
			expected: opt.DisableTypoToleranceOnWords([]string{}...),
		},
		{
			payload:  `"value1"`,
			expected: opt.DisableTypoToleranceOnWords("value1"),
		},
		{
			payload:  `"value1,value2,value3"`,
			expected: opt.DisableTypoToleranceOnWords("value1", "value2", "value3"),
		},
	} {
		var got opt.DisableTypoToleranceOnWordsOption
		err := json.Unmarshal([]byte(c.payload), &got)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), got.Get())
	}
}
