// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// AuthenticationUpdateResponse Response from the API when the Authentication is successfully updated.
type AuthenticationUpdateResponse struct {
	// The authentication UUID.
	AuthenticationID string `json:"authenticationID"`
	// An human readable name describing the object.
	Name string `json:"name"`
	// Date of last update (RFC3339 format).
	UpdatedAt string `json:"updatedAt"`
}

// NewAuthenticationUpdateResponse instantiates a new AuthenticationUpdateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAuthenticationUpdateResponse(authenticationID string, name string, updatedAt string) *AuthenticationUpdateResponse {
	this := &AuthenticationUpdateResponse{}
	this.AuthenticationID = authenticationID
	this.Name = name
	this.UpdatedAt = updatedAt
	return this
}

// NewAuthenticationUpdateResponseWithDefaults instantiates a new AuthenticationUpdateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAuthenticationUpdateResponseWithDefaults() *AuthenticationUpdateResponse {
	this := &AuthenticationUpdateResponse{}
	return this
}

// GetAuthenticationID returns the AuthenticationID field value.
func (o *AuthenticationUpdateResponse) GetAuthenticationID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthenticationID
}

// GetAuthenticationIDOk returns a tuple with the AuthenticationID field value
// and a boolean to check if the value has been set.
func (o *AuthenticationUpdateResponse) GetAuthenticationIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationID, true
}

// SetAuthenticationID sets field value.
func (o *AuthenticationUpdateResponse) SetAuthenticationID(v string) {
	o.AuthenticationID = v
}

// GetName returns the Name field value.
func (o *AuthenticationUpdateResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthenticationUpdateResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AuthenticationUpdateResponse) SetName(v string) {
	o.Name = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *AuthenticationUpdateResponse) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AuthenticationUpdateResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *AuthenticationUpdateResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o AuthenticationUpdateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["authenticationID"] = o.AuthenticationID
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationUpdateResponse) String() string {
	out := ""
	out += fmt.Sprintf("  authenticationID=%v\n", o.AuthenticationID)
	out += fmt.Sprintf("  name=%v\n", o.Name)
	out += fmt.Sprintf("  updatedAt=%v\n", o.UpdatedAt)
	return fmt.Sprintf("AuthenticationUpdateResponse {\n%s}", out)
}

type NullableAuthenticationUpdateResponse struct {
	value *AuthenticationUpdateResponse
	isSet bool
}

func (v NullableAuthenticationUpdateResponse) Get() *AuthenticationUpdateResponse {
	return v.value
}

func (v *NullableAuthenticationUpdateResponse) Set(val *AuthenticationUpdateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationUpdateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationUpdateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationUpdateResponse(val *AuthenticationUpdateResponse) *NullableAuthenticationUpdateResponse {
	return &NullableAuthenticationUpdateResponse{value: val, isSet: true}
}

func (v NullableAuthenticationUpdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationUpdateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
