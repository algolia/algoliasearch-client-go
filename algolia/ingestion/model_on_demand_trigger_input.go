// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// OnDemandTriggerInput Trigger information for manually-triggered tasks.
type OnDemandTriggerInput struct {
	Type OnDemandTriggerType `json:"type"`
}

// NewOnDemandTriggerInput instantiates a new OnDemandTriggerInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOnDemandTriggerInput(type_ OnDemandTriggerType) *OnDemandTriggerInput {
	this := &OnDemandTriggerInput{}
	this.Type = type_
	return this
}

// NewEmptyOnDemandTriggerInput return a pointer to an empty OnDemandTriggerInput object.
func NewEmptyOnDemandTriggerInput() *OnDemandTriggerInput {
	return &OnDemandTriggerInput{}
}

// GetType returns the Type field value.
func (o *OnDemandTriggerInput) GetType() OnDemandTriggerType {
	if o == nil {
		var ret OnDemandTriggerType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OnDemandTriggerInput) GetTypeOk() (*OnDemandTriggerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *OnDemandTriggerInput) SetType(v OnDemandTriggerType) *OnDemandTriggerInput {
	o.Type = v
	return o
}

func (o OnDemandTriggerInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	toSerialize["type"] = o.Type
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal OnDemandTriggerInput: %w", err)
	}

	return serialized, nil
}

func (o OnDemandTriggerInput) String() string {
	out := ""
	out += fmt.Sprintf("  type=%v\n", o.Type)
	return fmt.Sprintf("OnDemandTriggerInput {\n%s}", out)
}
