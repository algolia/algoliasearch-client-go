//+build ignore

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

var funcMap = template.FuncMap{
	"trimPrefix": strings.TrimPrefix,
	"title":      strings.Title,
}

func createTemplate(filename string) *template.Template {
	name := path.Base(filename)
	return template.Must(template.New(name).Funcs(funcMap).ParseFiles(filename))
}

func generateFile(tmpl *template.Template, data interface{}, filepath string) {
	var (
		b       bytes.Buffer
		content []byte
	)

	err := tmpl.Execute(&b, data)
	if err != nil {
		fmt.Printf("cannot execute template %s: %v", filepath, err)
		os.Exit(1)
	}

	content, err = format.Source(b.Bytes())
	if err != nil {
		fmt.Printf("cannot format generated code from template %s: %v", filepath, err)
		os.Exit(1)
	}

	os.Remove(filepath)

	if err = ioutil.WriteFile(filepath, content, 0644); err != nil {
		fmt.Printf("cannot write generated file from template %s: %v", filepath, err)
		os.Exit(1)
	}
}

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func generateFilename(jsonName string) (filename string) {
	filename = jsonName
	filename = matchFirstCap.ReplaceAllString(filename, "${1}_${2}")
	filename = matchAllCap.ReplaceAllString(filename, "${1}_${2}")
	filename = strings.ToLower(filename) + ".go"
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
		os.Exit(1)
	}
	return filenames
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

func convertInterfaceToString(defaultValue interface{}) string {
	var s string
	switch v := defaultValue.(type) {
	case bool:
		s = fmt.Sprintf("%t", v)
	case int:
		s = fmt.Sprintf("%d", v)
	case string:
		s = fmt.Sprintf("%q", v)
	case []string, map[string]string, map[string][]string:
		s = fmt.Sprintf("%#v", v)
	default:
		fmt.Printf("cannot convert interface to string: unhandled type %#v\n", defaultValue)
		os.Exit(1)
	}
	return s
}
