package search

import (
	"encoding/json"
	"fmt"
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
