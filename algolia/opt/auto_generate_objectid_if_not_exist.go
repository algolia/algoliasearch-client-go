package opt

import "encoding/json"

type AutoGenerateObjectIDIfNotExistOption struct {
	value bool
}

func AutoGenerateObjectIDIfNotExist(v bool) AutoGenerateObjectIDIfNotExistOption {
	return AutoGenerateObjectIDIfNotExistOption{v}
}

func (o AutoGenerateObjectIDIfNotExistOption) Get() bool {
	return o.value
}

func (o AutoGenerateObjectIDIfNotExistOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *AutoGenerateObjectIDIfNotExistOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.value)
}
