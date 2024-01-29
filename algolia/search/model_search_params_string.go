// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package search

import (
	"encoding/json"
	"fmt"
)

// SearchParamsString struct for SearchParamsString.
type SearchParamsString struct {
	// Search parameters as a URL-encoded query string.
	Params *string `json:"params,omitempty"`
}

type SearchParamsStringOption func(f *SearchParamsString)

func WithSearchParamsStringParams(val string) SearchParamsStringOption {
	return func(f *SearchParamsString) {
		f.Params = &val
	}
}

// NewSearchParamsString instantiates a new SearchParamsString object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSearchParamsString(opts ...SearchParamsStringOption) *SearchParamsString {
	this := &SearchParamsString{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptySearchParamsString return a pointer to an empty SearchParamsString object.
func NewEmptySearchParamsString() *SearchParamsString {
	return &SearchParamsString{}
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *SearchParamsString) GetParams() string {
	if o == nil || o.Params == nil {
		var ret string
		return ret
	}
	return *o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchParamsString) GetParamsOk() (*string, bool) {
	if o == nil || o.Params == nil {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *SearchParamsString) HasParams() bool {
	if o != nil && o.Params != nil {
		return true
	}

	return false
}

// SetParams gets a reference to the given string and assigns it to the Params field.
func (o *SearchParamsString) SetParams(v string) *SearchParamsString {
	o.Params = &v
	return o
}

func (o SearchParamsString) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Params != nil {
		toSerialize["params"] = o.Params
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal SearchParamsString: %w", err)
	}

	return serialized, nil
}

func (o SearchParamsString) String() string {
	out := ""
	out += fmt.Sprintf("  params=%v\n", o.Params)
	return fmt.Sprintf("SearchParamsString {\n%s}", out)
}

type NullableSearchParamsString struct {
	value *SearchParamsString
	isSet bool
}

func (v NullableSearchParamsString) Get() *SearchParamsString {
	return v.value
}

func (v *NullableSearchParamsString) Set(val *SearchParamsString) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchParamsString) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchParamsString) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchParamsString(val *SearchParamsString) *NullableSearchParamsString {
	return &NullableSearchParamsString{value: val, isSet: true}
}

func (v NullableSearchParamsString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableSearchParamsString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
