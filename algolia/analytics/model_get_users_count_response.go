// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package analytics

import (
	"encoding/json"
	"fmt"
)

// GetUsersCountResponse struct for GetUsersCountResponse.
type GetUsersCountResponse struct {
	// Number of unique users.
	Count int32 `json:"count"`
	// Daily number of unique users.
	Dates []DailyUsers `json:"dates"`
}

// NewGetUsersCountResponse instantiates a new GetUsersCountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetUsersCountResponse(count int32, dates []DailyUsers) *GetUsersCountResponse {
	this := &GetUsersCountResponse{}
	this.Count = count
	this.Dates = dates
	return this
}

// NewEmptyGetUsersCountResponse return a pointer to an empty GetUsersCountResponse object.
func NewEmptyGetUsersCountResponse() *GetUsersCountResponse {
	return &GetUsersCountResponse{}
}

// GetCount returns the Count field value.
func (o *GetUsersCountResponse) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *GetUsersCountResponse) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value.
func (o *GetUsersCountResponse) SetCount(v int32) *GetUsersCountResponse {
	o.Count = v
	return o
}

// GetDates returns the Dates field value.
func (o *GetUsersCountResponse) GetDates() []DailyUsers {
	if o == nil {
		var ret []DailyUsers
		return ret
	}

	return o.Dates
}

// GetDatesOk returns a tuple with the Dates field value
// and a boolean to check if the value has been set.
func (o *GetUsersCountResponse) GetDatesOk() ([]DailyUsers, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dates, true
}

// SetDates sets field value.
func (o *GetUsersCountResponse) SetDates(v []DailyUsers) *GetUsersCountResponse {
	o.Dates = v
	return o
}

func (o GetUsersCountResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	toSerialize["count"] = o.Count
	toSerialize["dates"] = o.Dates
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal GetUsersCountResponse: %w", err)
	}

	return serialized, nil
}

func (o GetUsersCountResponse) String() string {
	out := ""
	out += fmt.Sprintf("  count=%v\n", o.Count)
	out += fmt.Sprintf("  dates=%v\n", o.Dates)
	return fmt.Sprintf("GetUsersCountResponse {\n%s}", out)
}
