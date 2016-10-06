package algoliasearch

func buildKeysMap(keys []string) map[string]struct{} {
	m := make(map[string]struct{})
	for _, key := range keys {
		m[key] = struct{}{}
	}

	return m
}

func checkQuery(query Map, ignore ...string) error {
	ignoreKeysMap := buildKeysMap(ignore)

	for k, v := range query {
		if _, shouldIgnore := ignoreKeysMap[k]; shouldIgnore {
			continue
		}

		switch k {
		case "query",
			"queryType",
			"typoTolerance",
			"removeWordsIfNoResults",
			"restrictSearchableAttributes",
			"highlightPreTag",
			"highlightPostTag",
			"snippetEllipsisText",
			"filters",
			"analyticsTags",
			"optionalWords",
			"numericFilters",
			"tagFilters",
			"facets",
			"facetFilters",
			"aroundLatLng",
			"insideBoundingBox",
			"insidePolygon",
			"exactOnSingleWordQuery":
			if _, ok := v.(string); !ok {
				return invalidType(k, "string")
			}

		case "attributesToRetrieve",
			"disableTypoToleranceOnAttributes",
			"attributesToSnippet",
			"attributesToHighlight",
			"alternativesAsExact":
			if _, ok := v.([]string); !ok {
				return invalidType(k, "[]string")
			}

		case "minWordSizefor1Typo",
			"minWordSizefor2Typos",
			"minProximity",
			"page",
			"hitsPerPage",
			"getRankingInfo",
			"distinct",
			"maxValuesPerFacet",
			"aroundPrecision",
			"minimumAroundRadius":
			if _, ok := v.(int); !ok {
				return invalidType(k, "int")
			}

		case "allowTyposOnNumericTokens",
			"ignorePlurals",
			"advancedSyntax",
			"analytics",
			"synonyms",
			"replaceSynonymsInHighlight",
			"aroundLatLngViaIP",
			"facetingAfterDistinct":
			if _, ok := v.(bool); !ok {
				return invalidType(k, "bool")
			}

		case "removeStopWords":
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

		default:
		}

	}
	return nil
}
