package opt

import "encoding/json"

type TypoToleranceOption struct {
	valueBool   bool
	valueString string
}

func TypoTolerance(value bool) *TypoToleranceOption {
	return &TypoToleranceOption{valueBool: value}
}

func TypoToleranceMin() *TypoToleranceOption {
	return &TypoToleranceOption{valueString: "min"}
}

func TypoToleranceStrict() *TypoToleranceOption {
	return &TypoToleranceOption{valueString: "strict"}
}

func (o TypoToleranceOption) Get() (bool, string) {
	return o.valueBool, o.valueString
}

func (o TypoToleranceOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *TypoToleranceOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = "true"
		return nil
	}
	return json.Unmarshal(data, &o.value)
}
