package opt

import (
	"encoding/json"
)

type AlternativesAsExactOption struct {
	alternatives []string
}

func AlternativesAsExact(alternatives ...string) AlternativesAsExactOption {
	return AlternativesAsExactOption{alternatives}
}

func (o AlternativesAsExactOption) Get() []string {
	return o.alternatives
}

func (o AlternativesAsExactOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.alternatives)
}

func (o *AlternativesAsExactOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.alternatives)
}
