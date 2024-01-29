// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// TaskSortKeys Used to sort the Task list endpoint.
type TaskSortKeys string

// List of taskSortKeys.
const (
	TASKSORTKEYS_ENABLED      TaskSortKeys = "enabled"
	TASKSORTKEYS_TRIGGER_TYPE TaskSortKeys = "triggerType"
	TASKSORTKEYS_ACTION       TaskSortKeys = "action"
	TASKSORTKEYS_UPDATED_AT   TaskSortKeys = "updatedAt"
	TASKSORTKEYS_CREATED_AT   TaskSortKeys = "createdAt"
)

// All allowed values of TaskSortKeys enum.
var AllowedTaskSortKeysEnumValues = []TaskSortKeys{
	"enabled",
	"triggerType",
	"action",
	"updatedAt",
	"createdAt",
}

func (v *TaskSortKeys) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return fmt.Errorf("failed to unmarshal value '%s' for enum 'TaskSortKeys': %w", string(src), err)
	}
	enumTypeValue := TaskSortKeys(value)
	for _, existing := range AllowedTaskSortKeysEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TaskSortKeys", value)
}

// NewTaskSortKeysFromValue returns a pointer to a valid TaskSortKeys
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTaskSortKeysFromValue(v string) (*TaskSortKeys, error) {
	ev := TaskSortKeys(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TaskSortKeys: valid values are %v", v, AllowedTaskSortKeysEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TaskSortKeys) IsValid() bool {
	for _, existing := range AllowedTaskSortKeysEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to taskSortKeys value.
func (v TaskSortKeys) Ptr() *TaskSortKeys {
	return &v
}

type NullableTaskSortKeys struct {
	value *TaskSortKeys
	isSet bool
}

func (v NullableTaskSortKeys) Get() *TaskSortKeys {
	return v.value
}

func (v *NullableTaskSortKeys) Set(val *TaskSortKeys) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskSortKeys) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskSortKeys) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskSortKeys(val *TaskSortKeys) *NullableTaskSortKeys {
	return &NullableTaskSortKeys{value: val, isSet: true}
}

func (v NullableTaskSortKeys) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableTaskSortKeys) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
