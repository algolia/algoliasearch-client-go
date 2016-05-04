package algoliasearch

func checkGenerateSecuredAPIKey(params map[string]interface{}) error {
	if err := checkQuery(params, "userToken", "tagFilters"); err != nil {
		return err
	}

	if v, ok := params["userToken"]; ok {
		if _, ok := v.(string); !ok {
			return invalidParameter("userToken")
		}
	}

	if v, ok := params["tagFilters"]; ok {
		if _, ok := v.([]string); !ok {
			return invalidParameter("tagFilters")
		}
	}

	return nil
}

func checkKey(params map[string]interface{}) error {
	for k, v := range params {
		switch v.(type) {
		case []string:
			if k != "acl" &&
				k != "indexes" &&
				k != "referers" {
				return invalidParameter(k)
			}

		case string:
			if k != "description" &&
				k != "queryParameters" {
				return invalidParameter(k)
			}

		case int64:
			if k != "maxHitsPerQuery" &&
				k != "maxQueriesPerIPPerHour" &&
				k != "validity" {
				return invalidParameter(k)
			}

		default:
			return invalidParameter(k)
		}
	}

	return nil
}
