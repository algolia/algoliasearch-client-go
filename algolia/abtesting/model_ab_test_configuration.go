// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package abtesting

import (
	"encoding/json"
	"fmt"
)

// ABTestConfiguration A/B test configuration.
type ABTestConfiguration struct {
	Outliers                Outliers                 `json:"outliers"`
	EmptySearch             *EmptySearch             `json:"emptySearch,omitempty"`
	MinimumDetectableEffect *MinimumDetectableEffect `json:"minimumDetectableEffect,omitempty"`
}

type ABTestConfigurationOption func(f *ABTestConfiguration)

func WithABTestConfigurationEmptySearch(val EmptySearch) ABTestConfigurationOption {
	return func(f *ABTestConfiguration) {
		f.EmptySearch = &val
	}
}

func WithABTestConfigurationMinimumDetectableEffect(val MinimumDetectableEffect) ABTestConfigurationOption {
	return func(f *ABTestConfiguration) {
		f.MinimumDetectableEffect = &val
	}
}

// NewABTestConfiguration instantiates a new ABTestConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewABTestConfiguration(outliers Outliers, opts ...ABTestConfigurationOption) *ABTestConfiguration {
	this := &ABTestConfiguration{}
	this.Outliers = outliers
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyABTestConfiguration return a pointer to an empty ABTestConfiguration object.
func NewEmptyABTestConfiguration() *ABTestConfiguration {
	return &ABTestConfiguration{}
}

// GetOutliers returns the Outliers field value.
func (o *ABTestConfiguration) GetOutliers() Outliers {
	if o == nil {
		var ret Outliers
		return ret
	}

	return o.Outliers
}

// GetOutliersOk returns a tuple with the Outliers field value
// and a boolean to check if the value has been set.
func (o *ABTestConfiguration) GetOutliersOk() (*Outliers, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Outliers, true
}

// SetOutliers sets field value.
func (o *ABTestConfiguration) SetOutliers(v *Outliers) *ABTestConfiguration {
	o.Outliers = *v
	return o
}

// GetEmptySearch returns the EmptySearch field value if set, zero value otherwise.
func (o *ABTestConfiguration) GetEmptySearch() EmptySearch {
	if o == nil || o.EmptySearch == nil {
		var ret EmptySearch
		return ret
	}
	return *o.EmptySearch
}

// GetEmptySearchOk returns a tuple with the EmptySearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ABTestConfiguration) GetEmptySearchOk() (*EmptySearch, bool) {
	if o == nil || o.EmptySearch == nil {
		return nil, false
	}
	return o.EmptySearch, true
}

// HasEmptySearch returns a boolean if a field has been set.
func (o *ABTestConfiguration) HasEmptySearch() bool {
	if o != nil && o.EmptySearch != nil {
		return true
	}

	return false
}

// SetEmptySearch gets a reference to the given EmptySearch and assigns it to the EmptySearch field.
func (o *ABTestConfiguration) SetEmptySearch(v *EmptySearch) *ABTestConfiguration {
	o.EmptySearch = v
	return o
}

// GetMinimumDetectableEffect returns the MinimumDetectableEffect field value if set, zero value otherwise.
func (o *ABTestConfiguration) GetMinimumDetectableEffect() MinimumDetectableEffect {
	if o == nil || o.MinimumDetectableEffect == nil {
		var ret MinimumDetectableEffect
		return ret
	}
	return *o.MinimumDetectableEffect
}

// GetMinimumDetectableEffectOk returns a tuple with the MinimumDetectableEffect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ABTestConfiguration) GetMinimumDetectableEffectOk() (*MinimumDetectableEffect, bool) {
	if o == nil || o.MinimumDetectableEffect == nil {
		return nil, false
	}
	return o.MinimumDetectableEffect, true
}

// HasMinimumDetectableEffect returns a boolean if a field has been set.
func (o *ABTestConfiguration) HasMinimumDetectableEffect() bool {
	if o != nil && o.MinimumDetectableEffect != nil {
		return true
	}

	return false
}

// SetMinimumDetectableEffect gets a reference to the given MinimumDetectableEffect and assigns it to the MinimumDetectableEffect field.
func (o *ABTestConfiguration) SetMinimumDetectableEffect(v *MinimumDetectableEffect) *ABTestConfiguration {
	o.MinimumDetectableEffect = v
	return o
}

func (o ABTestConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["outliers"] = o.Outliers
	}
	if o.EmptySearch != nil {
		toSerialize["emptySearch"] = o.EmptySearch
	}
	if o.MinimumDetectableEffect != nil {
		toSerialize["minimumDetectableEffect"] = o.MinimumDetectableEffect
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal ABTestConfiguration: %w", err)
	}

	return serialized, nil
}

func (o ABTestConfiguration) String() string {
	out := ""
	out += fmt.Sprintf("  outliers=%v\n", o.Outliers)
	out += fmt.Sprintf("  emptySearch=%v\n", o.EmptySearch)
	out += fmt.Sprintf("  minimumDetectableEffect=%v\n", o.MinimumDetectableEffect)
	return fmt.Sprintf("ABTestConfiguration {\n%s}", out)
}

type NullableABTestConfiguration struct {
	value *ABTestConfiguration
	isSet bool
}

func (v NullableABTestConfiguration) Get() *ABTestConfiguration {
	return v.value
}

func (v *NullableABTestConfiguration) Set(val *ABTestConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableABTestConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableABTestConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableABTestConfiguration(val *ABTestConfiguration) *NullableABTestConfiguration {
	return &NullableABTestConfiguration{value: val, isSet: true}
}

func (v NullableABTestConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableABTestConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
