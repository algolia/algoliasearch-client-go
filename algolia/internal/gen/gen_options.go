// +build ignore

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

var opts = []struct {
	Name         string
	Type         string
	DefaultValue template.HTML
}{
	// bool
	{"AdvancedSyntax", "bool", "false"},
	{"AllowCompressionOfIntegerArray", "bool", "false"},
	{"AllowTyposOnNumericTokens", "bool", "true"},
	{"Analytics", "bool", "true"},
	{"AroundLatLngViaIP", "bool", "false"},
	{"AutoGenerateObjectIDIfNotExist", "bool", "false"},
	{"ClickAnalytics", "bool", "false"},
	{"CreateIfNotExists", "bool", "false"},
	{"EnablePersonalization", "bool", "false"},
	{"EnableRules", "bool", "false"},
	{"FacetingAfterDistinct", "bool", "false"},
	{"ForwardToReplicas", "bool", "false"},
	{"GetRankingInfo", "bool", "false"},
	{"PercentileComputation", "bool", "true"},
	{"ReplaceSynonymsInHighlight", "bool", "true"},
	{"RestrictHighlightAndSnippetArrays", "bool", "false"},
	{"SumOrFiltersScores", "bool", "false"},
	{"Synonyms", "bool", "true"},

	// int
	{"AroundPrecision", "int", "0"},
	{"Distinct", "int", "0"},
	{"HitsPerPage", "int", "20"},
	{"Length", "int", "0"},
	{"MaxFacetHits", "int", "10"},
	{"MaxValuesPerFacet", "int", "100"},
	{"MinimumAroundRadius", "int", "0"},
	{"MinProximity", "int", "1"},
	{"MinWordSizeFor1Typo", "int", "4"},
	{"MinWordSizeFor2Typos", "int", "8"},
	{"Offset", "int", "0"},
	{"Page", "int", "0"},
	{"PaginationLimitedTo", "int", "1000"},

	// string
	{"AroundLatLng", "string", `""`},
	{"AttributeForDistinct", "string", `""`},
	{"ExactOnSingleWordQuery", "string", `"attribute"`},
	{"Filters", "string", `"attribute"`},
	{"HighlightPostTag", "string", `"</em>"`},
	{"HighlightPreTag", "string", `"<em>"`},
	{"KeepDiacriticsOnCharacters", "string", `""`},
	{"QueryType", "string", `"prefixLast"`},
	{"RemoveWordsIfNoResults", "string", `"none"`},
	{"RuleContexts", "string", `"none"`},
	{"SeparatorsToIndex", "string", `""`},
	{"SnippetEllipsisText", "string", `"…"`},
	{"SortFacetValuesBy", "string", `"count"`},

	// []string
	{"AdvancedSyntaxFeatures", "[]string", "nil"},
	{"AlternativesAsExact", "[]string", "nil"},
	{"AnalyticsTags", "[]string", "nil"},
	{"AttributesForFaceting", "[]string", "nil"},
	{"AttributesToHighlight", "[]string", "nil"},
	{"AttributesToRetrieve", "[]string", `[]string{"*"}`},
	{"AttributesToSnippet", "[]string", "nil"},
	{"CamelCaseAttributes", "[]string", "nil"},
	{"CustomRanking", "[]string", "nil"},
	{"DisableExactOnAttributes", "[]string", "nil"},
	{"DisablePrefixOnAttributes", "[]string", "nil"},
	{"DisableTypoToleranceOnAttributes", "[]string", "nil"},
	{"DisableTypoToleranceOnWords", "[]string", "nil"},
	{"Facets", "[]string", "nil"},
	{"NumericAttributesForFiltering", "[]string", "nil"},
	{"OptionalWords", "[]string", "nil"},
	{"QueryLanguages", "[]string", "nil"},
	{"Ranking", "[]string", `[]string{"typo", "geo", "words", "filters", "proximity", "attribute", "exact", "custom"}`},
	{"Replicas", "[]string", "nil"},
	{"ResponseFields", "[]string", "nil"},
	{"RestrictSearchableAttributes", "[]string", "nil"},
	{"SearchableAttributes", "[]string", "nil"},
	{"UnretrievableAttributes", "[]string", "nil"},

	// map[string]string
	{"ExtraHeaders", "map[string]string", "nil"},
	{"ExtraURLParams", "map[string]string", "nil"},

	// map[string][]string
	{"DecompoundedAttributes", "map[string][]string", "nil"},
}

var funcMap = template.FuncMap{
	"trimPrefix": strings.TrimPrefix,
}

