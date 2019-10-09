package opt

import (
	"encoding/json"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/errs"
)

// DistinctOption is a wrapper for an Distinct option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type DistinctOption struct {
	value int
}

// Distinct returns an DistinctOption whose value is set to the given boolean.
func Distinct(enabled bool) *DistinctOption {
	if enabled {
		return &DistinctOption{value: 1}
	}
	return &DistinctOption{value: 0}
}

// DistinctOf returns an DistinctOption whose value is set to the given integer.
func DistinctOf(v int) *DistinctOption {
	return &DistinctOption{value: v}
}

// Get retrieves the actual value of the option parameter.
func (o *DistinctOption) Get() (bool, int) {
	if o == nil {
		return false, 0
	}
	return o.value == 1, o.value
}

// MarshalJSON implements the json.Marshaler interface for
// DistinctOption.
func (o DistinctOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

// UnmarshalJSON implements the json.Unmarshaler interface for
// DistinctOption.
func (o *DistinctOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	var i int
	if err := json.Unmarshal(data, &i); err == nil {
		o.value = i
		return nil
	}

	var b bool
	if err := json.Unmarshal(data, &b); err == nil {
		if b {
			o.value = 1
		} else {
			o.value = 0
		}
		return nil
	}

	return errs.ErrJSONDecode(data, "Distinct")
}

// Equal returns true if the given option is equal to the instance one. In case
// the given option is nil, we checked the instance one is set to the default
// value of the option.
func (o *DistinctOption) Equal(o2 *DistinctOption) bool {
	if o == nil {
		return o2 == nil || o2.value == 0
	}
	if o2 == nil {
		return o == nil || o.value == 0
	}
	return o.value == o2.value
}

// DistinctEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func DistinctEqual(o1, o2 *DistinctOption) bool {
	return o1.Equal(o2)
}
