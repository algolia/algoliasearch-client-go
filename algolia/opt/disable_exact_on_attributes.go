package opt

import (
	"encoding/json"
)

type DisableExactOnAttributesOption struct {
	attributes []string
}

func DisableExactOnAttributes(attributes ...string) DisableExactOnAttributesOption {
	return DisableExactOnAttributesOption{attributes}
}

func (o DisableExactOnAttributesOption) Get() []string {
	return o.attributes
}

func (o DisableExactOnAttributesOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.attributes)
}

func (o *DisableExactOnAttributesOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.attributes)
}
