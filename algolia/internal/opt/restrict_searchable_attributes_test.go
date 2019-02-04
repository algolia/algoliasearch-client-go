package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestRestrictSearchableAttributes(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected opt.RestrictSearchableAttributesOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.RestrictSearchableAttributesOption{},
		},
		{
			opts:     []interface{}{opt.RestrictSearchableAttributes()},
			expected: opt.RestrictSearchableAttributesOption{},
		},
		{
			opts:     []interface{}{opt.RestrictSearchableAttributes("attr1")},
			expected: opt.RestrictSearchableAttributes("attr1"),
		},
		{
			opts:     []interface{}{opt.RestrictSearchableAttributes("attr1", "attr2")},
			expected: opt.RestrictSearchableAttributes("attr1", "attr2"),
		},
	} {
		var (
			in  = ExtractRestrictSearchableAttributes(c.opts...)
			out opt.RestrictSearchableAttributesOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), out.Get())
	}
}
