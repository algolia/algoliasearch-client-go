package opt

import (
	"fmt"
	"reflect"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/debug"
)

// InsertOrReplaceOption inserts the given opt option into the given slice of
// options and return this new slice (original slice is left unchanged). If the
// option already existed in the original slice, it is replaced in the returned
// one.
func InsertOrReplaceOption(opts []interface{}, opt interface{}) []interface{} {
	if opts == nil && opt == nil {
		return nil
	}
	if opts == nil {
		return []interface{}{opt}
	}
	if opt == nil {
		return opts
	}

	t1 := reflect.Indirect(reflect.ValueOf(opt)).Type()

	var res []interface{}
	for _, o := range opts {
		t2 := reflect.Indirect(reflect.ValueOf(o)).Type()
		if t1 != t2 {
			res = append(res, o)
		}
	}
	return append(res, opt)
}

// InsertExtraURLParam inserts a new opt.ExtraURLParams into the given slice and
// returns this new slice (original slice is left unchanged).
func InsertExtraURLParam(opts []interface{}, k string, v interface{}) []interface{} {
	return append(opts, ExtraURLParams(map[string]string{k: convertInterfaceToString(v)}))
}

// InsertExtraHeader inserts a new opt.ExtraHeaders into the given slice and
// returns this new slice (original slice is left unchanged).
func InsertExtraHeader(opts []interface{}, k string, v interface{}) []interface{} {
	return append(opts, ExtraHeaders(map[string]string{k: convertInterfaceToString(v)}))
}

func convertInterfaceToString(itf interface{}) string {
	if itf == nil {
		return ""
	}

	switch v := itf.(type) {
	case string:
		return v
	case int:
		return fmt.Sprintf("%d", v)
	case bool:
		return fmt.Sprintf("%t", v)
	default:
		debug.Printf("cannot convert %#v to string", itf)
	}

	return ""
}
