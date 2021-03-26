//+build ignore

package main

import (
	"fmt"
	"os"
)

// Generates an Extract* function for each option found in the algolia/opt/
// directory in the algolia/internal/opt/ directory.

func main() {
	var (
		extractLiteralTemplate            = createTemplate("templates/extract/literal.go.tmpl")
		extractMapStringStringTemplate    = createTemplate("templates/extract/map_string_string.go.tmpl")
		extractMapStringInterfaceTemplate = createTemplate("templates/extract/map_string_interface.go.tmpl")
	)

	for _, opt := range options {
		filename := generateFilename(opt.Name)
		filepath := "../opt/" + filename

		if shouldBeGenerated(filepath) {
			switch opt.DefaultValue.(type) {
			case nil, bool, int, string, []string, map[string][]string, map[string]map[string]string, map[string]map[string]bool:
				generateFile(extractLiteralTemplate, opt.Name, filepath)
			case map[string]string:
				generateFile(extractMapStringStringTemplate, opt.Name, filepath)
			case map[string]interface{}:
				generateFile(extractMapStringInterfaceTemplate, opt.Name, filepath)
			default:
				fmt.Printf("cannot generate extract option file for %s: unhandled type %#v", opt.Name, opt.DefaultValue)
				os.Exit(1)
			}
		}
	}
}
