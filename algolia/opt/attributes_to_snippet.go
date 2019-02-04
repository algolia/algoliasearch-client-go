package opt

import (
	"encoding/json"
)

type AttributesToSnippetOption struct {
	attributes []string
}

func AttributesToSnippet(attributes ...string) AttributesToSnippetOption {
	return AttributesToSnippetOption{attributes}
}

func (o AttributesToSnippetOption) Get() []string {
	return o.attributes
}

func (o AttributesToSnippetOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.attributes)
}

func (o *AttributesToSnippetOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.attributes)
}
