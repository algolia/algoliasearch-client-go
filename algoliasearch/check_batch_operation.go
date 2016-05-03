package algoliasearch

import "fmt"

func checkBatchOperation(o BatchOperation) (err error) {
	switch o.Action {

	case "addObject", "updateObject", "partialUpdateObject", "partialUpdateObjectNoCreate":
		if _, ok := o.Body.(Object); !ok {
			err = fmt.Errorf("BatchOperation %s doesn't store an Object", o.Action)
		}

	case "deleteObject":
		if _, ok := o.Body.(string); !ok {
			err = fmt.Errorf("BatchOperation %s doesn't store a string", o.Action)
		}

	case "clear":
		// Doesn't require to check the Body field.

	default:
		err = fmt.Errorf("BatchOperation's Action field (%s) doesn't exist", o.Action)

	}

	return
}
