//+build ignore

package main

// Generates the settings.go file which contains the main Settings struct used
// by GetSettings/SetSettings along with its JSON serialization/deserialization
// methods. Old settings name are also correctly deserialized and used to fill
// the appropriate and more recent equilavent. For instance, if `slaves` is
// found in the JSON payload of the GetSettings response from the engine, it is
// automatically deserialized into the Replicas field of the Settings struct.
func main() {
	type setting struct {
		Name                          string
		BackwardCompatibleAlternative string
		DefaultValue                  string
	}

	var (
		settings         []setting
		settingsTemplate = createTemplate("templates/settings.go.tmpl")
		filepath         = "../../search/settings.go"
	)

	for _, opt := range options {
		if isSettings(opt.Kind) {
			settings = append(settings, setting{
				Name:                          opt.Name,
				BackwardCompatibleAlternative: opt.BackwardCompatibleAlternative,
				DefaultValue:                  convertInterfaceToString(opt.DefaultValue),
			})
		}
	}

	generateFile(settingsTemplate, settings, filepath)
}
