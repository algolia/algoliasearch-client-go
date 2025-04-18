// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package abtesting

import (
	"encoding/json"
	"fmt"
)

// ScheduleABTestResponse struct for ScheduleABTestResponse.
type ScheduleABTestResponse struct {
	// Unique scheduled A/B test identifier.
	AbTestScheduleID int32 `json:"abTestScheduleID"`
}

// NewScheduleABTestResponse instantiates a new ScheduleABTestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScheduleABTestResponse(abTestScheduleID int32) *ScheduleABTestResponse {
	this := &ScheduleABTestResponse{}
	this.AbTestScheduleID = abTestScheduleID
	return this
}

// NewEmptyScheduleABTestResponse return a pointer to an empty ScheduleABTestResponse object.
func NewEmptyScheduleABTestResponse() *ScheduleABTestResponse {
	return &ScheduleABTestResponse{}
}

// GetAbTestScheduleID returns the AbTestScheduleID field value.
func (o *ScheduleABTestResponse) GetAbTestScheduleID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AbTestScheduleID
}

// GetAbTestScheduleIDOk returns a tuple with the AbTestScheduleID field value
// and a boolean to check if the value has been set.
func (o *ScheduleABTestResponse) GetAbTestScheduleIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AbTestScheduleID, true
}

// SetAbTestScheduleID sets field value.
func (o *ScheduleABTestResponse) SetAbTestScheduleID(v int32) *ScheduleABTestResponse {
	o.AbTestScheduleID = v
	return o
}

func (o ScheduleABTestResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	toSerialize["abTestScheduleID"] = o.AbTestScheduleID
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal ScheduleABTestResponse: %w", err)
	}

	return serialized, nil
}

func (o ScheduleABTestResponse) String() string {
	out := ""
	out += fmt.Sprintf("  abTestScheduleID=%v\n", o.AbTestScheduleID)
	return fmt.Sprintf("ScheduleABTestResponse {\n%s}", out)
}
