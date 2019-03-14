package opt

import (
	"encoding/json"
)

type NumericFiltersOption struct {
	comp composableFilterOption
}

func NumericFilter(filter string) *NumericFiltersOption {
	return &NumericFiltersOption{composableFilter(filter)}
}

func NumericFilterOr(filters ...interface{}) *NumericFiltersOption {
	return &NumericFiltersOption{composableFilterOr(filters...)}
}

func NumericFilterAnd(filters ...interface{}) *NumericFiltersOption {
	return &NumericFiltersOption{composableFilterAnd(filters...)}
}

func (o NumericFiltersOption) Get() [][]string {
	return o.comp.Get()
}

func (o NumericFiltersOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.comp)
}

func (o *NumericFiltersOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.comp)
}

func (o *NumericFiltersOption) Equal(o2 *NumericFiltersOption) bool {
	return o.comp.Equal(&o2.comp)
}

func NumericFiltersEqual(o1, o2 *NumericFiltersOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
