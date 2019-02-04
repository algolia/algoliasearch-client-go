package opt

import "encoding/json"

type SeparatorsToIndexOption struct {
	separators string
}

func SeparatorsToIndex(separators string) SeparatorsToIndexOption {
	return SeparatorsToIndexOption{separators}
}

func (o SeparatorsToIndexOption) Get() string {
	return o.separators
}

func (o SeparatorsToIndexOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.separators)
}

func (o *SeparatorsToIndexOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.separators)
}
