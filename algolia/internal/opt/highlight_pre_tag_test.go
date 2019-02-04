package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestHighlightPreTag(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected opt.HighlightPreTagOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.HighlightPreTagOption{},
		},
		{
			opts:     []interface{}{opt.HighlightPreTag("")},
			expected: opt.HighlightPreTag(""),
		},
		{
			opts:     []interface{}{opt.HighlightPreTag("<em>")},
			expected: opt.HighlightPreTag("<em>"),
		},
	} {
		var (
			in  = ExtractHighlightPreTag(c.opts...)
			out opt.HighlightPreTagOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, c.expected, out)
	}
}
