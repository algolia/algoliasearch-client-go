package opt

import (
	"encoding/json"
	"reflect"
)

type RemoveStopWordsOption struct {
	removeStopWords bool
	languages       []string
}

func RemoveStopWords(v bool) *RemoveStopWordsOption {
	return &RemoveStopWordsOption{removeStopWords: v}
}

func RemoveStopWordsFor(languages ...string) *RemoveStopWordsOption {
	return &RemoveStopWordsOption{languages: languages}
}

func (o RemoveStopWordsOption) Get() (bool, []string) {
	return o.removeStopWords, o.languages
}

func (o RemoveStopWordsOption) MarshalJSON() ([]byte, error) {
	if len(o.languages) > 0 {
		return json.Marshal(o.languages)
	}
	return json.Marshal(o.removeStopWords)
}

func (o *RemoveStopWordsOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	if err := json.Unmarshal(data, &o.languages); err == nil {
		return nil
	}

	return json.Unmarshal(data, &o.removeStopWords)
}

func (o *RemoveStopWordsOption) Equal(o2 *RemoveStopWordsOption) bool {
	if o2 == nil {
		return o.removeStopWords == false && len(o.languages) == 0
	}
	return reflect.DeepEqual(o, o2)
}

func RemoveStopWordsEqual(o1, o2 *RemoveStopWordsOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
