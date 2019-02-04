package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestAlternativesAsExact(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected opt.AlternativesAsExactOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.AlternativesAsExactOption{},
		},
		{
			opts:     []interface{}{opt.AlternativesAsExact()},
			expected: opt.AlternativesAsExactOption{},
		},
		{
			opts:     []interface{}{opt.AlternativesAsExact("alt1")},
			expected: opt.AlternativesAsExact("alt1"),
		},
		{
			opts:     []interface{}{opt.AlternativesAsExact("alt1", "alt2")},
			expected: opt.AlternativesAsExact("alt1", "alt2"),
		},
	} {
		var (
			in  = ExtractAlternativesAsExact(c.opts...)
			out opt.AlternativesAsExactOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.ElementsMatch(t, c.expected.Get(), out.Get())
	}
}
