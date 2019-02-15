package opt

import "encoding/json"

type LengthOption struct {
	value int
}

func Length(v int) LengthOption {
	return LengthOption{v}
}

func (o LengthOption) Get() int {
	return o.value
}

func (o LengthOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *LengthOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = 0
		return nil
	}
	return json.Unmarshal(data, &o.value)
}