package opt

import "encoding/json"

type AttributesToHighlightOption struct {
	value []string
}

func AttributesToHighlight(v []string) AttributesToHighlightOption {
	return AttributesToHighlightOption{v}
}

func (o AttributesToHighlightOption) Get() []string {
	return o.value
}

func (o AttributesToHighlightOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *AttributesToHighlightOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = nil
		return nil
	}
	return json.Unmarshal(data, &o.value)
}