package algoliasearch

func checkQuery(query Map, ignore ...string) error {
Outer:
	for k, v := range query {
		// Continue if `k` is to be ignored.
		for _, s := range ignore {
			if s == k {
				continue Outer
			}
		}

		switch k {
		case "query",
			"queryType",
			"removeWordsIfNoResults",
			"highlightPreTag",
			"highlightPostTag",
			"snippetEllipsisText",
			"filters",
			"aroundLatLng",
			"exactOnSingleWordQuery":
			if _, ok := v.(string); !ok {
				return invalidType(k, "string")
			}

		case "attributesToRetrieve",
			"disableTypoToleranceOnAttributes",
			"attributesToSnippet",
			"attributesToHighlight",
			"alternativesAsExact",
			"responseFields",
			"disableExactOnAttributes":
			if _, ok := v.([]string); !ok {
				return invalidType(k, "[]string")
			}

		case "minWordSizefor1Typo",
			"minWordSizefor2Typos",
			"minProximity",
			"page",
			"hitsPerPage",
			"distinct",
			"maxValuesPerFacet",
			"aroundPrecision",
			"minimumAroundRadius",
			"maxFacetHits",
			"offset",
			"length":
			if _, ok := v.(int); !ok {
				return invalidType(k, "int")
			}

		case "allowTyposOnNumericTokens",
			"advancedSyntax",
			"analytics",
			"synonyms",
			"replaceSynonymsInHighlight",
			"aroundLatLngViaIP",
			"facetingAfterDistinct",
			"restrictHighlightAndSnippetArrays",
			"percentileComputation":
			if _, ok := v.(bool); !ok {
				return invalidType(k, "bool")
			}

		case "removeStopWords",
			"ignorePlurals":
			switch v.(type) {
			case []string, bool:
				// OK
			default:
				return invalidType(k, "[]string or bool")
			}

		case "aroundRadius":
			switch v.(type) {
			case int, string:
				// OK
			default:
				return invalidType(k, "int or string")
			}

		case "getRankingInfo":
			switch v.(type) {
			case int, bool:
				// OK
			default:
				return invalidType(k, "int or bool")
			}

		case "numericFilters",
			"tagFilters":
			switch v.(type) {
			case string, []interface{}:
				// OK
			default:
				return invalidType(k, "string or []interface{}")
			}

		case "analyticsTags",
			"restrictSearchableAttributes",
			"facets",
			"facetFilters",
			"optionalWords":
			switch v.(type) {
			case string, []string:
				// OK
			default:
				return invalidType(k, "string or []string")
			}

		case "insideBoundingBox",
			"insidePolygon":
			switch v.(type) {
			case string, [][]float64:
				// OK
			default:
				return invalidType(k, "string or [][]float64")
			}

		case "typoTolerance":
			switch v.(type) {
			case string, bool:
				// OK
			default:
				return invalidType(k, "string or bool")
			}

		default:
		}

	}
	return nil
}
