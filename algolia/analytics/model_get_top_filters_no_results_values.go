// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package analytics

import (
	"encoding/json"
	"fmt"
)

// GetTopFiltersNoResultsValues struct for GetTopFiltersNoResultsValues.
type GetTopFiltersNoResultsValues struct {
	// Number of occurrences.
	Count int32 `json:"count"`
	// Filters with no results.
	Values []GetTopFiltersNoResultsValue `json:"values"`
}

// NewGetTopFiltersNoResultsValues instantiates a new GetTopFiltersNoResultsValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetTopFiltersNoResultsValues(count int32, values []GetTopFiltersNoResultsValue) *GetTopFiltersNoResultsValues {
	this := &GetTopFiltersNoResultsValues{}
	this.Count = count
	this.Values = values
	return this
}

// NewEmptyGetTopFiltersNoResultsValues return a pointer to an empty GetTopFiltersNoResultsValues object.
func NewEmptyGetTopFiltersNoResultsValues() *GetTopFiltersNoResultsValues {
	return &GetTopFiltersNoResultsValues{}
}

// GetCount returns the Count field value.
func (o *GetTopFiltersNoResultsValues) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *GetTopFiltersNoResultsValues) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value.
func (o *GetTopFiltersNoResultsValues) SetCount(v int32) *GetTopFiltersNoResultsValues {
	o.Count = v
	return o
}

// GetValues returns the Values field value.
func (o *GetTopFiltersNoResultsValues) GetValues() []GetTopFiltersNoResultsValue {
	if o == nil {
		var ret []GetTopFiltersNoResultsValue
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *GetTopFiltersNoResultsValues) GetValuesOk() ([]GetTopFiltersNoResultsValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value.
func (o *GetTopFiltersNoResultsValues) SetValues(v []GetTopFiltersNoResultsValue) *GetTopFiltersNoResultsValues {
	o.Values = v
	return o
}

func (o GetTopFiltersNoResultsValues) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	toSerialize["count"] = o.Count
	toSerialize["values"] = o.Values
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal GetTopFiltersNoResultsValues: %w", err)
	}

	return serialized, nil
}

func (o GetTopFiltersNoResultsValues) String() string {
	out := ""
	out += fmt.Sprintf("  count=%v\n", o.Count)
	out += fmt.Sprintf("  values=%v\n", o.Values)
	return fmt.Sprintf("GetTopFiltersNoResultsValues {\n%s}", out)
}
