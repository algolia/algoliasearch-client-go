package opt

import "encoding/json"

type SnippetEllipsisTextOption struct {
	value string
}

func SnippetEllipsisText(v string) SnippetEllipsisTextOption {
	return SnippetEllipsisTextOption{v}
}

func (o SnippetEllipsisTextOption) Get() string {
	return o.value
}

func (o SnippetEllipsisTextOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *SnippetEllipsisTextOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.value)
}
