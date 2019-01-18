package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestCreateIfNotExists(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected opt.CreateIfNotExistsOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.CreateIfNotExists(true),
		},
		{
			opts:     []interface{}{opt.CreateIfNotExists(true)},
			expected: opt.CreateIfNotExists(true),
		},
		{
			opts:     []interface{}{opt.CreateIfNotExists(false)},
			expected: opt.CreateIfNotExists(false),
		},
		{
			opts: []interface{}{
				opt.CreateIfNotExists(false),
				opt.CreateIfNotExists(true),
			},
			expected: opt.CreateIfNotExists(false),
		},
		{
			opts: []interface{}{
				opt.CreateIfNotExists(true),
				opt.CreateIfNotExists(false),
			},
			expected: opt.CreateIfNotExists(true),
		},
	} {
		var (
			in  = ExtractCreateIfNotExists(c.opts...)
			out opt.CreateIfNotExistsOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, c.expected, out)
	}
}
