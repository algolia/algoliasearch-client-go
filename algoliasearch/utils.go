package algoliasearch

import "fmt"

func wrongType(expected, key string) error {
	return fmt.Errorf("Expected type `%s` for `%s`", expected, key)
}

func unknownField(field string) error {
	return fmt.Errorf("Unexpected field `%s`", field)
}
