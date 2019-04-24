package search

import (
	"encoding/json"
	"fmt"
)

type Placeholder struct {
	objectID     string
	Placeholder  string
	Replacements []string
}

func NewPlaceholder(objectID, placeholder string, replacements ...string) Placeholder {
	return Placeholder{objectID, placeholder, replacements}
}

func (s Placeholder) ObjectID() string  { return s.objectID }
func (s Placeholder) Type() SynonymType { return PlaceholderType }

type placeholder struct {
	ObjectID     string      `json:"objectID"`
	Type         SynonymType `json:"type"`
	Placeholder  string      `json:"placeholder"`
	Replacements []string    `json:"replacements"`
}

func (s Placeholder) MarshalJSON() ([]byte, error) {
	return json.Marshal(placeholder{s.ObjectID(), s.Type(), s.Placeholder, s.Replacements})
}

func (s *Placeholder) UnmarshalJSON(data []byte) error {
	if string(data) == jsonNull {
		return nil
	}

	var synonym placeholder
	err := json.Unmarshal(data, &synonym)
	if err != nil {
		return fmt.Errorf("cannot unmarshal Placeholder: %v", err)
	}

	if synonym.Type != s.Type() {
		return fmt.Errorf("cannot deserialize synonym of type %s into Placeholder", synonym.Type)
	}

	s.objectID = synonym.ObjectID
	s.Placeholder = synonym.Placeholder
	s.Replacements = synonym.Replacements
	return nil
}
