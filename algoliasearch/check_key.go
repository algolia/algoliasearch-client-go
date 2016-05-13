package algoliasearch

func checkGenerateSecuredAPIKey(params Map) error {
	if err := checkQuery(params, "userToken", "tagFilters"); err != nil {
		return err
	}

	if v, ok := params["userToken"]; ok {
		if _, ok := v.(string); !ok {
			return invalidType("userToken", "[]string")
		}
	}

	if v, ok := params["tagFilters"]; ok {
		if _, ok := v.(string); !ok {
			return invalidType("tagFilters", "[]string")
		}
	}

	return nil
}

func checkKey(params Map) error {
	for k, v := range params {
		switch k {
		case "acl", "indexes", "referers":
			if _, ok := v.([]string); !ok {
				return invalidType(k, "[]string")
			}

		case "description", "queryParameters":
			if _, ok := v.(string); !ok {
				return invalidType(k, "string")
			}

		case "maxHitsPerQuery", "maxQueriesPerIPPerHour", "validity":
			if _, ok := v.(int); !ok {
				return invalidType(k, "int")
			}

		default:
		}
	}

	return nil
}
