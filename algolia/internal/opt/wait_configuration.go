package opt

import (
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
)

// ExtractScopes returns the first found ScopesOption from the
// given variadic arguments or nil otherwise.
func ExtractWaitConfiguration(opts ...interface{}) *opt.WaitConfigurationOption {
	for _, o := range opts {
		if v, ok := o.(*opt.WaitConfigurationOption); ok {
			return v
		}
	}
	return opt.DefaultWaitConfiguration()
}
