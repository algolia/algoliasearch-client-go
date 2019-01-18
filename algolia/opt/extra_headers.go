package opt

type ExtraHeadersOption struct {
	headers map[string]string
}

func ExtraHeaders(headers map[string]string) ExtraHeadersOption {
	return ExtraHeadersOption{
		headers: headers,
	}
}

func (opt ExtraHeadersOption) Get() map[string]string {
	return opt.headers
}

// Because ExtraHeadersOption is not intended to be serialized/deserialized, the
// json.Marshaler and json.Unmarshaler need not to be implemented.
