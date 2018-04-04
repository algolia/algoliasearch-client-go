package algoliasearch

func checkGenerateSecuredAPIKey(params Map) error {
	if err := checkQuery(params,
		"restrictIndices",
		"restrictSources",
		"userToken",
		"validUntil",
	); err != nil {
		return err
	}

	for k, v := range params {
		switch k {
		case "restrictIndices", "restrictSources", "userToken":
			if _, ok := v.(string); !ok {
				return invalidType(k, "string")
			}

		case "validUntil":
			if _, ok := v.(int); !ok {
				return invalidType(k, "int")
			}

		default:
			// OK
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
			// OK
		}
	}

	return nil
}
