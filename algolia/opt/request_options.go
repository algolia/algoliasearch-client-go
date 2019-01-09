package opt

type RequestOptions struct {
	ExtraHeaders   map[string]string
	ExtraURLParams map[string]string
}

func ExtractRequestOptions(opts ...interface{}) (options RequestOptions) {
	for _, opt := range opts {
		v, ok := opt.(RequestOptions)
		if ok {
			return v
		}
	}
	return RequestOptions{}
}
