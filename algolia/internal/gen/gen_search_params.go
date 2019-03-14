//+build ignore

package main

// Generates the search_params.go file which contains the main QueryParams
// struct used by Search and other methods.
func main() {
	var (
		settings         []Option
		settingsTemplate = createTemplate("templates/search_params.go.tmpl")
		filepath         = "../../search/search_params.go"
	)

	for _, opt := range options {
		if isSearch(opt.Kind) {
			settings = append(settings, opt)
		}
	}

	generateFile(settingsTemplate, settings, filepath)
}
