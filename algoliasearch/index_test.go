package algoliasearch

import (
	"sort"
	"sync"
	"syscall"
	"testing"
)

// waitTask waits the task to be finished. If something went wrong, the
// `testing.T` variable is used to terminate the test case (call to `Fatal`).
func waitTask(t *testing.T, i Index, taskID int) {
	err := i.WaitTask(taskID)
	if err != nil {
		t.Fatalf("waitTask: Task %d not published: %s", taskID, err)
	}
}

// waitTasksAsync waits for the given `tasks` asynchronously. `waitTask` is
// caled for every taskID but everything is done concurrently.
func waitTasksAsync(t *testing.T, i Index, tasks []int) {
	var wg sync.WaitGroup

	for _, task := range tasks {
		wg.Add(1)

		go func(taskID int) {
			defer wg.Done()
			waitTask(t, i, taskID)
		}(task)
	}

	wg.Wait()
}

// initClient instantiates a new client according to the
// `ALGOLIA_APPLICATION_ID` and `ALGOLIA_API_KEY` environment variables.
func initClient(t *testing.T) Client {
	appID, haveAppID := syscall.Getenv("ALGOLIA_APPLICATION_ID")
	apiKey, haveAPIKey := syscall.Getenv("ALGOLIA_API_KEY")

	if !haveAPIKey || !haveAppID {
		t.Fatal("initClient: Missing ALGOLIA_APPLICATION_ID and/or ALGOLIA_API_KEY")
	}

	return NewClient(appID, apiKey)
}

// initIndex init the `c` client with the index called `name`. It also deletes
// the index if it was existing beforehand. It waits until the task is
// finished.
func initIndex(t *testing.T, c Client, name string) (i Index) {
	i = c.InitIndex(name).(*index)

	// List indices
	indexes, err := c.ListIndexes()
	if err != nil {
		t.Fatalf("initIndex: Cannot list existing indexes: %s", err)
	}

	// Delete index if it already exists
	for _, index := range indexes {
		if index.Name == name {
			res, err := i.Delete()
			if err != nil {
				t.Fatalf("initIndex: Cannot delete the index '%s': %s", name, err)
			}

			waitTask(t, i, res.TaskID)
		}
	}

	return
}

// initClientAndIndex is a wrapper for both the `initClient` and `initIndex`.
// Please check them for more detailed informations.
func initClientAndIndex(t *testing.T, name string) (c Client, i Index) {
	c = initClient(t)
	i = initIndex(t, c, name)

	return
}

// deleteIndex actually deletes the index and waits until the task is finished.
func deleteIndex(t *testing.T, i Index) {
	res, err := i.Delete()
	if err != nil {
		t.Fatalf("deleteIndex: Cannot delete the index: %s", err)
	}

	waitTask(t, i, res.TaskID)
}

// addOneObject is used to add a single dummy object to the index. This way, we
// make sure the index has been created (and not only initialized).
func addOneObject(t *testing.T, c Client, i Index) string {
	object := Object{"attribute": "value"}

	res, err := i.AddObject(object)
	if err != nil {
		t.Fatalf("addOneObject: Cannot add an object: %s", err)
	}

	waitTask(t, i, res.TaskID)

	return res.ObjectID
}

func TestDelete(t *testing.T) {
	c, i := initClientAndIndex(t, "TestDelete")

	addOneObject(t, c, i)

	res, err := i.Delete()
	if err != nil {
		t.Fatalf("TestDelete: Cannot delete the index: %s", err)
	}

	waitTask(t, i, res.TaskID)
}

func TestClear(t *testing.T) {
	c, i := initClientAndIndex(t, "TestClear")
	defer deleteIndex(t, i)

	objectID := addOneObject(t, c, i)

	_, err := i.GetObject(objectID, nil)
	if err != nil {
		t.Fatalf("TestClear: Cannot retrieve the object: %s, err")
	}

	res, err := i.Clear()
	if err != nil {
		t.Fatalf("TestClear: Cannot clear the index: %s, err")
	}
	waitTask(t, i, res.TaskID)

	_, err = i.GetObject(objectID, nil)
	if err == nil || err.Error() != "{\"message\":\"ObjectID does not exist\",\"status\":404}\n" {
		t.Fatalf("TestClear: Object %s should have been deleted after clear: %s", objectID, err)
	}
}

