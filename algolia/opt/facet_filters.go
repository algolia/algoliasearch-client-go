package opt

import (
	"encoding/json"
)

type FacetFiltersOption struct {
	comp composableFilterOption
}

func FacetFilter(filter string) *FacetFiltersOption {
	return &FacetFiltersOption{composableFilter(filter)}
}

func FacetFilterOr(filters ...interface{}) *FacetFiltersOption {
	return &FacetFiltersOption{composableFilterOr(filters...)}
}

func FacetFilterAnd(filters ...interface{}) *FacetFiltersOption {
	return &FacetFiltersOption{composableFilterAnd(filters...)}
}

func (o *FacetFiltersOption) Get() [][]string {
	if o == nil {
		return nil
	}
	return o.comp.Get()
}

func (o FacetFiltersOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.comp)
}

func (o *FacetFiltersOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.comp)
}

func (o *FacetFiltersOption) Equal(o2 *FacetFiltersOption) bool {
	return o.comp.Equal(&o2.comp)
}

func FacetFiltersEqual(o1, o2 *FacetFiltersOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
