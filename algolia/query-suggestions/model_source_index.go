// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package suggestions

import (
	"encoding/json"
	"fmt"
)

// SourceIndex Configuration of an Algolia index for Query Suggestions.
type SourceIndex struct {
	// Name of the Algolia index (case-sensitive) to use as source for query suggestions.
	IndexName string `json:"indexName"`
	// If true, Query Suggestions uses all replica indices to find popular searches. If false, only the primary index is used.
	Replicas      *bool    `json:"replicas,omitempty"`
	AnalyticsTags []string `json:"analyticsTags,omitempty"`
	Facets        []Facet  `json:"facets,omitempty"`
	// Minimum number of hits required to be included as a suggestion.  A search query must at least generate `minHits` search results to be included in the Query Suggestions index.
	MinHits *int32 `json:"minHits,omitempty"`
	// Minimum letters required to be included as a suggestion.  A search query must be at least `minLetters` long to be included in the Query Suggestions index.
	MinLetters *int32     `json:"minLetters,omitempty"`
	Generate   [][]string `json:"generate,omitempty"`
	External   []string   `json:"external,omitempty"`
}

type SourceIndexOption func(f *SourceIndex)

func WithSourceIndexReplicas(val bool) SourceIndexOption {
	return func(f *SourceIndex) {
		f.Replicas = &val
	}
}

func WithSourceIndexAnalyticsTags(val []string) SourceIndexOption {
	return func(f *SourceIndex) {
		f.AnalyticsTags = val
	}
}

func WithSourceIndexFacets(val []Facet) SourceIndexOption {
	return func(f *SourceIndex) {
		f.Facets = val
	}
}

func WithSourceIndexMinHits(val int32) SourceIndexOption {
	return func(f *SourceIndex) {
		f.MinHits = &val
	}
}

func WithSourceIndexMinLetters(val int32) SourceIndexOption {
	return func(f *SourceIndex) {
		f.MinLetters = &val
	}
}

func WithSourceIndexGenerate(val [][]string) SourceIndexOption {
	return func(f *SourceIndex) {
		f.Generate = val
	}
}

func WithSourceIndexExternal(val []string) SourceIndexOption {
	return func(f *SourceIndex) {
		f.External = val
	}
}

// NewSourceIndex instantiates a new SourceIndex object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSourceIndex(indexName string, opts ...SourceIndexOption) *SourceIndex {
	this := &SourceIndex{}
	this.IndexName = indexName
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptySourceIndex return a pointer to an empty SourceIndex object.
func NewEmptySourceIndex() *SourceIndex {
	return &SourceIndex{}
}

// GetIndexName returns the IndexName field value.
func (o *SourceIndex) GetIndexName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IndexName
}

// GetIndexNameOk returns a tuple with the IndexName field value
// and a boolean to check if the value has been set.
func (o *SourceIndex) GetIndexNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndexName, true
}

// SetIndexName sets field value.
func (o *SourceIndex) SetIndexName(v string) *SourceIndex {
	o.IndexName = v
	return o
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *SourceIndex) GetReplicas() bool {
	if o == nil || o.Replicas == nil {
		var ret bool
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceIndex) GetReplicasOk() (*bool, bool) {
	if o == nil || o.Replicas == nil {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *SourceIndex) HasReplicas() bool {
	if o != nil && o.Replicas != nil {
		return true
	}

	return false
}

// SetReplicas gets a reference to the given bool and assigns it to the Replicas field.
func (o *SourceIndex) SetReplicas(v bool) *SourceIndex {
	o.Replicas = &v
	return o
}

// GetAnalyticsTags returns the AnalyticsTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceIndex) GetAnalyticsTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AnalyticsTags
}

// GetAnalyticsTagsOk returns a tuple with the AnalyticsTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourceIndex) GetAnalyticsTagsOk() ([]string, bool) {
	if o == nil || o.AnalyticsTags == nil {
		return nil, false
	}
	return o.AnalyticsTags, true
}

// HasAnalyticsTags returns a boolean if a field has been set.
func (o *SourceIndex) HasAnalyticsTags() bool {
	if o != nil && o.AnalyticsTags != nil {
		return true
	}

	return false
}

// SetAnalyticsTags gets a reference to the given []string and assigns it to the AnalyticsTags field.
func (o *SourceIndex) SetAnalyticsTags(v []string) *SourceIndex {
	o.AnalyticsTags = v
	return o
}

// GetFacets returns the Facets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceIndex) GetFacets() []Facet {
	if o == nil {
		var ret []Facet
		return ret
	}
	return o.Facets
}

// GetFacetsOk returns a tuple with the Facets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourceIndex) GetFacetsOk() ([]Facet, bool) {
	if o == nil || o.Facets == nil {
		return nil, false
	}
	return o.Facets, true
}

// HasFacets returns a boolean if a field has been set.
func (o *SourceIndex) HasFacets() bool {
	if o != nil && o.Facets != nil {
		return true
	}

	return false
}

