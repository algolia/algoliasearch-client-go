// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package suggestions

import (
	"encoding/json"
	"fmt"
)

// IndexName struct for IndexName
type IndexName struct {
	// Index name to target.
	IndexName string `json:"indexName"`
}

// NewIndexName instantiates a new IndexName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexName(indexName string) *IndexName {
	this := &IndexName{}
	this.IndexName = indexName
	return this
}

// NewIndexNameWithDefaults instantiates a new IndexName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexNameWithDefaults() *IndexName {
	this := &IndexName{}
	return this
}

// GetIndexName returns the IndexName field value
func (o *IndexName) GetIndexName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IndexName
}

// GetIndexNameOk returns a tuple with the IndexName field value
// and a boolean to check if the value has been set.
func (o *IndexName) GetIndexNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndexName, true
}

// SetIndexName sets field value
func (o *IndexName) SetIndexName(v string) {
	o.IndexName = v
}

func (o IndexName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["indexName"] = o.IndexName
	}
	return json.Marshal(toSerialize)
}

func (o IndexName) String() string {
	out := ""
	out += fmt.Sprintf("  indexName=%v\n", o.IndexName)
	return fmt.Sprintf("IndexName {\n%s}", out)
}

type NullableIndexName struct {
	value *IndexName
	isSet bool
}

func (v NullableIndexName) Get() *IndexName {
	return v.value
}

func (v *NullableIndexName) Set(val *IndexName) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexName) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexName(val *IndexName) *NullableIndexName {
	return &NullableIndexName{value: val, isSet: true}
}

func (v NullableIndexName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}