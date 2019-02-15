package opt

import "encoding/json"

type AttributesToSnippetOption struct {
	value []string
}

func AttributesToSnippet(v []string) AttributesToSnippetOption {
	return AttributesToSnippetOption{v}
}

func (o AttributesToSnippetOption) Get() []string {
	return o.value
}

func (o AttributesToSnippetOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *AttributesToSnippetOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = nil
		return nil
	}
	return json.Unmarshal(data, &o.value)
}