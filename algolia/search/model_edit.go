// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// Edit struct for Edit
type Edit struct {
	Type *EditType `json:"type,omitempty"`
	// Text or patterns to remove from the query string.
	Delete *string `json:"delete,omitempty"`
	// Text that should be inserted in place of the removed text inside the query string.
	Insert *string `json:"insert,omitempty"`
}

type EditOption func(f *Edit)

func WithEditType(val EditType) EditOption {
	return func(f *Edit) {
		f.Type = &val
	}
}

func WithEditDelete(val string) EditOption {
	return func(f *Edit) {
		f.Delete = &val
	}
}

func WithEditInsert(val string) EditOption {
	return func(f *Edit) {
		f.Insert = &val
	}
}

// NewEdit instantiates a new Edit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdit(opts ...EditOption) *Edit {
	this := &Edit{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEditWithDefaults instantiates a new Edit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditWithDefaults() *Edit {
	this := &Edit{}
	return this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Edit) GetType() EditType {
	if o == nil || o.Type == nil {
		var ret EditType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Edit) GetTypeOk() (*EditType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Edit) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given EditType and assigns it to the Type field.
func (o *Edit) SetType(v EditType) {
	o.Type = &v
}

// GetDelete returns the Delete field value if set, zero value otherwise.
func (o *Edit) GetDelete() string {
	if o == nil || o.Delete == nil {
		var ret string
		return ret
	}
	return *o.Delete
}

// GetDeleteOk returns a tuple with the Delete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Edit) GetDeleteOk() (*string, bool) {
	if o == nil || o.Delete == nil {
		return nil, false
	}
	return o.Delete, true
}

// HasDelete returns a boolean if a field has been set.
func (o *Edit) HasDelete() bool {
	if o != nil && o.Delete != nil {
		return true
	}

	return false
}

// SetDelete gets a reference to the given string and assigns it to the Delete field.
func (o *Edit) SetDelete(v string) {
	o.Delete = &v
}

// GetInsert returns the Insert field value if set, zero value otherwise.
func (o *Edit) GetInsert() string {
	if o == nil || o.Insert == nil {
		var ret string
		return ret
	}
	return *o.Insert
}

// GetInsertOk returns a tuple with the Insert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Edit) GetInsertOk() (*string, bool) {
	if o == nil || o.Insert == nil {
		return nil, false
	}
	return o.Insert, true
}

// HasInsert returns a boolean if a field has been set.
func (o *Edit) HasInsert() bool {
	if o != nil && o.Insert != nil {
		return true
	}

	return false
}

// SetInsert gets a reference to the given string and assigns it to the Insert field.
func (o *Edit) SetInsert(v string) {
	o.Insert = &v
}

func (o Edit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Delete != nil {
		toSerialize["delete"] = o.Delete
	}
	if o.Insert != nil {
		toSerialize["insert"] = o.Insert
	}
	return json.Marshal(toSerialize)
}

func (o Edit) String() string {
	out := ""
	out += fmt.Sprintf("  type=%v\n", o.Type)
	out += fmt.Sprintf("  delete=%v\n", o.Delete)
	out += fmt.Sprintf("  insert=%v\n", o.Insert)
	return fmt.Sprintf("Edit {\n%s}", out)
}

type NullableEdit struct {
	value *Edit
	isSet bool
}

func (v NullableEdit) Get() *Edit {
	return v.value
}

func (v *NullableEdit) Set(val *Edit) {
	v.value = val
	v.isSet = true
}

func (v NullableEdit) IsSet() bool {
	return v.isSet
}

func (v *NullableEdit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdit(val *Edit) *NullableEdit {
	return &NullableEdit{value: val, isSet: true}
}

func (v NullableEdit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
