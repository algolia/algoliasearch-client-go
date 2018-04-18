package algoliasearch

import (
	"encoding/json"
	"fmt"
)

type CreateObjectRes struct {
	CreatedAt string `json:"createdAt"`
	ObjectID  string `json:"objectID"`
	TaskID    int    `json:"taskID"`
}

type UpdateObjectRes struct {
	ObjectID  string `json:"objectID"`
	TaskID    int    `json:"taskID"`
	UpdatedAt string `json:"updatedAt"`
}

type objects struct {
	Results []Object `json:"results"`
}

type Object Map

func (o Object) ObjectID() (objectID string, err error) {
	i, ok := o["objectID"]
	if !ok {
		err = fmt.Errorf("Cannot extract `objectID` field from Object `%s`", o)
		return
	}

	if objectID, ok = i.(string); !ok {
		err = fmt.Errorf("Cannot cast `objectID` field to string type from Object `%s`", o)
	}

	return
}

func (o Object) String() string {
	jsonContent, err := json.Marshal(&o)
	if err == nil {
		return string(jsonContent)
	}
	return fmt.Sprintf("%#v", o)
}
