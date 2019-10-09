// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"reflect"
)

// UserDataOption is a wrapper for an UserData option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type UserDataOption struct {
	value map[string]interface{}
}

// UserData wraps the given value into a UserDataOption.
func UserData(v map[string]interface{}) *UserDataOption {
	return &UserDataOption{v}
}

// Get retrieves the actual value of the option parameter.
func (o *UserDataOption) Get() map[string]interface{} {
	if o == nil {
		return map[string]interface{}{}
	}
	return o.value
}

// MarshalJSON implements the json.Marshaler interface for
// UserDataOption.
func (o UserDataOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

// UnmarshalJSON implements the json.Unmarshaler interface for
// UserDataOption.
func (o *UserDataOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = map[string]interface{}{}
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

// Equal returns true if the given option is equal to the instance one. In case
// the given option is nil, we checked the instance one is set to the default
// value of the option.
func (o *UserDataOption) Equal(o2 *UserDataOption) bool {
	if o == nil {
		return o2 == nil || reflect.DeepEqual(o2.value, map[string]interface{}{})
	}
	if o2 == nil {
		return o == nil || reflect.DeepEqual(o.value, map[string]interface{}{})
	}
	return reflect.DeepEqual(o.value, o2.value)
}

// UserDataEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func UserDataEqual(o1, o2 *UserDataOption) bool {
	return o1.Equal(o2)
}
