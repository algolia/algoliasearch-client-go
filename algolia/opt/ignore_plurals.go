package opt

import (
	"encoding/json"
	"reflect"
)

// IgnorePluralsOption is a wrapper for an IgnorePlurals option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type IgnorePluralsOption struct {
	ignorePlurals bool
	languages     []string
}

// IgnorePlurals returns an IgnorePluralsOption whose value is set to the given boolean.
func IgnorePlurals(v bool) *IgnorePluralsOption {
	return &IgnorePluralsOption{ignorePlurals: v}
}

// IgnorePluralsFor returns an IgnorePluralsOption whose value is set to the given list of
// languages.
func IgnorePluralsFor(languages ...string) *IgnorePluralsOption {
	return &IgnorePluralsOption{languages: languages}
}

// Get retrieves the actual value of the option parameter.
func (o *IgnorePluralsOption) Get() (bool, []string) {
	if o == nil {
		return false, nil
	}
	return o.ignorePlurals, o.languages
}

// MarshalJSON implements the json.Marshaler interface for
// IgnorePluralsOption.
func (o IgnorePluralsOption) MarshalJSON() ([]byte, error) {
	if len(o.languages) > 0 {
		return json.Marshal(o.languages)
	}
	return json.Marshal(o.ignorePlurals)
}

// UnmarshalJSON implements the json.Unmarshaler interface for
// IgnorePluralsOption.
func (o *IgnorePluralsOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	if err := json.Unmarshal(data, &o.languages); err == nil {
		return nil
	}

	return json.Unmarshal(data, &o.ignorePlurals)
}

// Equal returns true if the given option is equal to the instance one. In case
// the given option is nil, we checked the instance one is set to the default
// value of the option.
func (o *IgnorePluralsOption) Equal(o2 *IgnorePluralsOption) bool {
	if o == nil {
		return o2 == nil || !o2.ignorePlurals && len(o2.languages) == 0
	}
	if o2 == nil {
		return o == nil || !o.ignorePlurals && len(o.languages) == 0
	}
	return reflect.DeepEqual(o, o2)
}

// IgnorePluralsEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func IgnorePluralsEqual(o1, o2 *IgnorePluralsOption) bool {
	return o1.Equal(o2)
}
