package opt

import (
	"bytes"
	"encoding/json"
	"reflect"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/errs"
)

// UserDataOption is a wrapper for an UserData option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type UserDataOption struct {
	itf  interface{}
	data []byte
}

// UserData wraps the given value into a UserDataOption.
func UserData(v interface{}) *UserDataOption {
	return &UserDataOption{itf: v}
}

// Get retrieves the actual value of the option parameter as a map. If the value
// is not a map, nil is returned instead. In that case, use the Decode() method
// to appropriately decode the underlying userData into a different instance
// type.
func (o *UserDataOption) Get() map[string]interface{} {
	if o == nil {
		return map[string]interface{}{}
	}
	m, ok := o.itf.(map[string]interface{})
	if !ok {
		return map[string]interface{}{}
	}
	return m
}

func (o *UserDataOption) Decode(itf interface{}) error {
	if o == nil {
		return errs.ErrJSONDecode(nil, "UserDataOption")
	}
	if o.data == nil {
		data, err := json.Marshal(o.itf)
		if err != nil {
			return err
		}
		o.data = data
	}
	return json.Unmarshal(o.data, &itf)
}

// MarshalJSON implements the json.Marshaler interface for
// UserDataOption.
func (o UserDataOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.itf)
}

// UnmarshalJSON implements the json.Unmarshaler interface for
// UserDataOption.
func (o *UserDataOption) UnmarshalJSON(data []byte) error {
	o.data = data
	if string(data) == "null" {
		o.itf = map[string]interface{}{}
		return nil
	}
	return json.Unmarshal(data, &o.itf)
}

// Equal returns true if the given option is equal to the instance one. In case
// the given option is nil, we checked the instance one is set to the default
// value of the option.
func (o *UserDataOption) Equal(o2 *UserDataOption) bool {
	if o == nil {
		return o2 == nil || reflect.DeepEqual(o2.itf, map[string]interface{}{})
	}
	if o2 == nil {
		return o == nil || reflect.DeepEqual(o.itf, map[string]interface{}{})
	}
	data, err := json.Marshal(o.itf)
	data2, err2 := json.Marshal(o2.itf)
	return err == nil && err2 == nil && bytes.Equal(data, data2)
}

// UserDataEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func UserDataEqual(o1, o2 *UserDataOption) bool {
	return o1.Equal(o2)
}
