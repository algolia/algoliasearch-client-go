// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
)

// ExtractMinWordSizefor2Typos returns the first found MinWordSizefor2TyposOption from the
// given variadic arguments or nil otherwise.
func ExtractMinWordSizefor2Typos(opts ...interface{}) *opt.MinWordSizefor2TyposOption {
	for _, o := range opts {
		if v, ok := o.(*opt.MinWordSizefor2TyposOption); ok {
			return v
		}
	}
	return nil
}