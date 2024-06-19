// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// Mode Search mode the index will use to query for results.  This setting only applies to indices, for which Algolia enabled NeuralSearch for you.
type Mode string

// List of mode.
const (
	MODE_NEURAL_SEARCH  Mode = "neuralSearch"
	MODE_KEYWORD_SEARCH Mode = "keywordSearch"
)

// All allowed values of Mode enum.
var AllowedModeEnumValues = []Mode{
	"neuralSearch",
	"keywordSearch",
}

func (v *Mode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return fmt.Errorf("failed to unmarshal value '%s' for enum 'Mode': %w", string(src), err)
	}
	enumTypeValue := Mode(value)
	for _, existing := range AllowedModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Mode", value)
}

// NewModeFromValue returns a pointer to a valid Mode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewModeFromValue(v string) (*Mode, error) {
	ev := Mode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Mode: valid values are %v", v, AllowedModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v Mode) IsValid() bool {
	for _, existing := range AllowedModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to mode value.
func (v Mode) Ptr() *Mode {
	return &v
}
