package search

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_AutomaticFacetFilter_UnMarshalJSON(t *testing.T) {
	for _, c := range []struct {
		input    string
		expected AutomaticFacetFilter
	}{
		{
			`{"facet": "facet", "score": 42}`,
			AutomaticFacetFilter{Facet: "facet", Score: 42},
		},
		{
			`{"facet": "facet", "score": 42, "disjunctive": true}`,
			AutomaticFacetFilter{Facet: "facet", Score: 42, Disjunctive: true},
		},
		{
			`"facet"`,
			AutomaticFacetFilter{Facet: "facet"},
		},
	} {
		var actual AutomaticFacetFilter
		err := json.Unmarshal([]byte(c.input), &actual)
		require.NoError(t, err)
		require.Equal(t, c.expected, actual)
	}
}
