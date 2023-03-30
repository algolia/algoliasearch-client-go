package search

import (
	"encoding/json"
	"fmt"
)

type RuleParams struct {
	Query                         *RuleQuery             `json:"query,omitempty"`
	AutomaticFacetFilters         []AutomaticFacetFilter `json:"automaticFacetFilters,omitempty"`
	AutomaticOptionalFacetFilters []AutomaticFacetFilter `json:"automaticOptionalFacetFilters,omitempty"`
	RenderingContent              *RenderingContent      `json:"renderingContent,omitempty"`
	QueryParams
}

type RuleQuery struct {
	simpleQuery string
	objectQuery *RuleQueryObjectQuery
}

func NewRuleQuerySimple(query string) *RuleQuery {
	return &RuleQuery{simpleQuery: query}
}

func NewRuleQueryObject(object RuleQueryObjectQuery) *RuleQuery {
	return &RuleQuery{objectQuery: &object}
}

type RuleQueryObjectQuery struct {
	Edits []QueryEdit `json:"edits"`
}

func (q RuleQuery) Get() (string, *RuleQueryObjectQuery) {
	return q.simpleQuery, q.objectQuery
}

func (q RuleQuery) MarshalJSON() ([]byte, error) {
	if q.objectQuery == nil || len(q.objectQuery.Edits) == 0 {
		return json.Marshal(q.simpleQuery)
	}
	return json.Marshal(q.objectQuery)
}

func (q *RuleQuery) UnmarshalJSON(data []byte) error {
	if string(data) == jsonNull {
		return nil
	}

	var objectQuery RuleQueryObjectQuery
	if err := json.Unmarshal(data, &objectQuery); err == nil {
		q.objectQuery = &objectQuery
	}

	// Kept for backward-compatibility only
	var incrementalEdit struct {
		Remove []string `json:"remove"`
	}
	if err := json.Unmarshal(data, &incrementalEdit); err == nil {
		if q.objectQuery == nil {
			q.objectQuery = &RuleQueryObjectQuery{}
		}
		for _, word := range incrementalEdit.Remove {
			q.objectQuery.Edits = append(q.objectQuery.Edits, RemoveEdit(word))
		}
		return nil
	}

	var simpleQuery string
	if err := json.Unmarshal(data, &simpleQuery); err == nil {
		q.simpleQuery = simpleQuery
		return nil
	}

	return fmt.Errorf("cannot unmarshal query rule's query field")
}

type AutomaticFacetFilter struct {
	Facet       string `json:"facet"`
	Disjunctive bool   `json:"disjunctive"` // Defaults to false
	Score       int    `json:"score"`       // Defaults to 1
}

// AutomaticFacetFilter can be unmarshalled from a string or an object.
func (a *AutomaticFacetFilter) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}

	if string(data) == jsonNull {
		return nil
	}

	switch data[0] {
	case '"':
		var facet string
		if err := json.Unmarshal(data, &facet); err != nil {
			return err
		}
		*a = AutomaticFacetFilter{
			Facet: facet,
		}
		return nil
	case '{':
		var alias struct {
			Facet       string `json:"facet"`
			Disjunctive bool   `json:"disjunctive"`
			Score       int    `json:"score"`
		}
		if err := json.Unmarshal(data, &alias); err != nil {
			return err
		}
		*a = AutomaticFacetFilter{
			Facet:       alias.Facet,
			Disjunctive: alias.Disjunctive,
			Score:       alias.Score,
		}
		return nil
	default:
		return fmt.Errorf("cannot unmarshal automatic facet filter")
	}
}

type QueryEdit struct {
	Type   QueryEditType `json:"type"`
	Delete string        `json:"delete"`
	Insert string        `json:"insert,omitempty"`
}

type QueryEditType string

const (
	Remove  QueryEditType = "remove"
	Replace QueryEditType = "replace"
)

func RemoveEdit(word string) QueryEdit {
	return QueryEdit{
		Type:   Remove,
		Delete: word,
	}
}

func ReplaceEdit(old, new string) QueryEdit {
	return QueryEdit{
		Type:   Replace,
		Delete: old,
		Insert: new,
	}
}
