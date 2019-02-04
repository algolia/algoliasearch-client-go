package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestAnalyticsTags(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected opt.AnalyticsTagsOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.AnalyticsTagsOption{},
		},
		{
			opts:     []interface{}{opt.AnalyticsTags()},
			expected: opt.AnalyticsTagsOption{},
		},
		{
			opts:     []interface{}{opt.AnalyticsTags("tag1")},
			expected: opt.AnalyticsTags("tag1"),
		},
		{
			opts:     []interface{}{opt.AnalyticsTags("tag1", "tag2")},
			expected: opt.AnalyticsTags("tag1", "tag2"),
		},
	} {
		var (
			in  = ExtractAnalyticsTags(c.opts...)
			out opt.AnalyticsTagsOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), out.Get())
	}
}
