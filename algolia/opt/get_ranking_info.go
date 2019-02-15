package opt

import "encoding/json"

type GetRankingInfoOption struct {
	value bool
}

func GetRankingInfo(v bool) GetRankingInfoOption {
	return GetRankingInfoOption{v}
}

func (o GetRankingInfoOption) Get() bool {
	return o.value
}

func (o GetRankingInfoOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *GetRankingInfoOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = false
		return nil
	}
	return json.Unmarshal(data, &o.value)
}