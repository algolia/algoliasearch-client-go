package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestExactOnSingleWordQuery(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected opt.ExactOnSingleWordQueryOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.ExactOnSingleWordQueryOption{},
		},
		{
			opts:     []interface{}{opt.ExactOnSingleWordQuery("")},
			expected: opt.ExactOnSingleWordQuery(""),
		},
		{
			opts:     []interface{}{opt.ExactOnSingleWordQuery("attribute")},
			expected: opt.ExactOnSingleWordQuery("attribute"),
		},
	} {
		var (
			in  = ExtractExactOnSingleWordQuery(c.opts...)
			out opt.ExactOnSingleWordQueryOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, c.expected, out)
	}
}
