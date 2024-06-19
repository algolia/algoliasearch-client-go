// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// RunStatus Task run status.
type RunStatus string

// List of RunStatus.
const (
	RUN_STATUS_CREATED  RunStatus = "created"
	RUN_STATUS_STARTED  RunStatus = "started"
	RUN_STATUS_IDLED    RunStatus = "idled"
	RUN_STATUS_FINISHED RunStatus = "finished"
	RUN_STATUS_SKIPPED  RunStatus = "skipped"
)

// All allowed values of RunStatus enum.
var AllowedRunStatusEnumValues = []RunStatus{
	"created",
	"started",
	"idled",
	"finished",
	"skipped",
}

func (v *RunStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return fmt.Errorf("failed to unmarshal value '%s' for enum 'RunStatus': %w", string(src), err)
	}
	enumTypeValue := RunStatus(value)
	for _, existing := range AllowedRunStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RunStatus", value)
}

// NewRunStatusFromValue returns a pointer to a valid RunStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRunStatusFromValue(v string) (*RunStatus, error) {
	ev := RunStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RunStatus: valid values are %v", v, AllowedRunStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RunStatus) IsValid() bool {
	for _, existing := range AllowedRunStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RunStatus value.
func (v RunStatus) Ptr() *RunStatus {
	return &v
}
