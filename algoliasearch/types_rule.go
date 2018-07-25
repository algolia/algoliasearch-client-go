package algoliasearch

import (
	"encoding/json"
	"fmt"
	"time"
)

type Rule struct {
	Condition            RuleCondition   `json:"condition"`
	Consequence          RuleConsequence `json:"consequence"`
	Description          string          `json:"description,omitempty"`
	Enabled              bool            `json:"enabled"` // Defaults to true
	HighlightResult      Map             `json:"_highlightResult,omitempty"`
	ObjectID             string          `json:"objectID,omitempty"`
	Validity             []TimeRange     `json:"validity,omitempty"`
	isExplicitlyDisabled bool
}

func (r *Rule) Enable() {
	r.Enabled = true
}

func (r *Rule) Disable() {
	r.Enabled = false
	r.isExplicitlyDisabled = true
}

func (r *Rule) enableImplicitly() {
	if !r.isExplicitlyDisabled {
		r.Enable()
	}
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
	StartsWith RulePatternAnchoring = "startsWith"
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
	Hide     []HiddenObject   `json:"hide,omitempty"`
	UserData interface{}      `json:"userData,omitempty"`
}

// AutomaticFacetFilter
type AutomaticFacetFilter struct {
	Facet       string `json:"facet"`
	Disjunctive bool   `json:"disjunctive"` // Defaults to false
	Score       int    `json:"score"`
}

// QueryIncrementalEdit can be used as a value for the `query` key when used in
// the `RuleConsequence.Params` map. It is used to remove specific words from
// the original query string.
//
// Deprecated: Use `DeleteEdit` instead. More specifically, code previously
// written this way:
//
//  consequence := algoliasearch.RuleConsquence{
//  	Params: algoliasearch.Map{
//  		"query": algoliasearch.QueryIncrementalEdit{
//  			Remove: []string{"term1", "term2"},
//  		},
//  	},
//  }
//
// should now be written:
//
//  consequence := algoliasearch.RuleConsequence{
//  	Params: algoliasearch.Map{
//  		"query": algoliasearch.Map{
//  			"edits": []algoliasearch.Edit{
//  				algoliasearch.DeleteEdit("term1"),
//  				algoliasearch.DeleteEdit("term2"),
//  			},
//  		},
//  	},
//  }
//
type QueryIncrementalEdit struct {
	Remove []string `json:"remove"`
}

type Edit struct {
	Type   string `json:"type"`
	Delete string `json:"delete"`
	Insert string `json:"insert,omitempty"`
}

// DeleteEdit returns a new `Edit` instance used to remove the given `word`
// from an original query when used as a `RuleConsequence.Params`.
func DeleteEdit(word string) Edit {
	return Edit{
		Type:   "remove",
		Delete: word,
	}
}

// ReplaceEdit returns a new `Edit` instance used to replace the given `old`
// term with `new` in a query when used as a `RuleConsequence.Params`.
func ReplaceEdit(old, new string) Edit {
	return Edit{
		Type:   "replace",
		Delete: old,
		Insert: new,
	}
}

type PromotedObject struct {
	ObjectID string `json:"objectID"`
	Position int    `json:"position"`
}

type HiddenObject struct {
	ObjectID string `json:"objectID"`
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

type TimeRange struct {
	From  time.Time
	Until time.Time
}

type timeRangeResponse struct {
	From  int64 `json:"from"`
	Until int64 `json:"until"`
}

func (r TimeRange) MarshalJSON() ([]byte, error) {
	data := fmt.Sprintf(
		`{"from":%d,"until":%d}`,
		r.From.Unix(),
		r.Until.Unix(),
	)
	return []byte(data), nil
}

func (r *TimeRange) UnmarshalJSON(b []byte) error {
	var res timeRangeResponse

	err := json.Unmarshal(b, &res)
	if err != nil {
		return fmt.Errorf("cannot unmarshal integer values of time range: %s", err)
	}

	r.From = time.Unix(res.From, 0)
	r.Until = time.Unix(res.Until, 0)

	return nil
}

type SearchRulesRes struct {
	Hits    []Rule `json:"hits"`
	NbHits  int    `json:"nbHits"`
	Page    int    `json:"page"`
	NbPages int    `json:"nbPages"`
}
