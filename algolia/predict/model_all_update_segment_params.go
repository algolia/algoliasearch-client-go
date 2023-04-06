// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package predict

import (
	"encoding/json"
	"fmt"
)

// AllUpdateSegmentParams struct for AllUpdateSegmentParams
type AllUpdateSegmentParams struct {
	// The name or description of the segment.
	Name       *string                  `json:"name,omitempty"`
	Conditions *SegmentParentConditions `json:"conditions,omitempty"`
}

type AllUpdateSegmentParamsOption func(f *AllUpdateSegmentParams)

func WithAllUpdateSegmentParamsName(val string) AllUpdateSegmentParamsOption {
	return func(f *AllUpdateSegmentParams) {
		f.Name = &val
	}
}

func WithAllUpdateSegmentParamsConditions(val SegmentParentConditions) AllUpdateSegmentParamsOption {
	return func(f *AllUpdateSegmentParams) {
		f.Conditions = &val
	}
}

// NewAllUpdateSegmentParams instantiates a new AllUpdateSegmentParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllUpdateSegmentParams(opts ...AllUpdateSegmentParamsOption) *AllUpdateSegmentParams {
	this := &AllUpdateSegmentParams{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewAllUpdateSegmentParamsWithDefaults instantiates a new AllUpdateSegmentParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllUpdateSegmentParamsWithDefaults() *AllUpdateSegmentParams {
	this := &AllUpdateSegmentParams{}
	return this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AllUpdateSegmentParams) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllUpdateSegmentParams) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AllUpdateSegmentParams) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AllUpdateSegmentParams) SetName(v string) {
	o.Name = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *AllUpdateSegmentParams) GetConditions() SegmentParentConditions {
	if o == nil || o.Conditions == nil {
		var ret SegmentParentConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllUpdateSegmentParams) GetConditionsOk() (*SegmentParentConditions, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *AllUpdateSegmentParams) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given SegmentParentConditions and assigns it to the Conditions field.
func (o *AllUpdateSegmentParams) SetConditions(v SegmentParentConditions) {
	o.Conditions = &v
}

func (o AllUpdateSegmentParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	return json.Marshal(toSerialize)
}

func (o AllUpdateSegmentParams) String() string {
	out := ""
	out += fmt.Sprintf("  name=%v\n", o.Name)
	out += fmt.Sprintf("  conditions=%v\n", o.Conditions)
	return fmt.Sprintf("AllUpdateSegmentParams {\n%s}", out)
}

type NullableAllUpdateSegmentParams struct {
	value *AllUpdateSegmentParams
	isSet bool
}

func (v NullableAllUpdateSegmentParams) Get() *AllUpdateSegmentParams {
	return v.value
}

func (v *NullableAllUpdateSegmentParams) Set(val *AllUpdateSegmentParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAllUpdateSegmentParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAllUpdateSegmentParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllUpdateSegmentParams(val *AllUpdateSegmentParams) *NullableAllUpdateSegmentParams {
	return &NullableAllUpdateSegmentParams{value: val, isSet: true}
}

func (v NullableAllUpdateSegmentParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllUpdateSegmentParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
