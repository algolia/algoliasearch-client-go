// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package personalization

import (
	"encoding/json"
	"fmt"
)

// PersonalizationStrategyParams struct for PersonalizationStrategyParams.
type PersonalizationStrategyParams struct {
	// Scores associated with the events.
	EventScoring []EventScoring `json:"eventScoring"`
	// Scores associated with the facets.
	FacetScoring []FacetScoring `json:"facetScoring"`
	// The impact that personalization has on search results: a number between 0 (personalization disabled) and 100 (personalization fully enabled).
	PersonalizationImpact int32 `json:"personalizationImpact"`
}

// NewPersonalizationStrategyParams instantiates a new PersonalizationStrategyParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPersonalizationStrategyParams(eventScoring []EventScoring, facetScoring []FacetScoring, personalizationImpact int32) *PersonalizationStrategyParams {
	this := &PersonalizationStrategyParams{}
	this.EventScoring = eventScoring
	this.FacetScoring = facetScoring
	this.PersonalizationImpact = personalizationImpact
	return this
}

// NewPersonalizationStrategyParamsWithDefaults instantiates a new PersonalizationStrategyParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPersonalizationStrategyParamsWithDefaults() *PersonalizationStrategyParams {
	this := &PersonalizationStrategyParams{}
	return this
}

// GetEventScoring returns the EventScoring field value.
func (o *PersonalizationStrategyParams) GetEventScoring() []EventScoring {
	if o == nil {
		var ret []EventScoring
		return ret
	}

	return o.EventScoring
}

// GetEventScoringOk returns a tuple with the EventScoring field value
// and a boolean to check if the value has been set.
func (o *PersonalizationStrategyParams) GetEventScoringOk() ([]EventScoring, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventScoring, true
}

// SetEventScoring sets field value.
func (o *PersonalizationStrategyParams) SetEventScoring(v []EventScoring) {
	o.EventScoring = v
}

// GetFacetScoring returns the FacetScoring field value.
func (o *PersonalizationStrategyParams) GetFacetScoring() []FacetScoring {
	if o == nil {
		var ret []FacetScoring
		return ret
	}

	return o.FacetScoring
}

// GetFacetScoringOk returns a tuple with the FacetScoring field value
// and a boolean to check if the value has been set.
func (o *PersonalizationStrategyParams) GetFacetScoringOk() ([]FacetScoring, bool) {
	if o == nil {
		return nil, false
	}
	return o.FacetScoring, true
}

// SetFacetScoring sets field value.
func (o *PersonalizationStrategyParams) SetFacetScoring(v []FacetScoring) {
	o.FacetScoring = v
}

// GetPersonalizationImpact returns the PersonalizationImpact field value.
func (o *PersonalizationStrategyParams) GetPersonalizationImpact() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PersonalizationImpact
}

// GetPersonalizationImpactOk returns a tuple with the PersonalizationImpact field value
// and a boolean to check if the value has been set.
func (o *PersonalizationStrategyParams) GetPersonalizationImpactOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PersonalizationImpact, true
}

// SetPersonalizationImpact sets field value.
func (o *PersonalizationStrategyParams) SetPersonalizationImpact(v int32) {
	o.PersonalizationImpact = v
}

func (o PersonalizationStrategyParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["eventScoring"] = o.EventScoring
	}
	if true {
		toSerialize["facetScoring"] = o.FacetScoring
	}
	if true {
		toSerialize["personalizationImpact"] = o.PersonalizationImpact
	}
	return json.Marshal(toSerialize)
}

func (o PersonalizationStrategyParams) String() string {
	out := ""
	out += fmt.Sprintf("  eventScoring=%v\n", o.EventScoring)
	out += fmt.Sprintf("  facetScoring=%v\n", o.FacetScoring)
	out += fmt.Sprintf("  personalizationImpact=%v\n", o.PersonalizationImpact)
	return fmt.Sprintf("PersonalizationStrategyParams {\n%s}", out)
}

type NullablePersonalizationStrategyParams struct {
	value *PersonalizationStrategyParams
	isSet bool
}

func (v NullablePersonalizationStrategyParams) Get() *PersonalizationStrategyParams {
	return v.value
}

func (v *NullablePersonalizationStrategyParams) Set(val *PersonalizationStrategyParams) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonalizationStrategyParams) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonalizationStrategyParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonalizationStrategyParams(val *PersonalizationStrategyParams) *NullablePersonalizationStrategyParams {
	return &NullablePersonalizationStrategyParams{value: val, isSet: true}
}

func (v NullablePersonalizationStrategyParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonalizationStrategyParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
