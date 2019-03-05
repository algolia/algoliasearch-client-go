package opt

import (
	"reflect"
)

func InsertOrReplaceOption(opts []interface{}, opt interface{}) (res []interface{}) {
	if opts == nil {
		return nil
	}
	if opt == nil {
		return opts
	}

	t1 := reflect.Indirect(reflect.ValueOf(opt)).Type()

	for _, o := range opts {
		t2 := reflect.Indirect(reflect.ValueOf(o)).Type()
		if t1 != t2 {
			res = append(res, o)
		}
	}
	res = append(res, opt)
	return
}
