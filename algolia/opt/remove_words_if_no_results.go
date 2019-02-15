package opt

import "encoding/json"

type RemoveWordsIfNoResultsOption struct {
	value string
}

func RemoveWordsIfNoResults(v string) RemoveWordsIfNoResultsOption {
	return RemoveWordsIfNoResultsOption{v}
}

func (o RemoveWordsIfNoResultsOption) Get() string {
	return o.value
}

func (o RemoveWordsIfNoResultsOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *RemoveWordsIfNoResultsOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = &#34;none&#34;
		return nil
	}
	return json.Unmarshal(data, &o.value)
}