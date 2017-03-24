package algoliasearch

func checkSettings(settings Map) error {
	for k, v := range settings {
		switch k {
		case "attributesForFaceting",
			"attributesToIndex",
			"searchableAttributes",
			"numericAttributesToIndex",
			"numericAttributesForFiltering",
			"ranking",
			"customRanking",
			"slaves",
			"replicas",
			"unretrievableAttributes",
			"disableTypoToleranceOnAttributes",
			"disableTypoToleranceOnWords",
			"attributesToHighlight",
			"attributesToRetrieve",
			"attributesToSnippet",
			"responseFields",
			"disablePrefixOnAttributes",
			"disableExactOnAttributes",
			"alternativesAsExact":
			if _, ok := v.([]string); !ok {
				return invalidType(k, "[]string")
			}

		case "allowCompressionOfIntegerArray",
			"advancedSyntax",
			"allowTyposOnNumericTokens",
			"replaceSynonymsInHighlight",
			"forwardToSlaves",
			"forwardToReplicas",
			"restrictHighlightAndSnippetArrays":
			if _, ok := v.(bool); !ok {
				return invalidType(k, "bool")
			}

		case "hitsPerPage",
			"maxValuesPerFacet",
			"minProximity",
			"minWordSizefor1Typo",
			"minWordSizefor2Typos",
			"maxFacetHits",
			"paginationLimitedTo":
			if _, ok := v.(int); !ok {
				return invalidType(k, "int")
			}

		case "separatorsToIndex",
			"highlightPostTag",
			"highlightPreTag",
			"queryType",
			"snippetEllipsisText",
			"attributeForDistinct",
			"removeWordsIfNoResults",
			"exactOnSingleWordQuery":
			if _, ok := v.(string); !ok {
				return invalidType(k, "string")
			}

		case "typoTolerance":
			switch v.(type) {
			case string, bool:
				// OK
			default:
				return invalidType(k, "string or bool")
			}

		case "removeStopWords",
			"ignorePlurals":
			switch v.(type) {
			case []string, bool:
				// OK
			default:
				return invalidType(k, "[]string or bool")
			}

		case "distinct":
			switch v.(type) {
			case int, bool:
				// OK
			default:
				return invalidType(k, "int or bool")
			}

		case "optionalWords":
			switch v.(type) {
			case string, []string:
				// OK
			default:
				return invalidType(k, "string or []string")
			}

		default:
		}
	}

	return nil
}
