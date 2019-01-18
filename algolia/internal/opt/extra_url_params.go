package opt

import (
	"strings"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func ExtractExtraURLParams(opts ...interface{}) opt.ExtraURLParamsOption {
	merged := make(map[string]string)

	for _, o := range opts {
		if urlParams, ok := o.(opt.ExtraURLParamsOption); ok {
			for key, values := range urlParams.Get() {
				existingValues, ok := merged[key]
				if ok {
					merged[key] = strings.Join([]string{existingValues, values}, ",")
				} else {
					merged[key] = values
				}
			}
		}
	}

	return opt.ExtraURLParams(merged)
}
