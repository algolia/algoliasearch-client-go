package opt

import "github.com/algolia/algoliasearch-client-go/algolia/opt"

func ExtractAttributesToRetrieve(opts ...interface{}) *opt.AttributesToRetrieveOption {
	var (
		uniqMap   = make(map[string]bool)
		uniqSlice []string
	)

	for _, o := range opts {
		if v, ok := o.(opt.AttributesToRetrieveOption); ok {
			for _, attr := range v.Get() {
				uniqMap[attr] = true
			}
		}
	}

	if len(uniqMap) == 0 {
		return nil
	}

	for attr := range uniqMap {
		uniqSlice = append(uniqSlice, attr)
	}

	attributesToRetrieve := opt.AttributesToRetrieve(uniqSlice)
	return &attributesToRetrieve
}
