package opt

import (
	"encoding/json"
	"reflect"
	"strings"
)

// TypeOption is a wrapper for an Type option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type TypeOption struct {
	value []string
}

// Type wraps the given value into a TypeOption.
func Type(v ...string) *TypeOption {
	return &TypeOption{v}
}

// Get retrieves the actual value of the option parameter.
func (o *TypeOption) Get() []string {
	if o == nil {
		return []string{}
	}
	return o.value
}

// MarshalJSON implements the json.Marshaler interface for
// TypeOption.
func (o TypeOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(strings.Join(o.value, ","))
}

// UnmarshalJSON implements the json.Unmarshaler interface for
// TypeOption.
func (o *TypeOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = []string{}
		return nil
	}
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	if len(s) == 0 {
		o.value = []string{}
		return nil
	}
	o.value = strings.Split(s, ",")
	return nil
}

// Equal returns true if the given option is equal to the instance one. In case
// the given option is nil, we checked the instance one is set to the default
// value of the option.
func (o *TypeOption) Equal(o2 *TypeOption) bool {
	if o == nil {
		return o2 == nil || reflect.DeepEqual(o2.value, []string{})
	}
	if o2 == nil {
		return o == nil || reflect.DeepEqual(o.value, []string{})
	}
	return reflect.DeepEqual(o.value, o2.value)
}

// TypeEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func TypeEqual(o1, o2 *TypeOption) bool {
	return o1.Equal(o2)
}
