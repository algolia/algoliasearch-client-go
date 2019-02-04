package opt

import "encoding/json"

type AllowTyposOnNumericTokensOption struct {
	value bool
}

func AllowTyposOnNumericTokens(v bool) AllowTyposOnNumericTokensOption {
	return AllowTyposOnNumericTokensOption{v}
}

func (o AllowTyposOnNumericTokensOption) Get() bool {
	return o.value
}

func (o AllowTyposOnNumericTokensOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *AllowTyposOnNumericTokensOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.value)
}
