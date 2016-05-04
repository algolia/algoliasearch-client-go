package algoliasearch

import "fmt"

func checkSettings(settings map[string]interface{}) error {
	for k, v := range settings {
		switch v.(type) {
		case []Alternative:
			if k != "altCorrections" {
				return fmt.Errorf("SettingsMap field %s is invalid", k)
			}
		case [][]string:
			if k != "synonyms" {
				return fmt.Errorf("SettingsMap field %s is invalid", k)
			}
		case []string:
			if k != "attributesForDistinct" &&
				k != "attributesForFaceting" &&
				k != "attributesToIndex" &&
				k != "numericAttributesToIndex" &&
				k != "ranking" &&
				k != "slaves" &&
				k != "unretrievableAttributes" &&
				k != "disableTypoToleranceOnAttributes" &&
				k != "disableTypoToleranceOnWords" &&
				k != "attributesToHighlight" &&
				k != "attributesToRetrieve" &&
				k != "attributesToSnippet" &&
				k != "OptionalWords" {
				return fmt.Errorf("SettingsMap field %s is invalid", k)
			}
		case bool:
			if k != "allowCompressionOfIntegerArray" &&
				k != "advancedSyntax" &&
				k != "allowTyposOnNumericTokens" &&
				k != "ignorePlurals" &&
				k != "removeStopWords" &&
				k != "replaceSynonymsInHighlight" {
				return fmt.Errorf("SettingsMap field %s is invalid", k)
			}
		case int64:
			if k != "distinct" &&
				k != "hitsPerPage" &&
				k != "maxValuesPerFacet" &&
				k != "minProximity" &&
				k != "minWordSizefor1Typo" &&
				k != "minWordSizefor2Typos" {
				return fmt.Errorf("SettingsMap field %s is invalid", k)
			}
		case map[string][]string:
			if k != "placeholders" {
				return fmt.Errorf("SettingsMap field %s is invalid", k)
			}
		case string:
			if k != "separatorsToIndex" &&
				k != "highlightPostTag" &&
				k != "highlightPreTag" &&
				k != "queryType" &&
				k != "snippetEllipsisText" &&
				k != "typoTolerance" {
				return fmt.Errorf("SettingsMap field %s is invalid", k)
			}
			if k == "queryType" && v != "prefixAll" && v != "prefixLast" && v != "prefixNone" {
				return fmt.Errorf("SettingsMap value %s is invalid", v)
			}
		default:
			return fmt.Errorf("SettingsMap field %s is invalid", k)
		}
	}

	return nil
}
