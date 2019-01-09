package opt

type createIfNotExist struct{ value bool }

func CreateIfNotExist(v bool) createIfNotExist {
	return createIfNotExist{v}
}

func ExtractCreateIfNotExist(opts ...interface{}) bool {
	for _, opt := range opts {
		v, ok := opt.(createIfNotExist)
		if ok {
			return v.value
		}
	}
	return false
}
