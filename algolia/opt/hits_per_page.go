package opt

import "encoding/json"

type HitsPerPageOption struct {
	hitsPerPage int
}

func HitsPerPage(hitsPerPage int) HitsPerPageOption {
	return HitsPerPageOption{hitsPerPage}
}

func (o HitsPerPageOption) Get() int {
	return o.hitsPerPage
}

func (o HitsPerPageOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.hitsPerPage)
}

func (o *HitsPerPageOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.hitsPerPage)
}
