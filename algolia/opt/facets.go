package opt

import "encoding/json"

type FacetsOption struct {
	facets []string
}

func Facets(facets []string) FacetsOption {
	return FacetsOption{facets}
}

func (o FacetsOption) Get() []string {
	return o.facets
}

func (o FacetsOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.facets)
}

func (o *FacetsOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.facets)
}
