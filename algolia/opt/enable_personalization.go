package opt

import "encoding/json"

type EnablePersonalizationOption struct {
	value bool
}

func EnablePersonalization(v bool) EnablePersonalizationOption {
	return EnablePersonalizationOption{v}
}

func (o EnablePersonalizationOption) Get() bool {
	return o.value
}

func (o EnablePersonalizationOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *EnablePersonalizationOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.value)
}
