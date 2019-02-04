package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestEnablePersonalization(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected opt.EnablePersonalizationOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.EnablePersonalization(false),
		},
		{
			opts:     []interface{}{opt.EnablePersonalization(true)},
			expected: opt.EnablePersonalization(true),
		},
		{
			opts:     []interface{}{opt.EnablePersonalization(false)},
			expected: opt.EnablePersonalization(false),
		},
	} {
		var (
			in  = ExtractEnablePersonalization(c.opts...)
			out opt.EnablePersonalizationOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, c.expected, out)
	}
}
