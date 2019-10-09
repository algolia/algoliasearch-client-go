package opt

import (
	"encoding/json"
)

// OptionalFiltersOption is a wrapper for an OptionalFilters option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type OptionalFiltersOption struct {
	comp composableFilterOption
}

// OptionalFilter returns an OptionalFiltersOption whose value is set to the given
// string.
func OptionalFilter(filter string) *OptionalFiltersOption {
	return &OptionalFiltersOption{composableFilter(filter)}
}

// OptionalFilters returns an OptionalFiltersOption whose value is set as an OR of the
// given filters.
func OptionalFilterOr(filters ...interface{}) *OptionalFiltersOption {
	return &OptionalFiltersOption{composableFilterOr(filters...)}
}

// OptionalFilters returns an OptionalFiltersOption whose value is set as an AND of
// the given filters.
func OptionalFilterAnd(filters ...interface{}) *OptionalFiltersOption {
	return &OptionalFiltersOption{composableFilterAnd(filters...)}
}

// Get retrieves the actual value of the option parameter. The slice of slice
// represents filters that are ORed then ANDed. For instance, the following
// returned slice:
//
//              [][]string{ {"filter1", "filter2"}, {"filter3"} }
// means:
//
//              ("filter1" OR "filter2" ) AND "filter 3"
func (o *OptionalFiltersOption) Get() [][]string {
	if o == nil {
		return nil
	}
	return o.comp.Get()
}

// MarshalJSON implements the json.Marshaler interface for OptionalFiltersOption.
func (o OptionalFiltersOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.comp)
}

// UnmarshalJSON implements the json.Unmarshaler interface for OptionalFiltersOption.
func (o *OptionalFiltersOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.comp)
}

// Equal returns true if the given option is equal to the instance one. In case
// the given option is nil, we checked the instance one is set to the default
// value of the option.
func (o *OptionalFiltersOption) Equal(o2 *OptionalFiltersOption) bool {
	if o == nil {
		return o2 == nil || len(o2.comp.Get()) == 0
	}
	if o2 == nil {
		return o == nil || len(o.comp.Get()) == 0
	}
	return o.comp.Equal(&o2.comp)
}

// OptionalFiltersEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func OptionalFiltersEqual(o1, o2 *OptionalFiltersOption) bool {
	return o1.Equal(o2)
}
