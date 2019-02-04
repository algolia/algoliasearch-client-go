package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestDisableTypoToleranceOnAttributes(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected opt.DisableTypoToleranceOnAttributesOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.DisableTypoToleranceOnAttributesOption{},
		},
		{
			opts:     []interface{}{opt.DisableTypoToleranceOnAttributes()},
			expected: opt.DisableTypoToleranceOnAttributesOption{},
		},
		{
			opts:     []interface{}{opt.DisableTypoToleranceOnAttributes("attr1")},
			expected: opt.DisableTypoToleranceOnAttributes("attr1"),
		},
		{
			opts:     []interface{}{opt.DisableTypoToleranceOnAttributes("attr1", "attr2")},
			expected: opt.DisableTypoToleranceOnAttributes("attr1", "attr2"),
		},
	} {
		var (
			in  = ExtractDisableTypoToleranceOnAttributes(c.opts...)
			out opt.DisableTypoToleranceOnAttributesOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), out.Get())
	}
}
