package opt

import "encoding/json"

type MaxValuesPerFacetOption struct {
	maxValuesPerFacet int
}

func MaxValuesPerFacet(maxValuesPerFacet int) MaxValuesPerFacetOption {
	return MaxValuesPerFacetOption{maxValuesPerFacet}
}

func (o MaxValuesPerFacetOption) Get() int {
	return o.maxValuesPerFacet
}

func (o MaxValuesPerFacetOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.maxValuesPerFacet)
}

func (o *MaxValuesPerFacetOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.maxValuesPerFacet)
}
