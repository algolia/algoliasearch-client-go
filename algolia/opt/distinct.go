package opt

import (
	"encoding/json"

	"github.com/algolia/algoliasearch-client-go/algolia/errs"
)

type DistinctOption struct {
	value int
}

func Distinct(enabled bool) *DistinctOption {
	if enabled {
		return &DistinctOption{value: 1}
	}
	return &DistinctOption{value: 0}
}

func DistinctOf(v int) *DistinctOption {
	return &DistinctOption{value: v}
}

func (o *DistinctOption) Get() (bool, int) {
	if o == nil {
		return false, 0
	}
	return o.value == 1, o.value
}

func (o DistinctOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *DistinctOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	var i int
	if err := json.Unmarshal(data, &i); err == nil {
		o.value = i
		return nil
	}

	var b bool
	if err := json.Unmarshal(data, &b); err == nil {
		if b {
			o.value = 1
		} else {
			o.value = 0
		}
		return nil
	}

	return errs.ErrJSONDecode(data, "Distinct")
}

func (o *DistinctOption) Equal(o2 *DistinctOption) bool {
	if o2 == nil {
		return o.value == 0
	}
	return o.value == o2.value
}

func DistinctEqual(o1, o2 *DistinctOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
