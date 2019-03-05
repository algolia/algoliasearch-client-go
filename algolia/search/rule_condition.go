package search

type RuleCondition struct {
	Anchoring RulePatternAnchoring `json:"anchoring"`
	Pattern   string               `json:"pattern"`
	Context   string               `json:"context,omitempty"`
}

type RulePatternAnchoring string

const (
	Is         RulePatternAnchoring = "is"
	StartsWith RulePatternAnchoring = "startsWith"
	EndsWith   RulePatternAnchoring = "endsWith"
	Contains   RulePatternAnchoring = "contains"
)
