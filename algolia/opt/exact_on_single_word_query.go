package opt

import "encoding/json"

type ExactOnSingleWordQueryOption struct {
	value string
}

func ExactOnSingleWordQuery(v string) ExactOnSingleWordQueryOption {
	return ExactOnSingleWordQueryOption{v}
}

func (o ExactOnSingleWordQueryOption) Get() string {
	return o.value
}

func (o ExactOnSingleWordQueryOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *ExactOnSingleWordQueryOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.value)
}
