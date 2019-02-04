package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestDistinct(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected opt.DistinctOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.DistinctOption{},
		},
		{
			opts:     []interface{}{opt.Distinct(0)},
			expected: opt.Distinct(0),
		},
		{
			opts:     []interface{}{opt.Distinct(1)},
			expected: opt.Distinct(1),
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
		require.Equal(t, c.expected, out)
	}
}
