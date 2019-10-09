package opt

import (
	"encoding/json"
	"time"
)

// ValidUntilOption is a wrapper for an ValidUntil option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type ValidUntilOption struct {
	value time.Time
}

// ValidUntil returns an ValidUntilOption whose value is set to the given time.Time.
func ValidUntil(v time.Time) *ValidUntilOption {
	return &ValidUntilOption{v}
}

// Get retrieves the actual value of the option parameter.
func (o *ValidUntilOption) Get() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.value
}

// MarshalJSON implements the json.Marshaler interface for
// ValidUntilOption.
func (o ValidUntilOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

// UnmarshalJSON implements the json.Unmarshaler interface for
// ValidUntilOption.
func (o *ValidUntilOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = time.Time{}
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

// Equal returns true if the given option is equal to the instance one. In case
// the given option is nil, we checked the instance one is set to the default
// value of the option.
func (o *ValidUntilOption) Equal(o2 *ValidUntilOption) bool {
	if o == nil {
		return o2 == nil || o2.value.IsZero()
	}
	if o2 == nil {
		return o == nil || o.value.IsZero()
	}
	return o.value.Equal(o2.value)
}

// ValidUntilEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func ValidUntilEqual(o1, o2 *ValidUntilOption) bool {
	return o1.Equal(o2)
}
