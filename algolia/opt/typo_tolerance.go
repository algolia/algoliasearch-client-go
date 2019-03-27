package opt

import (
	"encoding/json"
	"fmt"

	"github.com/algolia/algoliasearch-client-go/algolia/errs"
)

type TypoToleranceOption struct {
	value string
}

func TypoTolerance(value bool) *TypoToleranceOption {
	return &TypoToleranceOption{value: fmt.Sprintf("%t", value)}
}

func TypoToleranceMin() *TypoToleranceOption {
	return &TypoToleranceOption{value: "min"}
}

func TypoToleranceStrict() *TypoToleranceOption {
	return &TypoToleranceOption{value: "strict"}
}

func (o TypoToleranceOption) Get() (bool, string) {
	if o.value == "true" {
		return true, ""
	}
	if o.value == "false" {
		return false, ""
	}
	return false, o.value
}

func (o TypoToleranceOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *TypoToleranceOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = "true"
		return nil
	}

	var v string
	if err := json.Unmarshal(data, &v); err == nil {
		o.value = v
		return nil
	}

	var b bool
	if err := json.Unmarshal(data, &b); err == nil {
		o.value = fmt.Sprintf("%t", b)
		return nil
	}

	return errs.ErrJSONDecode(data, "TypoTolerance")
}

func (o *TypoToleranceOption) Equal(o2 *TypoToleranceOption) bool {
	if o2 == nil {
		return o.value == "true"
	}
	return o.value == o2.value
}

func TypoToleranceEqual(o1, o2 *TypoToleranceOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
