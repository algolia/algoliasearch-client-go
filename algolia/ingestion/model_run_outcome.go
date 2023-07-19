// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// RunOutcome the model 'RunOutcome'
type RunOutcome string

// List of RunOutcome
const (
	RUNOUTCOME_SUCCESS    RunOutcome = "success"
	RUNOUTCOME_FAILURE    RunOutcome = "failure"
	RUNOUTCOME_PROCESSING RunOutcome = "processing"
)

// All allowed values of RunOutcome enum
var AllowedRunOutcomeEnumValues = []RunOutcome{
	"success",
	"failure",
	"processing",
}

func (v *RunOutcome) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RunOutcome(value)
	for _, existing := range AllowedRunOutcomeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RunOutcome", value)
}

// NewRunOutcomeFromValue returns a pointer to a valid RunOutcome
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRunOutcomeFromValue(v string) (*RunOutcome, error) {
	ev := RunOutcome(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RunOutcome: valid values are %v", v, AllowedRunOutcomeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RunOutcome) IsValid() bool {
	for _, existing := range AllowedRunOutcomeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RunOutcome value
func (v RunOutcome) Ptr() *RunOutcome {
	return &v
}

type NullableRunOutcome struct {
	value *RunOutcome
	isSet bool
}

func (v NullableRunOutcome) Get() *RunOutcome {
	return v.value
}

func (v *NullableRunOutcome) Set(val *RunOutcome) {
	v.value = val
	v.isSet = true
}

func (v NullableRunOutcome) IsSet() bool {
	return v.isSet
}

func (v *NullableRunOutcome) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunOutcome(val *RunOutcome) *NullableRunOutcome {
	return &NullableRunOutcome{value: val, isSet: true}
}

func (v NullableRunOutcome) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunOutcome) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}