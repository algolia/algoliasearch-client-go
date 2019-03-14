package opt

import (
	"encoding/json"
	"reflect"
)

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
		o.valueBool = true
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

func (o *TypoToleranceOption) Equal(o2 *TypoToleranceOption) bool {
	if o2 == nil {
		return o.valueBool == true && o.valueString == ""
	}
	return reflect.DeepEqual(o, o2)
}

func TypoToleranceEqual(o1, o2 *TypoToleranceOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
