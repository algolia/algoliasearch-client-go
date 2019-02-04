package opt

import (
	"encoding/json"
)

type RuleContextsOption struct {
	attributes []string
}

func RuleContexts(attributes ...string) RuleContextsOption {
	return RuleContextsOption{attributes}
}

func (o RuleContextsOption) Get() []string {
	return o.attributes
}

func (o RuleContextsOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.attributes)
}

func (o *RuleContextsOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.attributes)
}
