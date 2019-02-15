package opt

import "encoding/json"

type ReplaceSynonymsInHighlightOption struct {
	value bool
}

func ReplaceSynonymsInHighlight(v bool) ReplaceSynonymsInHighlightOption {
	return ReplaceSynonymsInHighlightOption{v}
}

func (o ReplaceSynonymsInHighlightOption) Get() bool {
	return o.value
}

func (o ReplaceSynonymsInHighlightOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *ReplaceSynonymsInHighlightOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = true
		return nil
	}
	return json.Unmarshal(data, &o.value)
}