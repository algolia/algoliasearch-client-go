package opt

import "encoding/json"

type HitsPerPageOption struct {
	value int
}

func HitsPerPage(v int) HitsPerPageOption {
	return HitsPerPageOption{v}
}

func (o HitsPerPageOption) Get() int {
	return o.value
}

func (o HitsPerPageOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *HitsPerPageOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = 20
		return nil
	}
	return json.Unmarshal(data, &o.value)
}