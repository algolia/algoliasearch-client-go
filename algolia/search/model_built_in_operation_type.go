// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// BuiltInOperationType How to change the attribute.
type BuiltInOperationType string

// List of builtInOperationType.
const (
	BUILT_IN_OPERATION_TYPE_INCREMENT      BuiltInOperationType = "Increment"
	BUILT_IN_OPERATION_TYPE_DECREMENT      BuiltInOperationType = "Decrement"
	BUILT_IN_OPERATION_TYPE_ADD            BuiltInOperationType = "Add"
	BUILT_IN_OPERATION_TYPE_REMOVE         BuiltInOperationType = "Remove"
	BUILT_IN_OPERATION_TYPE_ADD_UNIQUE     BuiltInOperationType = "AddUnique"
	BUILT_IN_OPERATION_TYPE_INCREMENT_FROM BuiltInOperationType = "IncrementFrom"
	BUILT_IN_OPERATION_TYPE_INCREMENT_SET  BuiltInOperationType = "IncrementSet"
)

// All allowed values of BuiltInOperationType enum.
var AllowedBuiltInOperationTypeEnumValues = []BuiltInOperationType{
	"Increment",
	"Decrement",
	"Add",
	"Remove",
	"AddUnique",
	"IncrementFrom",
	"IncrementSet",
}

func (v *BuiltInOperationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return fmt.Errorf("failed to unmarshal value '%s' for enum 'BuiltInOperationType': %w", string(src), err)
	}
	enumTypeValue := BuiltInOperationType(value)
	for _, existing := range AllowedBuiltInOperationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BuiltInOperationType", value)
}

// NewBuiltInOperationTypeFromValue returns a pointer to a valid BuiltInOperationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewBuiltInOperationTypeFromValue(v string) (*BuiltInOperationType, error) {
	ev := BuiltInOperationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BuiltInOperationType: valid values are %v", v, AllowedBuiltInOperationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v BuiltInOperationType) IsValid() bool {
	for _, existing := range AllowedBuiltInOperationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to builtInOperationType value.
func (v BuiltInOperationType) Ptr() *BuiltInOperationType {
	return &v
}

type NullableBuiltInOperationType struct {
	value *BuiltInOperationType
	isSet bool
}

func (v NullableBuiltInOperationType) Get() *BuiltInOperationType {
	return v.value
}

func (v *NullableBuiltInOperationType) Set(val *BuiltInOperationType) {
	v.value = val
	v.isSet = true
}

func (v NullableBuiltInOperationType) IsSet() bool {
	return v.isSet
}

func (v *NullableBuiltInOperationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuiltInOperationType(val *BuiltInOperationType) *NullableBuiltInOperationType {
	return &NullableBuiltInOperationType{value: val, isSet: true}
}

func (v NullableBuiltInOperationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableBuiltInOperationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
