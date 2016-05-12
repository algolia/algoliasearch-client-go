package algoliasearch

import "fmt"

type objects struct {
	Results []Object `json:"results"`
}

type Object map[string]interface{}

func (o Object) ObjectID() (objectID string, err error) {
	i, ok := o["objectID"]
	if !ok {
		err = fmt.Errorf("Cannot extract `objectID` field from Object")
		return
	}

	if objectID, ok = i.(string); !ok {
		err = fmt.Errorf("Cannot cast `objectID` field to string type")
	}

	return
}
