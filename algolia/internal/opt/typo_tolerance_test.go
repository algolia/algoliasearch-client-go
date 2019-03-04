package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestTypoTolerance(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.TypoToleranceOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.TypoToleranceTrue(),
		},
		{
			opts:     []interface{}{opt.TypoToleranceTrue()},
			expected: opt.TypoToleranceTrue(),
		},
		{
			opts:     []interface{}{opt.TypoToleranceFalse()},
			expected: opt.TypoToleranceFalse(),
		},
		{
			opts:     []interface{}{opt.TypoToleranceStrict()},
			expected: opt.TypoToleranceStrict(),
		},
		{
			opts:     []interface{}{opt.TypoToleranceMin()},
			expected: opt.TypoToleranceMin(),
		},
	} {
		var (
			in  = ExtractTypoTolerance(c.opts...)
			out opt.TypoToleranceOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, *c.expected, out)
	}
}
