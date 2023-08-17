package search

import (
	"fmt"
	"io"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/errs"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/wait"
)

// Account provides methods to interact with the Algolia Search API on multiple
// indices which belong to different Algolia applications.
type Account struct{}

// NewAccount instantiates a new account able to interact with the Algolia
// Search API on multiple indices which belong to different Algolia
// applications.
func NewAccount() *Account {
	return &Account{}
}

// CopyIndex copies the full content (objects, synonyms, rules, settings) of the
// given src index into the dst one. This method can only be used with indices
// which belong to different Algolia applications. To perform the same operation
// on indices which belong to the same Algolia application, use Client.CopyIndex
// which is optimized for this use-case.
func (a *Account) CopyIndex(src, dst *Index, opts ...interface{}) (*wait.Group, error) {
	if src.GetAppID() == dst.GetAppID() {
		return nil, errs.ErrSameAppID
	}

	if _, err := dst.GetSettings(); err == nil {
		return nil, errs.ErrIndexAlreadyExists
	}

	g := wait.NewGroup()

	// Copy synonyms
	{
		it, err := src.BrowseSynonyms()
		if err != nil {
			return nil, fmt.Errorf("cannot browse source index synonyms: %v", err)
		}

		var synonyms []Synonym

		for {
			synonym, err := it.Next()
			if err != nil {
				if err == io.EOF {
					break
				} else {
					return nil, fmt.Errorf("error while iterating source index synonyms: %v", err)
				}
			}
			synonyms = append(synonyms, synonym)
		}

		if synonyms != nil {
			res, err := dst.ReplaceAllSynonyms(synonyms)
			if err != nil {
				return nil, fmt.Errorf("error while replacing destination index synonyms: %v", err)
			}
			g.Collect(res)
		}
	}

	// Copy rules
	{
		it, err := src.BrowseRules()
		if err != nil {
			return nil, fmt.Errorf("cannot browse source index rules: %v", err)
		}

		var rules []Rule

		for {
			rule, err := it.Next()
			if err != nil {
				if err == io.EOF {
					break
				} else {
					return nil, fmt.Errorf("error while iterating source index rules: %v", err)
				}
			}
			rules = append(rules, *rule)
		}
		if rules != nil {
			res, err := dst.ReplaceAllRules(rules)
			if err != nil {
				return nil, fmt.Errorf("error while replacing destination index rules: %v", err)
			}
			g.Collect(res)
		}
	}

	// Copy settings
	{
		settings, err := src.GetSettings()
		if err != nil {
			return nil, fmt.Errorf("cannot retrieve source index settings: %v", err)
		}

		res, err := dst.SetSettings(settings)
		if err != nil {
			return nil, fmt.Errorf("cannot set destination index settings: %v", err)
		}
		g.Collect(res)

	}

	// Copy objects
	{
		it, err := src.BrowseObjects()
		if err != nil {
			return nil, fmt.Errorf("cannot browse source index objects: %v", err)
		}

		var objects []interface{}
		batchSize := 1000

		for {
			object, err := it.Next()
			if err != nil {
				if err == io.EOF {
					break
				} else {
					return nil, fmt.Errorf("error while iterating source index objects: %v", err)
				}
			}
			objects = append(objects, object)

			if len(objects) >= batchSize {
				res, err := dst.SaveObjects(objects)
				if err != nil {
					return nil, fmt.Errorf("error while saving batch of objects: %v", err)
				}
				g.Collect(res)
				objects = []interface{}{}
			}
		}

		// Send the last batch
		res, err := dst.SaveObjects(objects, opts...)
		if err != nil {
			return nil, fmt.Errorf("error while saving batch of objects: %v", err)
		}
		g.Collect(res)
	}

	return g, nil
}
