package algoliasearch

func checkGetLogs(params map[string]interface{}) error {
	for k, v := range params {
		switch v.(type) {
		case int64:
			if k != "length" &&
				k != "offset" {
				return invalidParameter(k)
			}

		case string:
			if k != "type" {
				return invalidParameter(k)
			}

		default:
			return invalidParameter(k)
		}
	}

	return nil
}
