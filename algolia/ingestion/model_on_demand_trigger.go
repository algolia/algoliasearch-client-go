// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// OnDemandTrigger Trigger information for manually-triggered tasks.
type OnDemandTrigger struct {
	Type OnDemandTriggerType `json:"type"`
	// The last time the scheduled task ran in RFC 3339 format.
	LastRun *string `json:"lastRun,omitempty"`
}

type OnDemandTriggerOption func(f *OnDemandTrigger)

func WithOnDemandTriggerLastRun(val string) OnDemandTriggerOption {
	return func(f *OnDemandTrigger) {
		f.LastRun = &val
	}
}

// NewOnDemandTrigger instantiates a new OnDemandTrigger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOnDemandTrigger(type_ OnDemandTriggerType, opts ...OnDemandTriggerOption) *OnDemandTrigger {
	this := &OnDemandTrigger{}
	this.Type = type_
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyOnDemandTrigger return a pointer to an empty OnDemandTrigger object.
func NewEmptyOnDemandTrigger() *OnDemandTrigger {
	return &OnDemandTrigger{}
}

// GetType returns the Type field value.
func (o *OnDemandTrigger) GetType() OnDemandTriggerType {
	if o == nil {
		var ret OnDemandTriggerType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OnDemandTrigger) GetTypeOk() (*OnDemandTriggerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *OnDemandTrigger) SetType(v OnDemandTriggerType) *OnDemandTrigger {
	o.Type = v
	return o
}

// GetLastRun returns the LastRun field value if set, zero value otherwise.
func (o *OnDemandTrigger) GetLastRun() string {
	if o == nil || o.LastRun == nil {
		var ret string
		return ret
	}
	return *o.LastRun
}

// GetLastRunOk returns a tuple with the LastRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnDemandTrigger) GetLastRunOk() (*string, bool) {
	if o == nil || o.LastRun == nil {
		return nil, false
	}
	return o.LastRun, true
}

// HasLastRun returns a boolean if a field has been set.
func (o *OnDemandTrigger) HasLastRun() bool {
	if o != nil && o.LastRun != nil {
		return true
	}

	return false
}

// SetLastRun gets a reference to the given string and assigns it to the LastRun field.
func (o *OnDemandTrigger) SetLastRun(v string) *OnDemandTrigger {
	o.LastRun = &v
	return o
}

func (o OnDemandTrigger) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	toSerialize["type"] = o.Type
	if o.LastRun != nil {
		toSerialize["lastRun"] = o.LastRun
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal OnDemandTrigger: %w", err)
	}

	return serialized, nil
}

func (o OnDemandTrigger) String() string {
	out := ""
	out += fmt.Sprintf("  type=%v\n", o.Type)
	out += fmt.Sprintf("  lastRun=%v\n", o.LastRun)
	return fmt.Sprintf("OnDemandTrigger {\n%s}", out)
}
