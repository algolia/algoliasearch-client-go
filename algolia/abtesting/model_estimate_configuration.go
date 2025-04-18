// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package abtesting

import (
	"encoding/json"
	"fmt"
)

// EstimateConfiguration A/B test configuration for estimating the sample size and duration using minimum detectable effect.
type EstimateConfiguration struct {
	Outliers                *Outliers               `json:"outliers,omitempty"`
	EmptySearch             *EmptySearch            `json:"emptySearch,omitempty"`
	MinimumDetectableEffect MinimumDetectableEffect `json:"minimumDetectableEffect"`
}

type EstimateConfigurationOption func(f *EstimateConfiguration)

func WithEstimateConfigurationOutliers(val Outliers) EstimateConfigurationOption {
	return func(f *EstimateConfiguration) {
		f.Outliers = &val
	}
}

func WithEstimateConfigurationEmptySearch(val EmptySearch) EstimateConfigurationOption {
	return func(f *EstimateConfiguration) {
		f.EmptySearch = &val
	}
}

// NewEstimateConfiguration instantiates a new EstimateConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEstimateConfiguration(minimumDetectableEffect MinimumDetectableEffect, opts ...EstimateConfigurationOption) *EstimateConfiguration {
	this := &EstimateConfiguration{}
	this.MinimumDetectableEffect = minimumDetectableEffect
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyEstimateConfiguration return a pointer to an empty EstimateConfiguration object.
func NewEmptyEstimateConfiguration() *EstimateConfiguration {
	return &EstimateConfiguration{}
}

// GetOutliers returns the Outliers field value if set, zero value otherwise.
func (o *EstimateConfiguration) GetOutliers() Outliers {
	if o == nil || o.Outliers == nil {
		var ret Outliers
		return ret
	}
	return *o.Outliers
}

// GetOutliersOk returns a tuple with the Outliers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimateConfiguration) GetOutliersOk() (*Outliers, bool) {
	if o == nil || o.Outliers == nil {
		return nil, false
	}
	return o.Outliers, true
}

// HasOutliers returns a boolean if a field has been set.
func (o *EstimateConfiguration) HasOutliers() bool {
	if o != nil && o.Outliers != nil {
		return true
	}

	return false
}

// SetOutliers gets a reference to the given Outliers and assigns it to the Outliers field.
func (o *EstimateConfiguration) SetOutliers(v *Outliers) *EstimateConfiguration {
	o.Outliers = v
	return o
}

// GetEmptySearch returns the EmptySearch field value if set, zero value otherwise.
func (o *EstimateConfiguration) GetEmptySearch() EmptySearch {
	if o == nil || o.EmptySearch == nil {
		var ret EmptySearch
		return ret
	}
	return *o.EmptySearch
}

// GetEmptySearchOk returns a tuple with the EmptySearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimateConfiguration) GetEmptySearchOk() (*EmptySearch, bool) {
	if o == nil || o.EmptySearch == nil {
		return nil, false
	}
	return o.EmptySearch, true
}

// HasEmptySearch returns a boolean if a field has been set.
func (o *EstimateConfiguration) HasEmptySearch() bool {
	if o != nil && o.EmptySearch != nil {
		return true
	}

	return false
}

// SetEmptySearch gets a reference to the given EmptySearch and assigns it to the EmptySearch field.
func (o *EstimateConfiguration) SetEmptySearch(v *EmptySearch) *EstimateConfiguration {
	o.EmptySearch = v
	return o
}

// GetMinimumDetectableEffect returns the MinimumDetectableEffect field value.
func (o *EstimateConfiguration) GetMinimumDetectableEffect() MinimumDetectableEffect {
	if o == nil {
		var ret MinimumDetectableEffect
		return ret
	}

	return o.MinimumDetectableEffect
}

// GetMinimumDetectableEffectOk returns a tuple with the MinimumDetectableEffect field value
// and a boolean to check if the value has been set.
func (o *EstimateConfiguration) GetMinimumDetectableEffectOk() (*MinimumDetectableEffect, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinimumDetectableEffect, true
}

// SetMinimumDetectableEffect sets field value.
func (o *EstimateConfiguration) SetMinimumDetectableEffect(v *MinimumDetectableEffect) *EstimateConfiguration {
	o.MinimumDetectableEffect = *v
	return o
}

func (o EstimateConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Outliers != nil {
		toSerialize["outliers"] = o.Outliers
	}
	if o.EmptySearch != nil {
		toSerialize["emptySearch"] = o.EmptySearch
	}
	toSerialize["minimumDetectableEffect"] = o.MinimumDetectableEffect
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal EstimateConfiguration: %w", err)
	}

	return serialized, nil
}

func (o EstimateConfiguration) String() string {
	out := ""
	out += fmt.Sprintf("  outliers=%v\n", o.Outliers)
	out += fmt.Sprintf("  emptySearch=%v\n", o.EmptySearch)
	out += fmt.Sprintf("  minimumDetectableEffect=%v\n", o.MinimumDetectableEffect)
	return fmt.Sprintf("EstimateConfiguration {\n%s}", out)
}
