// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// StreamingTrigger Trigger input for continuously running tasks.
type StreamingTrigger struct {
	Type StreamingTriggerType `json:"type"`
}

// NewStreamingTrigger instantiates a new StreamingTrigger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStreamingTrigger(type_ StreamingTriggerType) *StreamingTrigger {
	this := &StreamingTrigger{}
	this.Type = type_
	return this
}

// NewEmptyStreamingTrigger return a pointer to an empty StreamingTrigger object.
func NewEmptyStreamingTrigger() *StreamingTrigger {
	return &StreamingTrigger{}
}

// GetType returns the Type field value.
func (o *StreamingTrigger) GetType() StreamingTriggerType {
	if o == nil {
		var ret StreamingTriggerType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StreamingTrigger) GetTypeOk() (*StreamingTriggerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *StreamingTrigger) SetType(v StreamingTriggerType) *StreamingTrigger {
	o.Type = v
	return o
}

func (o StreamingTrigger) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	toSerialize["type"] = o.Type
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal StreamingTrigger: %w", err)
	}

	return serialized, nil
}

func (o StreamingTrigger) String() string {
	out := ""
	out += fmt.Sprintf("  type=%v\n", o.Type)
	return fmt.Sprintf("StreamingTrigger {\n%s}", out)
}
