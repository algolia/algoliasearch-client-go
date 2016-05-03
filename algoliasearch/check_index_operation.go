package algoliasearch

import "fmt"

func checkIndexOperation(o IndexOperation) error {
	if o.Operation != "copy" && o.Operation != "move" {
		return fmt.Errorf("Operation %s doesn't exist", o.Operation)
	}

	return nil
}
