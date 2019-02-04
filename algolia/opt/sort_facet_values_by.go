package opt

import "encoding/json"

type SortFacetValuesByOption struct {
	value string
}

func SortFacetValuesBy(v string) SortFacetValuesByOption {
	return SortFacetValuesByOption{v}
}

func (o SortFacetValuesByOption) Get() string {
	return o.value
}

func (o SortFacetValuesByOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *SortFacetValuesByOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.value)
}
