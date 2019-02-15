package opt

import "encoding/json"

type CreateIfNotExistsOption struct {
	value bool
}

func CreateIfNotExists(v bool) CreateIfNotExistsOption {
	return CreateIfNotExistsOption{v}
}

func (o CreateIfNotExistsOption) Get() bool {
	return o.value
}

func (o CreateIfNotExistsOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *CreateIfNotExistsOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = false
		return nil
	}
	return json.Unmarshal(data, &o.value)
}