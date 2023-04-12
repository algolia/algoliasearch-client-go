// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// Rule Rule object.
type Rule struct {
	// Unique identifier of the object.
	ObjectID string `json:"objectID"`
	// A list of conditions that should apply to activate a Rule. You can use up to 25 conditions per Rule.
	Conditions  []Condition  `json:"conditions,omitempty"`
	Consequence *Consequence `json:"consequence,omitempty"`
	// This field is intended for Rule management purposes, in particular to ease searching for Rules and presenting them to human readers. It's not interpreted by the API.
	Description *string `json:"description,omitempty"`
	// Whether the Rule is enabled. Disabled Rules remain in the index, but aren't applied at query time.
	Enabled *bool `json:"enabled,omitempty"`
	// By default, Rules are permanently valid. When validity periods are specified, the Rule applies only during those periods; it's ignored the rest of the time. The list must not be empty.
	Validity []TimeRange `json:"validity,omitempty"`
}

type RuleOption func(f *Rule)

func WithRuleConditions(val []Condition) RuleOption {
	return func(f *Rule) {
		f.Conditions = val
	}
}

func WithRuleConsequence(val Consequence) RuleOption {
	return func(f *Rule) {
		f.Consequence = &val
	}
}

func WithRuleDescription(val string) RuleOption {
	return func(f *Rule) {
		f.Description = &val
	}
}

func WithRuleEnabled(val bool) RuleOption {
	return func(f *Rule) {
		f.Enabled = &val
	}
}

func WithRuleValidity(val []TimeRange) RuleOption {
	return func(f *Rule) {
		f.Validity = val
	}
}

// NewRule instantiates a new Rule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRule(objectID string, opts ...RuleOption) *Rule {
	this := &Rule{}
	this.ObjectID = objectID
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewRuleWithDefaults instantiates a new Rule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleWithDefaults() *Rule {
	this := &Rule{}
	var enabled bool = true
	this.Enabled = &enabled
	return this
}

// GetObjectID returns the ObjectID field value
func (o *Rule) GetObjectID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectID
}

// GetObjectIDOk returns a tuple with the ObjectID field value
// and a boolean to check if the value has been set.
func (o *Rule) GetObjectIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectID, true
}

// SetObjectID sets field value
func (o *Rule) SetObjectID(v string) {
	o.ObjectID = v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *Rule) GetConditions() []Condition {
	if o == nil || o.Conditions == nil {
		var ret []Condition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rule) GetConditionsOk() ([]Condition, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *Rule) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []Condition and assigns it to the Conditions field.
func (o *Rule) SetConditions(v []Condition) {
	o.Conditions = v
}

// GetConsequence returns the Consequence field value if set, zero value otherwise.
func (o *Rule) GetConsequence() Consequence {
	if o == nil || o.Consequence == nil {
		var ret Consequence
		return ret
	}
	return *o.Consequence
}

// GetConsequenceOk returns a tuple with the Consequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rule) GetConsequenceOk() (*Consequence, bool) {
	if o == nil || o.Consequence == nil {
		return nil, false
	}
	return o.Consequence, true
}

// HasConsequence returns a boolean if a field has been set.
func (o *Rule) HasConsequence() bool {
	if o != nil && o.Consequence != nil {
		return true
	}

	return false
}

// SetConsequence gets a reference to the given Consequence and assigns it to the Consequence field.
func (o *Rule) SetConsequence(v Consequence) {
	o.Consequence = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Rule) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rule) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Rule) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Rule) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Rule) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rule) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Rule) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Rule) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetValidity returns the Validity field value if set, zero value otherwise.
func (o *Rule) GetValidity() []TimeRange {
	if o == nil || o.Validity == nil {
		var ret []TimeRange
		return ret
	}
	return o.Validity
}

// GetValidityOk returns a tuple with the Validity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rule) GetValidityOk() ([]TimeRange, bool) {
	if o == nil || o.Validity == nil {
		return nil, false
	}
	return o.Validity, true
}

// HasValidity returns a boolean if a field has been set.
func (o *Rule) HasValidity() bool {
	if o != nil && o.Validity != nil {
		return true
	}

	return false
}

// SetValidity gets a reference to the given []TimeRange and assigns it to the Validity field.
func (o *Rule) SetValidity(v []TimeRange) {
	o.Validity = v
}

func (o Rule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["objectID"] = o.ObjectID
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if o.Consequence != nil {
		toSerialize["consequence"] = o.Consequence
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Validity != nil {
		toSerialize["validity"] = o.Validity
	}
	return json.Marshal(toSerialize)
}

func (o Rule) String() string {
	out := ""
	out += fmt.Sprintf("  objectID=%v\n", o.ObjectID)
	out += fmt.Sprintf("  conditions=%v\n", o.Conditions)
	out += fmt.Sprintf("  consequence=%v\n", o.Consequence)
	out += fmt.Sprintf("  description=%v\n", o.Description)
	out += fmt.Sprintf("  enabled=%v\n", o.Enabled)
	out += fmt.Sprintf("  validity=%v\n", o.Validity)
	return fmt.Sprintf("Rule {\n%s}", out)
}

type NullableRule struct {
	value *Rule
	isSet bool
}

func (v NullableRule) Get() *Rule {
	return v.value
}

func (v *NullableRule) Set(val *Rule) {
	v.value = val
	v.isSet = true
}

func (v NullableRule) IsSet() bool {
	return v.isSet
}

func (v *NullableRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRule(val *Rule) *NullableRule {
	return &NullableRule{value: val, isSet: true}
}

func (v NullableRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}