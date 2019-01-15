package opt

type extraURLParams map[string]string

func ExtraURLParams(v map[string]string) extraURLParams { return extraURLParams(v) }

func ExtractExtraURLParams(opts ...interface{}) map[string]string {
	m := make(map[string]string)
	for _, opt := range opts {
		urlParams, ok := opt.(extraURLParams)
		if ok {
			for k, v := range urlParams {
				m[k] = v
			}
		}
	}
	return m
}