func stringSlicesAreEqual(s1, s2 []string) bool {
	sort.Strings(s1)
	sort.Strings(s2)

	for i, _ := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func settingsAreEqualByComparable(s1, s2 Settings) bool {
	return s1.AllowCompressionOfIntegerArray == s2.AllowCompressionOfIntegerArray &&
		s1.AttributeForDistinct == s2.AttributeForDistinct &&
		s1.SeparatorsToIndex == s2.SeparatorsToIndex &&
		s1.AdvancedSyntax == s2.AdvancedSyntax &&
		s1.AllowTyposOnNumericTokens == s2.AllowTyposOnNumericTokens &&
		s1.HighlightPostTag == s2.HighlightPostTag &&
		s1.HighlightPreTag == s2.HighlightPreTag &&
		s1.HitsPerPage == s2.HitsPerPage &&
		s1.IgnorePlurals == s2.IgnorePlurals &&
		s1.MaxValuesPerFacet == s2.MaxValuesPerFacet &&
		s1.MinProximity == s2.MinProximity &&
		s1.MinWordSizefor1Typo == s2.MinWordSizefor1Typo &&
		s1.MinWordSizefor2Typos == s2.MinWordSizefor2Typos &&
		s1.QueryType == s2.QueryType &&
		s1.ReplaceSynonymsInHighlight == s2.ReplaceSynonymsInHighlight &&
		s1.SnippetEllipsisText == s2.SnippetEllipsisText &&
		s1.TypoTolerance == s2.TypoTolerance
}

func settingsAreEqualByStringSlices(s1, s2 Settings) bool {
	return stringSlicesAreEqual(s1.AttributesForFaceting, s2.AttributesForFaceting) &&
		stringSlicesAreEqual(s1.AttributesToIndex, s2.AttributesToIndex) &&
		stringSlicesAreEqual(s1.CustomRanking, s2.CustomRanking) &&
		stringSlicesAreEqual(s1.NumericAttributesToIndex, s2.NumericAttributesToIndex) &&
		stringSlicesAreEqual(s1.Ranking, s2.Ranking) &&
		stringSlicesAreEqual(s1.Slaves, s2.Slaves) &&
		stringSlicesAreEqual(s1.UnretrievableAttributes, s2.UnretrievableAttributes) &&
		stringSlicesAreEqual(s1.DisableTypoToleranceOnAttributes, s2.DisableTypoToleranceOnAttributes) &&
		stringSlicesAreEqual(s1.DisableTypoToleranceOnWords, s2.DisableTypoToleranceOnWords) &&
		stringSlicesAreEqual(s1.AttributesToHighlight, s2.AttributesToHighlight) &&
		stringSlicesAreEqual(s1.AttributesToRetrieve, s2.AttributesToRetrieve) &&
		stringSlicesAreEqual(s1.AttributesToSnippet, s2.AttributesToSnippet) &&
		stringSlicesAreEqual(s1.OptionalWords, s2.OptionalWords)
}

func convertInterfaceSliceToStringSlice(in []interface{}) (out []string) {
	for i := 0; i < len(in); i++ {
		out = append(out, in[i].(string))
	}

	return
}

func settingsAreEqualByRemoveStopWords(t *testing.T, s1, s2 Settings) {
	itf1 := s1.RemoveStopWords
	itf2 := s2.RemoveStopWords

	is1, ok1 := itf1.([]interface{})
	ss2, ok2 := itf2.([]string)
	if ok1 && ok2 {
		ss1 := convertInterfaceSliceToStringSlice(is1)
		if stringSlicesAreEqual(ss1, ss2) {
			return
		} else {
			t.Fatalf("settingsAreEqualByRemoveStopWords: RemoveStopWords fields are different: %v != %v\n", ss1, ss2)
		}
	} else if !ok1 && !ok2 {
		b1, ok1 := itf1.(bool)
		b2, ok2 := itf2.(bool)

		if ok1 && ok2 {
			if b1 == b2 {
				return
			} else {
				t.Fatalf("settingsAreEqualByRemoveStopWords: RemoveStopWords fields are different: %t != %t\n", b1, b2)
			}
		}
	}

	t.Fatalf("settingsAreEqualByRemoveStopWords: RemoveStopWords fields are not typed as []string or bool: %v != %v\n", itf1, itf2)
}

func settingsAreEqualByDistinct(t *testing.T, s1, s2 Settings) {
	itf1 := s1.Distinct
	itf2 := s2.Distinct

	f1, ok1 := itf1.(float64)
	i2, ok2 := itf2.(int)
	if ok1 && ok2 {
		i1 := int(f1)
		if i1 == i2 {
			return
		} else {
			t.Fatalf("settingsAreEqualByDistinct: Distinct fields are different: %d != %d\n", i1, i2)
		}
	} else if !ok1 && !ok2 {
		b1, ok1 := itf1.(bool)
		b2, ok2 := itf2.(bool)

		if ok1 && ok2 {
			if b1 == b2 {
				return
			} else {
				t.Fatalf("settingsAreEqualByDistinct: Distinct fields are different: %t != %t\n", b1, b2)
			}
		}
	}

	t.Fatalf("settingsAreEqualByDistinct: Distinct fields are not typed as int or bool: %v != %v\n", itf1, itf2)
}

func settingsAreEqual(t *testing.T, s1, s2 Settings) {
	if !settingsAreEqualByComparable(s1, s2) {
		t.Fatalf("settingsAreEqual: Comparable fields are not equal:\n%v\n%v\n", s1, s2)
	}

	if !settingsAreEqualByStringSlices(s1, s2) {
		t.Fatalf("settingsAreEqual: String slice fields are not equal:\n%v\n%v\n", s1, s2)
	}

	settingsAreEqualByRemoveStopWords(t, s1, s2)
	settingsAreEqualByDistinct(t, s1, s2)
}

func setAndGetAndCompareSettings(t *testing.T, i Index, expectedSettings Settings, mapSettings Map) {
	res, err := i.SetSettings(mapSettings)
	if err != nil {
		t.Fatalf("TestSettings: Cannot set settings: %s", err)
	}
	waitTask(t, i, res.TaskID)

	settings, err := i.GetSettings()
	if err != nil {
		t.Fatalf("TestSettings: Cannot get settings: %s", err)
	}

	settingsAreEqual(t, settings, expectedSettings)
}

func TestSettings(t *testing.T) {
	_, i := initClientAndIndex(t, "TestSettings")
	defer deleteIndex(t, i)

	expectedSettings := Settings{
		AdvancedSyntax:                   true,
		AllowCompressionOfIntegerArray:   false,
		AllowTyposOnNumericTokens:        false,
		AttributeForDistinct:             "attribute",
		AttributesForFaceting:            []string{"attribute"},
		AttributesToHighlight:            []string{"attribute"},
		AttributesToIndex:                []string{"attribute"},
		AttributesToRetrieve:             []string{"attribute"},
		AttributesToSnippet:              []string{"attribute:20"},
		CustomRanking:                    []string{"asc(attribute)"},
		DisableTypoToleranceOnAttributes: []string{"attribute"},
		DisableTypoToleranceOnWords:      []string{"word"},
		Distinct:                         true,
		HighlightPostTag:                 "<p>",
		HighlightPreTag:                  "</p>",
		HitsPerPage:                      10,
		IgnorePlurals:                    true,
		MaxValuesPerFacet:                20,
		MinProximity:                     2,
		MinWordSizefor1Typo:              2,
		MinWordSizefor2Typos:             4,
		NumericAttributesToIndex:         []string{"attribute"},
		OptionalWords:                    []string{"optional", "words"},
		QueryType:                        "prefixAll",
		Ranking:                          []string{"typo", "geo", "words", "proximity", "attribute", "exact", "custom"},
		RemoveStopWords:                  []string{"en", "fr"},
		ReplaceSynonymsInHighlight:       false,
		SeparatorsToIndex:                "+#",
		Slaves:                           []string{},
		SnippetEllipsisText:              "...",
		TypoTolerance:                    "strict",
		UnretrievableAttributes:          []string{"unretrievable_attribute"},
	}

	mapSettings := Map{
		"advancedSyntax":                   true,
		"allowCompressionOfIntegerArray":   false,
		"allowTyposOnNumericTokens":        false,
		"attributeForDistinct":             "attribute",
		"attributesForFaceting":            []string{"attribute"},
		"attributesToHighlight":            []string{"attribute"},
		"attributesToIndex":                []string{"attribute"},
		"attributesToRetrieve":             []string{"attribute"},
		"attributesToSnippet":              []string{"attribute:20"},
		"customRanking":                    []string{"asc(attribute)"},
		"disableTypoToleranceOnAttributes": []string{"attribute"},
		"disableTypoToleranceOnWords":      []string{"word"},
		"distinct":                         true,
		"highlightPostTag":                 "<p>",
		"highlightPreTag":                  "</p>",
		"hitsPerPage":                      10,
		"ignorePlurals":                    true,
		"maxValuesPerFacet":                20,
		"minProximity":                     2,
		"minWordSizefor1Typo":              2,
		"minWordSizefor2Typos":             4,
		"numericAttributesToIndex":         []string{"attribute"},
		"optionalWords":                    []string{"optional", "words"},
		"queryType":                        "prefixAll",
		"ranking":                          []string{"typo", "geo", "words", "proximity", "attribute", "exact", "custom"},
		"removeStopWords":                  []string{"en", "fr"},
		"replaceSynonymsInHighlight":       false,
		"separatorsToIndex":                "+#",
		"slaves":                           []string{},
		"snippetEllipsisText":              "...",
		"typoTolerance":                    "strict",
		"unretrievableAttributes":          []string{"unretrievable_attribute"},
	}

	// Initial test
	setAndGetAndCompareSettings(t, i, expectedSettings, mapSettings)

	// Second test: change the values which can have a different type
	expectedSettings.RemoveStopWords = true
	mapSettings["removeStopWords"] = true
	expectedSettings.Distinct = 2
	mapSettings["distinct"] = 2
	setAndGetAndCompareSettings(t, i, expectedSettings, mapSettings)
}

func TestIndexing(t *testing.T) {
	_, i := initClientAndIndex(t, "TestIndexing")
	defer deleteIndex(t, i)

	var tasks []int

	// Set the settings
	{
		res, err := i.SetSettings(Map{
			"attributesToIndex":     []string{"company", "name"},
			"attributesForFaceting": []string{"company"},
			"customRanking":         []string{"asc(company)", "asc(name)"},
		})
		if err != nil {
			t.Fatalf("TestIndexing: Cannot set settings: %s", err)
		}
		tasks = append(tasks, res.TaskID)
	}

	// Add one object
	{
		object := Object{"name": "Facebook", "Company": "Mark Zuckerberg"}
		res, err := i.AddObject(object)
		if err != nil {
			t.Fatalf("TestIndexing: Cannot add one object: %s", err)
		}
		tasks = append(tasks, res.TaskID)
	}

	// Add multiple objects at once
	{
		objects := []Object{
			{"company": "Algolia", "name": "Julien Lemoine"},
			{"company": "Algolia", "name": "Nicolas Dessaigne"},
			{"company": "Amazon", "name": "Jeff Bezos"},
			{"company": "Apple", "name": "Steve Jobs"},
			{"company": "Apple", "name": "Steve Wozniak"},
			{"company": "Arista Networks", "name": "Jayshree Ullal"},
			{"company": "<GOOG>", "name": "Larry Page"},
			{"company": "<GOOG>", "name": "Rob Pike"},
			{"company": "<GOOG>", "name": "Sergueï Brin"},
			{"company": "Microsoft", "name": "Bill Gates"},
			{"company": "SpaceX", "name": "Elon Musk"},
			{"company": "Tesla", "name": "Elon Musk"},
			{"company": "Yahoo", "name": "Marissa Mayer"},
		}
		res, err := i.AddObjects(objects)
		if err != nil {
			t.Fatalf("TestIndexing: Cannot add multiple objects: %s", err)
		}
		tasks = append(tasks, res.TaskID)
	}

	// Add one synonym
	{
		synonym := NewSynonym("tesla", []string{"tesla", "tesla motors"})
		res, err := i.AddSynonym(synonym, true)
		if err != nil {
			t.Fatalf("TestIndexing: Cannot add one synonym: %s", err)
		}
		tasks = append(tasks, res.TaskID)
	}

	// Add multiple synonyms at once
	{
		synonyms := []Synonym{
			NewAltCorrectionSynomym("rob", []string{"robpike"}, "rob", AltCorrection1),
			NewAltCorrectionSynomym("pike", []string{"robpike"}, "pike", AltCorrection2),
			NewOneWaySynonym("julien", "speedblue", []string{"julien lemoine"}),
			NewPlaceholderSynonym("google_placeholder", "<GOOG>", []string{"Google", "GOOG"}),
		}
		res, err := i.BatchSynonyms(synonyms, false, false)
		if err != nil {
			t.Fatalf("TestIndexing: Cannot add multiple synonyms: %s", err)
		}
		tasks = append(tasks, res.TaskID)
	}

	waitTasksAsync(t, i, tasks)
}
