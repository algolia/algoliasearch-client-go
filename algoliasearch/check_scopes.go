package algoliasearch

import "fmt"

func checkScopes(scopes []string) error {
	for _, scope := range scopes {
		switch scope {
		case "rules",
			"settings",
			"synonyms":
			// OK
		default:
			return fmt.Errorf("Unknown scope \"%s\": allowed scopes are \"rules\", \"settings\" and \"synonyms\"", scope)
		}
	}
	return nil
}
