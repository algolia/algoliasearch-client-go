// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package recommend

import (
	"encoding/json"
	"fmt"
)

// TrendingItems struct for TrendingItems.
type TrendingItems struct {
	// Facet attribute. To be used in combination with `facetValue`. If specified, only recommendations matching the facet filter will be returned.
	FacetName string `json:"facetName"`
	// Facet value. To be used in combination with `facetName`. If specified, only recommendations matching the facet filter will be returned.
	FacetValue         string              `json:"facetValue"`
	Model              TrendingItemsModel  `json:"model"`
	FallbackParameters *SearchParamsObject `json:"fallbackParameters,omitempty"`
}

type TrendingItemsOption func(f *TrendingItems)

func WithTrendingItemsFallbackParameters(val SearchParamsObject) TrendingItemsOption {
	return func(f *TrendingItems) {
		f.FallbackParameters = &val
	}
}

// NewTrendingItems instantiates a new TrendingItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTrendingItems(facetName string, facetValue string, model TrendingItemsModel, opts ...TrendingItemsOption) *TrendingItems {
	this := &TrendingItems{}
	this.FacetName = facetName
	this.FacetValue = facetValue
	this.Model = model
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyTrendingItems return a pointer to an empty TrendingItems object.
func NewEmptyTrendingItems() *TrendingItems {
	return &TrendingItems{}
}

// GetFacetName returns the FacetName field value.
func (o *TrendingItems) GetFacetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FacetName
}

// GetFacetNameOk returns a tuple with the FacetName field value
// and a boolean to check if the value has been set.
func (o *TrendingItems) GetFacetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FacetName, true
}

// SetFacetName sets field value.
func (o *TrendingItems) SetFacetName(v string) *TrendingItems {
	o.FacetName = v
	return o
}

// GetFacetValue returns the FacetValue field value.
func (o *TrendingItems) GetFacetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FacetValue
}

// GetFacetValueOk returns a tuple with the FacetValue field value
// and a boolean to check if the value has been set.
func (o *TrendingItems) GetFacetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FacetValue, true
}

// SetFacetValue sets field value.
func (o *TrendingItems) SetFacetValue(v string) *TrendingItems {
	o.FacetValue = v
	return o
}

// GetModel returns the Model field value.
func (o *TrendingItems) GetModel() TrendingItemsModel {
	if o == nil {
		var ret TrendingItemsModel
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *TrendingItems) GetModelOk() (*TrendingItemsModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value.
func (o *TrendingItems) SetModel(v TrendingItemsModel) *TrendingItems {
	o.Model = v
	return o
}

// GetFallbackParameters returns the FallbackParameters field value if set, zero value otherwise.
func (o *TrendingItems) GetFallbackParameters() SearchParamsObject {
	if o == nil || o.FallbackParameters == nil {
		var ret SearchParamsObject
		return ret
	}
	return *o.FallbackParameters
}

// GetFallbackParametersOk returns a tuple with the FallbackParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrendingItems) GetFallbackParametersOk() (*SearchParamsObject, bool) {
	if o == nil || o.FallbackParameters == nil {
		return nil, false
	}
	return o.FallbackParameters, true
}

// HasFallbackParameters returns a boolean if a field has been set.
func (o *TrendingItems) HasFallbackParameters() bool {
	if o != nil && o.FallbackParameters != nil {
		return true
	}

	return false
}

// SetFallbackParameters gets a reference to the given SearchParamsObject and assigns it to the FallbackParameters field.
func (o *TrendingItems) SetFallbackParameters(v *SearchParamsObject) *TrendingItems {
	o.FallbackParameters = v
	return o
}

func (o TrendingItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["facetName"] = o.FacetName
	}
	if true {
		toSerialize["facetValue"] = o.FacetValue
	}
	if true {
		toSerialize["model"] = o.Model
	}
	if o.FallbackParameters != nil {
		toSerialize["fallbackParameters"] = o.FallbackParameters
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal TrendingItems: %w", err)
	}

	return serialized, nil
}

func (o TrendingItems) String() string {
	out := ""
	out += fmt.Sprintf("  facetName=%v\n", o.FacetName)
	out += fmt.Sprintf("  facetValue=%v\n", o.FacetValue)
	out += fmt.Sprintf("  model=%v\n", o.Model)
	out += fmt.Sprintf("  fallbackParameters=%v\n", o.FallbackParameters)
	return fmt.Sprintf("TrendingItems {\n%s}", out)
}
