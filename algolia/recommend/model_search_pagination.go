// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package recommend

import (
	"encoding/json"
	"fmt"
)

// SearchPagination struct for SearchPagination.
type SearchPagination struct {
	// Page of search results to retrieve.
	Page int32 `json:"page"`
	// Number of results (hits).
	NbHits int32 `json:"nbHits"`
	// Number of pages of results.
	NbPages int32 `json:"nbPages"`
	// Number of hits per page.
	HitsPerPage int32 `json:"hitsPerPage"`
}

// NewSearchPagination instantiates a new SearchPagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSearchPagination(page int32, nbHits int32, nbPages int32, hitsPerPage int32) *SearchPagination {
	this := &SearchPagination{}
	this.Page = page
	this.NbHits = nbHits
	this.NbPages = nbPages
	this.HitsPerPage = hitsPerPage
	return this
}

// NewEmptySearchPagination return a pointer to an empty SearchPagination object.
func NewEmptySearchPagination() *SearchPagination {
	return &SearchPagination{}
}

// GetPage returns the Page field value.
func (o *SearchPagination) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *SearchPagination) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value.
func (o *SearchPagination) SetPage(v int32) *SearchPagination {
	o.Page = v
	return o
}

// GetNbHits returns the NbHits field value.
func (o *SearchPagination) GetNbHits() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NbHits
}

// GetNbHitsOk returns a tuple with the NbHits field value
// and a boolean to check if the value has been set.
func (o *SearchPagination) GetNbHitsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NbHits, true
}

// SetNbHits sets field value.
func (o *SearchPagination) SetNbHits(v int32) *SearchPagination {
	o.NbHits = v
	return o
}

// GetNbPages returns the NbPages field value.
func (o *SearchPagination) GetNbPages() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NbPages
}

// GetNbPagesOk returns a tuple with the NbPages field value
// and a boolean to check if the value has been set.
func (o *SearchPagination) GetNbPagesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NbPages, true
}

// SetNbPages sets field value.
func (o *SearchPagination) SetNbPages(v int32) *SearchPagination {
	o.NbPages = v
	return o
}

// GetHitsPerPage returns the HitsPerPage field value.
func (o *SearchPagination) GetHitsPerPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.HitsPerPage
}

// GetHitsPerPageOk returns a tuple with the HitsPerPage field value
// and a boolean to check if the value has been set.
func (o *SearchPagination) GetHitsPerPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HitsPerPage, true
}

// SetHitsPerPage sets field value.
func (o *SearchPagination) SetHitsPerPage(v int32) *SearchPagination {
	o.HitsPerPage = v
	return o
}

func (o SearchPagination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["page"] = o.Page
	}
	if true {
		toSerialize["nbHits"] = o.NbHits
	}
	if true {
		toSerialize["nbPages"] = o.NbPages
	}
	if true {
		toSerialize["hitsPerPage"] = o.HitsPerPage
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal SearchPagination: %w", err)
	}

	return serialized, nil
}

func (o SearchPagination) String() string {
	out := ""
	out += fmt.Sprintf("  page=%v\n", o.Page)
	out += fmt.Sprintf("  nbHits=%v\n", o.NbHits)
	out += fmt.Sprintf("  nbPages=%v\n", o.NbPages)
	out += fmt.Sprintf("  hitsPerPage=%v\n", o.HitsPerPage)
	return fmt.Sprintf("SearchPagination {\n%s}", out)
}
