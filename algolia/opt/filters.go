package opt

type filters struct{ value string }

func Filters(v string) filters {
	return filters{v}
}

func ExtractFilters(opts ...interface{}) string {
	for _, opt := range opts {
		v, ok := opt.(filters)
		if ok {
			return v.value
		}
	}
	return ""
}
