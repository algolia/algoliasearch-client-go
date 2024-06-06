// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package recommend

import (
	"encoding/json"
	"fmt"
)

// AroundPrecisionFromValueInner Range object with lower and upper values in meters to define custom ranges.
type AroundPrecisionFromValueInner struct {
	// Lower boundary of a range in meters. The Geo ranking criterion considers all records within the range to be equal.
	From *int32 `json:"from,omitempty"`
	// Upper boundary of a range in meters. The Geo ranking criterion considers all records within the range to be equal.
	Value *int32 `json:"value,omitempty"`
}

type AroundPrecisionFromValueInnerOption func(f *AroundPrecisionFromValueInner)

func WithAroundPrecisionFromValueInnerFrom(val int32) AroundPrecisionFromValueInnerOption {
	return func(f *AroundPrecisionFromValueInner) {
		f.From = &val
	}
}

func WithAroundPrecisionFromValueInnerValue(val int32) AroundPrecisionFromValueInnerOption {
	return func(f *AroundPrecisionFromValueInner) {
		f.Value = &val
	}
}

// NewAroundPrecisionFromValueInner instantiates a new AroundPrecisionFromValueInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAroundPrecisionFromValueInner(opts ...AroundPrecisionFromValueInnerOption) *AroundPrecisionFromValueInner {
	this := &AroundPrecisionFromValueInner{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyAroundPrecisionFromValueInner return a pointer to an empty AroundPrecisionFromValueInner object.
func NewEmptyAroundPrecisionFromValueInner() *AroundPrecisionFromValueInner {
	return &AroundPrecisionFromValueInner{}
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *AroundPrecisionFromValueInner) GetFrom() int32 {
	if o == nil || o.From == nil {
		var ret int32
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AroundPrecisionFromValueInner) GetFromOk() (*int32, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *AroundPrecisionFromValueInner) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given int32 and assigns it to the From field.
func (o *AroundPrecisionFromValueInner) SetFrom(v int32) *AroundPrecisionFromValueInner {
	o.From = &v
	return o
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AroundPrecisionFromValueInner) GetValue() int32 {
	if o == nil || o.Value == nil {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AroundPrecisionFromValueInner) GetValueOk() (*int32, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AroundPrecisionFromValueInner) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *AroundPrecisionFromValueInner) SetValue(v int32) *AroundPrecisionFromValueInner {
	o.Value = &v
	return o
}

func (o AroundPrecisionFromValueInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal AroundPrecisionFromValueInner: %w", err)
	}

	return serialized, nil
}

func (o AroundPrecisionFromValueInner) String() string {
	out := ""
	out += fmt.Sprintf("  from=%v\n", o.From)
	out += fmt.Sprintf("  value=%v\n", o.Value)
	return fmt.Sprintf("AroundPrecisionFromValueInner {\n%s}", out)
}

type NullableAroundPrecisionFromValueInner struct {
	value *AroundPrecisionFromValueInner
	isSet bool
}

func (v NullableAroundPrecisionFromValueInner) Get() *AroundPrecisionFromValueInner {
	return v.value
}

func (v *NullableAroundPrecisionFromValueInner) Set(val *AroundPrecisionFromValueInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAroundPrecisionFromValueInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAroundPrecisionFromValueInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAroundPrecisionFromValueInner(val *AroundPrecisionFromValueInner) *NullableAroundPrecisionFromValueInner {
	return &NullableAroundPrecisionFromValueInner{value: val, isSet: true}
}

func (v NullableAroundPrecisionFromValueInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableAroundPrecisionFromValueInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
