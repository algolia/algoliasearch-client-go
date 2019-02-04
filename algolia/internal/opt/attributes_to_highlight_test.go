package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestAttributesToHighlight(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected opt.AttributesToHighlightOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.AttributesToHighlightOption{},
		},
		{
			opts:     []interface{}{opt.AttributesToHighlight()},
			expected: opt.AttributesToHighlightOption{},
		},
		{
			opts:     []interface{}{opt.AttributesToHighlight("attr1")},
			expected: opt.AttributesToHighlight("attr1"),
		},
		{
			opts:     []interface{}{opt.AttributesToHighlight("attr1", "attr2")},
			expected: opt.AttributesToHighlight("attr1", "attr2"),
		},
	} {
		var (
			in  = ExtractAttributesToHighlight(c.opts...)
			out opt.AttributesToHighlightOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), out.Get())
	}
}
