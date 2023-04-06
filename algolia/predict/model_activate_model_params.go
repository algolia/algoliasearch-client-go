// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package predict

import (
	"encoding/json"
	"fmt"
)

// ActivateModelParams struct for ActivateModelParams
type ActivateModelParams struct {
	Type ModelsToRetrieve `json:"type"`
	// The model’s instance name.
	Name string `json:"name"`
	// The data source ID, as returned by the (external) sources API.
	SourceID string `json:"sourceID"`
	// The index name.
	Index           string   `json:"index"`
	ModelAttributes []string `json:"modelAttributes,omitempty"`
}

type ActivateModelParamsOption func(f *ActivateModelParams)

func WithActivateModelParamsModelAttributes(val []string) ActivateModelParamsOption {
	return func(f *ActivateModelParams) {
		f.ModelAttributes = val
	}
}

// NewActivateModelParams instantiates a new ActivateModelParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivateModelParams(type_ ModelsToRetrieve, name string, sourceID string, index string, opts ...ActivateModelParamsOption) *ActivateModelParams {
	this := &ActivateModelParams{}
	this.Type = type_
	this.Name = name
	this.SourceID = sourceID
	this.Index = index
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewActivateModelParamsWithDefaults instantiates a new ActivateModelParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivateModelParamsWithDefaults() *ActivateModelParams {
	this := &ActivateModelParams{}
	return this
}

// GetType returns the Type field value
func (o *ActivateModelParams) GetType() ModelsToRetrieve {
	if o == nil {
		var ret ModelsToRetrieve
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ActivateModelParams) GetTypeOk() (*ModelsToRetrieve, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ActivateModelParams) SetType(v ModelsToRetrieve) {
	o.Type = v
}

// GetName returns the Name field value
func (o *ActivateModelParams) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ActivateModelParams) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ActivateModelParams) SetName(v string) {
	o.Name = v
}

// GetSourceID returns the SourceID field value
func (o *ActivateModelParams) GetSourceID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceID
}

// GetSourceIDOk returns a tuple with the SourceID field value
// and a boolean to check if the value has been set.
func (o *ActivateModelParams) GetSourceIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceID, true
}

// SetSourceID sets field value
func (o *ActivateModelParams) SetSourceID(v string) {
	o.SourceID = v
}

// GetIndex returns the Index field value
func (o *ActivateModelParams) GetIndex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *ActivateModelParams) GetIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *ActivateModelParams) SetIndex(v string) {
	o.Index = v
}

// GetModelAttributes returns the ModelAttributes field value if set, zero value otherwise.
func (o *ActivateModelParams) GetModelAttributes() []string {
	if o == nil || o.ModelAttributes == nil {
		var ret []string
		return ret
	}
	return o.ModelAttributes
}

// GetModelAttributesOk returns a tuple with the ModelAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivateModelParams) GetModelAttributesOk() ([]string, bool) {
	if o == nil || o.ModelAttributes == nil {
		return nil, false
	}
	return o.ModelAttributes, true
}

// HasModelAttributes returns a boolean if a field has been set.
func (o *ActivateModelParams) HasModelAttributes() bool {
	if o != nil && o.ModelAttributes != nil {
		return true
	}

	return false
}

// SetModelAttributes gets a reference to the given []string and assigns it to the ModelAttributes field.
func (o *ActivateModelParams) SetModelAttributes(v []string) {
	o.ModelAttributes = v
}

func (o ActivateModelParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["sourceID"] = o.SourceID
	}
	if true {
		toSerialize["index"] = o.Index
	}
	if o.ModelAttributes != nil {
		toSerialize["modelAttributes"] = o.ModelAttributes
	}
	return json.Marshal(toSerialize)
}

func (o ActivateModelParams) String() string {
	out := ""
	out += fmt.Sprintf("  type=%v\n", o.Type)
	out += fmt.Sprintf("  name=%v\n", o.Name)
	out += fmt.Sprintf("  sourceID=%v\n", o.SourceID)
	out += fmt.Sprintf("  index=%v\n", o.Index)
	out += fmt.Sprintf("  modelAttributes=%v\n", o.ModelAttributes)
	return fmt.Sprintf("ActivateModelParams {\n%s}", out)
}

type NullableActivateModelParams struct {
	value *ActivateModelParams
	isSet bool
}

func (v NullableActivateModelParams) Get() *ActivateModelParams {
	return v.value
}

func (v *NullableActivateModelParams) Set(val *ActivateModelParams) {
	v.value = val
	v.isSet = true
}

func (v NullableActivateModelParams) IsSet() bool {
	return v.isSet
}

func (v *NullableActivateModelParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivateModelParams(val *ActivateModelParams) *NullableActivateModelParams {
	return &NullableActivateModelParams{value: val, isSet: true}
}

func (v NullableActivateModelParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivateModelParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
