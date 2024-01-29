// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package analytics

import (
	"encoding/json"
	"fmt"
)

// TopSearch struct for TopSearch.
type TopSearch struct {
	// User query.
	Search string `json:"search"`
	// Number of tracked _and_ untracked searches (where the `clickAnalytics` parameter isn't `true`).
	Count int32 `json:"count"`
	// Number of hits the search query matched.
	NbHits int32 `json:"nbHits"`
}

// NewTopSearch instantiates a new TopSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTopSearch(search string, count int32, nbHits int32) *TopSearch {
	this := &TopSearch{}
	this.Search = search
	this.Count = count
	this.NbHits = nbHits
	return this
}

// NewEmptyTopSearch return a pointer to an empty TopSearch object.
func NewEmptyTopSearch() *TopSearch {
	return &TopSearch{}
}

// GetSearch returns the Search field value.
func (o *TopSearch) GetSearch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Search
}

// GetSearchOk returns a tuple with the Search field value
// and a boolean to check if the value has been set.
func (o *TopSearch) GetSearchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Search, true
}

// SetSearch sets field value.
func (o *TopSearch) SetSearch(v string) *TopSearch {
	o.Search = v
	return o
}

// GetCount returns the Count field value.
func (o *TopSearch) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *TopSearch) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value.
func (o *TopSearch) SetCount(v int32) *TopSearch {
	o.Count = v
	return o
}

// GetNbHits returns the NbHits field value.
func (o *TopSearch) GetNbHits() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NbHits
}

// GetNbHitsOk returns a tuple with the NbHits field value
// and a boolean to check if the value has been set.
func (o *TopSearch) GetNbHitsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NbHits, true
}

// SetNbHits sets field value.
func (o *TopSearch) SetNbHits(v int32) *TopSearch {
	o.NbHits = v
	return o
}

func (o TopSearch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["search"] = o.Search
	}
	if true {
		toSerialize["count"] = o.Count
	}
	if true {
		toSerialize["nbHits"] = o.NbHits
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal TopSearch: %w", err)
	}

	return serialized, nil
}

func (o TopSearch) String() string {
	out := ""
	out += fmt.Sprintf("  search=%v\n", o.Search)
	out += fmt.Sprintf("  count=%v\n", o.Count)
	out += fmt.Sprintf("  nbHits=%v\n", o.NbHits)
	return fmt.Sprintf("TopSearch {\n%s}", out)
}

type NullableTopSearch struct {
	value *TopSearch
	isSet bool
}

func (v NullableTopSearch) Get() *TopSearch {
	return v.value
}

func (v *NullableTopSearch) Set(val *TopSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableTopSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableTopSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopSearch(val *TopSearch) *NullableTopSearch {
	return &NullableTopSearch{value: val, isSet: true}
}

func (v NullableTopSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableTopSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
