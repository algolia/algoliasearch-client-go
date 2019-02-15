package opt

import (
	"encoding/json"
)

type ResponseFieldsOption struct {
	// TODO Check if "all" is represented by "*" or ["*"]
	fields []string
}

func ResponseFields(fields ...string) ResponseFieldsOption {
	return ResponseFieldsOption{fields}
}

func (o ResponseFieldsOption) Get() []string {
	return o.fields
}

func (o ResponseFieldsOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.fields)
}

func (o *ResponseFieldsOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.fields)
}
