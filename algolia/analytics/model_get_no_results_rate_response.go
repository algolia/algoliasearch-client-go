// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package analytics

import (
	"encoding/json"
	"fmt"
)

// GetNoResultsRateResponse struct for GetNoResultsRateResponse.
type GetNoResultsRateResponse struct {
	// No results rate, calculated as number of searches with zero results divided by the total number of searches.
	Rate float64 `json:"rate"`
	// Number of searches.
	Count int32 `json:"count"`
	// Number of searches without any results.
	NoResultCount int32 `json:"noResultCount"`
	// Daily no results rates.
	Dates []DailyNoResultsRates `json:"dates"`
}

// NewGetNoResultsRateResponse instantiates a new GetNoResultsRateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetNoResultsRateResponse(rate float64, count int32, noResultCount int32, dates []DailyNoResultsRates) *GetNoResultsRateResponse {
	this := &GetNoResultsRateResponse{}
	this.Rate = rate
	this.Count = count
	this.NoResultCount = noResultCount
	this.Dates = dates
	return this
}

// NewEmptyGetNoResultsRateResponse return a pointer to an empty GetNoResultsRateResponse object.
func NewEmptyGetNoResultsRateResponse() *GetNoResultsRateResponse {
	return &GetNoResultsRateResponse{}
}

// GetRate returns the Rate field value.
func (o *GetNoResultsRateResponse) GetRate() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Rate
}

// GetRateOk returns a tuple with the Rate field value
// and a boolean to check if the value has been set.
func (o *GetNoResultsRateResponse) GetRateOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rate, true
}

// SetRate sets field value.
func (o *GetNoResultsRateResponse) SetRate(v float64) *GetNoResultsRateResponse {
	o.Rate = v
	return o
}

// GetCount returns the Count field value.
func (o *GetNoResultsRateResponse) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *GetNoResultsRateResponse) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value.
func (o *GetNoResultsRateResponse) SetCount(v int32) *GetNoResultsRateResponse {
	o.Count = v
	return o
}

// GetNoResultCount returns the NoResultCount field value.
func (o *GetNoResultsRateResponse) GetNoResultCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NoResultCount
}

// GetNoResultCountOk returns a tuple with the NoResultCount field value
// and a boolean to check if the value has been set.
func (o *GetNoResultsRateResponse) GetNoResultCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NoResultCount, true
}

// SetNoResultCount sets field value.
func (o *GetNoResultsRateResponse) SetNoResultCount(v int32) *GetNoResultsRateResponse {
	o.NoResultCount = v
	return o
}

// GetDates returns the Dates field value.
func (o *GetNoResultsRateResponse) GetDates() []DailyNoResultsRates {
	if o == nil {
		var ret []DailyNoResultsRates
		return ret
	}

	return o.Dates
}

// GetDatesOk returns a tuple with the Dates field value
// and a boolean to check if the value has been set.
func (o *GetNoResultsRateResponse) GetDatesOk() ([]DailyNoResultsRates, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dates, true
}

// SetDates sets field value.
func (o *GetNoResultsRateResponse) SetDates(v []DailyNoResultsRates) *GetNoResultsRateResponse {
	o.Dates = v
	return o
}

func (o GetNoResultsRateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["rate"] = o.Rate
	}
	if true {
		toSerialize["count"] = o.Count
	}
	if true {
		toSerialize["noResultCount"] = o.NoResultCount
	}
	if true {
		toSerialize["dates"] = o.Dates
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal GetNoResultsRateResponse: %w", err)
	}

	return serialized, nil
}

func (o GetNoResultsRateResponse) String() string {
	out := ""
	out += fmt.Sprintf("  rate=%v\n", o.Rate)
	out += fmt.Sprintf("  count=%v\n", o.Count)
	out += fmt.Sprintf("  noResultCount=%v\n", o.NoResultCount)
	out += fmt.Sprintf("  dates=%v\n", o.Dates)
	return fmt.Sprintf("GetNoResultsRateResponse {\n%s}", out)
}

type NullableGetNoResultsRateResponse struct {
	value *GetNoResultsRateResponse
	isSet bool
}

func (v NullableGetNoResultsRateResponse) Get() *GetNoResultsRateResponse {
	return v.value
}

func (v *NullableGetNoResultsRateResponse) Set(val *GetNoResultsRateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNoResultsRateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNoResultsRateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNoResultsRateResponse(val *GetNoResultsRateResponse) *NullableGetNoResultsRateResponse {
	return &NullableGetNoResultsRateResponse{value: val, isSet: true}
}

func (v NullableGetNoResultsRateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableGetNoResultsRateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
