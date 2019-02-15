package opt

import "encoding/json"

type PercentileComputationOption struct {
	value bool
}

func PercentileComputation(v bool) PercentileComputationOption {
	return PercentileComputationOption{v}
}

func (o PercentileComputationOption) Get() bool {
	return o.value
}

func (o PercentileComputationOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *PercentileComputationOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = true
		return nil
	}
	return json.Unmarshal(data, &o.value)
}