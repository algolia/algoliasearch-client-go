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
		case "aroundLatLng",
			"exactOnSingleWordQuery",
			"filters",
			"highlightPostTag",
			"highlightPreTag",
			"query",
			"queryType",
			"removeWordsIfNoResults",
			"restrictSources",
			"snippetEllipsisText",
			"sortFacetValuesBy":
			if _, ok := v.(string); !ok {
				return invalidType(k, "string")
			}

		case "alternativesAsExact",
			"attributesToHighlight",
			"attributesToRetrieve",
			"attributesToSnippet",
			"disableExactOnAttributes",
			"disableTypoToleranceOnAttributes",
			"explain",
			"queryLanguages",
			"responseFields":
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
			"clickAnalytics",
			"enableRules",
			"synonyms",
			"replaceSynonymsInHighlight",
			"aroundLatLngViaIP",
			"facetingAfterDistinct",
			"restrictHighlightAndSnippetArrays",
			"percentileComputation",
			"sumOrFiltersScores":
			if _, ok := v.(bool); !ok {
				return invalidType(k, "bool")
			}

		case "ignorePlurals",
			"removeStopWords":
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

		case "facetFilters":
			switch v.(type) {
			case string, []string, [][]string, []interface{}:
				//OK
			default:
				return invalidType(k, "string, []string, [][]string or []interface{}")
			}

		case "analyticsTags",
			"restrictSearchableAttributes",
			"facets",
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
