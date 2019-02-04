package opt

import "encoding/json"

type RemoveStopWordsOption struct {
	removeStopWords bool
	languages       []string
}

func RemoveStopWords(languages ...string) RemoveStopWordsOption {
	return RemoveStopWordsOption{languages: languages}
}

func RemoveStopWordsTrue() RemoveStopWordsOption {
	return RemoveStopWordsOption{removeStopWords: true}
}

func RemoveStopWordsFalse() RemoveStopWordsOption {
	return RemoveStopWordsOption{removeStopWords: false}
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
