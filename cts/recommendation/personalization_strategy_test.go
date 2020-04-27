package recommendation

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/recommendation"
	"github.com/algolia/algoliasearch-client-go/v3/cts"
)

func TestPersonalizationStrategy(t *testing.T) {
	client := cts.InitRecommendationClient(t)

	strategy := recommendation.Strategy{
		EventsScoring: []recommendation.EventsScoring{
			{EventName: "Add to cart", EventType: "conversion", Score: 50},
			{EventName: "Purchase", EventType: "conversion", Score: 100},
		},
		FacetsScoring: []recommendation.FacetsScoring{
			{FacetName: "brand", Score: 100},
			{FacetName: "categories", Score: 10},
		},
		PersonalizationImpact: opt.PersonalizationImpact(0),
	}

	_, err := client.SetPersonalizationStrategy(strategy)
	require.NoError(t, err)

	got, err := client.GetPersonalizationStrategy()
	require.NoError(t, err)

	require.Equal(t, strategy, got)
}
