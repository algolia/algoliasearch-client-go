package opt

import (
	"encoding/json"
	"reflect"
)

type IgnorePluralsOption struct {
	ignorePlurals bool
	languages     []string
}

func IgnorePlurals(v bool) *IgnorePluralsOption {
	return &IgnorePluralsOption{ignorePlurals: v}
}

func IgnorePluralsFor(languages ...string) *IgnorePluralsOption {
	return &IgnorePluralsOption{languages: languages}
}

func (o IgnorePluralsOption) Get() (bool, []string) {
	return o.ignorePlurals, o.languages
}

func (o IgnorePluralsOption) MarshalJSON() ([]byte, error) {
	if len(o.languages) > 0 {
		return json.Marshal(o.languages)
	}
	return json.Marshal(o.ignorePlurals)
}

func (o *IgnorePluralsOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	if err := json.Unmarshal(data, &o.languages); err == nil {
		return nil
	}

	return json.Unmarshal(data, &o.ignorePlurals)
}

func (o *IgnorePluralsOption) Equal(o2 *IgnorePluralsOption) bool {
	if o2 == nil {
		return o.ignorePlurals == false && len(o.languages) == 0
	}
	return reflect.DeepEqual(o, o2)
}

func IgnorePluralsEqual(o1, o2 *IgnorePluralsOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
