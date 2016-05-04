package algoliasearch

import "fmt"

func invalidParameter(p string) error {
	return fmt.Errorf("`%s` doesn't exist or doesn't have the right type", p)
}
