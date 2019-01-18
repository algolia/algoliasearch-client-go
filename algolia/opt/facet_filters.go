package opt

import (
	"encoding/json"
)

type FacetFiltersOption struct {
	comp composableFilterOption
}

func FacetFilter(filter string) FacetFiltersOption {
	return FacetFiltersOption{composableFilter(filter)}
}

func FacetFilterOr(filters ...interface{}) FacetFiltersOption {
	return FacetFiltersOption{composableFilterOr(filters...)}
}

func FacetFilterAnd(filters ...interface{}) FacetFiltersOption {
	return FacetFiltersOption{composableFilterAnd(filters...)}
}

func (o FacetFiltersOption) Get() [][]string {
	return o.comp.Get()
}

func (o FacetFiltersOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.comp)
}

func (o *FacetFiltersOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.comp)
}
