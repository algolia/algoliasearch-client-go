// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// DeleteApiKeyResponse struct for DeleteApiKeyResponse.
type DeleteApiKeyResponse struct {
	// Date and time when the object was deleted, in RFC 3339 format.
	DeletedAt string `json:"deletedAt"`
}

// NewDeleteApiKeyResponse instantiates a new DeleteApiKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeleteApiKeyResponse(deletedAt string) *DeleteApiKeyResponse {
	this := &DeleteApiKeyResponse{}
	this.DeletedAt = deletedAt
	return this
}

// NewEmptyDeleteApiKeyResponse return a pointer to an empty DeleteApiKeyResponse object.
func NewEmptyDeleteApiKeyResponse() *DeleteApiKeyResponse {
	return &DeleteApiKeyResponse{}
}

// GetDeletedAt returns the DeletedAt field value.
func (o *DeleteApiKeyResponse) GetDeletedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value
// and a boolean to check if the value has been set.
func (o *DeleteApiKeyResponse) GetDeletedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeletedAt, true
}

// SetDeletedAt sets field value.
func (o *DeleteApiKeyResponse) SetDeletedAt(v string) *DeleteApiKeyResponse {
	o.DeletedAt = v
	return o
}

func (o DeleteApiKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["deletedAt"] = o.DeletedAt
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal DeleteApiKeyResponse: %w", err)
	}

	return serialized, nil
}

func (o DeleteApiKeyResponse) String() string {
	out := ""
	out += fmt.Sprintf("  deletedAt=%v\n", o.DeletedAt)
	return fmt.Sprintf("DeleteApiKeyResponse {\n%s}", out)
}

type NullableDeleteApiKeyResponse struct {
	value *DeleteApiKeyResponse
	isSet bool
}

func (v NullableDeleteApiKeyResponse) Get() *DeleteApiKeyResponse {
	return v.value
}

func (v *NullableDeleteApiKeyResponse) Set(val *DeleteApiKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteApiKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteApiKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteApiKeyResponse(val *DeleteApiKeyResponse) *NullableDeleteApiKeyResponse {
	return &NullableDeleteApiKeyResponse{value: val, isSet: true}
}

func (v NullableDeleteApiKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableDeleteApiKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
