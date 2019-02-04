package opt

import (
	"encoding/json"
)

type DisableTypoToleranceOnWordsOption struct {
	words []string
}

func DisableTypoToleranceOnWords(words ...string) DisableTypoToleranceOnWordsOption {
	return DisableTypoToleranceOnWordsOption{words}
}

func (o DisableTypoToleranceOnWordsOption) Get() []string {
	return o.words
}

func (o DisableTypoToleranceOnWordsOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.words)
}

func (o *DisableTypoToleranceOnWordsOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.words)
}
