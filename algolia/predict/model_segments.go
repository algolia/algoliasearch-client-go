// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package predict

import (
	"encoding/json"
	"fmt"
)

// Segments Segments that the user belongs to.
type Segments struct {
	// List of computed segments IDs.
	Computed []string `json:"computed"`
	// List of custom segments IDs.
	Custom []string `json:"custom"`
}

// NewSegments instantiates a new Segments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSegments(computed []string, custom []string) *Segments {
	this := &Segments{}
	this.Computed = computed
	this.Custom = custom
	return this
}

// NewSegmentsWithDefaults instantiates a new Segments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSegmentsWithDefaults() *Segments {
	this := &Segments{}
	return this
}

// GetComputed returns the Computed field value
func (o *Segments) GetComputed() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Computed
}

// GetComputedOk returns a tuple with the Computed field value
// and a boolean to check if the value has been set.
func (o *Segments) GetComputedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Computed, true
}

// SetComputed sets field value
func (o *Segments) SetComputed(v []string) {
	o.Computed = v
}

// GetCustom returns the Custom field value
func (o *Segments) GetCustom() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Custom
}

// GetCustomOk returns a tuple with the Custom field value
// and a boolean to check if the value has been set.
func (o *Segments) GetCustomOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Custom, true
}

// SetCustom sets field value
func (o *Segments) SetCustom(v []string) {
	o.Custom = v
}

func (o Segments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["computed"] = o.Computed
	}
	if true {
		toSerialize["custom"] = o.Custom
	}
	return json.Marshal(toSerialize)
}

func (o Segments) String() string {
	out := ""
	out += fmt.Sprintf("  computed=%v\n", o.Computed)
	out += fmt.Sprintf("  custom=%v\n", o.Custom)
	return fmt.Sprintf("Segments {\n%s}", out)
}

type NullableSegments struct {
	value *Segments
	isSet bool
}

func (v NullableSegments) Get() *Segments {
	return v.value
}

func (v *NullableSegments) Set(val *Segments) {
	v.value = val
	v.isSet = true
}

func (v NullableSegments) IsSet() bool {
	return v.isSet
}

func (v *NullableSegments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSegments(val *Segments) *NullableSegments {
	return &NullableSegments{value: val, isSet: true}
}

func (v NullableSegments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSegments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}