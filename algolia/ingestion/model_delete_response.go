// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// DeleteResponse struct for DeleteResponse.
type DeleteResponse struct {
	// Date of deletion (RFC3339 format).
	DeletedAt string `json:"deletedAt"`
}

// NewDeleteResponse instantiates a new DeleteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeleteResponse(deletedAt string) *DeleteResponse {
	this := &DeleteResponse{}
	this.DeletedAt = deletedAt
	return this
}

// NewEmptyDeleteResponse return a pointer to an empty DeleteResponse object.
func NewEmptyDeleteResponse() *DeleteResponse {
	return &DeleteResponse{}
}

// GetDeletedAt returns the DeletedAt field value.
func (o *DeleteResponse) GetDeletedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value
// and a boolean to check if the value has been set.
func (o *DeleteResponse) GetDeletedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeletedAt, true
}

// SetDeletedAt sets field value.
func (o *DeleteResponse) SetDeletedAt(v string) *DeleteResponse {
	o.DeletedAt = v
	return o
}

func (o DeleteResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["deletedAt"] = o.DeletedAt
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal DeleteResponse: %w", err)
	}

	return serialized, nil
}

func (o DeleteResponse) String() string {
	out := ""
	out += fmt.Sprintf("  deletedAt=%v\n", o.DeletedAt)
	return fmt.Sprintf("DeleteResponse {\n%s}", out)
}

type NullableDeleteResponse struct {
	value *DeleteResponse
	isSet bool
}

func (v NullableDeleteResponse) Get() *DeleteResponse {
	return v.value
}

func (v *NullableDeleteResponse) Set(val *DeleteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteResponse(val *DeleteResponse) *NullableDeleteResponse {
	return &NullableDeleteResponse{value: val, isSet: true}
}

func (v NullableDeleteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableDeleteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
