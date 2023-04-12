// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// Languages A dictionary language.
type Languages struct {
	Plurals   NullableDictionaryLanguage `json:"plurals"`
	Stopwords NullableDictionaryLanguage `json:"stopwords"`
	Compounds NullableDictionaryLanguage `json:"compounds"`
}

// NewLanguages instantiates a new Languages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLanguages(plurals NullableDictionaryLanguage, stopwords NullableDictionaryLanguage, compounds NullableDictionaryLanguage) *Languages {
	this := &Languages{}
	this.Plurals = plurals
	this.Stopwords = stopwords
	this.Compounds = compounds
	return this
}

// NewLanguagesWithDefaults instantiates a new Languages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLanguagesWithDefaults() *Languages {
	this := &Languages{}
	return this
}

// GetPlurals returns the Plurals field value
// If the value is explicit nil, the zero value for DictionaryLanguage will be returned
func (o *Languages) GetPlurals() DictionaryLanguage {
	if o == nil || o.Plurals.Get() == nil {
		var ret DictionaryLanguage
		return ret
	}

	return *o.Plurals.Get()
}

// GetPluralsOk returns a tuple with the Plurals field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Languages) GetPluralsOk() (*DictionaryLanguage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Plurals.Get(), o.Plurals.IsSet()
}

// SetPlurals sets field value
func (o *Languages) SetPlurals(v DictionaryLanguage) {
	o.Plurals.Set(&v)
}

// GetStopwords returns the Stopwords field value
// If the value is explicit nil, the zero value for DictionaryLanguage will be returned
func (o *Languages) GetStopwords() DictionaryLanguage {
	if o == nil || o.Stopwords.Get() == nil {
		var ret DictionaryLanguage
		return ret
	}

	return *o.Stopwords.Get()
}

// GetStopwordsOk returns a tuple with the Stopwords field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Languages) GetStopwordsOk() (*DictionaryLanguage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stopwords.Get(), o.Stopwords.IsSet()
}

// SetStopwords sets field value
func (o *Languages) SetStopwords(v DictionaryLanguage) {
	o.Stopwords.Set(&v)
}

// GetCompounds returns the Compounds field value
// If the value is explicit nil, the zero value for DictionaryLanguage will be returned
func (o *Languages) GetCompounds() DictionaryLanguage {
	if o == nil || o.Compounds.Get() == nil {
		var ret DictionaryLanguage
		return ret
	}

	return *o.Compounds.Get()
}

// GetCompoundsOk returns a tuple with the Compounds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Languages) GetCompoundsOk() (*DictionaryLanguage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Compounds.Get(), o.Compounds.IsSet()
}

// SetCompounds sets field value
func (o *Languages) SetCompounds(v DictionaryLanguage) {
	o.Compounds.Set(&v)
}

func (o Languages) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["plurals"] = o.Plurals.Get()
	}
	if true {
		toSerialize["stopwords"] = o.Stopwords.Get()
	}
	if true {
		toSerialize["compounds"] = o.Compounds.Get()
	}
	return json.Marshal(toSerialize)
}

func (o Languages) String() string {
	out := ""
	out += fmt.Sprintf("  plurals=%v\n", o.Plurals)
	out += fmt.Sprintf("  stopwords=%v\n", o.Stopwords)
	out += fmt.Sprintf("  compounds=%v\n", o.Compounds)
	return fmt.Sprintf("Languages {\n%s}", out)
}

type NullableLanguages struct {
	value *Languages
	isSet bool
}

func (v NullableLanguages) Get() *Languages {
	return v.value
}

func (v *NullableLanguages) Set(val *Languages) {
	v.value = val
	v.isSet = true
}

func (v NullableLanguages) IsSet() bool {
	return v.isSet
}

func (v *NullableLanguages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLanguages(val *Languages) *NullableLanguages {
	return &NullableLanguages{value: val, isSet: true}
}

func (v NullableLanguages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLanguages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}