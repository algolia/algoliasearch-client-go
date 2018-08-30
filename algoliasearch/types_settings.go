package algoliasearch

import (
	"fmt"
	"os"
	"reflect"
)

// Settings is the structure returned by `GetSettigs` to ease the use of the
// index settings.
type Settings struct {
	// Attributes
	SearchableAttributes    []string `json:"searchableAttributes"`
	AttributesForFaceting   []string `json:"attributesForFaceting"`
	UnretrievableAttributes []string `json:"unretrievableAttributes"`
	AttributesToRetrieve    []string `json:"attributesToRetrieve"`
	// Ranking
	Ranking       []string `json:"ranking"`
	CustomRanking []string `json:"customRanking"`
	Replicas      []string `json:"replicas"`
	// Faceting
	MaxValuesPerFacet int    `json:"maxValuesPerFacet"`
	SortFacetValuesBy string `json:"sortFacetValuesBy"`
	// Highlighting / Snippeting
	AttributesToHighlight             []string `json:"attributesToHighlight"`
	AttributesToSnippet               []string `json:"attributesToSnippet"`
	HighlightPreTag                   string   `json:"highlightPreTag"`
	HighlightPostTag                  string   `json:"highlightPostTag"`
	SnippetEllipsisText               string   `json:"snippetEllipsisText"`
	RestrictHighlightAndSnippetArrays bool     `json:"restrictHighlightAndSnippetArrays"`
	// Pagination
	HitsPerPage         int `json:"hitsPerPage"`
	PaginationLimitedTo int `json:"paginationLimitedTo"`
	// Typos
	MinWordSizefor1Typo              int         `json:"minWordSizefor1Typo"`
	MinWordSizefor2Typos             int         `json:"minWordSizefor2Typos"`
	TypoTolerance                    interface{} `json:"typoTolerance"` // string or bool
	AllowTyposOnNumericTokens        bool        `json:"allowTyposOnNumericTokens"`
	IgnorePlurals                    interface{} `json:"ignorePlurals"` // []interface{} (actually a []string) or bool
	DisableTypoToleranceOnAttributes []string    `json:"disableTypoToleranceOnAttributes"`
	DisableTypoToleranceOnWords      []string    `json:"disableTypoToleranceOnWords"`
	SeparatorsToIndex                string      `json:"separatorsToIndex"`
	// Query strategy
	QueryType                 string      `json:"queryType"`
	RemoveWordsIfNoResults    string      `json:"removeWordsIfNoResults"`
	AdvancedSyntax            bool        `json:"advancedSyntax"`
	OptionalWords             []string    `json:"optionalWords"`
	RemoveStopWords           interface{} `json:"removeStopWords"` // []interface{} (actually a []string) or bool
	DisablePrefixOnAttributes []string    `json:"disablePrefixOnAttributes"`
	DisableExactOnAttributes  []string    `json:"disableExactOnAttributes"`
	ExactOnSingleWordQuery    string      `json:"exactOnSingleWordQuery"`
	// Query rules
	EnableRules bool `json:"enableRules"`
	// Performance
	NumericAttributesForFiltering  []string `json:"numericAttributesForFiltering"`
	AllowCompressionOfIntegerArray bool     `json:"allowCompressionOfIntegerArray"`
	// Advanced
	AttributeForDistinct       string              `json:"attributeForDistinct"`
	Distinct                   interface{}         `json:"distinct"` // float64 (actually an int) or bool
	ReplaceSynonymsInHighlight bool                `json:"replaceSynonymsInHighlight"`
	MinProximity               int                 `json:"minProximity"`
	ResponseFields             []string            `json:"responseFields"`
	MaxFacetHits               int                 `json:"maxFacetHits"`
	CamelCaseAttributes        []string            `json:"camelCaseAttributes"`
	DecompoundedAttributes     map[string][]string `json:"decompoundedAttributes"`
	KeepDiacriticsOnCharacters string              `json:"keepDiacriticsOnCharacters"`
	// Backward-compatibility
	AttributesToIndex        []string `json:"attributesToIndex"`
	QueryLanguages           []string `json:"queryLanguages"`
	NumericAttributesToIndex []string `json:"numericAttributesToIndex"`
	Slaves                   []string `json:"slaves"`
}

// clean sets the zero-value fields of any `Settings struct` generated
// by `GetSettings`.
func (s *Settings) clean() {
	// interface{} values
	if s.TypoTolerance == nil {
		s.TypoTolerance = "true"
	}
	if s.IgnorePlurals == nil {
		s.IgnorePlurals = false
	}
	if s.RemoveStopWords == nil {
		s.RemoveStopWords = false
	}
	if s.Distinct == nil {
		s.Distinct = false
	}
	// integer values
	if s.MaxFacetHits == 0 {
		s.MaxFacetHits = 10
	}
	// string values
	if s.SortFacetValuesBy == "" {
		s.SortFacetValuesBy = "count"
	}
}

