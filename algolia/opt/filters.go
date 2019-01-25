package opt

import "encoding/json"

type FiltersOption struct {
	filters string
}

func Filters(filters string) FiltersOption {
	return FiltersOption{filters}
}

func (o FiltersOption) Get() string {
	return o.filters
}

func (o FiltersOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.filters)
}

func (o *FiltersOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.filters)
}
