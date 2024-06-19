// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// Condition struct for Condition.
type Condition struct {
	// Query pattern that triggers the rule.  You can use either a literal string, or a special pattern `{facet:ATTRIBUTE}`, where `ATTRIBUTE` is a facet name. The rule is triggered if the query matches the literal string or a value of the specified facet. For example, with `pattern: {facet:genre}`, the rule is triggered when users search for a genre, such as \"comedy\".
	Pattern   *string    `json:"pattern,omitempty"`
	Anchoring *Anchoring `json:"anchoring,omitempty"`
	// Whether the pattern should match plurals, synonyms, and typos.
	Alternatives *bool `json:"alternatives,omitempty"`
	// An additional restriction that only triggers the rule, when the search has the same value as `ruleContexts` parameter. For example, if `context: mobile`, the rule is only triggered when the search request has a matching `ruleContexts: mobile`. A rule context must only contain alphanumeric characters.
	Context *string `json:"context,omitempty"`
	// Filters that trigger the rule.  You can add add filters using the syntax `facet:value` so that the rule is triggered, when the specific filter is selected. You can use `filters` on its own or combine it with the `pattern` parameter.
	Filters *string `json:"filters,omitempty"`
}

type ConditionOption func(f *Condition)

func WithConditionPattern(val string) ConditionOption {
	return func(f *Condition) {
		f.Pattern = &val
	}
}

func WithConditionAnchoring(val Anchoring) ConditionOption {
	return func(f *Condition) {
		f.Anchoring = &val
	}
}

func WithConditionAlternatives(val bool) ConditionOption {
	return func(f *Condition) {
		f.Alternatives = &val
	}
}

func WithConditionContext(val string) ConditionOption {
	return func(f *Condition) {
		f.Context = &val
	}
}

func WithConditionFilters(val string) ConditionOption {
	return func(f *Condition) {
		f.Filters = &val
	}
}

// NewCondition instantiates a new Condition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCondition(opts ...ConditionOption) *Condition {
	this := &Condition{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyCondition return a pointer to an empty Condition object.
func NewEmptyCondition() *Condition {
	return &Condition{}
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *Condition) GetPattern() string {
	if o == nil || o.Pattern == nil {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Condition) GetPatternOk() (*string, bool) {
	if o == nil || o.Pattern == nil {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *Condition) HasPattern() bool {
	if o != nil && o.Pattern != nil {
		return true
	}

	return false
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *Condition) SetPattern(v string) *Condition {
	o.Pattern = &v
	return o
}

// GetAnchoring returns the Anchoring field value if set, zero value otherwise.
func (o *Condition) GetAnchoring() Anchoring {
	if o == nil || o.Anchoring == nil {
		var ret Anchoring
		return ret
	}
	return *o.Anchoring
}

// GetAnchoringOk returns a tuple with the Anchoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Condition) GetAnchoringOk() (*Anchoring, bool) {
	if o == nil || o.Anchoring == nil {
		return nil, false
	}
	return o.Anchoring, true
}

// HasAnchoring returns a boolean if a field has been set.
func (o *Condition) HasAnchoring() bool {
	if o != nil && o.Anchoring != nil {
		return true
	}

	return false
}

// SetAnchoring gets a reference to the given Anchoring and assigns it to the Anchoring field.
func (o *Condition) SetAnchoring(v Anchoring) *Condition {
	o.Anchoring = &v
	return o
}

// GetAlternatives returns the Alternatives field value if set, zero value otherwise.
func (o *Condition) GetAlternatives() bool {
	if o == nil || o.Alternatives == nil {
		var ret bool
		return ret
	}
	return *o.Alternatives
}

// GetAlternativesOk returns a tuple with the Alternatives field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Condition) GetAlternativesOk() (*bool, bool) {
	if o == nil || o.Alternatives == nil {
		return nil, false
	}
	return o.Alternatives, true
}

// HasAlternatives returns a boolean if a field has been set.
func (o *Condition) HasAlternatives() bool {
	if o != nil && o.Alternatives != nil {
		return true
	}

	return false
}

// SetAlternatives gets a reference to the given bool and assigns it to the Alternatives field.
func (o *Condition) SetAlternatives(v bool) *Condition {
	o.Alternatives = &v
	return o
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *Condition) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Condition) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *Condition) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *Condition) SetContext(v string) *Condition {
	o.Context = &v
	return o
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *Condition) GetFilters() string {
	if o == nil || o.Filters == nil {
		var ret string
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Condition) GetFiltersOk() (*string, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *Condition) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given string and assigns it to the Filters field.
func (o *Condition) SetFilters(v string) *Condition {
	o.Filters = &v
	return o
}

func (o Condition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Pattern != nil {
		toSerialize["pattern"] = o.Pattern
	}
	if o.Anchoring != nil {
		toSerialize["anchoring"] = o.Anchoring
	}
	if o.Alternatives != nil {
		toSerialize["alternatives"] = o.Alternatives
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal Condition: %w", err)
	}

	return serialized, nil
}

func (o Condition) String() string {
	out := ""
	out += fmt.Sprintf("  pattern=%v\n", o.Pattern)
	out += fmt.Sprintf("  anchoring=%v\n", o.Anchoring)
	out += fmt.Sprintf("  alternatives=%v\n", o.Alternatives)
	out += fmt.Sprintf("  context=%v\n", o.Context)
	out += fmt.Sprintf("  filters=%v\n", o.Filters)
	return fmt.Sprintf("Condition {\n%s}", out)
}