func main() {
	var (
		optionValueTemplate                           = template.Must(template.ParseFiles("templates/option_value.go.tmpl"))
		optionSliceTemplate                           = template.Must(template.New("option_slice.go.tmpl").Funcs(funcMap).ParseFiles("templates/option_slice.go.tmpl"))
		extractOptionTemplate                         = template.Must(template.ParseFiles("templates/extract_option.go.tmpl"))
		extractOptionTestBoolTemplate                 = template.Must(template.ParseFiles("templates/extract_option_bool_test.go.tmpl"))
		extractOptionTestIntTemplate                  = template.Must(template.ParseFiles("templates/extract_option_int_test.go.tmpl"))
		extractOptionTestStringTemplate               = template.Must(template.ParseFiles("templates/extract_option_string_test.go.tmpl"))
		extractOptionTestStringSliceTemplate          = template.Must(template.ParseFiles("templates/extract_option_string_slice_test.go.tmpl"))
		extractOptionTestMapStringStringTemplate      = template.Must(template.ParseFiles("templates/extract_option_map_string_string_test.go.tmpl"))
		extractOptionTestMapStringStringSliceTemplate = template.Must(template.ParseFiles("templates/extract_option_map_string_string_slice_test.go.tmpl"))
	)

	for _, opt := range opts {
		var err error

		// This step generate a single algolia/opt/NAME.go option file where name is set to use the
		// opt.Name field.

		filename := camelCaseToFilename(opt.Name)
		switch opt.Type {
		case "bool", "int", "string", "map[string]string", "map[string][]string":
			err = generateFile(optionValueTemplate, opt, "../../opt/"+filename)
		case "[]string":
			err = generateFile(optionSliceTemplate, opt, "../../opt/"+filename)
		default:
			err = fmt.Errorf("unhandled type %s", opt.Type)
		}
		if err != nil {
			fmt.Printf("error generating option file for %s: %v", opt.Name, err)
			return
		}

		// This step generate a single algolia/internal/opt/NAME_test.go option file where name is
		// set to opt.Name field. Those internal tests ensure that both the extract functions and
		// the JSON serialization/deserialization functions are working as expected.

		testFilename := strings.Replace(filename, ".go", "_test.go", -1)
		switch opt.Type {
		case "bool":
			err = generateFile(extractOptionTestBoolTemplate, opt, "../opt/"+testFilename)
		case "int":
			err = generateFile(extractOptionTestIntTemplate, opt, "../opt/"+testFilename)
		case "string":
			err = generateFile(extractOptionTestStringTemplate, opt, "../opt/"+testFilename)
		case "[]string":
			err = generateFile(extractOptionTestStringSliceTemplate, opt, "../opt/"+testFilename)
		case "map[string]string":
			err = generateFile(extractOptionTestMapStringStringTemplate, opt, "../opt/"+testFilename)
		case "map[string][]string":
			err = generateFile(extractOptionTestMapStringStringSliceTemplate, opt, "../opt/"+testFilename)
		default:
			err = fmt.Errorf("unhandled type %s", opt.Type)
		}
		if err != nil {
			fmt.Printf("error generating option test file for %s: %v", opt.Name, err)
			return
		}

	}

	for _, filename := range listFiles("../../opt") {

		// This step produce an extract function for each option found in the algolia/opt/
		// directory in the algolia/internal/opt/ directory.

		if strings.HasSuffix(filename, ".go") {
			// Some files have to be ignored because those are private types, no supposed to be used
			// directly. Hence, the extract function must no be generated for them.
			if strings.HasSuffix(filename, "composable_filter.go") {
				continue
			}
			optName := filenameToCamelCase(filename)
			err := generateFile(extractOptionTemplate, optName, "../opt/"+path.Base(filename))
			if err != nil {
				fmt.Printf("error generating extraction option file for %s: %v", optName, err)
				return
			}
		}
	}
}

func generateFile(tmpl *template.Template, data interface{}, filename string) error {
	var (
		b       bytes.Buffer
		content []byte
	)

	err := tmpl.Execute(&b, data)
	if err != nil {
		return fmt.Errorf("cannot execute template %s: %v", filename, err)
	}

	content, err = format.Source(b.Bytes())
	if err != nil {
		return fmt.Errorf("cannot format generated code from template %s: %v", filename, err)
	}

	os.Remove(filename)

	if err = ioutil.WriteFile(filename, content, 0644); err != nil {
		return fmt.Errorf("cannot write generated file from template %s: %v", filename, err)
	}

	return nil
}

func filenameToCamelCase(filename string) (camelCase string) {
	camelCase = path.Base(filename)
	camelCase = strings.TrimSuffix(camelCase, ".go")
	camelCase = strings.Replace(camelCase, "_id", "_ID", -1)
	camelCase = strings.Replace(camelCase, "_ip", "_IP", -1)
	camelCase = strings.Replace(camelCase, "_url", "_URL", -1)
	camelCase = strings.Replace(camelCase, "_", " ", -1)
	camelCase = strings.Title(camelCase)
	camelCase = strings.Replace(camelCase, " ", "", -1)
	return
}

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func camelCaseToFilename(camelCase string) (filename string) {
	filename = matchFirstCap.ReplaceAllString(camelCase, "${1}_${2}")
	filename = matchAllCap.ReplaceAllString(filename, "${1}_${2}")
	filename = strings.ToLower(filename)
	filename = filename + ".go"
	return
}

func listFiles(root string) []string {
	var filenames []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		filenames = append(filenames, path)
		return nil
	})

	if err != nil {
		fmt.Printf("error while list files from directory %s: %v\n", root, err)
		return nil
	}

	return filenames
}
