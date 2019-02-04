package opt

import (
	"encoding/json"
)

type OptionalWordsOption struct {
	words []string
}

func OptionalWords(words ...string) OptionalWordsOption {
	return OptionalWordsOption{words}
}

func (o OptionalWordsOption) Get() []string {
	return o.words
}

func (o OptionalWordsOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.words)
}

func (o *OptionalWordsOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.words)
}
