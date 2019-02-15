package algoliasearch

import (
	"fmt"
	"os"
	"reflect"
)

// Settings is the structure returned by `GetSettigs` to ease the use of the
// index settings.
type Settings struct {
	AdvancedSyntax                    bool                `json:"advancedSyntax"`
	AdvancedSyntaxFeatures            []string            `json:"advancedSyntaxFeatures"`
	AllowCompressionOfIntegerArray    bool                `json:"allowCompressionOfIntegerArray"`
	AllowTyposOnNumericTokens         bool                `json:"allowTyposOnNumericTokens"`
	AttributeForDistinct              string              `json:"attributeForDistinct"`
	AttributesForFaceting             []string            `json:"attributesForFaceting"`
	AttributesToHighlight             []string            `json:"attributesToHighlight"`
	AttributesToIndex                 []string            `json:"attributesToIndex"`
	AttributesToRetrieve              []string            `json:"attributesToRetrieve"`
	AttributesToSnippet               []string            `json:"attributesToSnippet"`
	CamelCaseAttributes               []string            `json:"camelCaseAttributes"`
	CustomRanking                     []string            `json:"customRanking"`
	DecompoundedAttributes            map[string][]string `json:"decompoundedAttributes"`
	DisableExactOnAttributes          []string            `json:"disableExactOnAttributes"`
	DisablePrefixOnAttributes         []string            `json:"disablePrefixOnAttributes"`
	DisableTypoToleranceOnAttributes  []string            `json:"disableTypoToleranceOnAttributes"`
	DisableTypoToleranceOnWords       []string            `json:"disableTypoToleranceOnWords"`
	Distinct                          interface{}         `json:"distinct"` // float64 (actually an int) or bool
	EnableRules                       bool                `json:"enableRules"`
	ExactOnSingleWordQuery            string              `json:"exactOnSingleWordQuery"`
	HighlightPostTag                  string              `json:"highlightPostTag"`
	HighlightPreTag                   string              `json:"highlightPreTag"`
	HitsPerPage                       int                 `json:"hitsPerPage"`
	IgnorePlurals                     interface{}         `json:"ignorePlurals"` // []interface{} (actually a []string) or bool
	KeepDiacriticsOnCharacters        string              `json:"keepDiacriticsOnCharacters"`
	MaxFacetHits                      int                 `json:"maxFacetHits"`
	MaxValuesPerFacet                 int                 `json:"maxValuesPerFacet"`
	MinProximity                      int                 `json:"minProximity"`
	MinWordSizefor1Typo               int                 `json:"minWordSizefor1Typo"`
	MinWordSizefor2Typos              int                 `json:"minWordSizefor2Typos"`
	NumericAttributesForFiltering     []string            `json:"numericAttributesForFiltering"`
	NumericAttributesToIndex          []string            `json:"numericAttributesToIndex"`
	OptionalWords                     []string            `json:"optionalWords"`
	PaginationLimitedTo               int                 `json:"paginationLimitedTo"`
	QueryLanguages                    []string            `json:"queryLanguages"`
	QueryType                         string              `json:"queryType"`
	Ranking                           []string            `json:"ranking"`
	RemoveStopWords                   interface{}         `json:"removeStopWords"` // []interface{} (actually a []string) or bool
	RemoveWordsIfNoResults            string              `json:"removeWordsIfNoResults"`
	ReplaceSynonymsInHighlight        bool                `json:"replaceSynonymsInHighlight"`
	Replicas                          []string            `json:"replicas"`
	ResponseFields                    []string            `json:"responseFields"`
	RestrictHighlightAndSnippetArrays bool                `json:"restrictHighlightAndSnippetArrays"`
	SearchableAttributes              []string            `json:"searchableAttributes"`
	SeparatorsToIndex                 string              `json:"separatorsToIndex"`
	Slaves                            []string            `json:"slaves"`
	SnippetEllipsisText               string              `json:"snippetEllipsisText"`
	SortFacetValuesBy                 string              `json:"sortFacetValuesBy"`
	TypoTolerance                     interface{}         `json:"typoTolerance"` // string or bool
	UnretrievableAttributes           []string            `json:"unretrievableAttributes"`
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
		"advancedSyntax":                    s.AdvancedSyntax,
		"advancedSyntaxFeatures":            s.AdvancedSyntaxFeatures,
		"allowCompressionOfIntegerArray":    s.AllowCompressionOfIntegerArray,
		"allowTyposOnNumericTokens":         s.AllowTyposOnNumericTokens,
		"attributeForDistinct":              s.AttributeForDistinct,
		"attributesForFaceting":             s.AttributesForFaceting,
		"attributesToHighlight":             s.AttributesToHighlight,
		"attributesToIndex":                 s.AttributesToIndex,
		"attributesToRetrieve":              s.AttributesToRetrieve,
		"attributesToSnippet":               s.AttributesToSnippet,
		"camelCaseAttributes":               s.CamelCaseAttributes,
		"customRanking":                     s.CustomRanking,
		"decompoundedAttributes":            s.DecompoundedAttributes,
		"disableExactOnAttributes":          s.DisableExactOnAttributes,
		"disablePrefixOnAttributes":         s.DisablePrefixOnAttributes,
		"disableTypoToleranceOnAttributes":  s.DisableTypoToleranceOnAttributes,
		"disableTypoToleranceOnWords":       s.DisableTypoToleranceOnWords,
		"enableRules":                       s.EnableRules,
		"exactOnSingleWordQuery":            s.ExactOnSingleWordQuery,
		"highlightPostTag":                  s.HighlightPostTag,
		"highlightPreTag":                   s.HighlightPreTag,
		"hitsPerPage":                       s.HitsPerPage,
		"keepDiacriticsOnCharacters":        s.KeepDiacriticsOnCharacters,
		"maxFacetHits":                      s.MaxFacetHits,
		"maxValuesPerFacet":                 s.MaxValuesPerFacet,
		"minProximity":                      s.MinProximity,
		"minWordSizefor1Typo":               s.MinWordSizefor1Typo,
		"minWordSizefor2Typos":              s.MinWordSizefor2Typos,
		"numericAttributesForFiltering":     s.NumericAttributesForFiltering,
		"numericAttributesToIndex":          s.NumericAttributesToIndex,
		"optionalWords":                     s.OptionalWords,
		"paginationLimitedTo":               s.PaginationLimitedTo,
		"queryType":                         s.QueryType,
		"ranking":                           s.Ranking,
		"removeWordsIfNoResults":            s.RemoveWordsIfNoResults,
		"replaceSynonymsInHighlight":        s.ReplaceSynonymsInHighlight,
		"replicas":                          s.Replicas,
		"responseFields":                    s.ResponseFields,
		"restrictHighlightAndSnippetArrays": s.RestrictHighlightAndSnippetArrays,
		"searchableAttributes":              s.SearchableAttributes,
		"separatorsToIndex":                 s.SeparatorsToIndex,
		"slaves":                            s.Slaves,
		"snippetEllipsisText":               s.SnippetEllipsisText,
		"sortFacetValuesBy":                 s.SortFacetValuesBy,
		"unretrievableAttributes":           s.UnretrievableAttributes,
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
