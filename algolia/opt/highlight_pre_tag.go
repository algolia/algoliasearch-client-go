package opt

import "encoding/json"

type HighlightPreTagOption struct {
	value string
}

func HighlightPreTag(v string) HighlightPreTagOption {
	return HighlightPreTagOption{v}
}

func (o HighlightPreTagOption) Get() string {
	return o.value
}

func (o HighlightPreTagOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *HighlightPreTagOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = &#34;&#34;
		return nil
	}
	return json.Unmarshal(data, &o.value)
}