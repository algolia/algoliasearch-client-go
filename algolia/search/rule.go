package search

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type Rule struct {
	Condition   RuleCondition   `json:"condition"`
	Consequence RuleConsequence `json:"consequence"`
	Description string          `json:"description,omitempty"`
	Enabled     bool            `json:"enabled"` // Defaults to true
	ObjectID    string          `json:"objectID,omitempty"`
	Validity    []TimeRange     `json:"validity,omitempty"`
}

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

func (r Rule) Equal(r2 Rule) bool {
	if !(r.ObjectID == r2.ObjectID &&
		r.Description == r2.Description &&
		r.Enabled == r2.Enabled &&
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

func (r TimeRange) Equal(r2 TimeRange) bool {
	return r.From.Equal(r2.From) && r.Until.Equal(r2.Until)
}
