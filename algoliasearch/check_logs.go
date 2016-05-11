package algoliasearch

func checkGetLogs(params map[string]interface{}) error {
	for k, v := range params {
		switch k {
		case "length", "offset":
			if _, ok := v.(int64); !ok {
				return invalidType(k, "int64")
			}

		case "type":
			if _, ok := v.(string); !ok {
				return invalidType(k, "string")
			}

		default:
		}
	}

	return nil
}
