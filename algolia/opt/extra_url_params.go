package opt

type ExtraURLParamsOption struct {
	urlParams map[string]string
}

func ExtraURLParams(urlParams map[string]string) ExtraURLParamsOption {
	return ExtraURLParamsOption{
		urlParams: urlParams,
	}
}

func (opt ExtraURLParamsOption) Get() map[string]string {
	return opt.urlParams
}

// Because ExtraURLParamsOption is not intended to be serialized/deserialized,
// the json.Marshaler and json.Unmarshaler need not to be implemented.
