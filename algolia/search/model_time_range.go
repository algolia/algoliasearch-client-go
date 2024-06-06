// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// TimeRange struct for TimeRange.
type TimeRange struct {
	// When the rule should start to be active, in Unix epoch time.
	From int32 `json:"from"`
	// When the rule should stop to be active, in Unix epoch time.
	Until int32 `json:"until"`
}

// NewTimeRange instantiates a new TimeRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeRange(from int32, until int32) *TimeRange {
	this := &TimeRange{}
	this.From = from
	this.Until = until
	return this
}

// NewEmptyTimeRange return a pointer to an empty TimeRange object.
func NewEmptyTimeRange() *TimeRange {
	return &TimeRange{}
}

// GetFrom returns the From field value.
func (o *TimeRange) GetFrom() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *TimeRange) GetFromOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value.
func (o *TimeRange) SetFrom(v int32) *TimeRange {
	o.From = v
	return o
}

// GetUntil returns the Until field value.
func (o *TimeRange) GetUntil() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Until
}

// GetUntilOk returns a tuple with the Until field value
// and a boolean to check if the value has been set.
func (o *TimeRange) GetUntilOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Until, true
}

// SetUntil sets field value.
func (o *TimeRange) SetUntil(v int32) *TimeRange {
	o.Until = v
	return o
}

func (o TimeRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["from"] = o.From
	}
	if true {
		toSerialize["until"] = o.Until
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal TimeRange: %w", err)
	}

	return serialized, nil
}

func (o TimeRange) String() string {
	out := ""
	out += fmt.Sprintf("  from=%v\n", o.From)
	out += fmt.Sprintf("  until=%v\n", o.Until)
	return fmt.Sprintf("TimeRange {\n%s}", out)
}

type NullableTimeRange struct {
	value *TimeRange
	isSet bool
}

func (v NullableTimeRange) Get() *TimeRange {
	return v.value
}

func (v *NullableTimeRange) Set(val *TimeRange) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeRange) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeRange(val *TimeRange) *NullableTimeRange {
	return &NullableTimeRange{value: val, isSet: true}
}

func (v NullableTimeRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableTimeRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
