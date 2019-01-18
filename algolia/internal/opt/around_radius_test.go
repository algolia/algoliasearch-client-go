package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestAroundRadius(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected opt.AroundRadiusOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.AroundRadiusOption{},
		},
		{
			opts:     []interface{}{opt.AroundRadius(0)},
			expected: opt.AroundRadius(0),
		},
		{
			opts:     []interface{}{opt.AroundRadius(42)},
			expected: opt.AroundRadius(42),
		},
		{
			opts:     []interface{}{opt.AroundRadius(42), opt.AroundRadius(43)},
			expected: opt.AroundRadius(42),
		},
		{
			opts:     []interface{}{opt.AroundRadiusAll()},
			expected: opt.AroundRadiusAll()},
		{
			opts:     []interface{}{opt.AroundRadius(42), opt.AroundRadiusAll()},
			expected: opt.AroundRadius(42),
		},
		{
			opts:     []interface{}{opt.AroundRadiusAll(), opt.AroundRadius(42)},
			expected: opt.AroundRadiusAll(),
		},
	} {
		var (
			in  = ExtractAroundRadius(c.opts...)
			out opt.AroundRadiusOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, c.expected, out)
	}
}
