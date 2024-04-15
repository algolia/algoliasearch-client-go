// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package search

import (
	"encoding/json"
	"fmt"
)

// SearchForFacetsOptions struct for SearchForFacetsOptions.
type SearchForFacetsOptions struct {
	// Facet name.
	Facet string `json:"facet"`
	// Index name (case-sensitive).
	IndexName string `json:"indexName"`
	// Text to search inside the facet's values.
	FacetQuery *string `json:"facetQuery,omitempty"`
	// Maximum number of facet values to return when [searching for facet values](https://www.algolia.com/doc/guides/managing-results/refine-results/faceting/#search-for-facet-values).
	MaxFacetHits *int32          `json:"maxFacetHits,omitempty"`
	Type         SearchTypeFacet `json:"type"`
}

type SearchForFacetsOptionsOption func(f *SearchForFacetsOptions)

func WithSearchForFacetsOptionsFacetQuery(val string) SearchForFacetsOptionsOption {
	return func(f *SearchForFacetsOptions) {
		f.FacetQuery = &val
	}
}

func WithSearchForFacetsOptionsMaxFacetHits(val int32) SearchForFacetsOptionsOption {
	return func(f *SearchForFacetsOptions) {
		f.MaxFacetHits = &val
	}
}

// NewSearchForFacetsOptions instantiates a new SearchForFacetsOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSearchForFacetsOptions(facet string, indexName string, type_ SearchTypeFacet, opts ...SearchForFacetsOptionsOption) *SearchForFacetsOptions {
	this := &SearchForFacetsOptions{}
	this.Facet = facet
	this.IndexName = indexName
	this.Type = type_
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptySearchForFacetsOptions return a pointer to an empty SearchForFacetsOptions object.
func NewEmptySearchForFacetsOptions() *SearchForFacetsOptions {
	return &SearchForFacetsOptions{}
}

// GetFacet returns the Facet field value.
func (o *SearchForFacetsOptions) GetFacet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Facet
}

// GetFacetOk returns a tuple with the Facet field value
// and a boolean to check if the value has been set.
func (o *SearchForFacetsOptions) GetFacetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Facet, true
}

// SetFacet sets field value.
func (o *SearchForFacetsOptions) SetFacet(v string) *SearchForFacetsOptions {
	o.Facet = v
	return o
}

// GetIndexName returns the IndexName field value.
func (o *SearchForFacetsOptions) GetIndexName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IndexName
}

// GetIndexNameOk returns a tuple with the IndexName field value
// and a boolean to check if the value has been set.
func (o *SearchForFacetsOptions) GetIndexNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndexName, true
}

// SetIndexName sets field value.
func (o *SearchForFacetsOptions) SetIndexName(v string) *SearchForFacetsOptions {
	o.IndexName = v
	return o
}

// GetFacetQuery returns the FacetQuery field value if set, zero value otherwise.
func (o *SearchForFacetsOptions) GetFacetQuery() string {
	if o == nil || o.FacetQuery == nil {
		var ret string
		return ret
	}
	return *o.FacetQuery
}

// GetFacetQueryOk returns a tuple with the FacetQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchForFacetsOptions) GetFacetQueryOk() (*string, bool) {
	if o == nil || o.FacetQuery == nil {
		return nil, false
	}
	return o.FacetQuery, true
}

// HasFacetQuery returns a boolean if a field has been set.
func (o *SearchForFacetsOptions) HasFacetQuery() bool {
	if o != nil && o.FacetQuery != nil {
		return true
	}

	return false
}

// SetFacetQuery gets a reference to the given string and assigns it to the FacetQuery field.
func (o *SearchForFacetsOptions) SetFacetQuery(v string) *SearchForFacetsOptions {
	o.FacetQuery = &v
	return o
}

// GetMaxFacetHits returns the MaxFacetHits field value if set, zero value otherwise.
func (o *SearchForFacetsOptions) GetMaxFacetHits() int32 {
	if o == nil || o.MaxFacetHits == nil {
		var ret int32
		return ret
	}
	return *o.MaxFacetHits
}

// GetMaxFacetHitsOk returns a tuple with the MaxFacetHits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchForFacetsOptions) GetMaxFacetHitsOk() (*int32, bool) {
	if o == nil || o.MaxFacetHits == nil {
		return nil, false
	}
	return o.MaxFacetHits, true
}

// HasMaxFacetHits returns a boolean if a field has been set.
func (o *SearchForFacetsOptions) HasMaxFacetHits() bool {
	if o != nil && o.MaxFacetHits != nil {
		return true
	}

	return false
}

// SetMaxFacetHits gets a reference to the given int32 and assigns it to the MaxFacetHits field.
func (o *SearchForFacetsOptions) SetMaxFacetHits(v int32) *SearchForFacetsOptions {
	o.MaxFacetHits = &v
	return o
}

// GetType returns the Type field value.
func (o *SearchForFacetsOptions) GetType() SearchTypeFacet {
	if o == nil {
		var ret SearchTypeFacet
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SearchForFacetsOptions) GetTypeOk() (*SearchTypeFacet, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SearchForFacetsOptions) SetType(v SearchTypeFacet) *SearchForFacetsOptions {
	o.Type = v
	return o
}

func (o SearchForFacetsOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["facet"] = o.Facet
	}
	if true {
		toSerialize["indexName"] = o.IndexName
	}
	if o.FacetQuery != nil {
		toSerialize["facetQuery"] = o.FacetQuery
	}
	if o.MaxFacetHits != nil {
		toSerialize["maxFacetHits"] = o.MaxFacetHits
	}
	if true {
		toSerialize["type"] = o.Type
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal SearchForFacetsOptions: %w", err)
	}

	return serialized, nil
}

func (o SearchForFacetsOptions) String() string {
	out := ""
	out += fmt.Sprintf("  facet=%v\n", o.Facet)
	out += fmt.Sprintf("  indexName=%v\n", o.IndexName)
	out += fmt.Sprintf("  facetQuery=%v\n", o.FacetQuery)
	out += fmt.Sprintf("  maxFacetHits=%v\n", o.MaxFacetHits)
	out += fmt.Sprintf("  type=%v\n", o.Type)
	return fmt.Sprintf("SearchForFacetsOptions {\n%s}", out)
}

type NullableSearchForFacetsOptions struct {
	value *SearchForFacetsOptions
	isSet bool
}

func (v NullableSearchForFacetsOptions) Get() *SearchForFacetsOptions {
	return v.value
}

func (v *NullableSearchForFacetsOptions) Set(val *SearchForFacetsOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchForFacetsOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchForFacetsOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchForFacetsOptions(val *SearchForFacetsOptions) *NullableSearchForFacetsOptions {
	return &NullableSearchForFacetsOptions{value: val, isSet: true}
}

func (v NullableSearchForFacetsOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableSearchForFacetsOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
