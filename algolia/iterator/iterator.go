package iterator

import "reflect"

type Iterator interface {
	Next(opts ...interface{}) (interface{}, error)
}

func New(itf interface{}) Iterator {
	// First, if the given interface already is an iterator.Iterator,
	// we return it as-is.
	if it, ok := itf.(Iterator); ok {
		return it
	}

	value := reflect.ValueOf(itf)
	switch value.Kind() {
	case reflect.Array | reflect.Slice:
		// In case the objects are sent as a slice or an array, we iterate over
		// them to produce a []interface{} which we then wrap on an
		// iterator.sliceIterator.
		var slice []interface{}
		for i := 0; i < value.Len(); i++ {
			slice = append(slice, value.Index(i).Interface())
		}
		return newIteratorFromSlice(slice)
	default:
		// Otherwise, the objects is most probably a single object. In that case,
		// we consider it as a single object and produce a slice with this object
		// as the sole element, which we also wrap in an iterator.Iterator.
		return newIteratorFromSlice([]interface{}{itf})
	}
}
