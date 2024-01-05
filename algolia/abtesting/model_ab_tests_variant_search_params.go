// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package abtesting

import (
	"encoding/json"
	"fmt"
)

// AbTestsVariantSearchParams struct for AbTestsVariantSearchParams.
type AbTestsVariantSearchParams struct {
	// A/B test index.
	Index string `json:"index"`
	// A/B test traffic percentage.
	TrafficPercentage int32 `json:"trafficPercentage"`
	// A/B test description.
	Description            *string                `json:"description,omitempty"`
	CustomSearchParameters map[string]interface{} `json:"customSearchParameters"`
}

type AbTestsVariantSearchParamsOption func(f *AbTestsVariantSearchParams)

func WithAbTestsVariantSearchParamsDescription(val string) AbTestsVariantSearchParamsOption {
	return func(f *AbTestsVariantSearchParams) {
		f.Description = &val
	}
}

// NewAbTestsVariantSearchParams instantiates a new AbTestsVariantSearchParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAbTestsVariantSearchParams(index string, trafficPercentage int32, customSearchParameters map[string]interface{}, opts ...AbTestsVariantSearchParamsOption) *AbTestsVariantSearchParams {
	this := &AbTestsVariantSearchParams{}
	this.Index = index
	this.TrafficPercentage = trafficPercentage
	this.CustomSearchParameters = customSearchParameters
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewAbTestsVariantSearchParamsWithDefaults instantiates a new AbTestsVariantSearchParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAbTestsVariantSearchParamsWithDefaults() *AbTestsVariantSearchParams {
	this := &AbTestsVariantSearchParams{}
	return this
}

// GetIndex returns the Index field value.
func (o *AbTestsVariantSearchParams) GetIndex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *AbTestsVariantSearchParams) GetIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value.
func (o *AbTestsVariantSearchParams) SetIndex(v string) {
	o.Index = v
}

// GetTrafficPercentage returns the TrafficPercentage field value.
func (o *AbTestsVariantSearchParams) GetTrafficPercentage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TrafficPercentage
}

// GetTrafficPercentageOk returns a tuple with the TrafficPercentage field value
// and a boolean to check if the value has been set.
func (o *AbTestsVariantSearchParams) GetTrafficPercentageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrafficPercentage, true
}

// SetTrafficPercentage sets field value.
func (o *AbTestsVariantSearchParams) SetTrafficPercentage(v int32) {
	o.TrafficPercentage = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AbTestsVariantSearchParams) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbTestsVariantSearchParams) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AbTestsVariantSearchParams) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AbTestsVariantSearchParams) SetDescription(v string) {
	o.Description = &v
}

// GetCustomSearchParameters returns the CustomSearchParameters field value.
func (o *AbTestsVariantSearchParams) GetCustomSearchParameters() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.CustomSearchParameters
}

// GetCustomSearchParametersOk returns a tuple with the CustomSearchParameters field value
// and a boolean to check if the value has been set.
func (o *AbTestsVariantSearchParams) GetCustomSearchParametersOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomSearchParameters, true
}

// SetCustomSearchParameters sets field value.
func (o *AbTestsVariantSearchParams) SetCustomSearchParameters(v map[string]interface{}) {
	o.CustomSearchParameters = v
}

func (o AbTestsVariantSearchParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["index"] = o.Index
	}
	if true {
		toSerialize["trafficPercentage"] = o.TrafficPercentage
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["customSearchParameters"] = o.CustomSearchParameters
	}
	return json.Marshal(toSerialize)
}

func (o AbTestsVariantSearchParams) String() string {
	out := ""
	out += fmt.Sprintf("  index=%v\n", o.Index)
	out += fmt.Sprintf("  trafficPercentage=%v\n", o.TrafficPercentage)
	out += fmt.Sprintf("  description=%v\n", o.Description)
	out += fmt.Sprintf("  customSearchParameters=%v\n", o.CustomSearchParameters)
	return fmt.Sprintf("AbTestsVariantSearchParams {\n%s}", out)
}

type NullableAbTestsVariantSearchParams struct {
	value *AbTestsVariantSearchParams
	isSet bool
}

func (v NullableAbTestsVariantSearchParams) Get() *AbTestsVariantSearchParams {
	return v.value
}

func (v *NullableAbTestsVariantSearchParams) Set(val *AbTestsVariantSearchParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAbTestsVariantSearchParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAbTestsVariantSearchParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAbTestsVariantSearchParams(val *AbTestsVariantSearchParams) *NullableAbTestsVariantSearchParams {
	return &NullableAbTestsVariantSearchParams{value: val, isSet: true}
}

func (v NullableAbTestsVariantSearchParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAbTestsVariantSearchParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
