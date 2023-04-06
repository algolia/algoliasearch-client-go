// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// AuthenticationWithInput struct for AuthenticationWithInput
type AuthenticationWithInput struct {
	// The authentication UUID.
	AuthenticationID string             `json:"authenticationID"`
	Type             AuthenticationType `json:"type"`
	// An human readable name describing the object.
	Name     string    `json:"name"`
	Platform *Platform `json:"platform,omitempty"`
	// Date of creation (RFC3339 format).
	CreatedAt string `json:"createdAt"`
	// Date of last update (RFC3339 format).
	UpdatedAt *string   `json:"updatedAt,omitempty"`
	Input     AuthInput `json:"input"`
}

type AuthenticationWithInputOption func(f *AuthenticationWithInput)

func WithAuthenticationWithInputPlatform(val Platform) AuthenticationWithInputOption {
	return func(f *AuthenticationWithInput) {
		f.Platform = &val
	}
}

func WithAuthenticationWithInputUpdatedAt(val string) AuthenticationWithInputOption {
	return func(f *AuthenticationWithInput) {
		f.UpdatedAt = &val
	}
}

// NewAuthenticationWithInput instantiates a new AuthenticationWithInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationWithInput(authenticationID string, type_ AuthenticationType, name string, createdAt string, input AuthInput, opts ...AuthenticationWithInputOption) *AuthenticationWithInput {
	this := &AuthenticationWithInput{}
	this.AuthenticationID = authenticationID
	this.Type = type_
	this.Name = name
	this.CreatedAt = createdAt
	this.Input = input
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewAuthenticationWithInputWithDefaults instantiates a new AuthenticationWithInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationWithInputWithDefaults() *AuthenticationWithInput {
	this := &AuthenticationWithInput{}
	return this
}

// GetAuthenticationID returns the AuthenticationID field value
func (o *AuthenticationWithInput) GetAuthenticationID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthenticationID
}

// GetAuthenticationIDOk returns a tuple with the AuthenticationID field value
// and a boolean to check if the value has been set.
func (o *AuthenticationWithInput) GetAuthenticationIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationID, true
}

// SetAuthenticationID sets field value
func (o *AuthenticationWithInput) SetAuthenticationID(v string) {
	o.AuthenticationID = v
}

// GetType returns the Type field value
func (o *AuthenticationWithInput) GetType() AuthenticationType {
	if o == nil {
		var ret AuthenticationType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AuthenticationWithInput) GetTypeOk() (*AuthenticationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AuthenticationWithInput) SetType(v AuthenticationType) {
	o.Type = v
}

// GetName returns the Name field value
func (o *AuthenticationWithInput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthenticationWithInput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuthenticationWithInput) SetName(v string) {
	o.Name = v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *AuthenticationWithInput) GetPlatform() Platform {
	if o == nil || o.Platform == nil {
		var ret Platform
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationWithInput) GetPlatformOk() (*Platform, bool) {
	if o == nil || o.Platform == nil {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *AuthenticationWithInput) HasPlatform() bool {
	if o != nil && o.Platform != nil {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given Platform and assigns it to the Platform field.
func (o *AuthenticationWithInput) SetPlatform(v Platform) {
	o.Platform = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *AuthenticationWithInput) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AuthenticationWithInput) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AuthenticationWithInput) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AuthenticationWithInput) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationWithInput) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AuthenticationWithInput) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *AuthenticationWithInput) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetInput returns the Input field value
func (o *AuthenticationWithInput) GetInput() AuthInput {
	if o == nil {
		var ret AuthInput
		return ret
	}

	return o.Input
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
func (o *AuthenticationWithInput) GetInputOk() (*AuthInput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Input, true
}

// SetInput sets field value
func (o *AuthenticationWithInput) SetInput(v AuthInput) {
	o.Input = v
}

func (o AuthenticationWithInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["authenticationID"] = o.AuthenticationID
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Platform != nil {
		toSerialize["platform"] = o.Platform
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if true {
		toSerialize["input"] = o.Input
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationWithInput) String() string {
	out := ""
	out += fmt.Sprintf("  authenticationID=%v\n", o.AuthenticationID)
	out += fmt.Sprintf("  type=%v\n", o.Type)
	out += fmt.Sprintf("  name=%v\n", o.Name)
	out += fmt.Sprintf("  platform=%v\n", o.Platform)
	out += fmt.Sprintf("  createdAt=%v\n", o.CreatedAt)
	out += fmt.Sprintf("  updatedAt=%v\n", o.UpdatedAt)
	out += fmt.Sprintf("  input=%v\n", o.Input)
	return fmt.Sprintf("AuthenticationWithInput {\n%s}", out)
}

type NullableAuthenticationWithInput struct {
	value *AuthenticationWithInput
	isSet bool
}

func (v NullableAuthenticationWithInput) Get() *AuthenticationWithInput {
	return v.value
}

func (v *NullableAuthenticationWithInput) Set(val *AuthenticationWithInput) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationWithInput) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationWithInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationWithInput(val *AuthenticationWithInput) *NullableAuthenticationWithInput {
	return &NullableAuthenticationWithInput{value: val, isSet: true}
}

func (v NullableAuthenticationWithInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationWithInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
