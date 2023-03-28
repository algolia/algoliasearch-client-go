// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"reflect"
	"strings"
)

// AdvancedSyntaxFeaturesOption is a wrapper for an AdvancedSyntaxFeatures option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type AdvancedSyntaxFeaturesOption struct {
	value []string
}

// AdvancedSyntaxFeatures wraps the given value into a AdvancedSyntaxFeaturesOption.
func AdvancedSyntaxFeatures(v ...string) *AdvancedSyntaxFeaturesOption {
	if v == nil {
		return &AdvancedSyntaxFeaturesOption{[]string{}}
	}
	return &AdvancedSyntaxFeaturesOption{v}
}

// Get retrieves the actual value of the option parameter.
func (o *AdvancedSyntaxFeaturesOption) Get() []string {
	if o == nil {
		return []string{"exactPhrase", "excludeWords"}
	}
	return o.value
}

// MarshalJSON implements the json.Marshaler interface for
// AdvancedSyntaxFeaturesOption.
func (o AdvancedSyntaxFeaturesOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

// UnmarshalJSON implements the json.Unmarshaler interface for
// AdvancedSyntaxFeaturesOption.
func (o *AdvancedSyntaxFeaturesOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = []string{"exactPhrase", "excludeWords"}
		return nil
	}
	var s string
	err := json.Unmarshal(data, &s)
	if err == nil {
		o.value = strings.Split(s, ",")
		if len(o.value) == 1 && o.value[0] == "" {
			o.value = []string{}
		}
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

// Equal returns true if the given option is equal to the instance one. In case
// the given option is nil, we checked the instance one is set to the default
// value of the option.
func (o *AdvancedSyntaxFeaturesOption) Equal(o2 *AdvancedSyntaxFeaturesOption) bool {
	if o == nil {
		return o2 == nil || reflect.DeepEqual(o2.value, []string{"exactPhrase", "excludeWords"})
	}
	if o2 == nil {
		return o == nil || reflect.DeepEqual(o.value, []string{"exactPhrase", "excludeWords"})
	}
	return reflect.DeepEqual(o.value, o2.value)
}

// AdvancedSyntaxFeaturesEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func AdvancedSyntaxFeaturesEqual(o1, o2 *AdvancedSyntaxFeaturesOption) bool {
	return o1.Equal(o2)
}
