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
				},
				"size": {
					"sortRemainingBy": "alpha"
				}
			}
		},
		"redirect": {
			"url": "https://algolia.com/doc"
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

	require.Nil(t, r.FacetOrdering.Values["size"].Order)

	require.Equal(t, *r.Redirect, Redirect{Url: "https://algolia.com/doc"})
}

func TestMarshalRenderingContent(t *testing.T) {
	tests := []struct {
		name     string
		input    RenderingContent
		expected string
	}{
		{
			name: "empty values",
			input: RenderingContent{
				FacetOrdering: &FacetOrdering{
					Facets: &FacetsOrder{
						Order: []string{"brand", "size", "color"},
					},
					Values: nil,
				},
			},
			expected: `{
				"facetOrdering": {
					"facets": {
						"order": [
							"brand",
							"size",
							"color"
						]
					}
				}
			}`,
		},
		{
			name: "values with one empty facet value `order`",
			input: RenderingContent{
				FacetOrdering: &FacetOrdering{
					Facets: &FacetsOrder{
						Order: []string{"brand", "size", "color"},
					},
					Values: map[string]FacetValuesOrder{
						"brand": NewFacetValuesOrder([]string{"Apple", "Sony", "Samsung"}, Alpha),
						"color": NewFacetValuesOrder(nil, Hidden),
					},
				},
			},
			expected: `{
				"facetOrdering": {
					"facets": {
						"order": [
							"brand",
							"size",
							"color"
						]
					},
					"values": {
						"brand": {
							"order": [
								"Apple",
								"Sony",
								"Samsung"
							],
							"sortRemainingBy": "alpha"
						},
						"color": {
							"sortRemainingBy": "hidden"
						}
					}
				}
			}`,
		},
		{
			name: "values with one empty facet value `order`",
			input: RenderingContent{
				FacetOrdering: nil,
			},
			expected: `{}`,
		},
		{
			name: "values with one empty facet value `order`",
			input: RenderingContent{
				Redirect: &Redirect{
					Url: "https://algolia.com/doc",
				},
			},
			expected: `{
				"redirect": {
					"url": "https://algolia.com/doc"
				}
			}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b, err := json.Marshal(test.input)
			require.NoError(t, err)
			require.JSONEq(t, test.expected, string(b))
		})
	}
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
