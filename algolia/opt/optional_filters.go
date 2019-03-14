package opt

import (
	"encoding/json"
)

type OptionalFiltersOption struct {
	comp composableFilterOption
}

func OptionalFilter(filter string) *OptionalFiltersOption {
	return &OptionalFiltersOption{composableFilter(filter)}
}

func OptionalFilterOr(filters ...interface{}) *OptionalFiltersOption {
	return &OptionalFiltersOption{composableFilterOr(filters...)}
}

func OptionalFilterAnd(filters ...interface{}) *OptionalFiltersOption {
	return &OptionalFiltersOption{composableFilterAnd(filters...)}
}

func (o OptionalFiltersOption) Get() [][]string {
	return o.comp.Get()
}

func (o OptionalFiltersOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.comp)
}

func (o *OptionalFiltersOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.comp)
}

func (o *OptionalFiltersOption) Equal(o2 *OptionalFiltersOption) bool {
	return o.comp.Equal(&o2.comp)
}

func OptionalFiltersEqual(o1, o2 *OptionalFiltersOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
