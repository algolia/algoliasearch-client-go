// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package composition

import (
	"encoding/json"
	"fmt"
)

// ResultsInjectedItemAppliedRulesInfoResponse struct for ResultsInjectedItemAppliedRulesInfoResponse.
type ResultsInjectedItemAppliedRulesInfoResponse struct {
	// Unique record identifier.
	ObjectID string `json:"objectID"`
}

// NewResultsInjectedItemAppliedRulesInfoResponse instantiates a new ResultsInjectedItemAppliedRulesInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewResultsInjectedItemAppliedRulesInfoResponse(objectID string) *ResultsInjectedItemAppliedRulesInfoResponse {
	this := &ResultsInjectedItemAppliedRulesInfoResponse{}
	this.ObjectID = objectID
	return this
}

// NewEmptyResultsInjectedItemAppliedRulesInfoResponse return a pointer to an empty ResultsInjectedItemAppliedRulesInfoResponse object.
func NewEmptyResultsInjectedItemAppliedRulesInfoResponse() *ResultsInjectedItemAppliedRulesInfoResponse {
	return &ResultsInjectedItemAppliedRulesInfoResponse{}
}

// GetObjectID returns the ObjectID field value.
func (o *ResultsInjectedItemAppliedRulesInfoResponse) GetObjectID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectID
}

// GetObjectIDOk returns a tuple with the ObjectID field value
// and a boolean to check if the value has been set.
func (o *ResultsInjectedItemAppliedRulesInfoResponse) GetObjectIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectID, true
}

// SetObjectID sets field value.
func (o *ResultsInjectedItemAppliedRulesInfoResponse) SetObjectID(v string) *ResultsInjectedItemAppliedRulesInfoResponse {
	o.ObjectID = v
	return o
}

func (o ResultsInjectedItemAppliedRulesInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	toSerialize["objectID"] = o.ObjectID
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal ResultsInjectedItemAppliedRulesInfoResponse: %w", err)
	}

	return serialized, nil
}

func (o ResultsInjectedItemAppliedRulesInfoResponse) String() string {
	out := ""
	out += fmt.Sprintf("  objectID=%v\n", o.ObjectID)
	return fmt.Sprintf("ResultsInjectedItemAppliedRulesInfoResponse {\n%s}", out)
}
