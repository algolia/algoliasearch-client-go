// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// AuthAlgoliaPartial Credentials for authenticating with Algolia.
type AuthAlgoliaPartial struct {
	// Algolia application ID.
	AppID *string `json:"appID,omitempty"`
	// Algolia API key with the ACL: `addObject`, `deleteObject`, `settings`, `editSettings`, `listIndexes`, `deleteIndex`. This field is `null` in the API response.
	ApiKey *string `json:"apiKey,omitempty"`
}

type AuthAlgoliaPartialOption func(f *AuthAlgoliaPartial)

func WithAuthAlgoliaPartialAppID(val string) AuthAlgoliaPartialOption {
	return func(f *AuthAlgoliaPartial) {
		f.AppID = &val
	}
}

func WithAuthAlgoliaPartialApiKey(val string) AuthAlgoliaPartialOption {
	return func(f *AuthAlgoliaPartial) {
		f.ApiKey = &val
	}
}

// NewAuthAlgoliaPartial instantiates a new AuthAlgoliaPartial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAuthAlgoliaPartial(opts ...AuthAlgoliaPartialOption) *AuthAlgoliaPartial {
	this := &AuthAlgoliaPartial{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyAuthAlgoliaPartial return a pointer to an empty AuthAlgoliaPartial object.
func NewEmptyAuthAlgoliaPartial() *AuthAlgoliaPartial {
	return &AuthAlgoliaPartial{}
}

// GetAppID returns the AppID field value if set, zero value otherwise.
func (o *AuthAlgoliaPartial) GetAppID() string {
	if o == nil || o.AppID == nil {
		var ret string
		return ret
	}
	return *o.AppID
}

// GetAppIDOk returns a tuple with the AppID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAlgoliaPartial) GetAppIDOk() (*string, bool) {
	if o == nil || o.AppID == nil {
		return nil, false
	}
	return o.AppID, true
}

// HasAppID returns a boolean if a field has been set.
func (o *AuthAlgoliaPartial) HasAppID() bool {
	if o != nil && o.AppID != nil {
		return true
	}

	return false
}

// SetAppID gets a reference to the given string and assigns it to the AppID field.
func (o *AuthAlgoliaPartial) SetAppID(v string) *AuthAlgoliaPartial {
	o.AppID = &v
	return o
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *AuthAlgoliaPartial) GetApiKey() string {
	if o == nil || o.ApiKey == nil {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAlgoliaPartial) GetApiKeyOk() (*string, bool) {
	if o == nil || o.ApiKey == nil {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *AuthAlgoliaPartial) HasApiKey() bool {
	if o != nil && o.ApiKey != nil {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *AuthAlgoliaPartial) SetApiKey(v string) *AuthAlgoliaPartial {
	o.ApiKey = &v
	return o
}

func (o AuthAlgoliaPartial) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.AppID != nil {
		toSerialize["appID"] = o.AppID
	}
	if o.ApiKey != nil {
		toSerialize["apiKey"] = o.ApiKey
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal AuthAlgoliaPartial: %w", err)
	}

	return serialized, nil
}

func (o AuthAlgoliaPartial) String() string {
	out := ""
	out += fmt.Sprintf("  appID=%v\n", o.AppID)
	out += fmt.Sprintf("  apiKey=%v\n", o.ApiKey)
	return fmt.Sprintf("AuthAlgoliaPartial {\n%s}", out)
}