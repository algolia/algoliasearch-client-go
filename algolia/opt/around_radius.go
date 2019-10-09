package opt

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/errs"
)

// AroundRadiusOption is a wrapper for an AroundRadius option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type AroundRadiusOption struct {
	meters int
	isAll  bool
}

// AroundRadius wraps the given value into an AroundRadiusOption.
func AroundRadius(meters int) *AroundRadiusOption {
	return &AroundRadiusOption{meters: meters}
}

// AroundRadiusAll returns an AroundRadiusOption whose value is set to "all".
func AroundRadiusAll() *AroundRadiusOption {
	return &AroundRadiusOption{isAll: true}
}

// Get retrieves the actual value of the option parameter.
func (o *AroundRadiusOption) Get() (int, string) {
	if o == nil || o.isAll {
		return 0, "all"
	}
	return o.meters, ""
}

// MarshalJSON implements the json.Marshaler interface for
// AroundRadiusOption.
func (o AroundRadiusOption) MarshalJSON() ([]byte, error) {
	if o.isAll {
		return []byte(`"all"`), nil
	}

	if o.meters != 0 {
		return []byte(fmt.Sprintf("%d", o.meters)), nil
	}

	return []byte("null"), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for
// AroundRadiusOption.
func (o *AroundRadiusOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	var meters int
	if err := json.Unmarshal(data, &meters); err == nil {
		o.meters = meters
		return nil
	}

	var all string
	if err := json.Unmarshal(data, &all); err == nil && all == "all" {
		o.isAll = true
		return nil
	}

	return errs.ErrJSONDecode(data, "AroundRadiusOption")
}

// Equal returns true if the given option is equal to the instance one. In case
// the given option is nil, we checked the instance one is set to the default
// value of the option.
func (o *AroundRadiusOption) Equal(o2 *AroundRadiusOption) bool {
	if o == nil {
		return o2 == nil || !o2.isAll && o2.meters == 0
	}
	if o2 == nil {
		return o == nil || !o.isAll && o.meters == 0
	}
	return reflect.DeepEqual(o, o2)
}

// AroundRadiusEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func AroundRadiusEqual(o1, o2 *AroundRadiusOption) bool {
	return o1.Equal(o2)
}
