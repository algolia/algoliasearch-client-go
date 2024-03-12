// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package search

import (
	"encoding/json"
	"fmt"
)

// SearchUserIdsResponse userIDs data.
type SearchUserIdsResponse struct {
	// User objects that match the query.
	Hits []UserHit `json:"hits"`
	// Number of results (hits).
	NbHits int32 `json:"nbHits"`
	// Page of search results to retrieve.
	Page int32 `json:"page"`
	// Maximum number of hits per page.
	HitsPerPage int32 `json:"hitsPerPage"`
	// Timestamp of the last update in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format.
	UpdatedAt string `json:"updatedAt"`
}

// NewSearchUserIdsResponse instantiates a new SearchUserIdsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSearchUserIdsResponse(hits []UserHit, nbHits int32, page int32, hitsPerPage int32, updatedAt string) *SearchUserIdsResponse {
	this := &SearchUserIdsResponse{}
	this.Hits = hits
	this.NbHits = nbHits
	this.Page = page
	this.HitsPerPage = hitsPerPage
	this.UpdatedAt = updatedAt
	return this
}

// NewEmptySearchUserIdsResponse return a pointer to an empty SearchUserIdsResponse object.
func NewEmptySearchUserIdsResponse() *SearchUserIdsResponse {
	return &SearchUserIdsResponse{}
}

// GetHits returns the Hits field value.
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

// SetHits sets field value.
func (o *SearchUserIdsResponse) SetHits(v []UserHit) *SearchUserIdsResponse {
	o.Hits = v
	return o
}

// GetNbHits returns the NbHits field value.
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

// SetNbHits sets field value.
func (o *SearchUserIdsResponse) SetNbHits(v int32) *SearchUserIdsResponse {
	o.NbHits = v
	return o
}

// GetPage returns the Page field value.
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

// SetPage sets field value.
func (o *SearchUserIdsResponse) SetPage(v int32) *SearchUserIdsResponse {
	o.Page = v
	return o
}

// GetHitsPerPage returns the HitsPerPage field value.
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

// SetHitsPerPage sets field value.
func (o *SearchUserIdsResponse) SetHitsPerPage(v int32) *SearchUserIdsResponse {
	o.HitsPerPage = v
	return o
}

// GetUpdatedAt returns the UpdatedAt field value.
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

// SetUpdatedAt sets field value.
func (o *SearchUserIdsResponse) SetUpdatedAt(v string) *SearchUserIdsResponse {
	o.UpdatedAt = v
	return o
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
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal SearchUserIdsResponse: %w", err)
	}

	return serialized, nil
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
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableSearchUserIdsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
