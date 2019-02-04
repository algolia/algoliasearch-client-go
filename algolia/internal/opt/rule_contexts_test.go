package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestRuleContexts(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected opt.RuleContextsOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.RuleContextsOption{},
		},
		{
			opts:     []interface{}{opt.RuleContexts()},
			expected: opt.RuleContextsOption{},
		},
		{
			opts:     []interface{}{opt.RuleContexts("mobile")},
			expected: opt.RuleContexts("mobile"),
		},
		{
			opts:     []interface{}{opt.RuleContexts("mobile", "iOS")},
			expected: opt.RuleContexts("mobile", "iOS"),
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
