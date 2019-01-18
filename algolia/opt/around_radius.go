package opt

import (
	"encoding/json"
	"fmt"

	"github.com/algolia/algoliasearch-client-go/algolia/errs"
)

type AroundRadiusOption struct {
	meters int
	isAll  bool
}

func AroundRadius(meters int) AroundRadiusOption {
	return AroundRadiusOption{meters: meters}
}

func AroundRadiusAll() AroundRadiusOption {
	return AroundRadiusOption{isAll: true}
}

func (o AroundRadiusOption) Get() (int, string) {
	if o.isAll {
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
