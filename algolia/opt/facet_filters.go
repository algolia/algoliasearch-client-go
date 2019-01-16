package opt

import (
	"bytes"
	"fmt"
	"strings"
)

type simpleFacetFilter struct {
	name  string
	value string
}

type orFacetFilters struct {
	ors []simpleFacetFilter
}

type andFacetFilters struct {
	ands []orFacetFilters
}

func (f simpleFacetFilter) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s:%s"`, f.name, f.value)), nil
}

func (f orFacetFilters) MarshalJSON() ([]byte, error) {
	var innerJSONs [][]byte
	for _, filter := range f.ors {
		innerJSON, err := filter.MarshalJSON()
		if err != nil {
			return nil, err
		}
		innerJSONs = append(innerJSONs, innerJSON)
	}
	fullJSON := bytes.Join(innerJSONs, []byte(","))
	return []byte(fmt.Sprintf("[[%s]]", fullJSON)), nil
}

func (f andFacetFilters) MarshalJSON() ([]byte, error) {
	var andJSONs []string

	for _, orFilters := range f.ands {
		var orJSONs []string

		for _, simpleFilter := range orFilters.ors {
			simpleFilterJSON, err := simpleFilter.MarshalJSON()
			if err != nil {
				return nil, err
			}
			orJSONs = append(orJSONs, string(simpleFilterJSON))
		}

		andJSONs = append(andJSONs, "["+strings.Join(orJSONs, ",")+"]")
	}

	return []byte("[" + strings.Join(andJSONs, ",") + "]"), nil
}

func FacetFilters(name, value string) simpleFacetFilter {
	return simpleFacetFilter{
		name:  name,
		value: value,
	}
}

func FacetFiltersOr(filters ...simpleFacetFilter) orFacetFilters {
	return orFacetFilters{filters}
}

func FacetFiltersAnd(filters ...interface{}) andFacetFilters {
	var ors []orFacetFilters

	for _, filter := range filters {
		switch f := filter.(type) {
		case simpleFacetFilter:
			ors = append(ors, orFacetFilters{[]simpleFacetFilter{f}})
		case orFacetFilters:
			ors = append(ors, f)
		}
	}

	return andFacetFilters{ors}
}

func ExtractFacetFilters(opts ...interface{}) interface{} {
	for _, opt := range opts {
		switch v := opt.(type) {
		case simpleFacetFilter, orFacetFilters, andFacetFilters:
			return v
		}
	}
	return nil
}
