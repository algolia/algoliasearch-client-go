package opt

import (
	"encoding/json"
)

// NumericFiltersOption is a wrapper for an NumericFilters option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type NumericFiltersOption struct {
	comp composableFilterOption
}

// NumericFilter returns an NumericFiltersOption whose value is set to the given
// string.
func NumericFilter(filter string) *NumericFiltersOption {
	return &NumericFiltersOption{composableFilter(filter)}
}

// NumericFilters returns an NumericFiltersOption whose value is set as an OR of the
// given filters.
func NumericFilterOr(filters ...interface{}) *NumericFiltersOption {
	return &NumericFiltersOption{composableFilterOr(filters...)}
}

// NumericFilters returns an NumericFiltersOption whose value is set as an AND of
// the given filters.
func NumericFilterAnd(filters ...interface{}) *NumericFiltersOption {
	return &NumericFiltersOption{composableFilterAnd(filters...)}
}

// Get retrieves the actual value of the option parameter. The slice of slice
// represents filters that are ORed then ANDed. For instance, the following
// returned slice:
//
//              [][]string{ {"filter1", "filter2"}, {"filter3"} }
// means:
//
//              ("filter1" OR "filter2" ) AND "filter 3"
func (o *NumericFiltersOption) Get() [][]string {
	if o == nil {
		return nil
	}
	return o.comp.Get()
}

// MarshalJSON implements the json.Marshaler interface for NumericFiltersOption.
func (o NumericFiltersOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.comp)
}

// UnmarshalJSON implements the json.Unmarshaler interface for NumericFiltersOption.
func (o *NumericFiltersOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.comp)
}

// Equal returns true if the given option is equal to the instance one. In case
// the given option is nil, we checked the instance one is set to the default
// value of the option.
func (o *NumericFiltersOption) Equal(o2 *NumericFiltersOption) bool {
	if o == nil {
		return o2 == nil || len(o2.comp.Get()) == 0
	}
	if o2 == nil {
		return o == nil || len(o.comp.Get()) == 0
	}
	return o.comp.Equal(&o2.comp)
}

// NumericFiltersEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func NumericFiltersEqual(o1, o2 *NumericFiltersOption) bool {
	return o1.Equal(o2)
}
