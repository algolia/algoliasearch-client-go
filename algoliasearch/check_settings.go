package algoliasearch

func checkSettings(settings Map) error {
	for k, v := range settings {
		switch k {
		case "alternativesAsExact",
			"attributesForFaceting",
			"attributesToHighlight",
			"attributesToIndex",
			"attributesToRetrieve",
			"attributesToSnippet",
			"camelCaseAttributes",
			"customRanking",
			"disableExactOnAttributes",
			"disablePrefixOnAttributes",
			"disableTypoToleranceOnAttributes",
			"disableTypoToleranceOnWords",
			"numericAttributesForFiltering",
			"numericAttributesToIndex",
			"ranking",
			"replicas",
			"responseFields",
			"searchableAttributes",
			"slaves",
			"unretrievableAttributes":
			if _, ok := v.([]string); !ok {
				return invalidType(k, "[]string")
			}

		case "allowCompressionOfIntegerArray",
			"advancedSyntax",
			"allowTyposOnNumericTokens",
			"enableRules",
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

		case "attributeForDistinct",
			"exactOnSingleWordQuery",
			"highlightPostTag",
			"highlightPreTag",
			"keepDiacriticsOnCharacters",
			"queryType",
			"removeWordsIfNoResults",
			"separatorsToIndex",
			"snippetEllipsisText",
			"sortFacetValuesBy":
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

		case "ignorePlurals",
			"queryLanguages",
			"removeStopWords":
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

		case "decompoundedAttributes":
			if _, ok := v.(map[string][]string); !ok {
				return invalidType(k, "map[string][]string")
			}

		default:
		}
	}

	return nil
}
