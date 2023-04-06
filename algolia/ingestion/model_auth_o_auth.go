// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// AuthOAuth Authentication input for OAuth login.
type AuthOAuth struct {
	// The OAuth endpoint URL.
	Url string `json:"url"`
	// The clientID.
	ClientId string `json:"client_id"`
	// The secret.
	ClientSecret string `json:"client_secret"`
}

// NewAuthOAuth instantiates a new AuthOAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthOAuth(url string, clientId string, clientSecret string) *AuthOAuth {
	this := &AuthOAuth{}
	this.Url = url
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	return this
}

// NewAuthOAuthWithDefaults instantiates a new AuthOAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthOAuthWithDefaults() *AuthOAuth {
	this := &AuthOAuth{}
	return this
}

// GetUrl returns the Url field value
func (o *AuthOAuth) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *AuthOAuth) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *AuthOAuth) SetUrl(v string) {
	o.Url = v
}

// GetClientId returns the ClientId field value
func (o *AuthOAuth) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *AuthOAuth) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *AuthOAuth) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *AuthOAuth) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *AuthOAuth) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *AuthOAuth) SetClientSecret(v string) {
	o.ClientSecret = v
}

func (o AuthOAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["client_secret"] = o.ClientSecret
	}
	return json.Marshal(toSerialize)
}

func (o AuthOAuth) String() string {
	out := ""
	out += fmt.Sprintf("  url=%v\n", o.Url)
	out += fmt.Sprintf("  client_id=%v\n", o.ClientId)
	out += fmt.Sprintf("  client_secret=%v\n", o.ClientSecret)
	return fmt.Sprintf("AuthOAuth {\n%s}", out)
}

type NullableAuthOAuth struct {
	value *AuthOAuth
	isSet bool
}

func (v NullableAuthOAuth) Get() *AuthOAuth {
	return v.value
}

func (v *NullableAuthOAuth) Set(val *AuthOAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthOAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthOAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthOAuth(val *AuthOAuth) *NullableAuthOAuth {
	return &NullableAuthOAuth{value: val, isSet: true}
}

func (v NullableAuthOAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthOAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
