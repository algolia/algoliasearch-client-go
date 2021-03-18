package search

import (
	"encoding/json"
	"fmt"
)

// Compound represents an entry in the Compounds dictionary.
type Compound struct {
	objectID      string
	language      string
	Word          string
	Decomposition []string
}

// Compound implements DictionaryEntry
func (c Compound) ObjectID() string { return c.objectID }
func (c Compound) Language() string { return c.language }

// NewCompound is a constructor of Compound
func NewCompound(objectID, language, word string, decomposition []string) Compound {
	return Compound{objectID, language, word, decomposition}
}

type compound struct {
	ObjectID      string   `json:"objectID"`
	Language      string   `json:"language"`
	Word          string   `json:"word"`
	Decomposition []string `json:"decomposition"`
}

func (c Compound) MarshalJSON() ([]byte, error) {
	return json.Marshal(compound{c.objectID, c.language, c.Word, c.Decomposition})
}

func (c *Compound) UnmarshalJSON(data []byte) error {
	if string(data) == jsonNull {
		return nil
	}

	var compound compound
	err := json.Unmarshal(data, &compound)
	if err != nil {
		return fmt.Errorf("cannot unmarshal Compound dictionary entry: %v", err)
	}

	c.objectID = compound.ObjectID
	c.language = compound.Language
	c.Word = compound.Word
	c.Decomposition = compound.Decomposition
	return nil
}
