package personalization

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/errs"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/personalization"
	"github.com/algolia/algoliasearch-client-go/v3/cts"
)

func TestPersonalizationStrategy(t *testing.T) {
	client := cts.InitPersonalizationClient(t)

	strategy := personalization.Strategy{
		EventsScoring: []personalization.EventsScoring{
			{EventName: "Add to cart", EventType: "conversion", Score: 50},
			{EventName: "Purchase", EventType: "conversion", Score: 100},
		},
		FacetsScoring: []personalization.FacetsScoring{
			{FacetName: "brand", Score: 100},
			{FacetName: "categories", Score: 10},
		},
		PersonalizationImpact: opt.PersonalizationImpact(0),
	}

	_, err := client.SetPersonalizationStrategy(strategy)
	if algoliaErr, ok := errs.IsAlgoliaErr(err); ok {
		require.Equal(t, 429, algoliaErr.Status)
		require.Equal(t, "Number of strategy saves exceeded for the day", algoliaErr.Message)
	} else {
		require.NoError(t, err)
	}

	got, err := client.GetPersonalizationStrategy()
	require.NoError(t, err)

	require.Equal(t, strategy, got)
}
