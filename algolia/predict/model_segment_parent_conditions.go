// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package predict

import (
	"encoding/json"
	"fmt"
)

// SegmentParentConditions The conditions that define which user profiles are included in the segment.  Can contain operands and a maximum of 1 level of nested conditions.
type SegmentParentConditions struct {
	Operator SegmentConditionOperator         `json:"operator"`
	Operands []SegmentParentConditionOperands `json:"operands"`
}

// NewSegmentParentConditions instantiates a new SegmentParentConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSegmentParentConditions(operator SegmentConditionOperator, operands []SegmentParentConditionOperands) *SegmentParentConditions {
	this := &SegmentParentConditions{}
	this.Operator = operator
	this.Operands = operands
	return this
}

// NewSegmentParentConditionsWithDefaults instantiates a new SegmentParentConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSegmentParentConditionsWithDefaults() *SegmentParentConditions {
	this := &SegmentParentConditions{}
	return this
}

// GetOperator returns the Operator field value
func (o *SegmentParentConditions) GetOperator() SegmentConditionOperator {
	if o == nil {
		var ret SegmentConditionOperator
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *SegmentParentConditions) GetOperatorOk() (*SegmentConditionOperator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *SegmentParentConditions) SetOperator(v SegmentConditionOperator) {
	o.Operator = v
}

// GetOperands returns the Operands field value
func (o *SegmentParentConditions) GetOperands() []SegmentParentConditionOperands {
	if o == nil {
		var ret []SegmentParentConditionOperands
		return ret
	}

	return o.Operands
}

// GetOperandsOk returns a tuple with the Operands field value
// and a boolean to check if the value has been set.
func (o *SegmentParentConditions) GetOperandsOk() ([]SegmentParentConditionOperands, bool) {
	if o == nil {
		return nil, false
	}
	return o.Operands, true
}

// SetOperands sets field value
func (o *SegmentParentConditions) SetOperands(v []SegmentParentConditionOperands) {
	o.Operands = v
}

func (o SegmentParentConditions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["operator"] = o.Operator
	}
	if true {
		toSerialize["operands"] = o.Operands
	}
	return json.Marshal(toSerialize)
}

func (o SegmentParentConditions) String() string {
	out := ""
	out += fmt.Sprintf("  operator=%v\n", o.Operator)
	out += fmt.Sprintf("  operands=%v\n", o.Operands)
	return fmt.Sprintf("SegmentParentConditions {\n%s}", out)
}

type NullableSegmentParentConditions struct {
	value *SegmentParentConditions
	isSet bool
}

func (v NullableSegmentParentConditions) Get() *SegmentParentConditions {
	return v.value
}

func (v *NullableSegmentParentConditions) Set(val *SegmentParentConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableSegmentParentConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableSegmentParentConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSegmentParentConditions(val *SegmentParentConditions) *NullableSegmentParentConditions {
	return &NullableSegmentParentConditions{value: val, isSet: true}
}

func (v NullableSegmentParentConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSegmentParentConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}