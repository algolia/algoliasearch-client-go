// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// Pagination Paginated API response.
type Pagination struct {
	// Number of pages in the API response.
	NbPages int32 `json:"nbPages"`
	// Page of the API response to retrieve.
	Page int32 `json:"page"`
	// Number of items in the API response.
	NbItems int32 `json:"nbItems"`
	// Number of items per page.
	ItemsPerPage int32 `json:"itemsPerPage"`
}

// NewPagination instantiates a new Pagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPagination(nbPages int32, page int32, nbItems int32, itemsPerPage int32) *Pagination {
	this := &Pagination{}
	this.NbPages = nbPages
	this.Page = page
	this.NbItems = nbItems
	this.ItemsPerPage = itemsPerPage
	return this
}

// NewEmptyPagination return a pointer to an empty Pagination object.
func NewEmptyPagination() *Pagination {
	return &Pagination{}
}

// GetNbPages returns the NbPages field value.
func (o *Pagination) GetNbPages() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NbPages
}

// GetNbPagesOk returns a tuple with the NbPages field value
// and a boolean to check if the value has been set.
func (o *Pagination) GetNbPagesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NbPages, true
}

// SetNbPages sets field value.
func (o *Pagination) SetNbPages(v int32) *Pagination {
	o.NbPages = v
	return o
}

// GetPage returns the Page field value.
func (o *Pagination) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *Pagination) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value.
func (o *Pagination) SetPage(v int32) *Pagination {
	o.Page = v
	return o
}

// GetNbItems returns the NbItems field value.
func (o *Pagination) GetNbItems() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NbItems
}

// GetNbItemsOk returns a tuple with the NbItems field value
// and a boolean to check if the value has been set.
func (o *Pagination) GetNbItemsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NbItems, true
}

// SetNbItems sets field value.
func (o *Pagination) SetNbItems(v int32) *Pagination {
	o.NbItems = v
	return o
}

// GetItemsPerPage returns the ItemsPerPage field value.
func (o *Pagination) GetItemsPerPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ItemsPerPage
}

// GetItemsPerPageOk returns a tuple with the ItemsPerPage field value
// and a boolean to check if the value has been set.
func (o *Pagination) GetItemsPerPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemsPerPage, true
}

// SetItemsPerPage sets field value.
func (o *Pagination) SetItemsPerPage(v int32) *Pagination {
	o.ItemsPerPage = v
	return o
}

func (o Pagination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	toSerialize["nbPages"] = o.NbPages
	toSerialize["page"] = o.Page
	toSerialize["nbItems"] = o.NbItems
	toSerialize["itemsPerPage"] = o.ItemsPerPage
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal Pagination: %w", err)
	}

	return serialized, nil
}

func (o Pagination) String() string {
	out := ""
	out += fmt.Sprintf("  nbPages=%v\n", o.NbPages)
	out += fmt.Sprintf("  page=%v\n", o.Page)
	out += fmt.Sprintf("  nbItems=%v\n", o.NbItems)
	out += fmt.Sprintf("  itemsPerPage=%v\n", o.ItemsPerPage)
	return fmt.Sprintf("Pagination {\n%s}", out)
}
