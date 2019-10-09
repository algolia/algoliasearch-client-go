// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"reflect"
)

// AttributesToHighlightOption is a wrapper for an AttributesToHighlight option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type AttributesToHighlightOption struct {
	value []string
}

// AttributesToHighlight wraps the given value into a AttributesToHighlightOption.
func AttributesToHighlight(v ...string) *AttributesToHighlightOption {
	return &AttributesToHighlightOption{v}
}

// Get retrieves the actual value of the option parameter.
func (o *AttributesToHighlightOption) Get() []string {
	if o == nil {
		return []string{}
	}
	return o.value
}

// MarshalJSON implements the json.Marshaler interface for
// AttributesToHighlightOption.
func (o AttributesToHighlightOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

// UnmarshalJSON implements the json.Unmarshaler interface for
// AttributesToHighlightOption.
func (o *AttributesToHighlightOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = []string{}
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

// Equal returns true if the given option is equal to the instance one. In case
// the given option is nil, we checked the instance one is set to the default
// value of the option.
func (o *AttributesToHighlightOption) Equal(o2 *AttributesToHighlightOption) bool {
	if o == nil {
		return o2 == nil || reflect.DeepEqual(o2.value, []string{})
	}
	if o2 == nil {
		return o == nil || reflect.DeepEqual(o.value, []string{})
	}
	return reflect.DeepEqual(o.value, o2.value)
}

// AttributesToHighlightEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func AttributesToHighlightEqual(o1, o2 *AttributesToHighlightOption) bool {
	return o1.Equal(o2)
}
