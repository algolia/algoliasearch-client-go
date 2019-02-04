package opt

import "github.com/algolia/algoliasearch-client-go/algolia/opt"

func ExtractRuleContexts(opts ...interface{}) *opt.RuleContextsOption {
	for _, o := range opts {
		if v, ok := o.(opt.RuleContextsOption); ok {
			return &v
		}
	}
	return nil
}
