package opt

import "encoding/json"

type DisableExactOnAttributesOption struct {
	value []string
}

func DisableExactOnAttributes(v []string) DisableExactOnAttributesOption {
	return DisableExactOnAttributesOption{v}
}

func (o DisableExactOnAttributesOption) Get() []string {
	return o.value
}

func (o DisableExactOnAttributesOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *DisableExactOnAttributesOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = nil
		return nil
	}
	return json.Unmarshal(data, &o.value)
}