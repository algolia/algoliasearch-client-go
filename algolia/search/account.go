package search

import (
	"fmt"

	"github.com/algolia/algoliasearch-client-go/algolia"
)

type Account struct{}

func NewAccount() *Account {
	return &Account{}
}

func (a *Account) CopyIndex(src, dst *Index, opts ...interface{}) (algolia.Waitable, error) {
	if src.GetAppID() == dst.GetAppID() {
		return nil, SameAppIDErr
	}

	if _, err := dst.GetSettings(); err == nil {
		return nil, IndexAlreadyExistsErr
	}

	await := algolia.Await()

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
				if err == NoMoreSynonymsErr {
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
			await.Collect(res)
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
				if err == NoMoreRulesErr {
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
			await.Collect(res)
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
		await.Collect(res)

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
				if err == NoMoreHitsErr {
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
				await.Collect(res)
				objects = []interface{}{}
			}
		}

		// Send the last batch
		res, err := dst.SaveObjects(objects, opts)
		if err != nil {
			return nil, fmt.Errorf("error while saving batch of objects: %v", err)
		}
		await.Collect(res)
		objects = []interface{}{}
	}

	return await, nil
}
