package opt

import "context"

func ExtractContext(opts ...interface{}) context.Context {
	for _, opt := range opts {
		v, ok := opt.(context.Context)
		if ok {
			return v
		}
	}
	return context.Background()
}
