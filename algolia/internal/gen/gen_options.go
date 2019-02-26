// +build ignore

package main

import (
	"bytes"
	"fmt"
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
	{"AllowTyposOnNumericTokens", "bool", "true"},
	{"Analytics", "bool", "true"},
	{"AroundLatLngViaIP", "bool", "false"},
	{"AutoGenerateObjectIDIfNotExist", "bool", "false"},
	{"CreateIfNotExists", "bool", "false"},
	{"EnablePersonalization", "bool", "false"},
	{"EnableRules", "bool", "false"},
	{"FacetingAfterDistinct", "bool", "false"},
	{"GetRankingInfo", "bool", "false"},
	{"PercentileComputation", "bool", "true"},
	{"Synonyms", "bool", "true"},
	{"SumOrFiltersScores", "bool", "false"},
	{"RestrictHighlightAndSnippetArrays", "bool", "false"},
	{"ReplaceSynonymsInHighlight", "bool", "true"},

	// int
	{"AroundPrecision", "int", "0"},
	{"Distinct", "int", "0"},
	{"HitsPerPage", "int", "20"},
	{"Length", "int", "0"},
	{"MaxFacetHits", "int", "10"},
	{"MaxValuesPerFacet", "int", "100"},
	{"MinProximity", "int", "1"},
	{"MinWordSizeFor1Typo", "int", "4"},
	{"MinWordSizeFor2Typos", "int", "8"},
	{"MinimumAroundRadius", "int", "0"},
	{"Offset", "int", "0"},
	{"Page", "int", "0"},
	{"PaginationLimitedTo", "int", "1000"},

	// string
	{"AroundLatLng", "string", `""`},
	{"ExactOnSingleWordQuery", "string", `"attribute"`},
	{"Filters", "string", `"attribute"`},
	{"HighlightPostTag", "string", `""`},
	{"HighlightPreTag", "string", `""`},
	{"QueryType", "string", `"prefixLast"`},
	{"RemoveWordsIfNoResults", "string", `"none"`},
	{"RuleContexts", "string", `"none"`},
	{"SortFacetValuesBy", "string", `"count"`},
	{"SnippetEllipsisText", "string", `"…"`},
	{"SeparatorsToIndex", "string", `""`},

	// []string
	{"AdvancedSyntaxFeatures", "[]string", "nil"},
	{"AlternativesAsExact", "[]string", "nil"},
	{"AnalyticsTags", "[]string", "nil"},
	{"AttributesToHighlight", "[]string", "nil"},
	{"AttributesToRetrieve", "[]string", "nil"},
	{"AttributesToSnippet", "[]string", "nil"},
	{"CamelCaseAttributes", "[]string", "nil"},
	{"DisableExactOnAttributes", "[]string", "nil"},
	{"DisableTypoToleranceOnAttributes", "[]string", "nil"},
	{"DisableTypoToleranceOnWords", "[]string", "nil"},
	{"Facets", "[]string", "nil"},
	{"OptionalWords", "[]string", "nil"},
	{"QueryLanguages", "[]string", "nil"},
	{"RestrictSearchableAttributes", "[]string", "nil"},

	// map[string]string
	{"ExtraHeaders", "map[string]string", "make(map[string]string)"},
	{"ExtraURLParams", "map[string]string", "make(map[string]string)"},
}

var funcMap = template.FuncMap{
	"trimPrefix": strings.TrimPrefix,
}

func main() {
	var (
		optionValueTemplate                  = template.Must(template.ParseFiles("templates/option_value.go.tmpl"))
		optionSliceTemplate                  = template.Must(template.New("option_slice.go.tmpl").Funcs(funcMap).ParseFiles("templates/option_slice.go.tmpl"))
		extractOptionTemplate                = template.Must(template.ParseFiles("templates/extract_option.go.tmpl"))
		extractOptionTestBoolTemplate        = template.Must(template.ParseFiles("templates/extract_option_bool_test.go.tmpl"))
		extractOptionTestIntTemplate         = template.Must(template.ParseFiles("templates/extract_option_int_test.go.tmpl"))
		extractOptionTestStringTemplate      = template.Must(template.ParseFiles("templates/extract_option_string_test.go.tmpl"))
		extractOptionTestStringSliceTemplate = template.Must(template.ParseFiles("templates/extract_option_string_slice_test.go.tmpl"))
	)

	for _, opt := range opts {
		var err error

		// This step generate a single algolia/opt/NAME.go option file where name is set to use the
		// opt.Name field.

		filename := camelCaseToFilename(opt.Name)
		switch opt.Type {
		case "bool", "int", "string", "map[string]string":
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
			//err = generateFile(extractOptionTestMapTemplate, opt, "../opt/"+testFilename)
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
	var b bytes.Buffer

	err := tmpl.Execute(&b, data)
	if err != nil {
		return fmt.Errorf("cannot execute template: %v", err)
	}

	os.Remove(filename)

	if err = ioutil.WriteFile(filename, b.Bytes(), 0755); err != nil {
		return fmt.Errorf("cannot generate file: %v", err)
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
