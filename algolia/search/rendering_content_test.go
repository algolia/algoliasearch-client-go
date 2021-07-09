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

func TestBuildRenderingContent(t *testing.T) {

	var r = RenderingContent{FacetOrdering: &FacetOrdering{
		Facets: &FacetsOrder{Order: []string{"size", "brand"}},
		Values: map[string]FacetValuesOrder{
			"brand": NewFacetValuesOrder([]string{"Uniqlo"}, Count),
			"size":  NewFacetValuesOrder([]string{"S", "M", "L"}, Hidden),
		},
	}}

	require.Equal(t, r.FacetOrdering.Facets.Order, []string{"size", "brand"})
	require.Equal(t, r.FacetOrdering.Values["brand"].Order, []string{"Uniqlo"})
	require.Equal(t, *r.FacetOrdering.Values["brand"].SortRemainingBy, Count)
	require.Equal(t, r.FacetOrdering.Values["size"].Order, []string{"S", "M", "L"})
	require.Equal(t, *r.FacetOrdering.Values["size"].SortRemainingBy, Hidden)

}
