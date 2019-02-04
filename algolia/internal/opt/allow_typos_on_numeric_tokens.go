package opt

import "github.com/algolia/algoliasearch-client-go/algolia/opt"

func ExtractAllowTyposOnNumericTokens(opts ...interface{}) *opt.AllowTyposOnNumericTokensOption {
	for _, o := range opts {
		if v, ok := o.(opt.AllowTyposOnNumericTokensOption); ok {
			return &v
		}
	}
	return nil
}
