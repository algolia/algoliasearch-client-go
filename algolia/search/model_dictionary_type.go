// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// DictionaryType the model 'DictionaryType'.
type DictionaryType string

// List of dictionaryType.
const (
	DICTIONARY_TYPE_PLURALS   DictionaryType = "plurals"
	DICTIONARY_TYPE_STOPWORDS DictionaryType = "stopwords"
	DICTIONARY_TYPE_COMPOUNDS DictionaryType = "compounds"
)

// All allowed values of DictionaryType enum.
var AllowedDictionaryTypeEnumValues = []DictionaryType{
	"plurals",
	"stopwords",
	"compounds",
}

func (v *DictionaryType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return fmt.Errorf("failed to unmarshal value '%s' for enum 'DictionaryType': %w", string(src), err)
	}
	enumTypeValue := DictionaryType(value)
	for _, existing := range AllowedDictionaryTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DictionaryType", value)
}

// NewDictionaryTypeFromValue returns a pointer to a valid DictionaryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDictionaryTypeFromValue(v string) (*DictionaryType, error) {
	ev := DictionaryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DictionaryType: valid values are %v", v, AllowedDictionaryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DictionaryType) IsValid() bool {
	for _, existing := range AllowedDictionaryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dictionaryType value.
func (v DictionaryType) Ptr() *DictionaryType {
	return &v
}

type NullableDictionaryType struct {
	value *DictionaryType
	isSet bool
}

func (v NullableDictionaryType) Get() *DictionaryType {
	return v.value
}

func (v *NullableDictionaryType) Set(val *DictionaryType) {
	v.value = val
	v.isSet = true
}

func (v NullableDictionaryType) IsSet() bool {
	return v.isSet
}

func (v *NullableDictionaryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDictionaryType(val *DictionaryType) *NullableDictionaryType {
	return &NullableDictionaryType{value: val, isSet: true}
}

func (v NullableDictionaryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableDictionaryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
