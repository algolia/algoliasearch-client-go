package search

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
)

// Rule represents an Algolia query rule.
type Rule struct {
	// Deprecated: Use `Conditions` instead to specify one or more condition(s)
	Condition   RuleCondition      `json:"condition"`
	Conditions  []RuleCondition    `json:"conditions,omitempty"`
	Consequence RuleConsequence    `json:"consequence"`
	Description string             `json:"description,omitempty"`
	Enabled     *opt.EnabledOption `json:"enabled,omitempty"`
	ObjectID    string             `json:"objectID,omitempty"`
	Validity    []TimeRange        `json:"validity,omitempty"`
	Tags        []string           `json:"tags,omitempty"`
	Scope       string             `json:"scope,omitempty"`
}

// TimeRange is a pair of begin/end time.Time used to represent a rule validity
// (used by Rule.Validity field).
type TimeRange struct {
	From  time.Time
	Until time.Time
}

func (r TimeRange) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"from":%d,"until":%d}`, r.From.Unix(), r.Until.Unix())), nil
}

func (r *TimeRange) UnmarshalJSON(b []byte) error {
	var res struct {
		From  int64 `json:"from"`
		Until int64 `json:"until"`
	}

	err := json.Unmarshal(b, &res)
	if err != nil {
		return fmt.Errorf("cannot unmarshal integer values of time range: %s", err)
	}

	r.From = time.Unix(res.From, 0)
	r.Until = time.Unix(res.Until, 0)
	return nil
}

// Equal returns true if the Rules are equal. It returns false otherwise.
func (r Rule) Equal(r2 Rule) bool {
	if !(r.ObjectID == r2.ObjectID &&
		r.Description == r2.Description &&
		r.Enabled.Equal(r2.Enabled) &&
		len(r.Validity) == len(r2.Validity) &&
		reflect.DeepEqual(r.Condition, r2.Condition) &&
		reflect.DeepEqual(r.Consequence, r2.Consequence)) {
		return false
	}

	for i, r := range r.Validity {
		if !r.Equal(r2.Validity[i]) {
			return false
		}
	}

	return true
}

func (r Rule) MarshalJSON() ([]byte, error) {
	type _Rule Rule
	type _RuleWithOptionalCondition struct {
		Condition *RuleCondition `json:"condition,omitempty"`
		_Rule
	}

	var tmp _RuleWithOptionalCondition
	tmp._Rule = _Rule(r)

	if r.Condition != (RuleCondition{}) {
		tmp.Condition = &r.Condition
	}
	return json.Marshal(tmp)
}

// Equal returns true if the TimeRanges are equal. It returns false otherwise.
func (r TimeRange) Equal(r2 TimeRange) bool {
	return r.From.Equal(r2.From) && r.Until.Equal(r2.Until)
}
