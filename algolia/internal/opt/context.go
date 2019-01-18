package opt

import "context"

func ExtractContext(opts ...interface{}) context.Context {
	for _, opt := range opts {
		if v, ok := opt.(context.Context); ok {
			return v
		}
	}
	return context.Background()
}
