package opt

import "encoding/json"

type SynonymsOption struct {
	value bool
}

func Synonyms(v bool) SynonymsOption {
	return SynonymsOption{v}
}

func (o SynonymsOption) Get() bool {
	return o.value
}

func (o SynonymsOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *SynonymsOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = true
		return nil
	}
	return json.Unmarshal(data, &o.value)
}