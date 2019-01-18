package opt

import (
	"strings"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func ExtractExtraHeaders(opts ...interface{}) opt.ExtraHeadersOption {
	merged := make(map[string]string)

	for _, o := range opts {
		if headers, ok := o.(opt.ExtraHeadersOption); ok {
			for key, values := range headers.Get() {
				existingValues, ok := merged[key]
				if ok {
					merged[key] = strings.Join([]string{existingValues, values}, ",")
				} else {
					merged[key] = values
				}
			}
		}
	}

	return opt.ExtraHeaders(merged)
}
