package opt

import (
	"encoding/json"
)

type TagFiltersOption struct {
	comp composableFilterOption
}

func TagFilter(filter string) *TagFiltersOption {
	return &TagFiltersOption{composableFilter(filter)}
}

func TagFilterOr(filters ...interface{}) *TagFiltersOption {
	return &TagFiltersOption{composableFilterOr(filters...)}
}

func TagFilterAnd(filters ...interface{}) *TagFiltersOption {
	return &TagFiltersOption{composableFilterAnd(filters...)}
}

func (o TagFiltersOption) Get() [][]string {
	return o.comp.Get()
}

func (o TagFiltersOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.comp)
}

func (o *TagFiltersOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.comp)
}
