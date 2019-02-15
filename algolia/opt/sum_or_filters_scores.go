package opt

import "encoding/json"

type sumOrFiltersScoresOption struct {
	value bool
}

func sumOrFiltersScores(v bool) sumOrFiltersScoresOption {
	return sumOrFiltersScoresOption{v}
}

func (o sumOrFiltersScoresOption) Get() bool {
	return o.value
}

func (o sumOrFiltersScoresOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *sumOrFiltersScoresOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = false
		return nil
	}
	return json.Unmarshal(data, &o.value)
}