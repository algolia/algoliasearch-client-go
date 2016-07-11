package algoliasearch

type Settings struct {
	// Indexing parameters
	AllowCompressionOfIntegerArray bool     `json:"allowCompressionOfIntegerArray"`
	AttributesForDistinct          []string `json:"attributesForDistinct"`
	AttributesForFaceting          []string `json:"attributesForFaceting"`
	AttributesToIndex              []string `json:"attributesToIndex"`
	CustomRanking                  []string `json:"customRanking"`
	NumericAttributesToIndex       []string `json:"numericAttributesToIndex"`
	Ranking                        []string `json:"ranking"`
	SeparatorsToIndex              string   `json:"separatorsToIndex"`
	Slaves                         []string `json:"slaves"`
	UnretrievableAttributes        []string `json:"unretrievableAttributes"`

	// Query expansion
	DisableTypoToleranceOnAttributes []string `json:"disableTypoToleranceOnAttributes"`
	DisableTypoToleranceOnWords      []string `json:"disableTypoToleranceOnWords"`

	// Default query parameters (can be overridden at query-time)
	AdvancedSyntax             bool        `json:"advancedSyntax"`
	AllowTyposOnNumericTokens  bool        `json:"allowTyposOnNumericTokens"`
	AttributesToHighlight      []string    `json:"attributesToHighlight"`
	AttributesToRetrieve       []string    `json:"attributesToRetrieve"`
	AttributesToSnippet        []string    `json:"attributesToSnippet"`
	Distinct                   int         `json:"distinct"`
	HighlightPostTag           string      `json:"highlightPostTag"`
	HighlightPreTag            string      `json:"highlightPreTag"`
	HitsPerPage                int         `json:"hitsPerPage"`
	IgnorePlurals              bool        `json:"ignorePlurals"`
	MaxValuesPerFacet          int         `json:"maxValuesPerFacet"`
	MinProximity               int         `json:"minProximity"`
	MinWordSizefor1Typo        int         `json:"minWordSizefor1Typo"`
	MinWordSizefor2Typos       int         `json:"minWordSizefor2Typos"`
	OptionalWords              []string    `json:"optionalWords"`
	QueryType                  string      `json:"queryType"`
	RemoveStopWords            interface{} `json:"removeStopWords"`
	ReplaceSynonymsInHighlight bool        `json:"replaceSynonymsInHighlight"`
	SnippetEllipsisText        string      `json:"snippetEllipsisText"`
	TypoTolerance              string      `json:"typoTolerance"`
}
