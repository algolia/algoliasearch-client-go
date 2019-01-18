package opt

import (
	"encoding/json"
)

type NumericFiltersOption struct {
	comp composableFilterOption
}

func NumericFilter(filter string) NumericFiltersOption {
	return NumericFiltersOption{composableFilter(filter)}
}

func NumericFilterOr(filters ...interface{}) NumericFiltersOption {
	return NumericFiltersOption{composableFilterOr(filters...)}
}

func NumericFilterAnd(filters ...interface{}) NumericFiltersOption {
	return NumericFiltersOption{composableFilterAnd(filters...)}
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
