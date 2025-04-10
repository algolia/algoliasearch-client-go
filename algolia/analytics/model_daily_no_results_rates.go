// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package analytics

import (
	"encoding/json"
	"fmt"
)

// DailyNoResultsRates struct for DailyNoResultsRates.
type DailyNoResultsRates struct {
	// Date in the format YYYY-MM-DD.
	Date string `json:"date"`
	// Number of searches without any results.
	NoResultCount int32 `json:"noResultCount"`
	// Number of searches.
	Count int32 `json:"count"`
	// No results rate: calculated as the number of searches with zero results divided by the total number of searches.
	Rate float64 `json:"rate"`
}

// NewDailyNoResultsRates instantiates a new DailyNoResultsRates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDailyNoResultsRates(date string, noResultCount int32, count int32, rate float64) *DailyNoResultsRates {
	this := &DailyNoResultsRates{}
	this.Date = date
	this.NoResultCount = noResultCount
	this.Count = count
	this.Rate = rate
	return this
}

// NewEmptyDailyNoResultsRates return a pointer to an empty DailyNoResultsRates object.
func NewEmptyDailyNoResultsRates() *DailyNoResultsRates {
	return &DailyNoResultsRates{}
}

// GetDate returns the Date field value.
func (o *DailyNoResultsRates) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *DailyNoResultsRates) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value.
func (o *DailyNoResultsRates) SetDate(v string) *DailyNoResultsRates {
	o.Date = v
	return o
}

// GetNoResultCount returns the NoResultCount field value.
func (o *DailyNoResultsRates) GetNoResultCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NoResultCount
}

// GetNoResultCountOk returns a tuple with the NoResultCount field value
// and a boolean to check if the value has been set.
func (o *DailyNoResultsRates) GetNoResultCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NoResultCount, true
}

// SetNoResultCount sets field value.
func (o *DailyNoResultsRates) SetNoResultCount(v int32) *DailyNoResultsRates {
	o.NoResultCount = v
	return o
}

// GetCount returns the Count field value.
func (o *DailyNoResultsRates) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *DailyNoResultsRates) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value.
func (o *DailyNoResultsRates) SetCount(v int32) *DailyNoResultsRates {
	o.Count = v
	return o
}

// GetRate returns the Rate field value.
func (o *DailyNoResultsRates) GetRate() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Rate
}

// GetRateOk returns a tuple with the Rate field value
// and a boolean to check if the value has been set.
func (o *DailyNoResultsRates) GetRateOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rate, true
}

// SetRate sets field value.
func (o *DailyNoResultsRates) SetRate(v float64) *DailyNoResultsRates {
	o.Rate = v
	return o
}

func (o DailyNoResultsRates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	toSerialize["date"] = o.Date
	toSerialize["noResultCount"] = o.NoResultCount
	toSerialize["count"] = o.Count
	toSerialize["rate"] = o.Rate
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal DailyNoResultsRates: %w", err)
	}

	return serialized, nil
}

func (o DailyNoResultsRates) String() string {
	out := ""
	out += fmt.Sprintf("  date=%v\n", o.Date)
	out += fmt.Sprintf("  noResultCount=%v\n", o.NoResultCount)
	out += fmt.Sprintf("  count=%v\n", o.Count)
	out += fmt.Sprintf("  rate=%v\n", o.Rate)
	return fmt.Sprintf("DailyNoResultsRates {\n%s}", out)
}
