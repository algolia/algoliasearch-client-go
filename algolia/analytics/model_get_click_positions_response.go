// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package analytics

import (
	"encoding/json"
	"fmt"
)

// GetClickPositionsResponse struct for GetClickPositionsResponse
type GetClickPositionsResponse struct {
	// Click positions.
	Positions []ClickPosition `json:"positions"`
}

// NewGetClickPositionsResponse instantiates a new GetClickPositionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetClickPositionsResponse(positions []ClickPosition) *GetClickPositionsResponse {
	this := &GetClickPositionsResponse{}
	this.Positions = positions
	return this
}

// NewGetClickPositionsResponseWithDefaults instantiates a new GetClickPositionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetClickPositionsResponseWithDefaults() *GetClickPositionsResponse {
	this := &GetClickPositionsResponse{}
	return this
}

// GetPositions returns the Positions field value
func (o *GetClickPositionsResponse) GetPositions() []ClickPosition {
	if o == nil {
		var ret []ClickPosition
		return ret
	}

	return o.Positions
}

// GetPositionsOk returns a tuple with the Positions field value
// and a boolean to check if the value has been set.
func (o *GetClickPositionsResponse) GetPositionsOk() ([]ClickPosition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Positions, true
}

// SetPositions sets field value
func (o *GetClickPositionsResponse) SetPositions(v []ClickPosition) {
	o.Positions = v
}

func (o GetClickPositionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["positions"] = o.Positions
	}
	return json.Marshal(toSerialize)
}

func (o GetClickPositionsResponse) String() string {
	out := ""
	out += fmt.Sprintf("  positions=%v\n", o.Positions)
	return fmt.Sprintf("GetClickPositionsResponse {\n%s}", out)
}

type NullableGetClickPositionsResponse struct {
	value *GetClickPositionsResponse
	isSet bool
}

func (v NullableGetClickPositionsResponse) Get() *GetClickPositionsResponse {
	return v.value
}

func (v *NullableGetClickPositionsResponse) Set(val *GetClickPositionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetClickPositionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetClickPositionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetClickPositionsResponse(val *GetClickPositionsResponse) *NullableGetClickPositionsResponse {
	return &NullableGetClickPositionsResponse{value: val, isSet: true}
}

func (v NullableGetClickPositionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetClickPositionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
