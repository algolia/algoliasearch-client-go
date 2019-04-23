package opt

import (
	"encoding/json"
	"fmt"

	"github.com/algolia/algoliasearch-client-go/algolia/errs"
)

// TypoToleranceOption is a wrapper for an TypoTolerance option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type TypoToleranceOption struct {
	value string
}

// TypoTolerance returns an TypoToleranceOption whose value is set to the given boolean.
func TypoTolerance(value bool) *TypoToleranceOption {
	return &TypoToleranceOption{value: fmt.Sprintf("%t", value)}
}

// TypoToleranceFor returns an TypoToleranceOption whose value is set to
// "min".
func TypoToleranceMin() *TypoToleranceOption {
	return &TypoToleranceOption{value: "min"}
}

// TypoToleranceFor returns an TypoToleranceOption whose value is set to
// "strict".
func TypoToleranceStrict() *TypoToleranceOption {
	return &TypoToleranceOption{value: "strict"}
}

// Get retrieves the actual value of the option parameter.
func (o *TypoToleranceOption) Get() (bool, string) {
	if o == nil {
		return true, ""
	}
	if o.value == "true" {
		return true, ""
	}
	if o.value == "false" {
		return false, ""
	}
	return false, o.value
}

// MarshalJSON implements the json.Marshaler interface for
// TypoToleranceOption.
func (o TypoToleranceOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

// UnmarshalJSON implements the json.Unmarshaler interface for
// TypoToleranceOption.
func (o *TypoToleranceOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = "true"
		return nil
	}

	var v string
	if err := json.Unmarshal(data, &v); err == nil {
		o.value = v
		return nil
	}

	var b bool
	if err := json.Unmarshal(data, &b); err == nil {
		o.value = fmt.Sprintf("%t", b)
		return nil
	}

	return errs.ErrJSONDecode(data, "TypoTolerance")
}

// Equal returns true if the given option is equal to the instance one. In case
// the given option is nil, we checked the instance one is set to the default
// value of the option.
func (o *TypoToleranceOption) Equal(o2 *TypoToleranceOption) bool {
	if o2 == nil {
		return o.value == "true"
	}
	return o.value == o2.value
}

// TypoToleranceEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func TypoToleranceEqual(o1, o2 *TypoToleranceOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
