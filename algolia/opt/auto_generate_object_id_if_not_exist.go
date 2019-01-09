package opt

type autoGenerateObjectIDIfNotExist struct{ value bool }

func AutoGenerateObjectIDIfNotExist(v bool) autoGenerateObjectIDIfNotExist {
	return autoGenerateObjectIDIfNotExist{v}
}

func ExtractAutoGenerateObjectIDIfNotExist(opts ...interface{}) bool {
	for _, opt := range opts {
		v, ok := opt.(autoGenerateObjectIDIfNotExist)
		if ok {
			return v.value
		}
	}
	return false
}