// SetFacets gets a reference to the given []Facet and assigns it to the Facets field.
func (o *SourceIndex) SetFacets(v []Facet) *SourceIndex {
	o.Facets = v
	return o
}

// GetMinHits returns the MinHits field value if set, zero value otherwise.
func (o *SourceIndex) GetMinHits() int32 {
	if o == nil || o.MinHits == nil {
		var ret int32
		return ret
	}
	return *o.MinHits
}

// GetMinHitsOk returns a tuple with the MinHits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceIndex) GetMinHitsOk() (*int32, bool) {
	if o == nil || o.MinHits == nil {
		return nil, false
	}
	return o.MinHits, true
}

// HasMinHits returns a boolean if a field has been set.
func (o *SourceIndex) HasMinHits() bool {
	if o != nil && o.MinHits != nil {
		return true
	}

	return false
}

// SetMinHits gets a reference to the given int32 and assigns it to the MinHits field.
func (o *SourceIndex) SetMinHits(v int32) *SourceIndex {
	o.MinHits = &v
	return o
}

// GetMinLetters returns the MinLetters field value if set, zero value otherwise.
func (o *SourceIndex) GetMinLetters() int32 {
	if o == nil || o.MinLetters == nil {
		var ret int32
		return ret
	}
	return *o.MinLetters
}

// GetMinLettersOk returns a tuple with the MinLetters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceIndex) GetMinLettersOk() (*int32, bool) {
	if o == nil || o.MinLetters == nil {
		return nil, false
	}
	return o.MinLetters, true
}

// HasMinLetters returns a boolean if a field has been set.
func (o *SourceIndex) HasMinLetters() bool {
	if o != nil && o.MinLetters != nil {
		return true
	}

	return false
}

// SetMinLetters gets a reference to the given int32 and assigns it to the MinLetters field.
func (o *SourceIndex) SetMinLetters(v int32) *SourceIndex {
	o.MinLetters = &v
	return o
}

// GetGenerate returns the Generate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceIndex) GetGenerate() [][]string {
	if o == nil {
		var ret [][]string
		return ret
	}
	return o.Generate
}

// GetGenerateOk returns a tuple with the Generate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourceIndex) GetGenerateOk() ([][]string, bool) {
	if o == nil || o.Generate == nil {
		return nil, false
	}
	return o.Generate, true
}

// HasGenerate returns a boolean if a field has been set.
func (o *SourceIndex) HasGenerate() bool {
	if o != nil && o.Generate != nil {
		return true
	}

	return false
}

// SetGenerate gets a reference to the given [][]string and assigns it to the Generate field.
func (o *SourceIndex) SetGenerate(v [][]string) *SourceIndex {
	o.Generate = v
	return o
}

// GetExternal returns the External field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceIndex) GetExternal() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.External
}

// GetExternalOk returns a tuple with the External field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SourceIndex) GetExternalOk() ([]string, bool) {
	if o == nil || o.External == nil {
		return nil, false
	}
	return o.External, true
}

// HasExternal returns a boolean if a field has been set.
func (o *SourceIndex) HasExternal() bool {
	if o != nil && o.External != nil {
		return true
	}

	return false
}

// SetExternal gets a reference to the given []string and assigns it to the External field.
func (o *SourceIndex) SetExternal(v []string) *SourceIndex {
	o.External = v
	return o
}

func (o SourceIndex) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	toSerialize["indexName"] = o.IndexName
	if o.Replicas != nil {
		toSerialize["replicas"] = o.Replicas
	}
	if o.AnalyticsTags != nil {
		toSerialize["analyticsTags"] = o.AnalyticsTags
	}
	if o.Facets != nil {
		toSerialize["facets"] = o.Facets
	}
	if o.MinHits != nil {
		toSerialize["minHits"] = o.MinHits
	}
	if o.MinLetters != nil {
		toSerialize["minLetters"] = o.MinLetters
	}
	if o.Generate != nil {
		toSerialize["generate"] = o.Generate
	}
	if o.External != nil {
		toSerialize["external"] = o.External
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal SourceIndex: %w", err)
	}

	return serialized, nil
}

func (o SourceIndex) String() string {
	out := ""
	out += fmt.Sprintf("  indexName=%v\n", o.IndexName)
	out += fmt.Sprintf("  replicas=%v\n", o.Replicas)
	out += fmt.Sprintf("  analyticsTags=%v\n", o.AnalyticsTags)
	out += fmt.Sprintf("  facets=%v\n", o.Facets)
	out += fmt.Sprintf("  minHits=%v\n", o.MinHits)
	out += fmt.Sprintf("  minLetters=%v\n", o.MinLetters)
	out += fmt.Sprintf("  generate=%v\n", o.Generate)
	out += fmt.Sprintf("  external=%v\n", o.External)
	return fmt.Sprintf("SourceIndex {\n%s}", out)
}
