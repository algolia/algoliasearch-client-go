// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package recommend

import (
	"encoding/json"
	"fmt"
)

// BaseTrendingRequest struct for BaseTrendingRequest
type BaseTrendingRequest struct {
	Model TrendingModels `json:"model"`
	// The facet name to use for trending models.
	FacetName *string `json:"facetName,omitempty"`
	// The facet value to use for trending models.
	FacetValue *string `json:"facetValue,omitempty"`
}

type BaseTrendingRequestOption func(f *BaseTrendingRequest)

func WithBaseTrendingRequestFacetName(val string) BaseTrendingRequestOption {
	return func(f *BaseTrendingRequest) {
		f.FacetName = &val
	}
}

func WithBaseTrendingRequestFacetValue(val string) BaseTrendingRequestOption {
	return func(f *BaseTrendingRequest) {
		f.FacetValue = &val
	}
}

// NewBaseTrendingRequest instantiates a new BaseTrendingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseTrendingRequest(model TrendingModels, opts ...BaseTrendingRequestOption) *BaseTrendingRequest {
	this := &BaseTrendingRequest{}
	this.Model = model
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewBaseTrendingRequestWithDefaults instantiates a new BaseTrendingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseTrendingRequestWithDefaults() *BaseTrendingRequest {
	this := &BaseTrendingRequest{}
	return this
}

// GetModel returns the Model field value
func (o *BaseTrendingRequest) GetModel() TrendingModels {
	if o == nil {
		var ret TrendingModels
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *BaseTrendingRequest) GetModelOk() (*TrendingModels, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *BaseTrendingRequest) SetModel(v TrendingModels) {
	o.Model = v
}

// GetFacetName returns the FacetName field value if set, zero value otherwise.
func (o *BaseTrendingRequest) GetFacetName() string {
	if o == nil || o.FacetName == nil {
		var ret string
		return ret
	}
	return *o.FacetName
}

// GetFacetNameOk returns a tuple with the FacetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTrendingRequest) GetFacetNameOk() (*string, bool) {
	if o == nil || o.FacetName == nil {
		return nil, false
	}
	return o.FacetName, true
}

// HasFacetName returns a boolean if a field has been set.
func (o *BaseTrendingRequest) HasFacetName() bool {
	if o != nil && o.FacetName != nil {
		return true
	}

	return false
}

// SetFacetName gets a reference to the given string and assigns it to the FacetName field.
func (o *BaseTrendingRequest) SetFacetName(v string) {
	o.FacetName = &v
}

// GetFacetValue returns the FacetValue field value if set, zero value otherwise.
func (o *BaseTrendingRequest) GetFacetValue() string {
	if o == nil || o.FacetValue == nil {
		var ret string
		return ret
	}
	return *o.FacetValue
}

// GetFacetValueOk returns a tuple with the FacetValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTrendingRequest) GetFacetValueOk() (*string, bool) {
	if o == nil || o.FacetValue == nil {
		return nil, false
	}
	return o.FacetValue, true
}

// HasFacetValue returns a boolean if a field has been set.
func (o *BaseTrendingRequest) HasFacetValue() bool {
	if o != nil && o.FacetValue != nil {
		return true
	}

	return false
}

// SetFacetValue gets a reference to the given string and assigns it to the FacetValue field.
func (o *BaseTrendingRequest) SetFacetValue(v string) {
	o.FacetValue = &v
}

func (o BaseTrendingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["model"] = o.Model
	}
	if o.FacetName != nil {
		toSerialize["facetName"] = o.FacetName
	}
	if o.FacetValue != nil {
		toSerialize["facetValue"] = o.FacetValue
	}
	return json.Marshal(toSerialize)
}

func (o BaseTrendingRequest) String() string {
	out := ""
	out += fmt.Sprintf("  model=%v\n", o.Model)
	out += fmt.Sprintf("  facetName=%v\n", o.FacetName)
	out += fmt.Sprintf("  facetValue=%v\n", o.FacetValue)
	return fmt.Sprintf("BaseTrendingRequest {\n%s}", out)
}

type NullableBaseTrendingRequest struct {
	value *BaseTrendingRequest
	isSet bool
}

func (v NullableBaseTrendingRequest) Get() *BaseTrendingRequest {
	return v.value
}

func (v *NullableBaseTrendingRequest) Set(val *BaseTrendingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseTrendingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseTrendingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseTrendingRequest(val *BaseTrendingRequest) *NullableBaseTrendingRequest {
	return &NullableBaseTrendingRequest{value: val, isSet: true}
}

func (v NullableBaseTrendingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseTrendingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
