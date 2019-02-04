package opt

import (
	"encoding/json"
)

type CamelCaseAttributesOption struct {
	attributes []string
}

func CamelCaseAttributes(attributes ...string) CamelCaseAttributesOption {
	return CamelCaseAttributesOption{attributes}
}

func (o CamelCaseAttributesOption) Get() []string {
	return o.attributes
}

func (o CamelCaseAttributesOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.attributes)
}

func (o *CamelCaseAttributesOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.attributes)
}
