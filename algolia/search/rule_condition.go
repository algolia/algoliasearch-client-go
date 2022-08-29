package search

import "encoding/json"

type RuleCondition struct {
	Anchoring    RulePatternAnchoring
	Pattern      string
	Context      string
	Alternatives *Alternatives
	Filters      string
}

func (c RuleCondition) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	if c.Anchoring != "" {
		m["anchoring"] = c.Anchoring
		m["pattern"] = c.Pattern
	}
	if c.Context != "" {
		m["context"] = c.Context
	}
	if c.Alternatives != nil {
		m["alternatives"] = *c.Alternatives
	}
	if c.Filters != "" {
		m["filters"] = c.Filters
	}
	return json.Marshal(m)
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
