// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// ReplaceSourceResponse struct for ReplaceSourceResponse.
type ReplaceSourceResponse struct {
	// Date and time when the object was updated, in RFC 3339 format.
	UpdatedAt string `json:"updatedAt"`
}

// NewReplaceSourceResponse instantiates a new ReplaceSourceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReplaceSourceResponse(updatedAt string) *ReplaceSourceResponse {
	this := &ReplaceSourceResponse{}
	this.UpdatedAt = updatedAt
	return this
}

// NewEmptyReplaceSourceResponse return a pointer to an empty ReplaceSourceResponse object.
func NewEmptyReplaceSourceResponse() *ReplaceSourceResponse {
	return &ReplaceSourceResponse{}
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *ReplaceSourceResponse) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ReplaceSourceResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *ReplaceSourceResponse) SetUpdatedAt(v string) *ReplaceSourceResponse {
	o.UpdatedAt = v
	return o
}

func (o ReplaceSourceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	toSerialize["updatedAt"] = o.UpdatedAt
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal ReplaceSourceResponse: %w", err)
	}

	return serialized, nil
}

func (o ReplaceSourceResponse) String() string {
	out := ""
	out += fmt.Sprintf("  updatedAt=%v\n", o.UpdatedAt)
	return fmt.Sprintf("ReplaceSourceResponse {\n%s}", out)
}
