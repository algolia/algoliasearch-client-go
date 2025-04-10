// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// SearchForFacetValuesResponse struct for SearchForFacetValuesResponse.
type SearchForFacetValuesResponse struct {
	// Matching facet values.
	FacetHits []FacetHits `json:"facetHits"`
	// Whether the facet count is exhaustive (true) or approximate (false). For more information, see [Why are my facet and hit counts not accurate](https://support.algolia.com/hc/en-us/articles/4406975248145-Why-are-my-facet-and-hit-counts-not-accurate-).
	ExhaustiveFacetsCount bool `json:"exhaustiveFacetsCount"`
	// Time the server took to process the request, in milliseconds.
	ProcessingTimeMS *int32 `json:"processingTimeMS,omitempty"`
}

type SearchForFacetValuesResponseOption func(f *SearchForFacetValuesResponse)

func WithSearchForFacetValuesResponseProcessingTimeMS(val int32) SearchForFacetValuesResponseOption {
	return func(f *SearchForFacetValuesResponse) {
		f.ProcessingTimeMS = &val
	}
}

// NewSearchForFacetValuesResponse instantiates a new SearchForFacetValuesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSearchForFacetValuesResponse(facetHits []FacetHits, exhaustiveFacetsCount bool, opts ...SearchForFacetValuesResponseOption) *SearchForFacetValuesResponse {
	this := &SearchForFacetValuesResponse{}
	this.FacetHits = facetHits
	this.ExhaustiveFacetsCount = exhaustiveFacetsCount
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptySearchForFacetValuesResponse return a pointer to an empty SearchForFacetValuesResponse object.
func NewEmptySearchForFacetValuesResponse() *SearchForFacetValuesResponse {
	return &SearchForFacetValuesResponse{}
}

// GetFacetHits returns the FacetHits field value.
func (o *SearchForFacetValuesResponse) GetFacetHits() []FacetHits {
	if o == nil {
		var ret []FacetHits
		return ret
	}

	return o.FacetHits
}

// GetFacetHitsOk returns a tuple with the FacetHits field value
// and a boolean to check if the value has been set.
func (o *SearchForFacetValuesResponse) GetFacetHitsOk() ([]FacetHits, bool) {
	if o == nil {
		return nil, false
	}
	return o.FacetHits, true
}

// SetFacetHits sets field value.
func (o *SearchForFacetValuesResponse) SetFacetHits(v []FacetHits) *SearchForFacetValuesResponse {
	o.FacetHits = v
	return o
}

// GetExhaustiveFacetsCount returns the ExhaustiveFacetsCount field value.
func (o *SearchForFacetValuesResponse) GetExhaustiveFacetsCount() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ExhaustiveFacetsCount
}

// GetExhaustiveFacetsCountOk returns a tuple with the ExhaustiveFacetsCount field value
// and a boolean to check if the value has been set.
func (o *SearchForFacetValuesResponse) GetExhaustiveFacetsCountOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExhaustiveFacetsCount, true
}

// SetExhaustiveFacetsCount sets field value.
func (o *SearchForFacetValuesResponse) SetExhaustiveFacetsCount(v bool) *SearchForFacetValuesResponse {
	o.ExhaustiveFacetsCount = v
	return o
}

// GetProcessingTimeMS returns the ProcessingTimeMS field value if set, zero value otherwise.
func (o *SearchForFacetValuesResponse) GetProcessingTimeMS() int32 {
	if o == nil || o.ProcessingTimeMS == nil {
		var ret int32
		return ret
	}
	return *o.ProcessingTimeMS
}

// GetProcessingTimeMSOk returns a tuple with the ProcessingTimeMS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchForFacetValuesResponse) GetProcessingTimeMSOk() (*int32, bool) {
	if o == nil || o.ProcessingTimeMS == nil {
		return nil, false
	}
	return o.ProcessingTimeMS, true
}

// HasProcessingTimeMS returns a boolean if a field has been set.
func (o *SearchForFacetValuesResponse) HasProcessingTimeMS() bool {
	if o != nil && o.ProcessingTimeMS != nil {
		return true
	}

	return false
}

// SetProcessingTimeMS gets a reference to the given int32 and assigns it to the ProcessingTimeMS field.
func (o *SearchForFacetValuesResponse) SetProcessingTimeMS(v int32) *SearchForFacetValuesResponse {
	o.ProcessingTimeMS = &v
	return o
}

func (o SearchForFacetValuesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	toSerialize["facetHits"] = o.FacetHits
	toSerialize["exhaustiveFacetsCount"] = o.ExhaustiveFacetsCount
	if o.ProcessingTimeMS != nil {
		toSerialize["processingTimeMS"] = o.ProcessingTimeMS
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal SearchForFacetValuesResponse: %w", err)
	}

	return serialized, nil
}

func (o SearchForFacetValuesResponse) String() string {
	out := ""
	out += fmt.Sprintf("  facetHits=%v\n", o.FacetHits)
	out += fmt.Sprintf("  exhaustiveFacetsCount=%v\n", o.ExhaustiveFacetsCount)
	out += fmt.Sprintf("  processingTimeMS=%v\n", o.ProcessingTimeMS)
	return fmt.Sprintf("SearchForFacetValuesResponse {\n%s}", out)
}
