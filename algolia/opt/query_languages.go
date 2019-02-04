package opt

import (
	"encoding/json"
)

type QueryLanguagesOption struct {
	attributes []string
}

func QueryLanguages(attributes ...string) QueryLanguagesOption {
	return QueryLanguagesOption{attributes}
}

func (o QueryLanguagesOption) Get() []string {
	return o.attributes
}

func (o QueryLanguagesOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.attributes)
}

func (o *QueryLanguagesOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.attributes)
}
