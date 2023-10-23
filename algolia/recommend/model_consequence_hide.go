// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package recommend

import (
	"encoding/json"
	"fmt"
)

// ConsequenceHide Unique identifier of the record to hide.
type ConsequenceHide struct {
	// Unique object identifier.
	ObjectID string `json:"objectID"`
}

// NewConsequenceHide instantiates a new ConsequenceHide object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsequenceHide(objectID string) *ConsequenceHide {
	this := &ConsequenceHide{}
	this.ObjectID = objectID
	return this
}

// NewConsequenceHideWithDefaults instantiates a new ConsequenceHide object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsequenceHideWithDefaults() *ConsequenceHide {
	this := &ConsequenceHide{}
	return this
}

// GetObjectID returns the ObjectID field value
func (o *ConsequenceHide) GetObjectID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectID
}

// GetObjectIDOk returns a tuple with the ObjectID field value
// and a boolean to check if the value has been set.
func (o *ConsequenceHide) GetObjectIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectID, true
}

// SetObjectID sets field value
func (o *ConsequenceHide) SetObjectID(v string) {
	o.ObjectID = v
}

func (o ConsequenceHide) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["objectID"] = o.ObjectID
	}
	return json.Marshal(toSerialize)
}

func (o ConsequenceHide) String() string {
	out := ""
	out += fmt.Sprintf("  objectID=%v\n", o.ObjectID)
	return fmt.Sprintf("ConsequenceHide {\n%s}", out)
}

type NullableConsequenceHide struct {
	value *ConsequenceHide
	isSet bool
}

func (v NullableConsequenceHide) Get() *ConsequenceHide {
	return v.value
}

func (v *NullableConsequenceHide) Set(val *ConsequenceHide) {
	v.value = val
	v.isSet = true
}

func (v NullableConsequenceHide) IsSet() bool {
	return v.isSet
}

func (v *NullableConsequenceHide) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsequenceHide(val *ConsequenceHide) *NullableConsequenceHide {
	return &NullableConsequenceHide{value: val, isSet: true}
}

func (v NullableConsequenceHide) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsequenceHide) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
