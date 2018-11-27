package algoliasearch

import "fmt"

type accountClient struct{}

func NewAccountClient() AccountClient {
	return &accountClient{}
}

func (a *accountClient) CopyIndex(src, dst Index) ([]int, error) {
	return a.CopyIndexWithRequestOptions(src, dst, nil)
}

func (a *accountClient) CopyIndexWithRequestOptions(src, dst Index, opts *RequestOptions) ([]int, error) {
	if src.GetAppID() == dst.GetAppID() {
		return nil, SameAppIDErr
	}

	if _, err := dst.GetSettingsWithRequestOptions(opts); err == nil {
		return nil, IndexAlreadyExistsErr
	}

	var taskIDs []int

	// Copy synonyms
	{
		it := NewSynonymIterator(src)

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
			synonyms = append(synonyms, *synonym)
		}

		if synonyms != nil {
			res, err := dst.ReplaceAllSynonymsWithRequestOptions(synonyms, opts)
			if err != nil {
				return nil, fmt.Errorf("error while replacing destination index synonyms: %v", err)
			}
			taskIDs = append(taskIDs, res.TaskID)
		}
	}

	// Copy rules
	{
		it := NewRuleIterator(src)

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
			res, err := dst.ReplaceAllRulesWithRequestOptions(rules, opts)
			if err != nil {
				return nil, fmt.Errorf("error while replacing destination index rules: %v", err)
			}
			taskIDs = append(taskIDs, res.TaskID)
		}
	}

	// Copy settings
	{
		settings, err := src.GetSettingsWithRequestOptions(opts)
		if err != nil {
			return nil, fmt.Errorf("cannot retrieve source index settings: %v", err)
		}

		res, err := dst.SetSettingsWithRequestOptions(settings.ToMap(), opts)
		if err != nil {
			return nil, fmt.Errorf("cannot set destination index settings: %v", err)
		}
		taskIDs = append(taskIDs, res.TaskID)

	}

	// Copy objects
	{
		it, err := src.BrowseAllWithRequestOptions(nil, opts)
		if err != nil {
			return nil, fmt.Errorf("cannot browse source index objects: %v", err)
		}

		var objects []Object
		batchSize := 1000

		for {
			res, err := it.Next()
			if err != nil {
				if err == NoMoreHitsErr {
					break
				} else {
					return nil, fmt.Errorf("error while iterating source index objects: %v", err)
				}
			}
			objects = append(objects, Object(res))

			if len(objects) >= batchSize {
				res, err := dst.AddObjectsWithRequestOptions(objects, opts)
				if err != nil {
					return nil, fmt.Errorf("error while saving batch of objects: %v", err)
				}
				taskIDs = append(taskIDs, res.TaskID)
				objects = []Object{}
			}
		}

		// Send the last batch
		res, err := dst.AddObjectsWithRequestOptions(objects, opts)
		if err != nil {
			return nil, fmt.Errorf("error while saving batch of objects: %v", err)
		}
		taskIDs = append(taskIDs, res.TaskID)
		objects = []Object{}
	}

	return taskIDs, nil
}
