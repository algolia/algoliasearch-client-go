// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package search

import (
	"encoding/json"
	"fmt"
)

// RemoveUserIdResponse struct for RemoveUserIdResponse.
type RemoveUserIdResponse struct {
	// Date and time when the object was deleted, in RFC 3339 format.
	DeletedAt string `json:"deletedAt"`
}

// NewRemoveUserIdResponse instantiates a new RemoveUserIdResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRemoveUserIdResponse(deletedAt string) *RemoveUserIdResponse {
	this := &RemoveUserIdResponse{}
	this.DeletedAt = deletedAt
	return this
}

// NewEmptyRemoveUserIdResponse return a pointer to an empty RemoveUserIdResponse object.
func NewEmptyRemoveUserIdResponse() *RemoveUserIdResponse {
	return &RemoveUserIdResponse{}
}

// GetDeletedAt returns the DeletedAt field value.
func (o *RemoveUserIdResponse) GetDeletedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value
// and a boolean to check if the value has been set.
func (o *RemoveUserIdResponse) GetDeletedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeletedAt, true
}

// SetDeletedAt sets field value.
func (o *RemoveUserIdResponse) SetDeletedAt(v string) *RemoveUserIdResponse {
	o.DeletedAt = v
	return o
}

func (o RemoveUserIdResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["deletedAt"] = o.DeletedAt
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal RemoveUserIdResponse: %w", err)
	}

	return serialized, nil
}

func (o RemoveUserIdResponse) String() string {
	out := ""
	out += fmt.Sprintf("  deletedAt=%v\n", o.DeletedAt)
	return fmt.Sprintf("RemoveUserIdResponse {\n%s}", out)
}

type NullableRemoveUserIdResponse struct {
	value *RemoveUserIdResponse
	isSet bool
}

func (v NullableRemoveUserIdResponse) Get() *RemoveUserIdResponse {
	return v.value
}

func (v *NullableRemoveUserIdResponse) Set(val *RemoveUserIdResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveUserIdResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveUserIdResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveUserIdResponse(val *RemoveUserIdResponse) *NullableRemoveUserIdResponse {
	return &NullableRemoveUserIdResponse{value: val, isSet: true}
}

func (v NullableRemoveUserIdResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableRemoveUserIdResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
