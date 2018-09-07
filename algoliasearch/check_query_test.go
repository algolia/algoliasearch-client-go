package algoliasearch

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckQuery(t *testing.T) {
	for _, m := range []Map{
		Map{"facetFilters": "filter"},
		Map{"facetFilters": []string{"f1", "f2"}},
		Map{"facetFilters": [][]string{[]string{"f1", "f2"}, []string{"f3"}}},
		Map{"facetFilters": []interface{}{[]string{"f1", "f2"}, "f3"}},
	} {
		require.NoError(t, checkQuery(m), "should accept the following query parameter: %#v", m)
	}
}
