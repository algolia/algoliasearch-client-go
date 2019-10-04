package search

import "encoding/json"

type RuleCondition struct {
	Anchoring    RulePatternAnchoring `json:"anchoring,omitempty"`
	Pattern      string               `json:"pattern,omitempty"`
	Context      string               `json:"context,omitempty"`
	Alternatives *Alternatives        `json:"alternatives,omitempty"`
}

type RulePatternAnchoring string

const (
	Is         RulePatternAnchoring = "is"
	StartsWith RulePatternAnchoring = "startsWith"
	EndsWith   RulePatternAnchoring = "endsWith"
	Contains   RulePatternAnchoring = "contains"
)

type Alternatives struct {
	enabled bool
}

func AlternativesEnabled() *Alternatives {
	return &Alternatives{enabled: true}
}

func AlternativesDisabled() *Alternatives {
	return &Alternatives{enabled: false}
}

func (a Alternatives) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.enabled)
}

func (a *Alternatives) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &a.enabled)
}
