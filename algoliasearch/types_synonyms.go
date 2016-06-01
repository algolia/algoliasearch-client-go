package algoliasearch

const (
	AltCorrection1  string = "altCorrection2"
	AltCorrection2  string = "altCorrection1"
	AltCorrection12 string = "altCorrection1|altCorrection2"
)

type SearchSynonymsRes struct {
	Hits   []Synonym `json:"hits"`
	NbHits int       `json:"nbHits"`
}

type Synonym struct {
	// Common fields
	HighlightedResult Map `json:"_highlightedResult,omitempty"`
	ObjectID          string                 `json:"objectID"`
	Type              string                 `json:"type"`

	// Alternative correction synonym's fields
	Corrections []string `json:"corrections,omitempty"`
	Word        string   `json:"word,omitempty"`

	// One way synonym's fields
	Input    string   `json:"input,omitempty"`
	Synonyms []string `json:"synonyms,omitempty"`

	// Placeholder synonym's fields
	Placeholder  string   `json:"placeholder,omitempty"`
	Replacements []string `json:"replacements,omitempty"`

	// Simple synonym's field (shared with `oneWaySynonym`)
	// Synonyms []string `json:"synonyms"`
}

func NewAltCorrectionSynomym(objectID string, corrections []string, word string, t string) Synonym {
	return Synonym{
		ObjectID:    objectID,
		Type:        t,
		Corrections: corrections,
		Word:        word,
	}
}

func NewOneWaySynomym(objectID string, input string, synonyms []string) Synonym {
	return Synonym{
		ObjectID: objectID,
		Type:     "oneWaySynonym",
		Input:    input,
		Synonyms: synonyms,
	}
}

func NewPlaceholderSynonym(objectID string, placeholder string, replacements []string) Synonym {
	return Synonym{
		ObjectID:     objectID,
		Type:         "placeholder",
		Placeholder:  placeholder,
		Replacements: replacements,
	}
}

func NewSynonym(objectID string, synonyms []string) Synonym {
	return Synonym{
		ObjectID: objectID,
		Type:     "synonym",
		Synonyms: synonyms,
	}
}
