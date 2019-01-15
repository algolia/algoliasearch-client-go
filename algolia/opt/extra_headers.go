package opt

type extraHeaders map[string]string

func ExtraHeaders(v map[string]string) extraHeaders { return extraHeaders(v) }

func ExtractExtraHeaders(opts ...interface{}) map[string]string {
	m := make(map[string]string)
	for _, opt := range opts {
		headers, ok := opt.(extraHeaders)
		if ok {
			for k, v := range headers {
				m[k] = v
			}
		}
	}
	return m
}
