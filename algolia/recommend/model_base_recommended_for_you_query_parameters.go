// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package recommend

import (
	"encoding/json"
	"fmt"
)

// BaseRecommendedForYouQueryParameters struct for BaseRecommendedForYouQueryParameters.
type BaseRecommendedForYouQueryParameters struct {
	// Unique pseudonymous or anonymous user identifier.  This helps with analytics and click and conversion events. For more information, see [user token](https://www.algolia.com/doc/guides/sending-events/concepts/usertoken/).
	UserToken string `json:"userToken"`
}

// NewBaseRecommendedForYouQueryParameters instantiates a new BaseRecommendedForYouQueryParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBaseRecommendedForYouQueryParameters(userToken string) *BaseRecommendedForYouQueryParameters {
	this := &BaseRecommendedForYouQueryParameters{}
	this.UserToken = userToken
	return this
}

// NewEmptyBaseRecommendedForYouQueryParameters return a pointer to an empty BaseRecommendedForYouQueryParameters object.
func NewEmptyBaseRecommendedForYouQueryParameters() *BaseRecommendedForYouQueryParameters {
	return &BaseRecommendedForYouQueryParameters{}
}

// GetUserToken returns the UserToken field value.
func (o *BaseRecommendedForYouQueryParameters) GetUserToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value
// and a boolean to check if the value has been set.
func (o *BaseRecommendedForYouQueryParameters) GetUserTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserToken, true
}

// SetUserToken sets field value.
func (o *BaseRecommendedForYouQueryParameters) SetUserToken(v string) *BaseRecommendedForYouQueryParameters {
	o.UserToken = v
	return o
}

func (o BaseRecommendedForYouQueryParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["userToken"] = o.UserToken
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal BaseRecommendedForYouQueryParameters: %w", err)
	}

	return serialized, nil
}

func (o BaseRecommendedForYouQueryParameters) String() string {
	out := ""
	out += fmt.Sprintf("  userToken=%v\n", o.UserToken)
	return fmt.Sprintf("BaseRecommendedForYouQueryParameters {\n%s}", out)
}

type NullableBaseRecommendedForYouQueryParameters struct {
	value *BaseRecommendedForYouQueryParameters
	isSet bool
}

func (v NullableBaseRecommendedForYouQueryParameters) Get() *BaseRecommendedForYouQueryParameters {
	return v.value
}

func (v *NullableBaseRecommendedForYouQueryParameters) Set(val *BaseRecommendedForYouQueryParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseRecommendedForYouQueryParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseRecommendedForYouQueryParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseRecommendedForYouQueryParameters(val *BaseRecommendedForYouQueryParameters) *NullableBaseRecommendedForYouQueryParameters {
	return &NullableBaseRecommendedForYouQueryParameters{value: val, isSet: true}
}

func (v NullableBaseRecommendedForYouQueryParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableBaseRecommendedForYouQueryParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
