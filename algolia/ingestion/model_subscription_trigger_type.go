// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// SubscriptionTriggerType Task runs after receiving subscribed event.
type SubscriptionTriggerType string

// List of SubscriptionTriggerType.
const (
	SUBSCRIPTIONTRIGGERTYPE_SUBSCRIPTION SubscriptionTriggerType = "subscription"
)

// All allowed values of SubscriptionTriggerType enum.
var AllowedSubscriptionTriggerTypeEnumValues = []SubscriptionTriggerType{
	"subscription",
}

func (v *SubscriptionTriggerType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return fmt.Errorf("failed to unmarshal value '%s' for enum 'SubscriptionTriggerType': %w", string(src), err)
	}
	enumTypeValue := SubscriptionTriggerType(value)
	for _, existing := range AllowedSubscriptionTriggerTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SubscriptionTriggerType", value)
}

// NewSubscriptionTriggerTypeFromValue returns a pointer to a valid SubscriptionTriggerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSubscriptionTriggerTypeFromValue(v string) (*SubscriptionTriggerType, error) {
	ev := SubscriptionTriggerType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SubscriptionTriggerType: valid values are %v", v, AllowedSubscriptionTriggerTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SubscriptionTriggerType) IsValid() bool {
	for _, existing := range AllowedSubscriptionTriggerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SubscriptionTriggerType value.
func (v SubscriptionTriggerType) Ptr() *SubscriptionTriggerType {
	return &v
}

type NullableSubscriptionTriggerType struct {
	value *SubscriptionTriggerType
	isSet bool
}

func (v NullableSubscriptionTriggerType) Get() *SubscriptionTriggerType {
	return v.value
}

func (v *NullableSubscriptionTriggerType) Set(val *SubscriptionTriggerType) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionTriggerType) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionTriggerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionTriggerType(val *SubscriptionTriggerType) *NullableSubscriptionTriggerType {
	return &NullableSubscriptionTriggerType{value: val, isSet: true}
}

func (v NullableSubscriptionTriggerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableSubscriptionTriggerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
