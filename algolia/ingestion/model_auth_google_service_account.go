// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// AuthGoogleServiceAccount Credentials for authenticating with a Google service account, such as BigQuery.
type AuthGoogleServiceAccount struct {
	// Email address of the Google service account.
	ClientEmail string `json:"clientEmail"`
	// Private key of the Google service account. This field is `null` in the API response.
	PrivateKey string `json:"privateKey"`
}

// NewAuthGoogleServiceAccount instantiates a new AuthGoogleServiceAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAuthGoogleServiceAccount(clientEmail string, privateKey string) *AuthGoogleServiceAccount {
	this := &AuthGoogleServiceAccount{}
	this.ClientEmail = clientEmail
	this.PrivateKey = privateKey
	return this
}

// NewEmptyAuthGoogleServiceAccount return a pointer to an empty AuthGoogleServiceAccount object.
func NewEmptyAuthGoogleServiceAccount() *AuthGoogleServiceAccount {
	return &AuthGoogleServiceAccount{}
}

// GetClientEmail returns the ClientEmail field value.
func (o *AuthGoogleServiceAccount) GetClientEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientEmail
}

// GetClientEmailOk returns a tuple with the ClientEmail field value
// and a boolean to check if the value has been set.
func (o *AuthGoogleServiceAccount) GetClientEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientEmail, true
}

// SetClientEmail sets field value.
func (o *AuthGoogleServiceAccount) SetClientEmail(v string) *AuthGoogleServiceAccount {
	o.ClientEmail = v
	return o
}

// GetPrivateKey returns the PrivateKey field value.
func (o *AuthGoogleServiceAccount) GetPrivateKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value
// and a boolean to check if the value has been set.
func (o *AuthGoogleServiceAccount) GetPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateKey, true
}

// SetPrivateKey sets field value.
func (o *AuthGoogleServiceAccount) SetPrivateKey(v string) *AuthGoogleServiceAccount {
	o.PrivateKey = v
	return o
}

func (o AuthGoogleServiceAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["clientEmail"] = o.ClientEmail
	}
	if true {
		toSerialize["privateKey"] = o.PrivateKey
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal AuthGoogleServiceAccount: %w", err)
	}

	return serialized, nil
}

func (o AuthGoogleServiceAccount) String() string {
	out := ""
	out += fmt.Sprintf("  clientEmail=%v\n", o.ClientEmail)
	out += fmt.Sprintf("  privateKey=%v\n", o.PrivateKey)
	return fmt.Sprintf("AuthGoogleServiceAccount {\n%s}", out)
}

type NullableAuthGoogleServiceAccount struct {
	value *AuthGoogleServiceAccount
	isSet bool
}

func (v NullableAuthGoogleServiceAccount) Get() *AuthGoogleServiceAccount {
	return v.value
}

func (v *NullableAuthGoogleServiceAccount) Set(val *AuthGoogleServiceAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthGoogleServiceAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthGoogleServiceAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthGoogleServiceAccount(val *AuthGoogleServiceAccount) *NullableAuthGoogleServiceAccount {
	return &NullableAuthGoogleServiceAccount{value: val, isSet: true}
}

func (v NullableAuthGoogleServiceAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableAuthGoogleServiceAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
