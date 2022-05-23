package opt

import "encoding/json"

// ReRankingApplyFilterOption is a wrapper for an ReRankingApplyFilter option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type ReRankingApplyFilterOption struct {
	comp composableFilterOption
}

// ReRankingApplyFilter wraps the given value into a ReRankingApplyFilterOption.
func ReRankingApplyFilter(v string) *ReRankingApplyFilterOption {
	return &ReRankingApplyFilterOption{composableFilter(v)}
}

// ReRankingApplyFilterOr returns an ReRankingApplyFilterOption whose value is set as an OR of the
// given filters.
func ReRankingApplyFilterOr(filters ...interface{}) *ReRankingApplyFilterOption {
	return &ReRankingApplyFilterOption{composableFilterOr(filters...)}
}

// ReRankingApplyFilterAnd returns an ReRankingApplyFilterOption whose value is set as an AND of
// the given filters.
func ReRankingApplyFilterAnd(filters ...interface{}) *ReRankingApplyFilterOption {
	return &ReRankingApplyFilterOption{composableFilterAnd(filters...)}
}

// Get retrieves the actual value of the option parameter.
func (o *ReRankingApplyFilterOption) Get() [][]string {
	if o == nil {
		return nil
	}
	return o.comp.Get()
}

// MarshalJSON implements the json.Marshaler interface for
// ReRankingApplyFilterOption.
func (o ReRankingApplyFilterOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.comp)
}

// UnmarshalJSON implements the json.Unmarshaler interface for
// ReRankingApplyFilterOption.
func (o *ReRankingApplyFilterOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.comp)
}

// Equal returns true if the given option is equal to the instance one. In case
// the given option is nil, we checked the instance one is set to the default
// value of the option.
func (o *ReRankingApplyFilterOption) Equal(o2 *ReRankingApplyFilterOption) bool {
	if o == nil {
		return o2 == nil || len(o2.comp.Get()) == 0
	}
	if o2 == nil {
		return o == nil || len(o.comp.Get()) == 0
	}
	return o.comp.Equal(&o2.comp)
}

// ReRankingApplyFilterEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func ReRankingApplyFilterEqual(o1, o2 *ReRankingApplyFilterOption) bool {
	return o1.Equal(o2)
}
