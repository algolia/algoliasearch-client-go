package opt

import "encoding/json"

type AroundPrecisionOption struct {
	meters int
}

func AroundPrecision(meters int) AroundPrecisionOption {
	return AroundPrecisionOption{meters}
}

func (o AroundPrecisionOption) Get() int {
	return o.meters
}

func (o AroundPrecisionOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.meters)
}

func (o *AroundPrecisionOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.meters)
}
