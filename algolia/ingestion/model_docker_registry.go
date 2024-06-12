// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// DockerRegistry Container registry name from where to pull the image.
type DockerRegistry string

// List of DockerRegistry.
const (
	DOCKER_REGISTRY_DOCKERHUB DockerRegistry = "dockerhub"
	DOCKER_REGISTRY_GHCR      DockerRegistry = "ghcr"
)

// All allowed values of DockerRegistry enum.
var AllowedDockerRegistryEnumValues = []DockerRegistry{
	"dockerhub",
	"ghcr",
}

func (v *DockerRegistry) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return fmt.Errorf("failed to unmarshal value '%s' for enum 'DockerRegistry': %w", string(src), err)
	}
	enumTypeValue := DockerRegistry(value)
	for _, existing := range AllowedDockerRegistryEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DockerRegistry", value)
}

// NewDockerRegistryFromValue returns a pointer to a valid DockerRegistry
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDockerRegistryFromValue(v string) (*DockerRegistry, error) {
	ev := DockerRegistry(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DockerRegistry: valid values are %v", v, AllowedDockerRegistryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DockerRegistry) IsValid() bool {
	for _, existing := range AllowedDockerRegistryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DockerRegistry value.
func (v DockerRegistry) Ptr() *DockerRegistry {
	return &v
}

type NullableDockerRegistry struct {
	value *DockerRegistry
	isSet bool
}

func (v NullableDockerRegistry) Get() *DockerRegistry {
	return v.value
}

func (v *NullableDockerRegistry) Set(val *DockerRegistry) {
	v.value = val
	v.isSet = true
}

func (v NullableDockerRegistry) IsSet() bool {
	return v.isSet
}

func (v *NullableDockerRegistry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDockerRegistry(val *DockerRegistry) *NullableDockerRegistry {
	return &NullableDockerRegistry{value: val, isSet: true}
}

func (v NullableDockerRegistry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableDockerRegistry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
