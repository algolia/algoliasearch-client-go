package opt

import "encoding/json"

type MinProximityOption struct {
	proximity int
}

func MinProximity(proximity int) MinProximityOption {
	return MinProximityOption{proximity}
}

func (o MinProximityOption) Get() int {
	return o.proximity
}

func (o MinProximityOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.proximity)
}

func (o *MinProximityOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.proximity)
}
