// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// DestinationCreateResponse API response for creating a new destination.
type DestinationCreateResponse struct {
	// Universally unique identifier (UUID) of a destination resource.
	DestinationID string `json:"destinationID"`
	// Descriptive name for the resource.
	Name string `json:"name"`
	// Date of creation in RFC3339 format.
	CreatedAt string `json:"createdAt"`
}

// NewDestinationCreateResponse instantiates a new DestinationCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDestinationCreateResponse(destinationID string, name string, createdAt string) *DestinationCreateResponse {
	this := &DestinationCreateResponse{}
	this.DestinationID = destinationID
	this.Name = name
	this.CreatedAt = createdAt
	return this
}

// NewEmptyDestinationCreateResponse return a pointer to an empty DestinationCreateResponse object.
func NewEmptyDestinationCreateResponse() *DestinationCreateResponse {
	return &DestinationCreateResponse{}
}

// GetDestinationID returns the DestinationID field value.
func (o *DestinationCreateResponse) GetDestinationID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationID
}

// GetDestinationIDOk returns a tuple with the DestinationID field value
// and a boolean to check if the value has been set.
func (o *DestinationCreateResponse) GetDestinationIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationID, true
}

// SetDestinationID sets field value.
func (o *DestinationCreateResponse) SetDestinationID(v string) *DestinationCreateResponse {
	o.DestinationID = v
	return o
}

// GetName returns the Name field value.
func (o *DestinationCreateResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DestinationCreateResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DestinationCreateResponse) SetName(v string) *DestinationCreateResponse {
	o.Name = v
	return o
}

// GetCreatedAt returns the CreatedAt field value.
func (o *DestinationCreateResponse) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DestinationCreateResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *DestinationCreateResponse) SetCreatedAt(v string) *DestinationCreateResponse {
	o.CreatedAt = v
	return o
}

func (o DestinationCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["destinationID"] = o.DestinationID
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal DestinationCreateResponse: %w", err)
	}

	return serialized, nil
}

func (o DestinationCreateResponse) String() string {
	out := ""
	out += fmt.Sprintf("  destinationID=%v\n", o.DestinationID)
	out += fmt.Sprintf("  name=%v\n", o.Name)
	out += fmt.Sprintf("  createdAt=%v\n", o.CreatedAt)
	return fmt.Sprintf("DestinationCreateResponse {\n%s}", out)
}

type NullableDestinationCreateResponse struct {
	value *DestinationCreateResponse
	isSet bool
}

func (v NullableDestinationCreateResponse) Get() *DestinationCreateResponse {
	return v.value
}

func (v *NullableDestinationCreateResponse) Set(val *DestinationCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDestinationCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDestinationCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestinationCreateResponse(val *DestinationCreateResponse) *NullableDestinationCreateResponse {
	return &NullableDestinationCreateResponse{value: val, isSet: true}
}

func (v NullableDestinationCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableDestinationCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
