// +build ignore

package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

// Generates all option files and their associated tests (serialization,
// deserialization and extraction using the Extract* functions) in the following
// directories:
//
//   - algolia/opt/*.go               (option files)
//   - algolia/internal/opt/*_test.go (option test files)
//
func main() {
	var (
		optionTestBoolTemplate                     = createTemplate("templates/option_test/bool.go.tmpl")
		optionTestIntTemplate                      = createTemplate("templates/option_test/int.go.tmpl")
		optionTestStringTemplate                   = createTemplate("templates/option_test/string.go.tmpl")
		optionTestStringSliceTemplate              = createTemplate("templates/option_test/string_slice.go.tmpl")
		optionTestMapStringStringTemplate          = createTemplate("templates/option_test/map_string_string.go.tmpl")
		optionTestMapStringStringSliceTemplate     = createTemplate("templates/option_test/map_string_string_slice.go.tmpl")
		optionTestMapStringMapStringStringTemplate = createTemplate("templates/option_test/map_string_map_string_string.go.tmpl")
		optionTestMapStringMapStringBoolTemplate   = createTemplate("templates/option_test/map_string_map_string_bool.go.tmpl")
		optionTestMapStringInterfaceTemplate       = createTemplate("templates/option_test/map_string_interface_slice.go.tmpl")

		optionLiteralTemplate     = createTemplate("templates/option/literal.go.tmpl")
		optionMapTemplate         = createTemplate("templates/option/map.go.tmpl")
		optionStringSliceTemplate = createTemplate("templates/option/string_slice.go.tmpl")
	)

	for _, opt := range options {
		if opt.DefaultValue == nil {
			continue
		}

		data := struct {
			Name         string
			Type         string
			DefaultValue string
		}{
			strings.Title(opt.Name),
			reflect.TypeOf(opt.DefaultValue).String(),
			convertInterfaceToString(opt.DefaultValue),
		}

		filename := generateFilename(opt.Name)
		filepath := "../../opt/" + filename
		testFilepath := "../opt/" + strings.Replace(filename, ".go", "_test.go", -1)

		// 1.
		// This step generates a single algolia/opt/NAME.go option file where
		// NAME is set to use the opt.Name field.

		if shouldBeGenerated(filepath) {
			switch opt.DefaultValue.(type) {
			case bool, int, string:
				generateFile(optionLiteralTemplate, data, filepath)
			case map[string]string, map[string][]string, map[string]interface{}, map[string]map[string]string, map[string]map[string]bool:
				generateFile(optionMapTemplate, data, filepath)
			case []string:
				generateFile(optionStringSliceTemplate, data, filepath)
			default:
				fmt.Printf("cannot generate option file for %s: unhandled type %#v", opt.Name, opt.DefaultValue)
				os.Exit(1)
			}
		}

		// 2.
		// This step generates a single algolia/internal/opt/NAME_test.go option
		// test file where NAME is set to opt.Name field. Those internal tests
		// ensure that both the Extract* functions and the JSON
		// serialization/deserialization functions are working as expected.

		if shouldBeGenerated(testFilepath) {
			switch opt.DefaultValue.(type) {
			case bool:
				generateFile(optionTestBoolTemplate, data, testFilepath)
			case int:
				generateFile(optionTestIntTemplate, data, testFilepath)
			case string:
				generateFile(optionTestStringTemplate, data, testFilepath)
			case []string:
				generateFile(optionTestStringSliceTemplate, data, testFilepath)
			case map[string]string:
				generateFile(optionTestMapStringStringTemplate, data, testFilepath)
			case map[string][]string:
				generateFile(optionTestMapStringStringSliceTemplate, data, testFilepath)
			case map[string]map[string]string:
				generateFile(optionTestMapStringMapStringStringTemplate, data, testFilepath)
			case map[string]interface{}:
				generateFile(optionTestMapStringInterfaceTemplate, data, testFilepath)
			case map[string]map[string]bool:
				generateFile(optionTestMapStringMapStringBoolTemplate, data, testFilepath)
			default:
				fmt.Printf("cannot generate option test file for %s: unhandled type %#v", opt.Name, opt.DefaultValue)
				os.Exit(1)
			}
		}
	}
}
