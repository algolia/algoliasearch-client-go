package search

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_AutomaticFacetFilter_UnmarshalJSON(t *testing.T) {
	for _, c := range []struct {
		input    string
		expected AutomaticFacetFilter
		err      string
	}{
		{
			input:    `{"facet": "facet", "score": 42}`,
			expected: AutomaticFacetFilter{Facet: "facet", Score: 42},
		},
		{
			input:    `{"facet": "facet", "score": 42, "disjunctive": true}`,
			expected: AutomaticFacetFilter{Facet: "facet", Score: 42, Disjunctive: true},
		},
		{
			input:    `"facet"`,
			expected: AutomaticFacetFilter{Facet: "facet"},
		},
		{
			input: `[]`,
			err:   `cannot unmarshal automatic facet filter`,
		},
	} {
		var actual AutomaticFacetFilter
		err := json.Unmarshal([]byte(c.input), &actual)
		if c.err != "" {
			require.EqualError(t, err, c.err)
			continue
		}
		require.NoError(t, err)
		require.Equal(t, c.expected, actual)
	}
}
