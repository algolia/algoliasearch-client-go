// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package recommend

import (
	"encoding/json"
	"fmt"
)

// TrendingFacets struct for TrendingFacets.
type TrendingFacets struct {
	// Facet attribute for which to retrieve trending facet values.
	FacetName          any                 `json:"facetName"`
	Model              TrendingFacetsModel `json:"model"`
	FallbackParameters *FallbackParams     `json:"fallbackParameters,omitempty"`
}

type TrendingFacetsOption func(f *TrendingFacets)

func WithTrendingFacetsFallbackParameters(val FallbackParams) TrendingFacetsOption {
	return func(f *TrendingFacets) {
		f.FallbackParameters = &val
	}
}

// NewTrendingFacets instantiates a new TrendingFacets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTrendingFacets(facetName any, model TrendingFacetsModel, opts ...TrendingFacetsOption) *TrendingFacets {
	this := &TrendingFacets{}
	this.FacetName = facetName
	this.Model = model
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyTrendingFacets return a pointer to an empty TrendingFacets object.
func NewEmptyTrendingFacets() *TrendingFacets {
	return &TrendingFacets{}
}

// GetFacetName returns the FacetName field value.
// If the value is explicit nil, the zero value for any will be returned.
func (o *TrendingFacets) GetFacetName() any {
	if o == nil {
		var ret any
		return ret
	}

	return o.FacetName
}

// GetFacetNameOk returns a tuple with the FacetName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TrendingFacets) GetFacetNameOk() (*any, bool) {
	if o == nil || o.FacetName == nil {
		return nil, false
	}
	return &o.FacetName, true
}

// SetFacetName sets field value.
func (o *TrendingFacets) SetFacetName(v any) *TrendingFacets {
	o.FacetName = v
	return o
}

// GetModel returns the Model field value.
func (o *TrendingFacets) GetModel() TrendingFacetsModel {
	if o == nil {
		var ret TrendingFacetsModel
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *TrendingFacets) GetModelOk() (*TrendingFacetsModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value.
func (o *TrendingFacets) SetModel(v TrendingFacetsModel) *TrendingFacets {
	o.Model = v
	return o
}

// GetFallbackParameters returns the FallbackParameters field value if set, zero value otherwise.
func (o *TrendingFacets) GetFallbackParameters() FallbackParams {
	if o == nil || o.FallbackParameters == nil {
		var ret FallbackParams
		return ret
	}
	return *o.FallbackParameters
}

// GetFallbackParametersOk returns a tuple with the FallbackParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrendingFacets) GetFallbackParametersOk() (*FallbackParams, bool) {
	if o == nil || o.FallbackParameters == nil {
		return nil, false
	}
	return o.FallbackParameters, true
}

// HasFallbackParameters returns a boolean if a field has been set.
func (o *TrendingFacets) HasFallbackParameters() bool {
	if o != nil && o.FallbackParameters != nil {
		return true
	}

	return false
}

// SetFallbackParameters gets a reference to the given FallbackParams and assigns it to the FallbackParameters field.
func (o *TrendingFacets) SetFallbackParameters(v *FallbackParams) *TrendingFacets {
	o.FallbackParameters = v
	return o
}

func (o TrendingFacets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.FacetName != nil {
		toSerialize["facetName"] = o.FacetName
	}
	if true {
		toSerialize["model"] = o.Model
	}
	if o.FallbackParameters != nil {
		toSerialize["fallbackParameters"] = o.FallbackParameters
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal TrendingFacets: %w", err)
	}

	return serialized, nil
}

func (o TrendingFacets) String() string {
	out := ""
	out += fmt.Sprintf("  facetName=%v\n", o.FacetName)
	out += fmt.Sprintf("  model=%v\n", o.Model)
	out += fmt.Sprintf("  fallbackParameters=%v\n", o.FallbackParameters)
	return fmt.Sprintf("TrendingFacets {\n%s}", out)
}