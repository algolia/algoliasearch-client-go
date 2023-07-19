// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// SourceUpdateResponse struct for SourceUpdateResponse
type SourceUpdateResponse struct {
	// The source UUID.
	SourceID string `json:"sourceID" validate:"required"`
	Name     string `json:"name" validate:"required"`
	// Date of last update (RFC3339 format).
	UpdatedAt string `json:"updatedAt" validate:"required"`
}

// NewSourceUpdateResponse instantiates a new SourceUpdateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceUpdateResponse(sourceID string, name string, updatedAt string) *SourceUpdateResponse {
	this := &SourceUpdateResponse{}
	this.SourceID = sourceID
	this.Name = name
	this.UpdatedAt = updatedAt
	return this
}

// NewSourceUpdateResponseWithDefaults instantiates a new SourceUpdateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceUpdateResponseWithDefaults() *SourceUpdateResponse {
	this := &SourceUpdateResponse{}
	return this
}

// GetSourceID returns the SourceID field value
func (o *SourceUpdateResponse) GetSourceID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceID
}

// GetSourceIDOk returns a tuple with the SourceID field value
// and a boolean to check if the value has been set.
func (o *SourceUpdateResponse) GetSourceIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceID, true
}

// SetSourceID sets field value
func (o *SourceUpdateResponse) SetSourceID(v string) {
	o.SourceID = v
}

// GetName returns the Name field value
func (o *SourceUpdateResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SourceUpdateResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SourceUpdateResponse) SetName(v string) {
	o.Name = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *SourceUpdateResponse) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *SourceUpdateResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *SourceUpdateResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o SourceUpdateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["sourceID"] = o.SourceID
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

func (o SourceUpdateResponse) String() string {
	out := ""
	out += fmt.Sprintf("  sourceID=%v\n", o.SourceID)
	out += fmt.Sprintf("  name=%v\n", o.Name)
	out += fmt.Sprintf("  updatedAt=%v\n", o.UpdatedAt)
	return fmt.Sprintf("SourceUpdateResponse {\n%s}", out)
}

type NullableSourceUpdateResponse struct {
	value *SourceUpdateResponse
	isSet bool
}

func (v NullableSourceUpdateResponse) Get() *SourceUpdateResponse {
	return v.value
}

func (v *NullableSourceUpdateResponse) Set(val *SourceUpdateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceUpdateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceUpdateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceUpdateResponse(val *SourceUpdateResponse) *NullableSourceUpdateResponse {
	return &NullableSourceUpdateResponse{value: val, isSet: true}
}

func (v NullableSourceUpdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceUpdateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}