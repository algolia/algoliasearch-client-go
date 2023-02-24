package opt

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/errs"
)

type composableFilterOption struct {
	filters [][]string
}

func composableFilter(filter string) composableFilterOption {
	ors := []string{filter}
	ands := [][]string{ors}
	return composableFilterOption{ands}
}

func composableFilterOr(filters ...interface{}) composableFilterOption {
	var ors []string

	for _, or := range filters {
		switch v := or.(type) {
		case string:
			ors = append(ors, v)
		case composableFilterGet:
			if len(v.Get()) == 1 && len(v.Get()[0]) == 1 {
				ors = append(ors, v.Get()[0][0])
			}
		}
	}
	return composableFilterOption{[][]string{ors}}
}

func composableFilterAnd(filters ...interface{}) composableFilterOption {
	var ands [][]string

	for _, and := range filters {
		switch v := and.(type) {
		case string:
			ands = append(ands, []string{v})
		case []string:
			ands = append(ands, v)
		case composableFilterGet:
			if len(v.Get()) == 1 {
				ands = append(ands, v.Get()[0])
			}
		}
	}

	return composableFilterOption{ands}
}

type composableFilterGet interface {
	Get() [][]string
}

func (o *composableFilterOption) Get() [][]string {
	if o == nil {
		return nil
	}
	return o.filters
}

func (o composableFilterOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.filters)
}

func (o *composableFilterOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	var (
		ok      = false
		filter  string
		ors     []string
		ands    [][]string
		options []interface{}
	)

	// Handles legacy one-string format as filters.
	// Adds outer groups as `ands`, adds inner groups as `ors` if there are any.
	//"A:1,B:2"       => [["A:1"],["B:2"]]
	//"(A:1,B:2),C:3" => [["A:1","B:2"],["C:3"]]
	//"(A:1,B:2)"     => [["A:1","B:2"]]
	if json.Unmarshal(data, &filter) == nil {
		ok = true
		replacer := strings.NewReplacer("(", " ", ")", " ")

		var start, count int
		for i, c := range filter {
			switch c {
			case '(':
				count++
			case ')':
				count--
			case ',':
				if count == 0 {
					// remove parentheses and split filters by comma
					ors = strings.Split(replacer.Replace(filter[start:i]), ",")
					ands = append(ands, ors)
					start = i + 1
				}
			}
		}

		// add last chunk of the filter
		ors = strings.Split(replacer.Replace(filter[start:]), ",")
		ands = append(ands, ors)
	}

	if !ok && json.Unmarshal(data, &options) == nil {
		ok = true

		for _, option := range options {
			switch v := option.(type) {
			case []interface{}:
				ors = []string{}

				for _, val := range v {
					ors = append(ors, fmt.Sprint(val))
				}

				ands = append(ands, ors)
			case string:
				ands = append(ands, []string{v})
			default:
				return errs.ErrJSONDecode(data, "composableFilterOption (string or []string or [][]string)")
			}
		}
	}

	if !ok {
		return errs.ErrJSONDecode(data, "composableFilterOption (string or []string or [][]string)")
	}

	var cleanANDs [][]string
	for _, ors := range ands {
		var cleanORs []string
		for _, filter := range ors {
			filter = strings.Trim(filter, " ")
			if len(filter) > 0 {
				cleanORs = append(cleanORs, filter)
			}
		}
		if len(cleanORs) > 0 {
			cleanANDs = append(cleanANDs, cleanORs)
		}
	}
	o.filters = cleanANDs
	return nil
}

func (o *composableFilterOption) Equal(o2 *composableFilterOption) bool {
	if o == nil {
		return o2 == nil || len(o2.filters) == 0
	}
	if o2 == nil {
		return o == nil || len(o.filters) == 0
	}
	return reflect.DeepEqual(o, o2)
}
