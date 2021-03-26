package search

import (
	"encoding/json"
	"fmt"
)

// DictionaryEntry represents an entry in a given dictionary.
type DictionaryEntry interface {
	ObjectID() string
	Language() string
}

type rawDictionaryEntry struct{ impl DictionaryEntry }

func (d *rawDictionaryEntry) UnmarshalJSON(data []byte) error {
	var (
		rawDictEntry map[string]interface{}
		err          error
	)

	if err = json.Unmarshal(data, &rawDictEntry); err != nil {
		return fmt.Errorf("cannot unmarshal dictionary entry: error when unmarshalling %v", err)
	}

	_, ok := rawDictEntry["decomposition"]
	if ok {
		var compound Compound
		err = json.Unmarshal(data, &compound)
		d.impl = compound
		return err
	}

	_, ok = rawDictEntry["words"]
	if ok {
		var plural Plural
		err = json.Unmarshal(data, &plural)
		d.impl = plural
		return err
	}

	_, ok = rawDictEntry["word"]
	if ok {
		var stopWord Stopword
		err = json.Unmarshal(data, &stopWord)
		d.impl = stopWord
		return err
	}

	return fmt.Errorf("cannot unmarshal dictionary entry: incorrect input")

}
