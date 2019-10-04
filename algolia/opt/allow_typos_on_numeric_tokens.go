// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

// AllowTyposOnNumericTokensOption is a wrapper for an AllowTyposOnNumericTokens option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type AllowTyposOnNumericTokensOption struct {
	value bool
}

// AllowTyposOnNumericTokens wraps the given value into a AllowTyposOnNumericTokensOption.
func AllowTyposOnNumericTokens(v bool) *AllowTyposOnNumericTokensOption {
	return &AllowTyposOnNumericTokensOption{v}
}

// Get retrieves the actual value of the option parameter.
func (o *AllowTyposOnNumericTokensOption) Get() bool {
	if o == nil {
		return true
	}
	return o.value
}

// MarshalJSON implements the json.Marshaler interface for
// AllowTyposOnNumericTokensOption.
func (o AllowTyposOnNumericTokensOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

// UnmarshalJSON implements the json.Unmarshaler interface for
// AllowTyposOnNumericTokensOption.
func (o *AllowTyposOnNumericTokensOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = true
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

// Equal returns true if the given option is equal to the instance one. In case
// the given option is nil, we checked the instance one is set to the default
// value of the option.
func (o *AllowTyposOnNumericTokensOption) Equal(o2 *AllowTyposOnNumericTokensOption) bool {
	if o == nil {
		return o2 == nil || o2.value == true
	}
	if o2 == nil {
		return o == nil || o.value == true
	}
	return o.value == o2.value
}

// AllowTyposOnNumericTokensEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func AllowTyposOnNumericTokensEqual(o1, o2 *AllowTyposOnNumericTokensOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
