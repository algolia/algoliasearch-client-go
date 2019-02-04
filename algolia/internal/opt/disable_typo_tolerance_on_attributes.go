package opt

import "github.com/algolia/algoliasearch-client-go/algolia/opt"

func ExtractDisableTypoToleranceOnAttributes(opts ...interface{}) *opt.DisableTypoToleranceOnAttributesOption {
	for _, o := range opts {
		if v, ok := o.(opt.DisableTypoToleranceOnAttributesOption); ok {
			return &v
		}
	}
	return nil
}
