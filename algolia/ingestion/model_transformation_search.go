// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// TransformationSearch struct for TransformationSearch.
type TransformationSearch struct {
	TransformationIDs []string `json:"transformationIDs"`
}

// NewTransformationSearch instantiates a new TransformationSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTransformationSearch(transformationIDs []string) *TransformationSearch {
	this := &TransformationSearch{}
	this.TransformationIDs = transformationIDs
	return this
}

// NewEmptyTransformationSearch return a pointer to an empty TransformationSearch object.
func NewEmptyTransformationSearch() *TransformationSearch {
	return &TransformationSearch{}
}

// GetTransformationIDs returns the TransformationIDs field value.
func (o *TransformationSearch) GetTransformationIDs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TransformationIDs
}

// GetTransformationIDsOk returns a tuple with the TransformationIDs field value
// and a boolean to check if the value has been set.
func (o *TransformationSearch) GetTransformationIDsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransformationIDs, true
}

// SetTransformationIDs sets field value.
func (o *TransformationSearch) SetTransformationIDs(v []string) *TransformationSearch {
	o.TransformationIDs = v
	return o
}

func (o TransformationSearch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	toSerialize["transformationIDs"] = o.TransformationIDs
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal TransformationSearch: %w", err)
	}

	return serialized, nil
}

func (o TransformationSearch) String() string {
	out := ""
	out += fmt.Sprintf("  transformationIDs=%v\n", o.TransformationIDs)
	return fmt.Sprintf("TransformationSearch {\n%s}", out)
}
