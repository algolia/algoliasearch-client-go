// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package recommend

import (
	"encoding/json"
	"fmt"
)

// AutomaticFacetFilter Filter or optional filter to be applied to the search.
type AutomaticFacetFilter struct {
	// Facet name to be applied as filter. The name must match placeholders in the `pattern` parameter. For example, with `pattern: {facet:genre}`, `automaticFacetFilters` must be `genre`.
	Facet string `json:"facet"`
	// Filter scores to give different weights to individual filters.
	Score *int32 `json:"score,omitempty"`
	// Whether the filter is disjunctive or conjunctive.  If true the filter has multiple matches, multiple occurrences are combined with the logical `OR` operation. If false, multiple occurrences are combined with the logical `AND` operation.
	Disjunctive *bool `json:"disjunctive,omitempty"`
}

type AutomaticFacetFilterOption func(f *AutomaticFacetFilter)

func WithAutomaticFacetFilterScore(val int32) AutomaticFacetFilterOption {
	return func(f *AutomaticFacetFilter) {
		f.Score = &val
	}
}

func WithAutomaticFacetFilterDisjunctive(val bool) AutomaticFacetFilterOption {
	return func(f *AutomaticFacetFilter) {
		f.Disjunctive = &val
	}
}

// NewAutomaticFacetFilter instantiates a new AutomaticFacetFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAutomaticFacetFilter(facet string, opts ...AutomaticFacetFilterOption) *AutomaticFacetFilter {
	this := &AutomaticFacetFilter{}
	this.Facet = facet
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyAutomaticFacetFilter return a pointer to an empty AutomaticFacetFilter object.
func NewEmptyAutomaticFacetFilter() *AutomaticFacetFilter {
	return &AutomaticFacetFilter{}
}

// GetFacet returns the Facet field value.
func (o *AutomaticFacetFilter) GetFacet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Facet
}

// GetFacetOk returns a tuple with the Facet field value
// and a boolean to check if the value has been set.
func (o *AutomaticFacetFilter) GetFacetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Facet, true
}

// SetFacet sets field value.
func (o *AutomaticFacetFilter) SetFacet(v string) *AutomaticFacetFilter {
	o.Facet = v
	return o
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *AutomaticFacetFilter) GetScore() int32 {
	if o == nil || o.Score == nil {
		var ret int32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomaticFacetFilter) GetScoreOk() (*int32, bool) {
	if o == nil || o.Score == nil {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *AutomaticFacetFilter) HasScore() bool {
	if o != nil && o.Score != nil {
		return true
	}

	return false
}

// SetScore gets a reference to the given int32 and assigns it to the Score field.
func (o *AutomaticFacetFilter) SetScore(v int32) *AutomaticFacetFilter {
	o.Score = &v
	return o
}

// GetDisjunctive returns the Disjunctive field value if set, zero value otherwise.
func (o *AutomaticFacetFilter) GetDisjunctive() bool {
	if o == nil || o.Disjunctive == nil {
		var ret bool
		return ret
	}
	return *o.Disjunctive
}

// GetDisjunctiveOk returns a tuple with the Disjunctive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomaticFacetFilter) GetDisjunctiveOk() (*bool, bool) {
	if o == nil || o.Disjunctive == nil {
		return nil, false
	}
	return o.Disjunctive, true
}

// HasDisjunctive returns a boolean if a field has been set.
func (o *AutomaticFacetFilter) HasDisjunctive() bool {
	if o != nil && o.Disjunctive != nil {
		return true
	}

	return false
}

// SetDisjunctive gets a reference to the given bool and assigns it to the Disjunctive field.
func (o *AutomaticFacetFilter) SetDisjunctive(v bool) *AutomaticFacetFilter {
	o.Disjunctive = &v
	return o
}

func (o AutomaticFacetFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["facet"] = o.Facet
	}
	if o.Score != nil {
		toSerialize["score"] = o.Score
	}
	if o.Disjunctive != nil {
		toSerialize["disjunctive"] = o.Disjunctive
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal AutomaticFacetFilter: %w", err)
	}

	return serialized, nil
}

func (o AutomaticFacetFilter) String() string {
	out := ""
	out += fmt.Sprintf("  facet=%v\n", o.Facet)
	out += fmt.Sprintf("  score=%v\n", o.Score)
	out += fmt.Sprintf("  disjunctive=%v\n", o.Disjunctive)
	return fmt.Sprintf("AutomaticFacetFilter {\n%s}", out)
}

type NullableAutomaticFacetFilter struct {
	value *AutomaticFacetFilter
	isSet bool
}

func (v NullableAutomaticFacetFilter) Get() *AutomaticFacetFilter {
	return v.value
}

func (v *NullableAutomaticFacetFilter) Set(val *AutomaticFacetFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableAutomaticFacetFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableAutomaticFacetFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutomaticFacetFilter(val *AutomaticFacetFilter) *NullableAutomaticFacetFilter {
	return &NullableAutomaticFacetFilter{value: val, isSet: true}
}

func (v NullableAutomaticFacetFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableAutomaticFacetFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}