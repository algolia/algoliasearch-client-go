// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package analytics

import (
	"encoding/json"
	"fmt"

	"github.com/algolia/algoliasearch-client-go/v4/algolia/utils"
)

// DailyAverageClicks struct for DailyAverageClicks.
type DailyAverageClicks struct {
	// Average position of a clicked search result in the list of search results. If null, Algolia didn't receive any search requests with `clickAnalytics` set to true.
	Average utils.NullableFloat64 `json:"average"`
	// Number of clicks associated with this search.
	ClickCount int32 `json:"clickCount"`
	// Date in the format YYYY-MM-DD.
	Date string `json:"date"`
}

// NewDailyAverageClicks instantiates a new DailyAverageClicks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDailyAverageClicks(average utils.NullableFloat64, clickCount int32, date string) *DailyAverageClicks {
	this := &DailyAverageClicks{}
	this.Average = average
	this.ClickCount = clickCount
	this.Date = date
	return this
}

// NewEmptyDailyAverageClicks return a pointer to an empty DailyAverageClicks object.
func NewEmptyDailyAverageClicks() *DailyAverageClicks {
	return &DailyAverageClicks{}
}

// GetAverage returns the Average field value.
// If the value is explicit nil, the zero value for float64 will be returned.
func (o *DailyAverageClicks) GetAverage() float64 {
	if o == nil || o.Average.Get() == nil {
		var ret float64
		return ret
	}

	return *o.Average.Get()
}

// GetAverageOk returns a tuple with the Average field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DailyAverageClicks) GetAverageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Average.Get(), o.Average.IsSet()
}

// SetAverage sets field value.
func (o *DailyAverageClicks) SetAverage(v float64) *DailyAverageClicks {
	o.Average.Set(&v)
	return o
}

// GetClickCount returns the ClickCount field value.
func (o *DailyAverageClicks) GetClickCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ClickCount
}

// GetClickCountOk returns a tuple with the ClickCount field value
// and a boolean to check if the value has been set.
func (o *DailyAverageClicks) GetClickCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClickCount, true
}

// SetClickCount sets field value.
func (o *DailyAverageClicks) SetClickCount(v int32) *DailyAverageClicks {
	o.ClickCount = v
	return o
}

// GetDate returns the Date field value.
func (o *DailyAverageClicks) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *DailyAverageClicks) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value.
func (o *DailyAverageClicks) SetDate(v string) *DailyAverageClicks {
	o.Date = v
	return o
}

func (o DailyAverageClicks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["average"] = o.Average.Get()
	}
	if true {
		toSerialize["clickCount"] = o.ClickCount
	}
	if true {
		toSerialize["date"] = o.Date
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal DailyAverageClicks: %w", err)
	}

	return serialized, nil
}

func (o DailyAverageClicks) String() string {
	out := ""
	out += fmt.Sprintf("  average=%v\n", o.Average)
	out += fmt.Sprintf("  clickCount=%v\n", o.ClickCount)
	out += fmt.Sprintf("  date=%v\n", o.Date)
	return fmt.Sprintf("DailyAverageClicks {\n%s}", out)
}

type NullableDailyAverageClicks struct {
	value *DailyAverageClicks
	isSet bool
}

func (v NullableDailyAverageClicks) Get() *DailyAverageClicks {
	return v.value
}

func (v *NullableDailyAverageClicks) Set(val *DailyAverageClicks) {
	v.value = val
	v.isSet = true
}

func (v NullableDailyAverageClicks) IsSet() bool {
	return v.isSet
}

func (v *NullableDailyAverageClicks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDailyAverageClicks(val *DailyAverageClicks) *NullableDailyAverageClicks {
	return &NullableDailyAverageClicks{value: val, isSet: true}
}

func (v NullableDailyAverageClicks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableDailyAverageClicks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
