package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestHighlightPostTag(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected opt.HighlightPostTagOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.HighlightPostTagOption{},
		},
		{
			opts:     []interface{}{opt.HighlightPostTag("")},
			expected: opt.HighlightPostTag(""),
		},
		{
			opts:     []interface{}{opt.HighlightPostTag("</em>")},
			expected: opt.HighlightPostTag("</em>"),
		},
	} {
		var (
			in  = ExtractHighlightPostTag(c.opts...)
			out opt.HighlightPostTagOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, c.expected, out)
	}
}
