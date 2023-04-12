package opt

import (
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
)

// ExtractWaitConfiguration returns the first found wait configuration from....
// given variadic arguments or nil otherwise.
func ExtractWaitConfiguration(opts ...interface{}) *opt.WaitConfigurationOption {
	for _, o := range opts {
		if v, ok := o.(*opt.WaitConfigurationOption); ok {
			return v
		}
	}
	return opt.DefaultWaitConfiguration()
}
