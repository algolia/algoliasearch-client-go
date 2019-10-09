// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"reflect"
)

// ScopesOption is a wrapper for an Scopes option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type ScopesOption struct {
	value []string
}

// Scopes wraps the given value into a ScopesOption.
func Scopes(v ...string) *ScopesOption {
	return &ScopesOption{v}
}

// Get retrieves the actual value of the option parameter.
func (o *ScopesOption) Get() []string {
	if o == nil {
		return []string{}
	}
	return o.value
}

// MarshalJSON implements the json.Marshaler interface for
// ScopesOption.
func (o ScopesOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

// UnmarshalJSON implements the json.Unmarshaler interface for
// ScopesOption.
func (o *ScopesOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = []string{}
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

// Equal returns true if the given option is equal to the instance one. In case
// the given option is nil, we checked the instance one is set to the default
// value of the option.
func (o *ScopesOption) Equal(o2 *ScopesOption) bool {
	if o == nil {
		return o2 == nil || reflect.DeepEqual(o2.value, []string{})
	}
	if o2 == nil {
		return o == nil || reflect.DeepEqual(o.value, []string{})
	}
	return reflect.DeepEqual(o.value, o2.value)
}

// ScopesEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func ScopesEqual(o1, o2 *ScopesOption) bool {
	return o1.Equal(o2)
}
