package opt

import "github.com/algolia/algoliasearch-client-go/algolia/opt"

func ExtractAnalyticsTags(opts ...interface{}) *opt.AnalyticsTagsOption {
	for _, o := range opts {
		if v, ok := o.(opt.AnalyticsTagsOption); ok {
			return &v
		}
	}

	return nil
}
