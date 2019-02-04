package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestDisableExactOnAttributes(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected opt.DisableExactOnAttributesOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.DisableExactOnAttributesOption{},
		},
		{
			opts:     []interface{}{opt.DisableExactOnAttributes()},
			expected: opt.DisableExactOnAttributesOption{},
		},
		{
			opts:     []interface{}{opt.DisableExactOnAttributes("attr1")},
			expected: opt.DisableExactOnAttributes("attr1"),
		},
		{
			opts:     []interface{}{opt.DisableExactOnAttributes("attr1", "attr2")},
			expected: opt.DisableExactOnAttributes("attr1", "attr2"),
		},
	} {
		var (
			in  = ExtractDisableExactOnAttributes(c.opts...)
			out opt.DisableExactOnAttributesOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), out.Get())
	}
}
