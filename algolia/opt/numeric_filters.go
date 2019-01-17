package opt

import (
	"bytes"
	"fmt"
	"strings"
)

type simpleNumericFilter struct {
	filter string
}

type orNumericFilters struct {
	ors []simpleNumericFilter
}

type andNumericFilters struct {
	ands []orNumericFilters
}

func (f simpleNumericFilter) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, f.filter)), nil
}

func (f orNumericFilters) MarshalJSON() ([]byte, error) {
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

func (f andNumericFilters) MarshalJSON() ([]byte, error) {
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

func NumericFilters(filter string) simpleNumericFilter {
	return simpleNumericFilter{
		filter: filter,
	}
}

func NumericFiltersOr(filters ...simpleNumericFilter) orNumericFilters {
	return orNumericFilters{filters}
}

func NumericFiltersAnd(filters ...interface{}) andNumericFilters {
	var ors []orNumericFilters

	for _, filter := range filters {
		switch f := filter.(type) {
		case simpleNumericFilter:
			ors = append(ors, orNumericFilters{[]simpleNumericFilter{f}})
		case orNumericFilters:
			ors = append(ors, f)
		}
	}

	return andNumericFilters{ors}
}

func ExtractNumericFilters(opts ...interface{}) interface{} {
	for _, opt := range opts {
		switch v := opt.(type) {
		case simpleNumericFilter, orNumericFilters, andNumericFilters:
			return v
		}
	}
	return nil
}
