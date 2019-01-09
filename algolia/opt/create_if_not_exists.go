package opt

type createIfNotExists struct{ value bool }

func CreateIfNotExists(v bool) createIfNotExists {
	return createIfNotExists{v}
}

func ExtractCreateIfNotExists(opts ...interface{}) bool {
	for _, opt := range opts {
		v, ok := opt.(createIfNotExists)
		if ok {
			return v.value
		}
	}
	return false
}
