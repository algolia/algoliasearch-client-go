// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package abtesting

import (
	"encoding/json"
	"fmt"
)

// EstimateABTestResponse struct for EstimateABTestResponse.
type EstimateABTestResponse struct {
	// Estimated number of days needed to reach the sample sizes required for detecting the configured effect. This value is based on historical traffic.
	DurationDays *int64 `json:"durationDays,omitempty"`
	// Sample size estimates for each variant. The first element is the control variant. Each element is the estimated number of searches required to achieve the desired statistical significance.
	SampleSizes []int64 `json:"sampleSizes,omitempty"`
}

type EstimateABTestResponseOption func(f *EstimateABTestResponse)

func WithEstimateABTestResponseDurationDays(val int64) EstimateABTestResponseOption {
	return func(f *EstimateABTestResponse) {
		f.DurationDays = &val
	}
}

func WithEstimateABTestResponseSampleSizes(val []int64) EstimateABTestResponseOption {
	return func(f *EstimateABTestResponse) {
		f.SampleSizes = val
	}
}

// NewEstimateABTestResponse instantiates a new EstimateABTestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEstimateABTestResponse(opts ...EstimateABTestResponseOption) *EstimateABTestResponse {
	this := &EstimateABTestResponse{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyEstimateABTestResponse return a pointer to an empty EstimateABTestResponse object.
func NewEmptyEstimateABTestResponse() *EstimateABTestResponse {
	return &EstimateABTestResponse{}
}

// GetDurationDays returns the DurationDays field value if set, zero value otherwise.
func (o *EstimateABTestResponse) GetDurationDays() int64 {
	if o == nil || o.DurationDays == nil {
		var ret int64
		return ret
	}
	return *o.DurationDays
}

// GetDurationDaysOk returns a tuple with the DurationDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimateABTestResponse) GetDurationDaysOk() (*int64, bool) {
	if o == nil || o.DurationDays == nil {
		return nil, false
	}
	return o.DurationDays, true
}

// HasDurationDays returns a boolean if a field has been set.
func (o *EstimateABTestResponse) HasDurationDays() bool {
	if o != nil && o.DurationDays != nil {
		return true
	}

	return false
}

// SetDurationDays gets a reference to the given int64 and assigns it to the DurationDays field.
func (o *EstimateABTestResponse) SetDurationDays(v int64) *EstimateABTestResponse {
	o.DurationDays = &v
	return o
}

// GetSampleSizes returns the SampleSizes field value if set, zero value otherwise.
func (o *EstimateABTestResponse) GetSampleSizes() []int64 {
	if o == nil || o.SampleSizes == nil {
		var ret []int64
		return ret
	}
	return o.SampleSizes
}

// GetSampleSizesOk returns a tuple with the SampleSizes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimateABTestResponse) GetSampleSizesOk() ([]int64, bool) {
	if o == nil || o.SampleSizes == nil {
		return nil, false
	}
	return o.SampleSizes, true
}

// HasSampleSizes returns a boolean if a field has been set.
func (o *EstimateABTestResponse) HasSampleSizes() bool {
	if o != nil && o.SampleSizes != nil {
		return true
	}

	return false
}

// SetSampleSizes gets a reference to the given []int64 and assigns it to the SampleSizes field.
func (o *EstimateABTestResponse) SetSampleSizes(v []int64) *EstimateABTestResponse {
	o.SampleSizes = v
	return o
}

func (o EstimateABTestResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.DurationDays != nil {
		toSerialize["durationDays"] = o.DurationDays
	}
	if o.SampleSizes != nil {
		toSerialize["sampleSizes"] = o.SampleSizes
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal EstimateABTestResponse: %w", err)
	}

	return serialized, nil
}

func (o EstimateABTestResponse) String() string {
	out := ""
	out += fmt.Sprintf("  durationDays=%v\n", o.DurationDays)
	out += fmt.Sprintf("  sampleSizes=%v\n", o.SampleSizes)
	return fmt.Sprintf("EstimateABTestResponse {\n%s}", out)
}
