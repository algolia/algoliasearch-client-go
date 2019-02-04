package opt

import (
	"encoding/json"
)

type DisableTypoToleranceOnAttributesOption struct {
	attributes []string
}

func DisableTypoToleranceOnAttributes(attributes ...string) DisableTypoToleranceOnAttributesOption {
	return DisableTypoToleranceOnAttributesOption{attributes}
}

func (o DisableTypoToleranceOnAttributesOption) Get() []string {
	return o.attributes
}

func (o DisableTypoToleranceOnAttributesOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.attributes)
}

func (o *DisableTypoToleranceOnAttributesOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.attributes)
}
