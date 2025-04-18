// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// AuthenticationUpdateResponse API response for a successful update of an authentication resource.
type AuthenticationUpdateResponse struct {
	// Universally unique identifier (UUID) of an authentication resource.
	AuthenticationID string `json:"authenticationID"`
	// Descriptive name for the resource.
	Name string `json:"name"`
	// Date of last update in RFC 3339 format.
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

// NewEmptyAuthenticationUpdateResponse return a pointer to an empty AuthenticationUpdateResponse object.
func NewEmptyAuthenticationUpdateResponse() *AuthenticationUpdateResponse {
	return &AuthenticationUpdateResponse{}
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
func (o *AuthenticationUpdateResponse) SetAuthenticationID(v string) *AuthenticationUpdateResponse {
	o.AuthenticationID = v
	return o
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
func (o *AuthenticationUpdateResponse) SetName(v string) *AuthenticationUpdateResponse {
	o.Name = v
	return o
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
func (o *AuthenticationUpdateResponse) SetUpdatedAt(v string) *AuthenticationUpdateResponse {
	o.UpdatedAt = v
	return o
}

func (o AuthenticationUpdateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	toSerialize["authenticationID"] = o.AuthenticationID
	toSerialize["name"] = o.Name
	toSerialize["updatedAt"] = o.UpdatedAt
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal AuthenticationUpdateResponse: %w", err)
	}

	return serialized, nil
}

func (o AuthenticationUpdateResponse) String() string {
	out := ""
	out += fmt.Sprintf("  authenticationID=%v\n", o.AuthenticationID)
	out += fmt.Sprintf("  name=%v\n", o.Name)
	out += fmt.Sprintf("  updatedAt=%v\n", o.UpdatedAt)
	return fmt.Sprintf("AuthenticationUpdateResponse {\n%s}", out)
}
