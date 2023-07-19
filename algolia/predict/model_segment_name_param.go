// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package predict

import (
	"encoding/json"
	"fmt"
)

// SegmentNameParam struct for SegmentNameParam
type SegmentNameParam struct {
	// The name or description of the segment.
	Name *string `json:"name,omitempty"`
}

type SegmentNameParamOption func(f *SegmentNameParam)

func WithSegmentNameParamName(val string) SegmentNameParamOption {
	return func(f *SegmentNameParam) {
		f.Name = &val
	}
}

// NewSegmentNameParam instantiates a new SegmentNameParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSegmentNameParam(opts ...SegmentNameParamOption) *SegmentNameParam {
	this := &SegmentNameParam{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewSegmentNameParamWithDefaults instantiates a new SegmentNameParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSegmentNameParamWithDefaults() *SegmentNameParam {
	this := &SegmentNameParam{}
	return this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SegmentNameParam) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentNameParam) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SegmentNameParam) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SegmentNameParam) SetName(v string) {
	o.Name = &v
}

func (o SegmentNameParam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

func (o SegmentNameParam) String() string {
	out := ""
	out += fmt.Sprintf("  name=%v\n", o.Name)
	return fmt.Sprintf("SegmentNameParam {\n%s}", out)
}

type NullableSegmentNameParam struct {
	value *SegmentNameParam
	isSet bool
}

func (v NullableSegmentNameParam) Get() *SegmentNameParam {
	return v.value
}

func (v *NullableSegmentNameParam) Set(val *SegmentNameParam) {
	v.value = val
	v.isSet = true
}

func (v NullableSegmentNameParam) IsSet() bool {
	return v.isSet
}

func (v *NullableSegmentNameParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSegmentNameParam(val *SegmentNameParam) *NullableSegmentNameParam {
	return &NullableSegmentNameParam{value: val, isSet: true}
}

func (v NullableSegmentNameParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSegmentNameParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}