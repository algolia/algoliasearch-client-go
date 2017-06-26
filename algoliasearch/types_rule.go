package algoliasearch

type Rule struct {
	ObjectID    string          `json:"objectID,omitempty"`
	Condition   RuleCondition   `json:"condition"`
	Consequence RuleConsequence `json:"consequence"`
	Description string          `json:"description,omitempty"`
}

// RuleCondition is the part of an Algolia Rule which describes the condition
// for the rule. The `Context` is optional, hence, it will get ignored if an
// empty string is used to set it.
type RuleCondition struct {
	Anchoring RulePatternAnchoring `json:"anchoring"`
	Pattern   string               `json:"pattern"`
	Context   string               `json:"context,omitempty"`
}

type RulePatternAnchoring string

const (
	Is         RulePatternAnchoring = "is"
	StarstWith RulePatternAnchoring = "startsWith"
	EndsWith   RulePatternAnchoring = "endsWith"
	Contains   RulePatternAnchoring = "contains"
)

// NewSimpleRuleCondition generates a RuleCondition where only the `Anchoring`
// and `Pattern` fields are specified. The optional `Context` field is then
// excluded.
func NewSimpleRuleCondition(anchoring RulePatternAnchoring, pattern string) RuleCondition {
	return NewRuleCondition(anchoring, pattern, "")
}

// NewRuleCondition generates a RuleCondition where all the possible fields can
// be specified.
func NewRuleCondition(anchoring RulePatternAnchoring, pattern, context string) RuleCondition {
	return RuleCondition{
		Anchoring: anchoring,
		Pattern:   pattern,
		Context:   context,
	}
}

type RuleConsequence struct {
	Params   Map              `json:"params,omitempty"`
	Promote  []PromotedObject `json:"promote,omitempty"`
	UserData interface{}      `json:"userData,omitempty"`
}

type QueryIncrementalEdit struct {
	Remove []string `json:"remove"`
}

type PromotedObject struct {
	ObjectID string `json:"objectID"`
	Position int    `json:"position"`
}

type SaveRuleRes struct {
	TaskID    int    `json:"taskID"`
	UpdatedAt string `json:"updatedAt"`
}

type BatchRulesRes struct {
	TaskID    int    `json:"taskID"`
	UpdatedAt string `json:"updatedAt"`
}

type DeleteRuleRes struct {
	TaskID    int    `json:"taskID"`
	UpdatedAt string `json:"updatedAt"`
}

type ClearRulesRes struct {
	TaskID    int    `json:"taskID"`
	UpdatedAt string `json:"updatedAt"`
}

type SearchRulesRes struct {
	Hits    []Rule `json:"hits"`
	NbHits  int    `json:"nbHits"`
	Page    int    `json:"page"`
	NbPages int    `json:"nbPages"`
}
