// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// SearchUserIdsResponse userIDs data.
type SearchUserIdsResponse struct {
	// List of user object matching the query.
	Hits []UserHit `json:"hits" validate:"required"`
	// Number of hits that the search query matched.
	NbHits int32 `json:"nbHits" validate:"required"`
	// Specify the page to retrieve.
	Page int32 `json:"page" validate:"required"`
	// Maximum number of hits in a page. Minimum is 1, maximum is 1000.
	HitsPerPage int32 `json:"hitsPerPage" validate:"required"`
	// Date of last update (ISO-8601 format).
	UpdatedAt string `json:"updatedAt" validate:"required"`
}

// NewSearchUserIdsResponse instantiates a new SearchUserIdsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchUserIdsResponse(hits []UserHit, nbHits int32, page int32, hitsPerPage int32, updatedAt string) *SearchUserIdsResponse {
	this := &SearchUserIdsResponse{}
	this.Hits = hits
	this.NbHits = nbHits
	this.Page = page
	this.HitsPerPage = hitsPerPage
	this.UpdatedAt = updatedAt
	return this
}

// NewSearchUserIdsResponseWithDefaults instantiates a new SearchUserIdsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchUserIdsResponseWithDefaults() *SearchUserIdsResponse {
	this := &SearchUserIdsResponse{}
	var page int32 = 0
	this.Page = page
	var hitsPerPage int32 = 20
	this.HitsPerPage = hitsPerPage
	return this
}

// GetHits returns the Hits field value
func (o *SearchUserIdsResponse) GetHits() []UserHit {
	if o == nil {
		var ret []UserHit
		return ret
	}

	return o.Hits
}

// GetHitsOk returns a tuple with the Hits field value
// and a boolean to check if the value has been set.
func (o *SearchUserIdsResponse) GetHitsOk() ([]UserHit, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hits, true
}

// SetHits sets field value
func (o *SearchUserIdsResponse) SetHits(v []UserHit) {
	o.Hits = v
}

// GetNbHits returns the NbHits field value
func (o *SearchUserIdsResponse) GetNbHits() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NbHits
}

// GetNbHitsOk returns a tuple with the NbHits field value
// and a boolean to check if the value has been set.
func (o *SearchUserIdsResponse) GetNbHitsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NbHits, true
}

// SetNbHits sets field value
func (o *SearchUserIdsResponse) SetNbHits(v int32) {
	o.NbHits = v
}

// GetPage returns the Page field value
func (o *SearchUserIdsResponse) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *SearchUserIdsResponse) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *SearchUserIdsResponse) SetPage(v int32) {
	o.Page = v
}

// GetHitsPerPage returns the HitsPerPage field value
func (o *SearchUserIdsResponse) GetHitsPerPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.HitsPerPage
}

// GetHitsPerPageOk returns a tuple with the HitsPerPage field value
// and a boolean to check if the value has been set.
func (o *SearchUserIdsResponse) GetHitsPerPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HitsPerPage, true
}

// SetHitsPerPage sets field value
func (o *SearchUserIdsResponse) SetHitsPerPage(v int32) {
	o.HitsPerPage = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *SearchUserIdsResponse) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *SearchUserIdsResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *SearchUserIdsResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o SearchUserIdsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["hits"] = o.Hits
	}
	if true {
		toSerialize["nbHits"] = o.NbHits
	}
	if true {
		toSerialize["page"] = o.Page
	}
	if true {
		toSerialize["hitsPerPage"] = o.HitsPerPage
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

func (o SearchUserIdsResponse) String() string {
	out := ""
	out += fmt.Sprintf("  hits=%v\n", o.Hits)
	out += fmt.Sprintf("  nbHits=%v\n", o.NbHits)
	out += fmt.Sprintf("  page=%v\n", o.Page)
	out += fmt.Sprintf("  hitsPerPage=%v\n", o.HitsPerPage)
	out += fmt.Sprintf("  updatedAt=%v\n", o.UpdatedAt)
	return fmt.Sprintf("SearchUserIdsResponse {\n%s}", out)
}

type NullableSearchUserIdsResponse struct {
	value *SearchUserIdsResponse
	isSet bool
}

func (v NullableSearchUserIdsResponse) Get() *SearchUserIdsResponse {
	return v.value
}

func (v *NullableSearchUserIdsResponse) Set(val *SearchUserIdsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchUserIdsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchUserIdsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchUserIdsResponse(val *SearchUserIdsResponse) *NullableSearchUserIdsResponse {
	return &NullableSearchUserIdsResponse{value: val, isSet: true}
}

func (v NullableSearchUserIdsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchUserIdsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}