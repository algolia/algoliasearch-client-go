//+build ignore

package main

import (
	"path"
	"strings"
)

// Generates an Extract* function for each option found in the algolia/opt/
// directory in the algolia/internal/opt/ directory.

func main() {
	var (
		extractLiteralTemplate = createTemplate("templates/extract/literal.go.tmpl")
		extractMapTemplate     = createTemplate("templates/extract/map.go.tmpl")
	)

	for _, filename := range listFiles("../../opt") {
		if !strings.HasSuffix(filename, ".go") {
			continue
		}

		// Some files have to be ignored because those are private types, not
		// supposed to be used directly. Hence, the extract function must not
		// be generated for them.

		if strings.HasSuffix(filename, "composable_filter.go") {
			continue
		}

		filepath := "../opt/" + path.Base(filename)
		optName := filenameToCamelCase(filename)

		if strings.HasSuffix(filename, "extra_headers.go") ||
			strings.HasSuffix(filename, "extra_url_params.go") {
			generateFile(extractMapTemplate, optName, filepath)
		} else {
			generateFile(extractLiteralTemplate, optName, filepath)
		}
	}
}
