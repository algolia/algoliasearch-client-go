package search

import (
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/algolia/algoliasearch-client-go/v3/cts"
	"github.com/stretchr/testify/require"
)

func TestIndexSettingsRenderingContent(t *testing.T) {
	t.Parallel()
	client, _, indexName1 := cts.InitSearchClient1AndIndex(t)

	{
		renderingContent := search.RenderingContent{
			FacetOrdering: &search.FacetOrdering{
				Facets: &search.FacetsOrder{
					Order: []string{"brand", "size", "color"},
				},
				Values: map[string]search.FacetValuesOrder{
					"brand": search.NewFacetValuesOrder([]string{"Apple", "Sony", "Samsung"}, search.Alpha),
					"color": search.NewFacetValuesOrder(nil, search.Hidden),
				},
			}}

		client := client.InitIndex(indexName1)
		client.SetSettings(search.Settings{RenderingContent: &renderingContent})
		settings, _ := client.GetSettings()

		require.True(t, settings.RenderingContent == &renderingContent)
	}
	{
		renderingContent := search.RenderingContent{
			FacetOrdering: &search.FacetOrdering{
				Facets: &search.FacetsOrder{
					Order: []string{"brand", "size", "color"},
				},
				Values: map[string]search.FacetValuesOrder{},
			}}

		client := client.InitIndex(indexName1)
		client.SetSettings(search.Settings{RenderingContent: &renderingContent})
		settings, _ := client.GetSettings()

		require.True(t, settings.RenderingContent.FacetOrdering.Values == nil)
	}
	{
		renderingContent := search.RenderingContent{
			FacetOrdering: &search.FacetOrdering{
				Values: map[string]search.FacetValuesOrder{
					"brand": search.NewFacetValuesOrder([]string{"Apple", "Sony", "Samsung"}, search.Alpha),
					"color": search.NewFacetValuesOrder(nil, search.Hidden),
				},
			}}

		client := client.InitIndex(indexName1)
		client.SetSettings(search.Settings{RenderingContent: &renderingContent})
		settings, _ := client.GetSettings()

		require.True(t, settings.RenderingContent.FacetOrdering == nil)
	}
}
