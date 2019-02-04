package opt

import "github.com/algolia/algoliasearch-client-go/algolia/opt"

func ExtractAnalytics(opts ...interface{}) *opt.AnalyticsOption {
	for _, o := range opts {
		if v, ok := o.(opt.AnalyticsOption); ok {
			return &v
		}
	}
	return nil
}
