// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package insights

import (
	"encoding/json"
	"fmt"
)

// InsightsEvents struct for InsightsEvents.
type InsightsEvents struct {
	// List of click and conversion events.  An event is an object representing a user interaction. Events have attributes that describe the interaction, such as an event name, a type, or a user token.  **All** events must be valid, otherwise the API returns an error.
	Events []EventsItems `json:"events"`
}

// NewInsightsEvents instantiates a new InsightsEvents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInsightsEvents(events []EventsItems) *InsightsEvents {
	this := &InsightsEvents{}
	this.Events = events
	return this
}

// NewEmptyInsightsEvents return a pointer to an empty InsightsEvents object.
func NewEmptyInsightsEvents() *InsightsEvents {
	return &InsightsEvents{}
}

// GetEvents returns the Events field value.
func (o *InsightsEvents) GetEvents() []EventsItems {
	if o == nil {
		var ret []EventsItems
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *InsightsEvents) GetEventsOk() ([]EventsItems, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value.
func (o *InsightsEvents) SetEvents(v []EventsItems) *InsightsEvents {
	o.Events = v
	return o
}

func (o InsightsEvents) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["events"] = o.Events
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal InsightsEvents: %w", err)
	}

	return serialized, nil
}

func (o InsightsEvents) String() string {
	out := ""
	out += fmt.Sprintf("  events=%v\n", o.Events)
	return fmt.Sprintf("InsightsEvents {\n%s}", out)
}

type NullableInsightsEvents struct {
	value *InsightsEvents
	isSet bool
}

func (v NullableInsightsEvents) Get() *InsightsEvents {
	return v.value
}

func (v *NullableInsightsEvents) Set(val *InsightsEvents) {
	v.value = val
	v.isSet = true
}

func (v NullableInsightsEvents) IsSet() bool {
	return v.isSet
}

func (v *NullableInsightsEvents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInsightsEvents(val *InsightsEvents) *NullableInsightsEvents {
	return &NullableInsightsEvents{value: val, isSet: true}
}

func (v NullableInsightsEvents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableInsightsEvents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
