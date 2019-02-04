package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestAttributesToSnippet(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected opt.AttributesToSnippetOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.AttributesToSnippetOption{},
		},
		{
			opts:     []interface{}{opt.AttributesToSnippet()},
			expected: opt.AttributesToSnippetOption{},
		},
		{
			opts:     []interface{}{opt.AttributesToSnippet("attr1")},
			expected: opt.AttributesToSnippet("attr1"),
		},
		{
			opts:     []interface{}{opt.AttributesToSnippet("attr1", "attr2")},
			expected: opt.AttributesToSnippet("attr1", "attr2"),
		},
	} {
		var (
			in  = ExtractAttributesToSnippet(c.opts...)
			out opt.AttributesToSnippetOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), out.Get())
	}
}
