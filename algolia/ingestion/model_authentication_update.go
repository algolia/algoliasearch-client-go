// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// AuthenticationUpdate Request body for updating an authentication resource.
type AuthenticationUpdate struct {
	Type *AuthenticationType `json:"type,omitempty"`
	// Descriptive name for the resource.
	Name     *string           `json:"name,omitempty"`
	Platform NullablePlatform  `json:"platform,omitempty"`
	Input    *AuthInputPartial `json:"input,omitempty"`
}

type AuthenticationUpdateOption func(f *AuthenticationUpdate)

func WithAuthenticationUpdateType(val AuthenticationType) AuthenticationUpdateOption {
	return func(f *AuthenticationUpdate) {
		f.Type = &val
	}
}

func WithAuthenticationUpdateName(val string) AuthenticationUpdateOption {
	return func(f *AuthenticationUpdate) {
		f.Name = &val
	}
}

func WithAuthenticationUpdatePlatform(val NullablePlatform) AuthenticationUpdateOption {
	return func(f *AuthenticationUpdate) {
		f.Platform = val
	}
}

func WithAuthenticationUpdateInput(val AuthInputPartial) AuthenticationUpdateOption {
	return func(f *AuthenticationUpdate) {
		f.Input = &val
	}
}

// NewAuthenticationUpdate instantiates a new AuthenticationUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAuthenticationUpdate(opts ...AuthenticationUpdateOption) *AuthenticationUpdate {
	this := &AuthenticationUpdate{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyAuthenticationUpdate return a pointer to an empty AuthenticationUpdate object.
func NewEmptyAuthenticationUpdate() *AuthenticationUpdate {
	return &AuthenticationUpdate{}
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AuthenticationUpdate) GetType() AuthenticationType {
	if o == nil || o.Type == nil {
		var ret AuthenticationType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationUpdate) GetTypeOk() (*AuthenticationType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AuthenticationUpdate) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given AuthenticationType and assigns it to the Type field.
func (o *AuthenticationUpdate) SetType(v AuthenticationType) *AuthenticationUpdate {
	o.Type = &v
	return o
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AuthenticationUpdate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AuthenticationUpdate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AuthenticationUpdate) SetName(v string) *AuthenticationUpdate {
	o.Name = &v
	return o
}

// GetPlatform returns the Platform field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticationUpdate) GetPlatform() Platform {
	if o == nil || o.Platform.Get() == nil {
		var ret Platform
		return ret
	}
	return *o.Platform.Get()
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AuthenticationUpdate) GetPlatformOk() (*Platform, bool) {
	if o == nil {
		return nil, false
	}
	return o.Platform.Get(), o.Platform.IsSet()
}

// HasPlatform returns a boolean if a field has been set.
func (o *AuthenticationUpdate) HasPlatform() bool {
	if o != nil && o.Platform.IsSet() {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given NullablePlatform and assigns it to the Platform field.
func (o *AuthenticationUpdate) SetPlatform(v Platform) *AuthenticationUpdate {
	o.Platform.Set(&v)
	return o
}

// SetPlatformNil sets the value for Platform to be an explicit nil.
func (o *AuthenticationUpdate) SetPlatformNil() {
	o.Platform.Set(nil)
}

// UnsetPlatform ensures that no value is present for Platform, not even an explicit nil.
func (o *AuthenticationUpdate) UnsetPlatform() {
	o.Platform.Unset()
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *AuthenticationUpdate) GetInput() AuthInputPartial {
	if o == nil || o.Input == nil {
		var ret AuthInputPartial
		return ret
	}
	return *o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationUpdate) GetInputOk() (*AuthInputPartial, bool) {
	if o == nil || o.Input == nil {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *AuthenticationUpdate) HasInput() bool {
	if o != nil && o.Input != nil {
		return true
	}

	return false
}

// SetInput gets a reference to the given AuthInputPartial and assigns it to the Input field.
func (o *AuthenticationUpdate) SetInput(v *AuthInputPartial) *AuthenticationUpdate {
	o.Input = v
	return o
}

func (o AuthenticationUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Platform.IsSet() {
		toSerialize["platform"] = o.Platform.Get()
	}
	if o.Input != nil {
		toSerialize["input"] = o.Input
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal AuthenticationUpdate: %w", err)
	}

	return serialized, nil
}

func (o AuthenticationUpdate) String() string {
	out := ""
	out += fmt.Sprintf("  type=%v\n", o.Type)
	out += fmt.Sprintf("  name=%v\n", o.Name)
	out += fmt.Sprintf("  platform=%v\n", o.Platform)
	out += fmt.Sprintf("  input=%v\n", o.Input)
	return fmt.Sprintf("AuthenticationUpdate {\n%s}", out)
}

type NullableAuthenticationUpdate struct {
	value *AuthenticationUpdate
	isSet bool
}

func (v NullableAuthenticationUpdate) Get() *AuthenticationUpdate {
	return v.value
}

func (v *NullableAuthenticationUpdate) Set(val *AuthenticationUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationUpdate(val *AuthenticationUpdate) *NullableAuthenticationUpdate {
	return &NullableAuthenticationUpdate{value: val, isSet: true}
}

func (v NullableAuthenticationUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableAuthenticationUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
