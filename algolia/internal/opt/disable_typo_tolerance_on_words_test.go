package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestDisableTypoToleranceOnWords(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected opt.DisableTypoToleranceOnWordsOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.DisableTypoToleranceOnWordsOption{},
		},
		{
			opts:     []interface{}{opt.DisableTypoToleranceOnWords()},
			expected: opt.DisableTypoToleranceOnWordsOption{},
		},
		{
			opts:     []interface{}{opt.DisableTypoToleranceOnWords("word1")},
			expected: opt.DisableTypoToleranceOnWords("word1"),
		},
		{
			opts:     []interface{}{opt.DisableTypoToleranceOnWords("word1", "word2")},
			expected: opt.DisableTypoToleranceOnWords("word1", "word2"),
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
