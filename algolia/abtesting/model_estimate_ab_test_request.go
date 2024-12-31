// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package abtesting

import (
	"encoding/json"
	"fmt"
)

// EstimateABTestRequest struct for EstimateABTestRequest.
type EstimateABTestRequest struct {
	Configuration EstimateConfiguration `json:"configuration"`
	// A/B test variants.
	Variants []AddABTestsVariant `json:"variants"`
}

// NewEstimateABTestRequest instantiates a new EstimateABTestRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEstimateABTestRequest(configuration EstimateConfiguration, variants []AddABTestsVariant) *EstimateABTestRequest {
	this := &EstimateABTestRequest{}
	this.Configuration = configuration
	this.Variants = variants
	return this
}

// NewEmptyEstimateABTestRequest return a pointer to an empty EstimateABTestRequest object.
func NewEmptyEstimateABTestRequest() *EstimateABTestRequest {
	return &EstimateABTestRequest{}
}

// GetConfiguration returns the Configuration field value.
func (o *EstimateABTestRequest) GetConfiguration() EstimateConfiguration {
	if o == nil {
		var ret EstimateConfiguration
		return ret
	}

	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value
// and a boolean to check if the value has been set.
func (o *EstimateABTestRequest) GetConfigurationOk() (*EstimateConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Configuration, true
}

// SetConfiguration sets field value.
func (o *EstimateABTestRequest) SetConfiguration(v *EstimateConfiguration) *EstimateABTestRequest {
	o.Configuration = *v
	return o
}

// GetVariants returns the Variants field value.
func (o *EstimateABTestRequest) GetVariants() []AddABTestsVariant {
	if o == nil {
		var ret []AddABTestsVariant
		return ret
	}

	return o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value
// and a boolean to check if the value has been set.
func (o *EstimateABTestRequest) GetVariantsOk() ([]AddABTestsVariant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Variants, true
}

// SetVariants sets field value.
func (o *EstimateABTestRequest) SetVariants(v []AddABTestsVariant) *EstimateABTestRequest {
	o.Variants = v
	return o
}

func (o EstimateABTestRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["configuration"] = o.Configuration
	}
	if true {
		toSerialize["variants"] = o.Variants
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal EstimateABTestRequest: %w", err)
	}

	return serialized, nil
}

func (o EstimateABTestRequest) String() string {
	out := ""
	out += fmt.Sprintf("  configuration=%v\n", o.Configuration)
	out += fmt.Sprintf("  variants=%v\n", o.Variants)
	return fmt.Sprintf("EstimateABTestRequest {\n%s}", out)
}