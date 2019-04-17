package opt

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/algolia/algoliasearch-client-go/algolia/errs"
)

type AroundRadiusOption struct {
	meters int
	isAll  bool
}

func AroundRadius(meters int) *AroundRadiusOption {
	return &AroundRadiusOption{meters: meters}
}

func AroundRadiusAll() *AroundRadiusOption {
	return &AroundRadiusOption{isAll: true}
}

func (o *AroundRadiusOption) Get() (int, string) {
	if o == nil || o.isAll {
		return 0, "all"
	}
	return o.meters, ""
}

func (o AroundRadiusOption) MarshalJSON() ([]byte, error) {
	if o.isAll {
		return []byte(`"all"`), nil
	}

	if o.meters != 0 {
		return []byte(fmt.Sprintf("%d", o.meters)), nil
	}

	return []byte("null"), nil
}

func (o *AroundRadiusOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	var meters int
	if err := json.Unmarshal(data, &meters); err == nil {
		o.meters = meters
		return nil
	}

	var all string
	if err := json.Unmarshal(data, &all); err == nil && all == "all" {
		o.isAll = true
		return nil
	}

	return errs.ErrJSONDecode(data, "AroundRadiusOption")
}

func (o *AroundRadiusOption) Equal(o2 *AroundRadiusOption) bool {
	if o2 == nil {
		return !o.isAll && o.meters == 0
	}
	return reflect.DeepEqual(o, o2)
}

func AroundRadiusEqual(o1, o2 *AroundRadiusOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
