package opt

import "encoding/json"

type MinimumAroundRadiusOption struct {
	meters int
}

func MinimumAroundRadius(meters int) MinimumAroundRadiusOption {
	return MinimumAroundRadiusOption{meters}
}

func (o MinimumAroundRadiusOption) Get() int {
	return o.meters
}

func (o MinimumAroundRadiusOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.meters)
}

func (o *MinimumAroundRadiusOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.meters)
}
