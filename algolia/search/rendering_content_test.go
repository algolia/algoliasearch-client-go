package search

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnmarshalRenderingContent(t *testing.T) {

	payload := `{
		"facetOrdering": {
			"facets": {
				"order": ["brand", "size", "color"]
			},
			"values": {
				"brand": {
					"order": ["Apple", "Sony", "Samsung"],
					"sortRemainingBy": "alpha"
				},
				"color": {
					"order": ["red", "green"]
				}
			}
		}
	}`
	var r RenderingContent
	err := json.Unmarshal([]byte(payload), &r)
	require.NoError(t, err)
	require.Equal(t, r.FacetOrdering.Facets.Order, []string{"brand", "size", "color"})
	require.Equal(t, r.FacetOrdering.Values["brand"].Order, []string{"Apple", "Sony", "Samsung"})
	require.Equal(t, *r.FacetOrdering.Values["brand"].SortRemainingBy, Alpha)
	require.Equal(t, r.FacetOrdering.Values["color"].Order, []string{"red", "green"})
	require.Nil(t, r.FacetOrdering.Values["color"].SortRemainingBy)
}
