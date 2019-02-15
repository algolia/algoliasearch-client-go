package opt

import "encoding/json"

type FacetingAfterDistinctOption struct {
	value bool
}

func FacetingAfterDistinct(v bool) FacetingAfterDistinctOption {
	return FacetingAfterDistinctOption{v}
}

func (o FacetingAfterDistinctOption) Get() bool {
	return o.value
}

func (o FacetingAfterDistinctOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *FacetingAfterDistinctOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = false
		return nil
	}
	return json.Unmarshal(data, &o.value)
}