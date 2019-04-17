//+build ignore

package main

import "strings"

func main() {
	type option struct {
		VarName  string
		TypeName string
	}

	var (
		allOptions                []option
		optionGettersTestTemplate = createTemplate("templates/option_getters_test.go.tmpl")
		filepath                  = "../../opt/option_getters_test.go"
	)

	for _, opt := range options {
		varName := opt.Name
		typeName := strings.Title(opt.Name) + "Option"

		if varName == "type" {
			varName = "typeVar"
		}

		allOptions = append(allOptions, option{
			VarName:  varName,
			TypeName: typeName,
		})
	}

	generateFile(optionGettersTestTemplate, allOptions, filepath)
}
