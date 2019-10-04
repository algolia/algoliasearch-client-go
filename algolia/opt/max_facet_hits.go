// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

// MaxFacetHitsOption is a wrapper for an MaxFacetHits option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type MaxFacetHitsOption struct {
	value int
}

// MaxFacetHits wraps the given value into a MaxFacetHitsOption.
func MaxFacetHits(v int) *MaxFacetHitsOption {
	return &MaxFacetHitsOption{v}
}

// Get retrieves the actual value of the option parameter.
func (o *MaxFacetHitsOption) Get() int {
	if o == nil {
		return 10
	}
	return o.value
}

// MarshalJSON implements the json.Marshaler interface for
// MaxFacetHitsOption.
func (o MaxFacetHitsOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

// UnmarshalJSON implements the json.Unmarshaler interface for
// MaxFacetHitsOption.
func (o *MaxFacetHitsOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = 10
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

// Equal returns true if the given option is equal to the instance one. In case
// the given option is nil, we checked the instance one is set to the default
// value of the option.
func (o *MaxFacetHitsOption) Equal(o2 *MaxFacetHitsOption) bool {
	if o == nil {
		return o2 == nil || o2.value == 10
	}
	if o2 == nil {
		return o == nil || o.value == 10
	}
	return o.value == o2.value
}

// MaxFacetHitsEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func MaxFacetHitsEqual(o1, o2 *MaxFacetHitsOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
