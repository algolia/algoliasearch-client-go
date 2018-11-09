package algoliasearch

func checkRule(rule Rule) error {
	return checkRules([]Rule{rule})
}

func checkRules(rules []Rule) error {
	for _, rule := range rules {
		consequenceParams := rule.Consequence.Params

		if rule.ObjectID == "" {
			return emptyField("Rule.ObjectID")
		}

		// The elements of the RuleConsequence's Params map are first checked
		// as query parameters (as any regular query parameters is considered
		// valid) at the exception of the "query" field, if any, which is
		// ignored. This special case is mandatory as `query` is expected to be
		// a string in the context of a regular search query but it may also be
		// a `QueryIncrementalEdit` in the context of a RuleConsequence Params.
		if err := checkQuery(consequenceParams, "query"); err != nil {
			return err
		}

		// Now that that RuleConsequence's Params map elements have been
		// type-checked as regular search parameters, they are now checked as
		// RuleConsequence parameters.
		for k, v := range consequenceParams {
			switch k {

			case "query":
				switch v.(type) {
				case string, QueryIncrementalEdit, Edit, Map, map[string]interface{}:
					// OK
				default:
					return invalidType(k, "string, QueryIncrementalEdit, Edit or Map")
				}

			case "automaticFacetFilters",
				"automaticOptionalFacetFilters":
				switch v.(type) {
				case []string, []AutomaticFacetFilter:
					// OK
				default:
					return invalidType(k, "[]string or []AutomaticFacetFilter")
				}

			default:
				// OK
			}
		}
	}

	return nil
}

func checkSearchRulesParams(params Map) error {
	for k, v := range params {
		switch k {

		case "query",
			"context":
			if _, ok := v.(string); !ok {
				return invalidType(k, "string")
			}

		case "anchoring":
			if _, ok := v.(RulePatternAnchoring); !ok {
				return invalidType(k, "RulePatternAnchoring")
			}

		case "page",
			"hitsPerPage":
			if _, ok := v.(int); !ok {
				return invalidType(k, "int")
			}

		default:
			// OK

		}
	}

	return nil
}
