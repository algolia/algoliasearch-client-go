package opt

import (
	"encoding/json"
	"reflect"
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

func (o composableFilterOption) Get() [][]string {
	return o.filters
}

func (o composableFilterOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.filters)
}

func (o *composableFilterOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.filters)
}

func (o *composableFilterOption) Equal(o2 *composableFilterOption) bool {
	if o2 == nil {
		return len(o.filters) == 0
	}
	return reflect.DeepEqual(o, o2)
}
