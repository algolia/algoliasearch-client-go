// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"reflect"
)

// AttributesToRetrieveOption is a wrapper for an AttributesToRetrieve option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type AttributesToRetrieveOption struct {
	value []string
}

// AttributesToRetrieve wraps the given value into a AttributesToRetrieveOption.
func AttributesToRetrieve(v ...string) *AttributesToRetrieveOption {
	return &AttributesToRetrieveOption{v}
}

// Get retrieves the actual value of the option parameter.
func (o *AttributesToRetrieveOption) Get() []string {
	if o == nil {
		return []string{"*"}
	}
	return o.value
}

// MarshalJSON implements the json.Marshaler interface for
// AttributesToRetrieveOption.
func (o AttributesToRetrieveOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

// UnmarshalJSON implements the json.Unmarshaler interface for
// AttributesToRetrieveOption.
func (o *AttributesToRetrieveOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = []string{"*"}
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

// Equal returns true if the given option is equal to the instance one. In case
// the given option is nil, we checked the instance one is set to the default
// value of the option.
func (o *AttributesToRetrieveOption) Equal(o2 *AttributesToRetrieveOption) bool {
	if o == nil {
		return o2 == nil || reflect.DeepEqual(o2.value, []string{"*"})
	}
	if o2 == nil {
		return o == nil || reflect.DeepEqual(o.value, []string{"*"})
	}
	return reflect.DeepEqual(o.value, o2.value)
}

// AttributesToRetrieveEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func AttributesToRetrieveEqual(o1, o2 *AttributesToRetrieveOption) bool {
	return o1.Equal(o2)
}
