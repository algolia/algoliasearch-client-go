package iterator

type sliceIterator struct {
	slice []interface{}
	pos   int
}

func newIteratorFromSlice(slice []interface{}) *sliceIterator {
	return &sliceIterator{
		slice: slice,
		pos:   -1,
	}
}

func (it *sliceIterator) Next(opts ...interface{}) (interface{}, error) {
	it.pos++
	if it.slice == nil || len(it.slice) == 0 || it.pos >= len(it.slice) {
		return nil, nil
	}
	return it.slice[it.pos], nil
}
