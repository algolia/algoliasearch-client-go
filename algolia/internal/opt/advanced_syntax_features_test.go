package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestAdvancedSyntaxFeatures(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected opt.AdvancedSyntaxFeaturesOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.AdvancedSyntaxFeaturesOption{},
		},
		{
			opts:     []interface{}{opt.AdvancedSyntaxFeatures()},
			expected: opt.AdvancedSyntaxFeaturesOption{},
		},
		{
			opts:     []interface{}{opt.AdvancedSyntaxFeatures("exactPhrase")},
			expected: opt.AdvancedSyntaxFeatures("exactPhrase"),
		},
		{
			opts:     []interface{}{opt.AdvancedSyntaxFeatures("excludeWords", "exactPhrase")},
			expected: opt.AdvancedSyntaxFeatures("excludeWords", "exactPhrase"),
		},
	} {
		var (
			in  = ExtractAdvancedSyntaxFeatures(c.opts...)
			out opt.AdvancedSyntaxFeaturesOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), out.Get())
	}
}