// ToMap produces a `Map` corresponding to the `Settings struct`. It should
// only be used when it's needed to pass a `Settings struct` to `SetSettings`,
// typically when one needs to copy settings between two indices.
func (s *Settings) ToMap() Map {
	// Guarantee that zero-value fields are correctly set.
	s.clean()

	// Add all fields except:
	//  - TypoTolerance string or bool
	//  - IgnorePlurals []interface{} or bool
	//  - RemoveStopWords []interface{} or bool
	//  - Distinct float64 or bool
	m := Map{
		// Attributes
		"searchableAttributes":    s.SearchableAttributes,
		"attributesForFaceting":   s.AttributesForFaceting,
		"unretrievableAttributes": s.UnretrievableAttributes,
		"attributesToRetrieve":    s.AttributesToRetrieve,
		// Ranking
		"ranking":       s.Ranking,
		"customRanking": s.CustomRanking,
		"replicas":      s.Replicas,
		// Faceting
		"maxValuesPerFacet": s.MaxValuesPerFacet,
		"sortFacetValuesBy": s.SortFacetValuesBy,
		// Highlighting / Snippeting
		"attributesToHighlight":             s.AttributesToHighlight,
		"attributesToSnippet":               s.AttributesToSnippet,
		"highlightPreTag":                   s.HighlightPreTag,
		"highlightPostTag":                  s.HighlightPostTag,
		"snippetEllipsisText":               s.SnippetEllipsisText,
		"restrictHighlightAndSnippetArrays": s.RestrictHighlightAndSnippetArrays,
		// Pagination
		"hitsPerPage":         s.HitsPerPage,
		"paginationLimitedTo": s.PaginationLimitedTo,
		// Typos
		"minWordSizefor1Typo":              s.MinWordSizefor1Typo,
		"minWordSizefor2Typos":             s.MinWordSizefor2Typos,
		"allowTyposOnNumericTokens":        s.AllowTyposOnNumericTokens,
		"disableTypoToleranceOnAttributes": s.DisableTypoToleranceOnAttributes,
		"disableTypoToleranceOnWords":      s.DisableTypoToleranceOnWords,
		"separatorsToIndex":                s.SeparatorsToIndex,
		// Query strategy
		"queryType":                 s.QueryType,
		"removeWordsIfNoResults":    s.RemoveWordsIfNoResults,
		"advancedSyntax":            s.AdvancedSyntax,
		"optionalWords":             s.OptionalWords,
		"disablePrefixOnAttributes": s.DisablePrefixOnAttributes,
		"disableExactOnAttributes":  s.DisableExactOnAttributes,
		"exactOnSingleWordQuery":    s.ExactOnSingleWordQuery,
		// Query rules
		"enableRules": s.EnableRules,
		// Performance
		"numericAttributesForFiltering":  s.NumericAttributesForFiltering,
		"allowCompressionOfIntegerArray": s.AllowCompressionOfIntegerArray,
		// Advanced
		"attributeForDistinct":       s.AttributeForDistinct,
		"replaceSynonymsInHighlight": s.ReplaceSynonymsInHighlight,
		"minProximity":               s.MinProximity,
		"responseFields":             s.ResponseFields,
		"maxFacetHits":               s.MaxFacetHits,
		"camelCaseAttributes":        s.CamelCaseAttributes,
		"decompoundedAttributes":     s.DecompoundedAttributes,
		"keepDiacriticsOnCharacters": s.KeepDiacriticsOnCharacters,
	}

	// Remove empty string slices to avoid creating null-valued fields in the
	// JSON settings sent to the API
	var sliceAttributesToRemove []string
	for attr, value := range m {
		switch v := value.(type) {
		case []string:
			if len(v) == 0 {
				sliceAttributesToRemove = append(sliceAttributesToRemove, attr)
			}
		}
	}
	for _, attr := range sliceAttributesToRemove {
		delete(m, attr)
	}

	// Handle `TypoTolerance` separately as it may be either a `bool` or a
	// string.
	switch v := s.TypoTolerance.(type) {
	case bool:
		m["typoTolerance"] = v
	case string:
		if v == "true" {
			m["typoTolerance"] = true
		} else if v == "false" {
			m["typoTolerance"] = false
		} else {
			m["typoTolerance"] = v
		}
	}

	// Handle `IgnorePlurals` separately as it may be either a `bool` or a
	// `[]interface{}` which is in fact a `[]string`.
	switch v := s.IgnorePlurals.(type) {
	case bool:
		m["ignorePlurals"] = v
	case []interface{}:
		var languages []string
		for _, itf := range v {
			lang, ok := itf.(string)
			if ok {
				languages = append(languages, lang)
			} else {
				fmt.Fprintln(os.Stderr, "Settings.ToMap(): `ignorePlurals` slice doesn't only contain strings")
			}
		}
		if len(languages) > 0 {
			m["ignorePlurals"] = languages
		}
	default:
		fmt.Fprintf(os.Stderr, "Settings.ToMap(): Wrong type for `ignorePlurals`: %v\n", reflect.TypeOf(s.IgnorePlurals))
	}

	// Handle `RemoveStopWords` separately as it may be either a `bool` or a
	// `[]interface{}` which is in fact a `[]string`.
	switch v := s.RemoveStopWords.(type) {
	case bool:
		m["removeStopWords"] = v
	case []interface{}:
		var languages []string
		for _, itf := range v {
			lang, ok := itf.(string)
			if ok {
				languages = append(languages, lang)
			} else {
				fmt.Fprintln(os.Stderr, "Settings.ToMap(): `removeStopWords` slice doesn't only contain strings")
			}
		}
		if len(languages) > 0 {
			m["removeStopWords"] = languages
		}
	default:
		fmt.Fprintf(os.Stderr, "Settings.ToMap(): Wrong type for `removeStopWords`: %v\n", reflect.TypeOf(s.RemoveStopWords))
	}

	// Handle `Distinct` separately as it may be either a `bool` or a `float64`
	// which is in fact a `int`.
	switch v := s.Distinct.(type) {
	case bool:
		m["distinct"] = v
	case float64:
		m["distinct"] = int(v)
	}

	return m
}
