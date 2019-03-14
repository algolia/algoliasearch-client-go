package opt

import (
	"reflect"
)

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
