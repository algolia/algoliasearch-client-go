package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestIgnorePlurals(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.IgnorePluralsOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.IgnorePlurals(false),
		},
		{
			opts:     []interface{}{opt.IgnorePlurals(true)},
			expected: opt.IgnorePlurals(true),
		},
		{
			opts:     []interface{}{opt.IgnorePlurals(false)},
			expected: opt.IgnorePlurals(false),
		},
		{
			opts:     []interface{}{opt.IgnorePluralsFor("fr", "en")},
			expected: opt.IgnorePluralsFor("fr", "en"),
		},
	} {
		var (
			in  = ExtractIgnorePlurals(c.opts...)
			out opt.IgnorePluralsOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, *c.expected, out)
	}
}
