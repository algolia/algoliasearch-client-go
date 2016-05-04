package algoliasearch

import "fmt"

type checkFunc func(interface{}, string) error

type checkMap map[string]checkFunc

var checkQueryMap = checkMap{
	"query":                        isString,
	"queryType":                    checkQueryType,
	"typoTolerance":                checkTypoTolerance,
	"removeWordsIfNoResults":       checkRemoveWordsIfNoResults,
	"restrictSearchableAttributes": isString,
	"highlightPreTag":              isString,
	"highlightPostTag":             isString,
	"snippetEllipsisText":          isString,
	"filters":                      isString,

	"minWordSizefor1Typo":  isInt64,
	"minWordSizefor2Typos": isInt64,
	"minProximity":         isInt64,
	"page":                 isInt64,
	"hitsPerPage":          isInt64,
	"getRankingInfo":       isInt64,
	"distinct":             isInt64,
	"maxValuesPerFacet":    isInt64,
	"aroundRadius":         isInt64,
	"aroundPrecision":      isInt64,
	"minimumAroundRadius":  isInt64,

	"allowTyposOnNumericTokens":  isBool,
	"ignorePlurals":              isBool,
	"advancedSyntax":             isBool,
	"analytics":                  isBool,
	"synonyms":                   isBool,
	"replaceSynonymsInHighlight": isBool,
	"removeStopWords":            isBool,
	"aroundLatLngViaIP":          isBool,

	"analyticsTags":                    isStringSlice,
	"optionalWords":                    isStringSlice,
	"disableTypoToleranceOnAttributes": isStringSlice,
	"attributesToRetrieve":             isStringSlice,
	"attributesToHighlight":            isStringSlice,
	"numericFilters":                   isStringSlice,
	"tagFilters":                       isStringSlice,
	"facets":                           isStringSlice,
	"facetFilters":                     isStringSlice,

	"attributesToSnippet": isStringIntMap,

	"aroundLatLng":      checkAroundLatLng,
	"insideBoundingBox": checkInsideBoundingBox,
	"insidePolygon":     checkInsidePolygon,
}

func checkParams(params map[string]interface{}, m checkMap, ignore ...string) error {
outer:
	for k, v := range params {
		for _, s := range ignore {
			if s == k {
				continue outer
			}
		}

		f, ok := m[k]
		if !ok {
			return invalidField(k)
		}

		if err := f(v, k); err != nil {
			return err
		}
	}

	return nil
}

func invalidType(expected, key string) error {
	return fmt.Errorf("Expected type `%s` for `%s`", expected, key)
}

func invalidField(field string) error {
	return fmt.Errorf("Invalid field `%s`", field)
}

func invalidValue(v string, field string) error {
	return fmt.Errorf("Invalid value `%s` for field `%s`", v, field)
}

func isString(i interface{}, k string) error {
	if _, ok := i.(string); !ok {
		return wrongType("string", k)
	}

	return nil
}

func isInt64(i interface{}, k string) error {
	if _, ok := i.(int64); !ok {
		return wrongType("int64", k)
	}

	return nil
}

func isBool(i interface{}, k string) error {
	if _, ok := i.(bool); !ok {
		return wrongType("bool", k)
	}

	return nil
}

func isStringSlice(i interface{}, k string) error {
	if _, ok := i.([]string); !ok {
		return wrongType("[]string", k)
	}

	return nil
}

func isStringIntMap(i interface{}, k string) error {
	if _, ok := i.(map[string]int); !ok {
		return wrongType("map[string]int", k)
	}

	return nil
}

func isFloat64Slice(i interface{}, k string) error {
	if _, ok := i.([]float64); !ok {
		return wrongType("[]float64", k)
	}

	return nil
}

func isOneOf(s, k string, values []string) error {
	for _, a := range values {
		if a == s {
			return nil
		}
	}

	return invalidValue(s, k)
}

func checkQueryType(i interface{}, k string) error {
	if err := isString(i); err != nil {
		return err
	}

	return isOneOf(i.(string), k, []string{"prefixAll", "prefixLast", "prefixAll"})
}

func checkTypoTolerance(i interface{}, k string) error {
	if err := isString(i); err != nil {
		return err
	}

	return isOneOf(i.(string), k, []string{"false", "true", "min", "strict"})
}

func checkRemoveWordsIfNoResults(i interface{}, k string) error {
	if err := isString(i); err != nil {
		return err
	}

	return isOneOf(i.(string), k, []string{"lastWords", "firstWords", "allOptional", "none"})
}

func checkAroundLatLng(i interface{}, k string) error {
	if err := isFloat64Slice(i, k); !ok {
		return err
	}

	l := len(i.([]float64))
	if l != 2 {
		return fmt.Errorf("Invalid number of dimensions for `%s`. Expected `2` got `%d`", k, l)
	}

	return nil
}

func checkInsideBoundingBox(i interface{}, k string) error {
	if err := isFloat64Slice(i, k); !ok {
		return err
	}

	l := len(i.([]float64))
	if l == 0 || l%4 != 0 {
		return fmt.Errorf("Invalid number of dimensions for `%s`. Expected a multiple of `4` got `%d`", k, l)
	}

	return nil
}

func checkInsidePolygon(i interface{}, k string) error {
	if err := isFloat64Slice(i, k); !ok {
		return err
	}

	l := len(i.([]float64))
	if l < 6 {
		return fmt.Errorf("Invalid number of dimensions for `%s`. Expected at least `6` got `%d`", k, l)
	}

	return nil
}

func checkGetLogs(params map[string]interface{}) error {
	for k, v := range params {
		switch k {
		case "length", "offset":
			if _, ok := v.(int64); !ok {
				return wrongType("int64", k)
			}

		case "type":
			if _, ok := v.(string); !ok {
				return wrongType("string", k)
			}

		default:
			return unknownField(k)
		}
	}

	return nil
}

func checkGenerateSecuredAPIKey(params map[string]interface{}) error {
	for k, v := range params {
		switch k {
		case "userToken":
			if _, ok := v.(string); !ok {
				return wrongType("string", k)
			}

		case "tagFilters":
			if _, ok := v.([]string); !ok {
				return wrongType("[]string", k)
			}
		}
	}

	// TODO: Plug query check function here.

	return nil
}

func checkKeyReq(params map[string]interface{}) error {
	for k, v := range params {
		switch k {
		case "acl", "indexes", "referers":
			if _, ok := v.([]string); !ok {
				return wrongType("[]string", k)
			}

		case "description", "queryParameters":
			if _, ok := v.(string); !ok {
				return wrongType("string", k)
			}

		case "maxHitsPerQuery", "maxQueriesPerIPPerHour", "validity":
			if _, ok := v.(int64); !ok {
				return wrongType("int64", k)
			}

		default:
			return unknownField(k)
		}
	}

	return nil
}

func checkBatchAction(action string) error {
	actions := map[string]struct{}{
		"addObject": struct{}{}, "updateObject": struct{}{},
		"partialUpdateObject": struct{}{}, "deleteObject": struct{}{},
		"partialUpdateObjectNoCreate": struct{}{}, "clear": struct{}{},
	}

	_, ok := actions[action]
	if !ok {
		return fmt.Errorf("Unknown action `%s`", action)
	}
	return nil
}
