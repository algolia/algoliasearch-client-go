package algoliasearch

import "fmt"

func checkKey(k Key) error {
	for _, access := range k.ACL {
		if access != "search" && access != "browse" && access != "addObject" &&
			access != "deleteObject" && access != "deleteIndex" && access != "settings" &&
			access != "editSettings" && access != "analytics" && access != "listIndexes" {
			return fmt.Errorf("ACL element %s doesn't exist", access)
		}
	}

	return nil
}
