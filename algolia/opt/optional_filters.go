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
