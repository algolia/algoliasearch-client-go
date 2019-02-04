package opt

import "github.com/algolia/algoliasearch-client-go/algolia/opt"

func ExtractExactOnSingleWordQuery(opts ...interface{}) *opt.ExactOnSingleWordQueryOption {
	for _, o := range opts {
		if v, ok := o.(opt.ExactOnSingleWordQueryOption); ok {
			return &v
		}
	}
	return nil
}
