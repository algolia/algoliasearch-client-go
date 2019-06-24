package search

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func TestSerializationNestedArraysInMultipleQueries(t *testing.T) {
	queries := []IndexedQuery{
		NewIndexedQuery("test",
			opt.FacetFilterAnd(
				opt.FacetFilterOr(
					opt.FacetFilter("facet1:true"),
					opt.FacetFilter("facet2:true"),
				),
				opt.FacetFilter("facet3:true"),
			),
			opt.InsideBoundingBox([][4]float64{
				{1.0, 2.0, 3.0, 4.0},
				{5.0, 6.0, 7.0, 8.0},
			}),
		),
	}

	reqIn := newMultipleQueriesReq(queries, "none")

	data, err := json.Marshal(reqIn)
	require.NoError(t, err)

	var reqOut multipleQueriesReq
	err = json.Unmarshal(data, &reqOut)
	require.NoError(t, err)

	require.Len(t, reqOut.Requests, 1)
	params := reqOut.Requests[0].Params

	values, err := url.ParseQuery(params)
	require.NoError(t, err)
	require.Len(t, values, 3)

	// Check facetFilters
	{
		var (
			got      [][]string
			expected = [][]string{
				{"facet1:true", "facet2:true"},
				{"facet3:true"},
			}
		)

		err = json.Unmarshal([]byte(values.Get("facetFilters")), &got)
		require.NoError(t, err)
		require.Equal(t, expected, got)
	}

	// Check insideBoundingBox
	{
		var (
			got      [][]float64
			expected = [][]float64{
				{1.0, 2.0, 3.0, 4.0},
				{5.0, 6.0, 7.0, 8.0},
			}
		)

		err = json.Unmarshal([]byte(values.Get("insideBoundingBox")), &got)
		require.NoError(t, err)
		require.Equal(t, expected, got)
	}
}
