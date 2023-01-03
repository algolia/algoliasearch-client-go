// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestRuleContexts(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.RuleContextsOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.RuleContexts([]string{}...),
		},
		{
			opts:     []interface{}{opt.RuleContexts("value1")},
			expected: opt.RuleContexts("value1"),
		},
		{
			opts:     []interface{}{opt.RuleContexts("value1", "value2", "value3")},
			expected: opt.RuleContexts("value1", "value2", "value3"),
		},
	} {
		var (
			in  = ExtractRuleContexts(c.opts...)
			out opt.RuleContextsOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), out.Get())
	}
}

func TestRuleContexts_CommaSeparatedString(t *testing.T) {
	for _, c := range []struct {
		payload  string
		expected *opt.RuleContextsOption
	}{
		{
			payload:  `""`,
			expected: opt.RuleContexts([]string{}...),
		},
		{
			payload:  `"value1"`,
			expected: opt.RuleContexts("value1"),
		},
		{
			payload:  `"value1,value2,value3"`,
			expected: opt.RuleContexts("value1", "value2", "value3"),
		},
	} {
		var got opt.RuleContextsOption
		err := json.Unmarshal([]byte(c.payload), &got)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), got.Get())
	}
}