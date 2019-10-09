package analytics

import (
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/analytics"
	"github.com/stretchr/testify/require"
)

func checkABTestsAreEqual(t *testing.T, a analytics.ABTest, b analytics.ABTestResponse) {
	require.Equal(t, a.Name, b.Name)
	require.Equal(t, a.EndAt.Unix(), b.EndAt.Unix())
	require.Equal(t, len(a.Variants), len(b.Variants))

	var responseVariants []analytics.Variant
	for _, v := range b.Variants {
		responseVariants = append(responseVariants, analytics.Variant{
			Index:                  v.Index,
			TrafficPercentage:      v.TrafficPercentage,
			Description:            v.Description,
			CustomSearchParameters: v.CustomSearchParameters,
		})
	}

	require.Equal(t, len(a.Variants), len(responseVariants))

	found := 0
	for _, v1 := range a.Variants {
		for _, v2 := range responseVariants {
			if v1.Index == v2.Index &&
				v1.TrafficPercentage == v2.TrafficPercentage &&
				v1.Description == v2.Description &&
				v1.CustomSearchParameters.Equal(v2.CustomSearchParameters) {
				found++
				break
			}
		}
	}
	require.Equal(t, len(a.Variants), found)
}
