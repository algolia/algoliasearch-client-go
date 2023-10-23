// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package recommend

import (
	"encoding/json"
	"fmt"
)

// GetRecommendationsParams Recommend parameters.
type GetRecommendationsParams struct {
	// Request parameters depend on the model (recommendations or trending).
	Requests []RecommendationsRequest `json:"requests"`
}

// NewGetRecommendationsParams instantiates a new GetRecommendationsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRecommendationsParams(requests []RecommendationsRequest) *GetRecommendationsParams {
	this := &GetRecommendationsParams{}
	this.Requests = requests
	return this
}

// NewGetRecommendationsParamsWithDefaults instantiates a new GetRecommendationsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRecommendationsParamsWithDefaults() *GetRecommendationsParams {
	this := &GetRecommendationsParams{}
	return this
}

// GetRequests returns the Requests field value
func (o *GetRecommendationsParams) GetRequests() []RecommendationsRequest {
	if o == nil {
		var ret []RecommendationsRequest
		return ret
	}

	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *GetRecommendationsParams) GetRequestsOk() ([]RecommendationsRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Requests, true
}

// SetRequests sets field value
func (o *GetRecommendationsParams) SetRequests(v []RecommendationsRequest) {
	o.Requests = v
}

func (o GetRecommendationsParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["requests"] = o.Requests
	}
	return json.Marshal(toSerialize)
}

func (o GetRecommendationsParams) String() string {
	out := ""
	out += fmt.Sprintf("  requests=%v\n", o.Requests)
	return fmt.Sprintf("GetRecommendationsParams {\n%s}", out)
}

type NullableGetRecommendationsParams struct {
	value *GetRecommendationsParams
	isSet bool
}

func (v NullableGetRecommendationsParams) Get() *GetRecommendationsParams {
	return v.value
}

func (v *NullableGetRecommendationsParams) Set(val *GetRecommendationsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRecommendationsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRecommendationsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRecommendationsParams(val *GetRecommendationsParams) *NullableGetRecommendationsParams {
	return &NullableGetRecommendationsParams{value: val, isSet: true}
}

func (v NullableGetRecommendationsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRecommendationsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
