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

func (o *TagFiltersOption) Equal(o2 *TagFiltersOption) bool {
	return o.comp.Equal(&o2.comp)
}

func TagFiltersEqual(o1, o2 *TagFiltersOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
