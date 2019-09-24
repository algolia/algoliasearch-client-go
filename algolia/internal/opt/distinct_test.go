package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestDistinct(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.DistinctOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.Distinct(false),
		},
		{
			opts:     []interface{}{opt.Distinct(false)},
			expected: opt.Distinct(false),
		},
		{
			opts:     []interface{}{opt.Distinct(true)},
			expected: opt.Distinct(true),
		},
		{
			opts:     []interface{}{opt.DistinctOf(1)},
			expected: opt.DistinctOf(1),
		},
		{
			opts:     []interface{}{opt.DistinctOf(2)},
			expected: opt.DistinctOf(2),
		},
	} {
		var (
			in  = ExtractDistinct(c.opts...)
			out opt.DistinctOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, *c.expected, out)
	}
}
