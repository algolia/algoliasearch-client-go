//+build ignore

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

var settings = []struct {
	Name                          string
	BackwardCompatibleAlternative string
	OverrideJSONTag               string
}{ // Attributes
	{"SearchableAttributes", "AttributesToRetrieve", ""},
	{"AttributesForFaceting", "", ""},
	{"UnretrievableAttributes", "", ""},
	{"AttributesToRetrieve", "", ""},

	// Ranking
	{"Ranking", "", ""},
	{"CustomRanking", "", ""},
	{"Replicas", "Slaves", ""},

	// Faceting
	{"MaxValuesPerFacet", "", ""},
	{"SortFacetValuesBy", "", ""},

	// Highlighting - snippeting
	{"AttributesToHighlight", "", ""},
	{"AttributesToSnippet", "", ""},
	{"HighlightPreTag", "", ""},
	{"HighlightPostTag", "", ""},
	{"SnippetEllipsisText", "", ""},
	{"RestrictHighlightAndSnippetArrays", "", ""},

	// Pagination
	{"HitsPerPage", "", ""},
	{"PaginationLimitedTo", "", ""},

	// Typos
	{"MinWordSizeFor1Typo", "", "minWordSizefor1Typo"},
	{"MinWordSizeFor2Typos", "", "minWordSizefor2Typos"},
	{"TypoTolerance", "", ""},
	{"AllowTyposOnNumericTokens", "", ""},
	{"DisableTypoToleranceOnAttributes", "", ""},
	{"DisableTypoToleranceOnWords", "", ""},
	{"SeparatorsToIndex", "", ""},

	// Languages
	{"IgnorePlurals", "", ""},
	{"RemoveStopWords", "", ""},
	{"CamelCaseAttributes", "", ""},
	{"DecompoundedAttributes", "", ""},
	{"KeepDiacriticsOnCharacters", "", ""},
	{"QueryLanguages", "", ""},

	// Query strategy
	{"QueryType", "", ""},
	{"RemoveWordsIfNoResults", "", ""},
	{"AdvancedSyntax", "", ""},
	{"OptionalWords", "", ""},
	{"DisablePrefixOnAttributes", "", ""},
	{"DisableExactOnAttributes", "", ""},
	{"ExactOnSingleWordQuery", "", ""},
	{"AlternativesAsExact", "", ""},

	// Query rules
	{"EnableRules", "", ""},

	// Performance
	{"NumericAttributesForFiltering", "NumericAttributesToIndex", ""},
	{"AllowCompressionOfIntegerArray", "", ""},

	// Advanced
	{"AttributeForDistinct", "", ""},
	{"Distinct", "", ""},
	{"ReplaceSynonymsInHighlight", "", ""},
	{"MinProximity", "", ""},
	{"ResponseFields", "", ""},
}

var funcMap = template.FuncMap{
	"lowerFirstLetter": lowerFirstLetter,
}

func lowerFirstLetter(s string) string {
	if len(s) == 0 {
		return ""
	}
	return strings.ToLower(s[0:1]) + s[1:]
}

func main() {
	var (
		settingsTemplate = template.Must(template.New("settings.go.tmpl").Funcs(funcMap).ParseFiles("templates/settings.go.tmpl"))
		filename         = "../../search/settings.go"
		b                bytes.Buffer
		content          []byte
	)

	err := settingsTemplate.Execute(&b, settings)
	if err != nil {
		fmt.Printf("cannot execute template %s: %v", settingsTemplate.Name(), err)
		return
	}

	content, err = format.Source(b.Bytes())
	if err != nil {
		fmt.Printf("cannot format generated code from template %s: %v", settingsTemplate.Name(), err)
		return
	}

	os.Remove(filename)

	if err = ioutil.WriteFile(filename, content, 0644); err != nil {
		fmt.Printf("cannot write generated file from template %s: %v", settingsTemplate.Name(), err)
		return
	}
}
