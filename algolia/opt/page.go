package opt

import "encoding/json"

type PageOption struct {
	page int
}

func Page(page int) PageOption {
	return PageOption{page}
}

func (o PageOption) Get() int {
	return o.page
}

func (o PageOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.page)
}

func (o *PageOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.page)
}
