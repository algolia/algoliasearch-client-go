package opt

import (
	"encoding/json"
)

type AttributesToHighlightOption struct {
	attributes []string
}

func AttributesToHighlight(attributes ...string) AttributesToHighlightOption {
	return AttributesToHighlightOption{attributes}
}

func (o AttributesToHighlightOption) Get() []string {
	return o.attributes
}

func (o AttributesToHighlightOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.attributes)
}

func (o *AttributesToHighlightOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.attributes)
}
