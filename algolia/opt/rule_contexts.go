package opt

import "encoding/json"

type RuleContextsOption struct {
	value string
}

func RuleContexts(v string) RuleContextsOption {
	return RuleContextsOption{v}
}

func (o RuleContextsOption) Get() string {
	return o.value
}

func (o RuleContextsOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *RuleContextsOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = &#34;none&#34;
		return nil
	}
	return json.Unmarshal(data, &o.value)
}