// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package composition

import (
	"encoding/json"
	"fmt"
)

// SearchForFacetValuesParams struct for SearchForFacetValuesParams.
type SearchForFacetValuesParams struct {
	// Search query.
	Query *string `json:"query,omitempty"`
	// Maximum number of facet values to return when [searching for facet values](https://www.algolia.com/doc/guides/managing-results/refine-results/faceting/#search-for-facet-values).
	MaxFacetHits *int32  `json:"maxFacetHits,omitempty"`
	SearchQuery  *Params `json:"searchQuery,omitempty"`
}

type SearchForFacetValuesParamsOption func(f *SearchForFacetValuesParams)

func WithSearchForFacetValuesParamsQuery(val string) SearchForFacetValuesParamsOption {
	return func(f *SearchForFacetValuesParams) {
		f.Query = &val
	}
}

func WithSearchForFacetValuesParamsMaxFacetHits(val int32) SearchForFacetValuesParamsOption {
	return func(f *SearchForFacetValuesParams) {
		f.MaxFacetHits = &val
	}
}

func WithSearchForFacetValuesParamsSearchQuery(val Params) SearchForFacetValuesParamsOption {
	return func(f *SearchForFacetValuesParams) {
		f.SearchQuery = &val
	}
}

// NewSearchForFacetValuesParams instantiates a new SearchForFacetValuesParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSearchForFacetValuesParams(opts ...SearchForFacetValuesParamsOption) *SearchForFacetValuesParams {
	this := &SearchForFacetValuesParams{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptySearchForFacetValuesParams return a pointer to an empty SearchForFacetValuesParams object.
func NewEmptySearchForFacetValuesParams() *SearchForFacetValuesParams {
	return &SearchForFacetValuesParams{}
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *SearchForFacetValuesParams) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchForFacetValuesParams) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *SearchForFacetValuesParams) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *SearchForFacetValuesParams) SetQuery(v string) *SearchForFacetValuesParams {
	o.Query = &v
	return o
}

// GetMaxFacetHits returns the MaxFacetHits field value if set, zero value otherwise.
func (o *SearchForFacetValuesParams) GetMaxFacetHits() int32 {
	if o == nil || o.MaxFacetHits == nil {
		var ret int32
		return ret
	}
	return *o.MaxFacetHits
}

// GetMaxFacetHitsOk returns a tuple with the MaxFacetHits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchForFacetValuesParams) GetMaxFacetHitsOk() (*int32, bool) {
	if o == nil || o.MaxFacetHits == nil {
		return nil, false
	}
	return o.MaxFacetHits, true
}

// HasMaxFacetHits returns a boolean if a field has been set.
func (o *SearchForFacetValuesParams) HasMaxFacetHits() bool {
	if o != nil && o.MaxFacetHits != nil {
		return true
	}

	return false
}

// SetMaxFacetHits gets a reference to the given int32 and assigns it to the MaxFacetHits field.
func (o *SearchForFacetValuesParams) SetMaxFacetHits(v int32) *SearchForFacetValuesParams {
	o.MaxFacetHits = &v
	return o
}

// GetSearchQuery returns the SearchQuery field value if set, zero value otherwise.
func (o *SearchForFacetValuesParams) GetSearchQuery() Params {
	if o == nil || o.SearchQuery == nil {
		var ret Params
		return ret
	}
	return *o.SearchQuery
}

// GetSearchQueryOk returns a tuple with the SearchQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchForFacetValuesParams) GetSearchQueryOk() (*Params, bool) {
	if o == nil || o.SearchQuery == nil {
		return nil, false
	}
	return o.SearchQuery, true
}

// HasSearchQuery returns a boolean if a field has been set.
func (o *SearchForFacetValuesParams) HasSearchQuery() bool {
	if o != nil && o.SearchQuery != nil {
		return true
	}

	return false
}

// SetSearchQuery gets a reference to the given Params and assigns it to the SearchQuery field.
func (o *SearchForFacetValuesParams) SetSearchQuery(v *Params) *SearchForFacetValuesParams {
	o.SearchQuery = v
	return o
}

func (o SearchForFacetValuesParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.MaxFacetHits != nil {
		toSerialize["maxFacetHits"] = o.MaxFacetHits
	}
	if o.SearchQuery != nil {
		toSerialize["searchQuery"] = o.SearchQuery
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal SearchForFacetValuesParams: %w", err)
	}

	return serialized, nil
}

func (o SearchForFacetValuesParams) String() string {
	out := ""
	out += fmt.Sprintf("  query=%v\n", o.Query)
	out += fmt.Sprintf("  maxFacetHits=%v\n", o.MaxFacetHits)
	out += fmt.Sprintf("  searchQuery=%v\n", o.SearchQuery)
	return fmt.Sprintf("SearchForFacetValuesParams {\n%s}", out)
}
