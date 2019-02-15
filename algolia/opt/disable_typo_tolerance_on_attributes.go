package opt

import "encoding/json"

type DisableTypoToleranceOnAttributesOption struct {
	value []string
}

func DisableTypoToleranceOnAttributes(v []string) DisableTypoToleranceOnAttributesOption {
	return DisableTypoToleranceOnAttributesOption{v}
}

func (o DisableTypoToleranceOnAttributesOption) Get() []string {
	return o.value
}

func (o DisableTypoToleranceOnAttributesOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *DisableTypoToleranceOnAttributesOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = nil
		return nil
	}
	return json.Unmarshal(data, &o.value)
}