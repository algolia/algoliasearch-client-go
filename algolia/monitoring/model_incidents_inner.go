// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package monitoring

import (
	"encoding/json"
	"fmt"
)

// IncidentsInner struct for IncidentsInner.
type IncidentsInner struct {
	// Timestamp, measured in milliseconds since the Unix epoch.
	T *int64    `json:"t,omitempty"`
	V *Incident `json:"v,omitempty"`
}

type IncidentsInnerOption func(f *IncidentsInner)

func WithIncidentsInnerT(val int64) IncidentsInnerOption {
	return func(f *IncidentsInner) {
		f.T = &val
	}
}

func WithIncidentsInnerV(val Incident) IncidentsInnerOption {
	return func(f *IncidentsInner) {
		f.V = &val
	}
}

// NewIncidentsInner instantiates a new IncidentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentsInner(opts ...IncidentsInnerOption) *IncidentsInner {
	this := &IncidentsInner{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyIncidentsInner return a pointer to an empty IncidentsInner object.
func NewEmptyIncidentsInner() *IncidentsInner {
	return &IncidentsInner{}
}

// GetT returns the T field value if set, zero value otherwise.
func (o *IncidentsInner) GetT() int64 {
	if o == nil || o.T == nil {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentsInner) GetTOk() (*int64, bool) {
	if o == nil || o.T == nil {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *IncidentsInner) HasT() bool {
	if o != nil && o.T != nil {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *IncidentsInner) SetT(v int64) *IncidentsInner {
	o.T = &v
	return o
}

// GetV returns the V field value if set, zero value otherwise.
func (o *IncidentsInner) GetV() Incident {
	if o == nil || o.V == nil {
		var ret Incident
		return ret
	}
	return *o.V
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentsInner) GetVOk() (*Incident, bool) {
	if o == nil || o.V == nil {
		return nil, false
	}
	return o.V, true
}

// HasV returns a boolean if a field has been set.
func (o *IncidentsInner) HasV() bool {
	if o != nil && o.V != nil {
		return true
	}

	return false
}

// SetV gets a reference to the given Incident and assigns it to the V field.
func (o *IncidentsInner) SetV(v *Incident) *IncidentsInner {
	o.V = v
	return o
}

func (o IncidentsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.T != nil {
		toSerialize["t"] = o.T
	}
	if o.V != nil {
		toSerialize["v"] = o.V
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal IncidentsInner: %w", err)
	}

	return serialized, nil
}

func (o IncidentsInner) String() string {
	out := ""
	out += fmt.Sprintf("  t=%v\n", o.T)
	out += fmt.Sprintf("  v=%v\n", o.V)
	return fmt.Sprintf("IncidentsInner {\n%s}", out)
}
