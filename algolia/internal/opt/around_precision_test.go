package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestAroundPrecision(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.AroundPrecisionOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.AroundPrecision(opt.AroundPrecisionRange{From: 0, Value: 1}),
		},
		{
			opts:     []interface{}{opt.AroundPrecision(opt.AroundPrecisionRange{From: 0, Value: 0})},
			expected: opt.AroundPrecision(opt.AroundPrecisionRange{From: 0, Value: 0}),
		},
		{
			opts:     []interface{}{opt.AroundPrecision(opt.AroundPrecisionRange{From: 0, Value: 1})},
			expected: opt.AroundPrecision(opt.AroundPrecisionRange{From: 0, Value: 1}),
		},
		{
			opts:     []interface{}{opt.AroundPrecision(opt.AroundPrecisionRange{From: 0, Value: -42})},
			expected: opt.AroundPrecision(opt.AroundPrecisionRange{From: 0, Value: -42}),
		},
		{
			opts:     []interface{}{opt.AroundPrecision(opt.AroundPrecisionRange{From: 10, Value: 20})},
			expected: opt.AroundPrecision(opt.AroundPrecisionRange{From: 10, Value: 20}),
		},
		{
			opts: []interface{}{opt.AroundPrecision(
				opt.AroundPrecisionRange{From: 10, Value: 20},
				opt.AroundPrecisionRange{From: 30, Value: 40},
			)},
			expected: opt.AroundPrecision(
				opt.AroundPrecisionRange{From: 10, Value: 20},
				opt.AroundPrecisionRange{From: 30, Value: 40},
			),
		},
	} {
		var (
			in  = ExtractAroundPrecision(c.opts...)
			out opt.AroundPrecisionOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, *c.expected, out)
		require.True(t, c.expected.Equal(&out))
		require.True(t, out.Equal(c.expected))
	}
}
