// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// AuthOAuthPartial Credentials for authenticating with OAuth 2.0.
type AuthOAuthPartial struct {
	// URL for the OAuth endpoint.
	Url *string `json:"url,omitempty"`
	// Client ID.
	ClientId *string `json:"client_id,omitempty"`
	// Client secret. This field is `null` in the API response.
	ClientSecret *string `json:"client_secret,omitempty"`
	// OAuth scope.
	Scope *string `json:"scope,omitempty"`
}

type AuthOAuthPartialOption func(f *AuthOAuthPartial)

func WithAuthOAuthPartialUrl(val string) AuthOAuthPartialOption {
	return func(f *AuthOAuthPartial) {
		f.Url = &val
	}
}

func WithAuthOAuthPartialClientId(val string) AuthOAuthPartialOption {
	return func(f *AuthOAuthPartial) {
		f.ClientId = &val
	}
}

func WithAuthOAuthPartialClientSecret(val string) AuthOAuthPartialOption {
	return func(f *AuthOAuthPartial) {
		f.ClientSecret = &val
	}
}

func WithAuthOAuthPartialScope(val string) AuthOAuthPartialOption {
	return func(f *AuthOAuthPartial) {
		f.Scope = &val
	}
}

// NewAuthOAuthPartial instantiates a new AuthOAuthPartial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAuthOAuthPartial(opts ...AuthOAuthPartialOption) *AuthOAuthPartial {
	this := &AuthOAuthPartial{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyAuthOAuthPartial return a pointer to an empty AuthOAuthPartial object.
func NewEmptyAuthOAuthPartial() *AuthOAuthPartial {
	return &AuthOAuthPartial{}
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AuthOAuthPartial) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthOAuthPartial) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AuthOAuthPartial) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AuthOAuthPartial) SetUrl(v string) *AuthOAuthPartial {
	o.Url = &v
	return o
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *AuthOAuthPartial) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthOAuthPartial) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *AuthOAuthPartial) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *AuthOAuthPartial) SetClientId(v string) *AuthOAuthPartial {
	o.ClientId = &v
	return o
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *AuthOAuthPartial) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthOAuthPartial) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *AuthOAuthPartial) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *AuthOAuthPartial) SetClientSecret(v string) *AuthOAuthPartial {
	o.ClientSecret = &v
	return o
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *AuthOAuthPartial) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthOAuthPartial) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *AuthOAuthPartial) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *AuthOAuthPartial) SetScope(v string) *AuthOAuthPartial {
	o.Scope = &v
	return o
}

func (o AuthOAuthPartial) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.ClientSecret != nil {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal AuthOAuthPartial: %w", err)
	}

	return serialized, nil
}

func (o AuthOAuthPartial) String() string {
	out := ""
	out += fmt.Sprintf("  url=%v\n", o.Url)
	out += fmt.Sprintf("  client_id=%v\n", o.ClientId)
	out += fmt.Sprintf("  client_secret=%v\n", o.ClientSecret)
	out += fmt.Sprintf("  scope=%v\n", o.Scope)
	return fmt.Sprintf("AuthOAuthPartial {\n%s}", out)
}