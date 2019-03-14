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
		extractLiteralTemplate = createTemplate("templates/extract/literal.go.tmpl")
		extractMapTemplate     = createTemplate("templates/extract/map.go.tmpl")
	)

	for _, opt := range options {
		filename := generateFilename(opt.Name)
		filepath := "../opt/" + filename

		if shouldBeGenerated(filepath) {
			switch opt.DefaultValue.(type) {
			case nil, bool, int, string, []string:
				generateFile(extractLiteralTemplate, opt.Name, filepath)
			case map[string]string, map[string][]string:
				generateFile(extractMapTemplate, opt.Name, filepath)
			default:
				fmt.Printf("cannot generate extract option file for %s: unhandled type %#v", opt.Name, opt.DefaultValue)
				os.Exit(1)
			}
		}
	}
}
