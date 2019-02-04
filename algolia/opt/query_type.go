package opt

import "encoding/json"

type QueryTypeOption struct {
	value string
}

func QueryType_PrefixLast(v string) QueryTypeOption {
	return QueryTypeOption{v}
}

func (o QueryTypeOption) Get() string {
	return o.value
}

func (o QueryTypeOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *QueryTypeOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.value)
}
