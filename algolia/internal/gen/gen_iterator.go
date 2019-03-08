//+build ignore

package main

import "fmt"

func main() {
	iteratorTemplate := createTemplate("templates/iterator.go.tmpl")

	for _, kind := range []string{
		"rule",
	} {
		filepath := fmt.Sprintf("../../search/%s_iterator.go", kind)
		generateFile(iteratorTemplate, kind, filepath)
	}
}
