// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package analytics

import (
	"encoding/json"
	"fmt"
)

// ClickPosition struct for ClickPosition
type ClickPosition struct {
	// Range of positions with the following pattern: - Positions from 1 to 10 included are displayed in separated groups. - Positions from 11 to 20 included are grouped together. - Positions from 21 and up are grouped together.
	Position []int32 `json:"position" validate:"required"`
	// The number of click event.
	ClickCount int32 `json:"clickCount" validate:"required"`
}

// NewClickPosition instantiates a new ClickPosition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClickPosition(position []int32, clickCount int32) *ClickPosition {
	this := &ClickPosition{}
	this.Position = position
	this.ClickCount = clickCount
	return this
}

// NewClickPositionWithDefaults instantiates a new ClickPosition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClickPositionWithDefaults() *ClickPosition {
	this := &ClickPosition{}
	return this
}

// GetPosition returns the Position field value
func (o *ClickPosition) GetPosition() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *ClickPosition) GetPositionOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Position, true
}

// SetPosition sets field value
func (o *ClickPosition) SetPosition(v []int32) {
	o.Position = v
}

// GetClickCount returns the ClickCount field value
func (o *ClickPosition) GetClickCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ClickCount
}

// GetClickCountOk returns a tuple with the ClickCount field value
// and a boolean to check if the value has been set.
func (o *ClickPosition) GetClickCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClickCount, true
}

// SetClickCount sets field value
func (o *ClickPosition) SetClickCount(v int32) {
	o.ClickCount = v
}

func (o ClickPosition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["position"] = o.Position
	}
	if true {
		toSerialize["clickCount"] = o.ClickCount
	}
	return json.Marshal(toSerialize)
}

func (o ClickPosition) String() string {
	out := ""
	out += fmt.Sprintf("  position=%v\n", o.Position)
	out += fmt.Sprintf("  clickCount=%v\n", o.ClickCount)
	return fmt.Sprintf("ClickPosition {\n%s}", out)
}

type NullableClickPosition struct {
	value *ClickPosition
	isSet bool
}

func (v NullableClickPosition) Get() *ClickPosition {
	return v.value
}

func (v *NullableClickPosition) Set(val *ClickPosition) {
	v.value = val
	v.isSet = true
}

func (v NullableClickPosition) IsSet() bool {
	return v.isSet
}

func (v *NullableClickPosition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClickPosition(val *ClickPosition) *NullableClickPosition {
	return &NullableClickPosition{value: val, isSet: true}
}

func (v NullableClickPosition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClickPosition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}