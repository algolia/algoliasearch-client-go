// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package recommend

import (
	"encoding/json"
	"fmt"
)

// Anchoring Which part of the search query the pattern should match:  - `startsWith`. The pattern must match the begginning of the query. - `endsWith`. The pattern must match the end of the query. - `is`. The pattern must match the query exactly. - `contains`. The pattern must match anywhere in the query.  Empty queries are only allowed as pattern with `anchoring: is`.
type Anchoring string

// List of anchoring.
const (
	ANCHORING_IS          Anchoring = "is"
	ANCHORING_STARTS_WITH Anchoring = "startsWith"
	ANCHORING_ENDS_WITH   Anchoring = "endsWith"
	ANCHORING_CONTAINS    Anchoring = "contains"
)

// All allowed values of Anchoring enum.
var AllowedAnchoringEnumValues = []Anchoring{
	"is",
	"startsWith",
	"endsWith",
	"contains",
}

func (v *Anchoring) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return fmt.Errorf("failed to unmarshal value '%s' for enum 'Anchoring': %w", string(src), err)
	}
	enumTypeValue := Anchoring(value)
	for _, existing := range AllowedAnchoringEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Anchoring", value)
}

// NewAnchoringFromValue returns a pointer to a valid Anchoring
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAnchoringFromValue(v string) (*Anchoring, error) {
	ev := Anchoring(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Anchoring: valid values are %v", v, AllowedAnchoringEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v Anchoring) IsValid() bool {
	for _, existing := range AllowedAnchoringEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to anchoring value.
func (v Anchoring) Ptr() *Anchoring {
	return &v
}

type NullableAnchoring struct {
	value *Anchoring
	isSet bool
}

func (v NullableAnchoring) Get() *Anchoring {
	return v.value
}

func (v *NullableAnchoring) Set(val *Anchoring) {
	v.value = val
	v.isSet = true
}

func (v NullableAnchoring) IsSet() bool {
	return v.isSet
}

func (v *NullableAnchoring) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnchoring(val *Anchoring) *NullableAnchoring {
	return &NullableAnchoring{value: val, isSet: true}
}

func (v NullableAnchoring) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableAnchoring) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}