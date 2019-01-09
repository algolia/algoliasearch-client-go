package opt

type requestOptions struct {
	ExtraHeaders   map[string]string
	ExtraURLParams map[string]string
}

func newRequestOptions() requestOptions {
	return requestOptions{
		ExtraHeaders:   make(map[string]string),
		ExtraURLParams: make(map[string]string),
	}
}

func RequestOptions(extraHeaders, extraURLParams map[string]string) requestOptions {
	opts := newRequestOptions()
	if extraHeaders != nil {
		opts.ExtraHeaders = extraHeaders
	}
	if extraURLParams != nil {
		opts.ExtraURLParams = extraURLParams
	}
	return opts
}

func ExtractRequestOptions(opts ...interface{}) (options requestOptions) {
	for _, opt := range opts {
		v, ok := opt.(requestOptions)
		if ok {
			return v
		}
	}
	return requestOptions{}
}
