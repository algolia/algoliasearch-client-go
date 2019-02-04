package opt

import "encoding/json"

type SumOrFiltersScoresOption struct {
	sumOrFiltersScores bool
}

func SumOrFiltersScores(sumOrFiltersScores bool) SumOrFiltersScoresOption {
	return SumOrFiltersScoresOption{sumOrFiltersScores}
}

func (o SumOrFiltersScoresOption) Get() bool {
	return o.sumOrFiltersScores
}

func (o SumOrFiltersScoresOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.sumOrFiltersScores)
}

func (o *SumOrFiltersScoresOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.sumOrFiltersScores)
}
