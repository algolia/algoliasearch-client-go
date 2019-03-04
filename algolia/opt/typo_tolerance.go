package opt

import "encoding/json"

type TypoToleranceOption struct {
	// TODO: support both string and bool
	value string
}

func TypoToleranceTrue() *TypoToleranceOption {
	return &TypoToleranceOption{"true"}
}

func TypoToleranceFalse() *TypoToleranceOption {
	return &TypoToleranceOption{"false"}
}

func TypoToleranceMin() *TypoToleranceOption {
	return &TypoToleranceOption{"min"}
}

func TypoToleranceStrict() *TypoToleranceOption {
	return &TypoToleranceOption{"strict"}
}

func (o TypoToleranceOption) Get() string {
	return o.value
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
