// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package recommend

import (
	"encoding/json"
	"fmt"
)

// RecommendationsHits struct for RecommendationsHits.
type RecommendationsHits struct {
	Hits []RecommendationsHit `json:"hits"`
	// Text to search for in an index.
	Query *string `json:"query,omitempty"`
	// URL-encoded string of all search parameters.
	Params *string `json:"params,omitempty"`
}

type RecommendationsHitsOption func(f *RecommendationsHits)

func WithRecommendationsHitsQuery(val string) RecommendationsHitsOption {
	return func(f *RecommendationsHits) {
		f.Query = &val
	}
}

func WithRecommendationsHitsParams(val string) RecommendationsHitsOption {
	return func(f *RecommendationsHits) {
		f.Params = &val
	}
}

// NewRecommendationsHits instantiates a new RecommendationsHits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRecommendationsHits(hits []RecommendationsHit, opts ...RecommendationsHitsOption) *RecommendationsHits {
	this := &RecommendationsHits{}
	this.Hits = hits
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyRecommendationsHits return a pointer to an empty RecommendationsHits object.
func NewEmptyRecommendationsHits() *RecommendationsHits {
	return &RecommendationsHits{}
}

// GetHits returns the Hits field value.
func (o *RecommendationsHits) GetHits() []RecommendationsHit {
	if o == nil {
		var ret []RecommendationsHit
		return ret
	}

	return o.Hits
}

// GetHitsOk returns a tuple with the Hits field value
// and a boolean to check if the value has been set.
func (o *RecommendationsHits) GetHitsOk() ([]RecommendationsHit, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hits, true
}

// SetHits sets field value.
func (o *RecommendationsHits) SetHits(v []RecommendationsHit) *RecommendationsHits {
	o.Hits = v
	return o
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *RecommendationsHits) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsHits) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *RecommendationsHits) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *RecommendationsHits) SetQuery(v string) *RecommendationsHits {
	o.Query = &v
	return o
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *RecommendationsHits) GetParams() string {
	if o == nil || o.Params == nil {
		var ret string
		return ret
	}
	return *o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsHits) GetParamsOk() (*string, bool) {
	if o == nil || o.Params == nil {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *RecommendationsHits) HasParams() bool {
	if o != nil && o.Params != nil {
		return true
	}

	return false
}

// SetParams gets a reference to the given string and assigns it to the Params field.
func (o *RecommendationsHits) SetParams(v string) *RecommendationsHits {
	o.Params = &v
	return o
}

func (o RecommendationsHits) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["hits"] = o.Hits
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.Params != nil {
		toSerialize["params"] = o.Params
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal RecommendationsHits: %w", err)
	}

	return serialized, nil
}

func (o RecommendationsHits) String() string {
	out := ""
	out += fmt.Sprintf("  hits=%v\n", o.Hits)
	out += fmt.Sprintf("  query=%v\n", o.Query)
	out += fmt.Sprintf("  params=%v\n", o.Params)
	return fmt.Sprintf("RecommendationsHits {\n%s}", out)
}

type NullableRecommendationsHits struct {
	value *RecommendationsHits
	isSet bool
}

func (v NullableRecommendationsHits) Get() *RecommendationsHits {
	return v.value
}

func (v *NullableRecommendationsHits) Set(val *RecommendationsHits) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommendationsHits) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommendationsHits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommendationsHits(val *RecommendationsHits) *NullableRecommendationsHits {
	return &NullableRecommendationsHits{value: val, isSet: true}
}

func (v NullableRecommendationsHits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableRecommendationsHits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
