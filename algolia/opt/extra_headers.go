package opt

import "encoding/json"

type ExtraHeadersOption struct {
	value map[string]string
}

func ExtraHeaders(v map[string]string) ExtraHeadersOption {
	return ExtraHeadersOption{v}
}

func (o ExtraHeadersOption) Get() map[string]string {
	return o.value
}

func (o ExtraHeadersOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *ExtraHeadersOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = make(map[string]string)
		return nil
	}
	return json.Unmarshal(data, &o.value)
}