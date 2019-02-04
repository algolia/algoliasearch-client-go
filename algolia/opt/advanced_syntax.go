package opt

import "encoding/json"

type AdvancedSyntaxOption struct {
	value bool
}

func AdvancedSyntax(v bool) AdvancedSyntaxOption {
	return AdvancedSyntaxOption{v}
}

func (o AdvancedSyntaxOption) Get() bool {
	return o.value
}

func (o AdvancedSyntaxOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *AdvancedSyntaxOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.value)
}
