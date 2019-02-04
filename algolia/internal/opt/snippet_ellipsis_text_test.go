package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestSnippetEllipsisText(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected opt.SnippetEllipsisTextOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.SnippetEllipsisTextOption{},
		},
		{
			opts:     []interface{}{opt.SnippetEllipsisText("")},
			expected: opt.SnippetEllipsisText(""),
		},
		{
			opts:     []interface{}{opt.SnippetEllipsisText("...")},
			expected: opt.SnippetEllipsisText("..."),
		},
	} {
		var (
			in  = ExtractSnippetEllipsisText(c.opts...)
			out opt.SnippetEllipsisTextOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, c.expected, out)
	}
}
