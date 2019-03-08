// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

type CamelCaseAttributesOption struct {
	value []string
}

func CamelCaseAttributes(v ...string) *CamelCaseAttributesOption {
	return &CamelCaseAttributesOption{v}
}

func (o CamelCaseAttributesOption) Get() []string {
	return o.value
}

func (o CamelCaseAttributesOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *CamelCaseAttributesOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = []string{}
		return nil
	}
	return json.Unmarshal(data, &o.value)
}