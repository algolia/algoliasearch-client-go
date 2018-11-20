package algoliasearch

import (
	"os"
	"reflect"
	"sync"
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

// addOneObject is used to add a single dummy object to the index. This way, we
// make sure the index has been created (and not only initialized).
func addOneObject(t *testing.T, i Index) string {
	object := Object{"attribute": "value"}

	res, err := i.AddObject(object)
	if err != nil {
		t.Fatalf("addOneObject: Cannot add an object: %s", err)
	}

	waitTask(t, i, res.TaskID)

	return res.ObjectID
}

// initClient instantiates a new client according to the
// `ALGOLIA_APPLICATION_ID_1` and `ALGOLIA_ADMIN_KEY_1` environment variables.
func initClient(t *testing.T) Client {
	appID := os.Getenv("ALGOLIA_APPLICATION_ID_1")
	apiKey := os.Getenv("ALGOLIA_ADMIN_KEY_1")

	if appID == "" || apiKey == "" {
		t.Fatal("initClient: Missing ALGOLIA_APPLICATION_ID_1 and/or ALGOLIA_ADMIN_KEY_1")
	}

	return NewClient(appID, apiKey)
}

// initClient instantiates a new client according to the
// `ALGOLIA_APPLICATION_ID_2` and `ALGOLIA_ADMIN_KEY_2` environment variables.
func initClient2(t *testing.T) Client {
	appID := os.Getenv("ALGOLIA_APPLICATION_ID_2")
	apiKey := os.Getenv("ALGOLIA_ADMIN_KEY_2")

	if appID == "" || apiKey == "" {
		t.Fatal("initClient: Missing ALGOLIA_APPLICATION_ID_2 and/or ALGOLIA_ADMIN_KEY_2")
	}

	return NewClient(appID, apiKey)
}

// initMCMClient is the same as initClient but read different env vars
func initMCMClient(t *testing.T) Client {
	appID := os.Getenv("ALGOLIA_APPLICATION_ID_MCM")
	apiKey := os.Getenv("ALGOLIA_ADMIN_KEY_MCM")

	if appID == "" || apiKey == "" {
		t.Fatal("initClient: Missing ALGOLIA_APPLICATION_ID_MCM and/or ALGOLIA_ADMIN_KEY_MCM")
	}

	return NewClient(appID, apiKey)
}

// initClientWithHosts instantiates a new client according to the
// `ALGOLIA_APPLICATION_ID_1` and `ALGOLIA_ADMIN_KEY_1` environment variables and set
// one of the host to specifically timeout.
func initClientWithTimeoutHosts(t *testing.T) Client {
	appID := os.Getenv("ALGOLIA_APPLICATION_ID_1")
	apiKey := os.Getenv("ALGOLIA_ADMIN_KEY_1")

	if appID == "" || apiKey == "" {
		t.Fatal("initClient: Missing ALGOLIA_APPLICATION_ID_1 and/or ALGOLIA_ADMIN_KEY_1")
	}

	return NewClientWithHosts(appID, apiKey, []string{
		"algolia.biz",
		appID + "-1.algolianet.com",
		appID + "-2.algolianet.com",
		appID + "-2.algolianet.com",
	})
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

func initClientAndAnalytics(t *testing.T) (c Client, a Analytics) {
	c = initClient(t)
	a = c.InitAnalytics()
	return
}

// addObjectsAndSynonyms populates the given Index with several records, set
// the settings accordingly and add several synonyms. This helper is needed as
// we need to test both Synonym methods and the SynonymIterator in two
// different tests.
func addObjectsAndSynonyms(t *testing.T, i Index, testName string) []Synonym {
	var tasks []int

	t.Log(testName + ": Set the settings")
	{
		res, err := i.SetSettings(Map{
			"searchableAttributes": []string{"company"},
			"customRanking":        []string{"asc(company)"},
		})
		if err != nil {
			t.Fatalf(testName+": Cannot set settings: %s", err)
		}
		tasks = append(tasks, res.TaskID)
	}

	t.Log(testName + ": Add multiple objects at once")
	{
		objects := []Object{
			{"company": "<GOOG>"},
			{"company": "Algolia"},
			{"company": "Amazon"},
			{"company": "Apple"},
			{"company": "Arista Networks"},
			{"company": "Microsoft"},
			{"company": "SpaceX"},
			{"company": "Tesla"},
			{"company": "Yahoo"},
		}
		res, err := i.AddObjects(objects)
		if err != nil {
			t.Fatalf(testName+": Cannot add multiple objects: %s", err)
		}
		tasks = append(tasks, res.TaskID)
	}

	synonyms := []Synonym{
		NewAltCorrectionSynonym("rob", []string{"robpike"}, "rob", AltCorrection1),
		NewAltCorrectionSynonym("pike", []string{"robpike"}, "pike", AltCorrection2),
		NewOneWaySynonym("julien", "speedblue", []string{"julien lemoine"}),
		NewPlaceholderSynonym("google_placeholder", "<GOOG>", []string{"Google", "GOOG"}),
	}

	t.Log(testName + ": Add multiple synonyms at once")
	{
		res, err := i.BatchSynonyms(synonyms, false, false)
		if err != nil {
			t.Fatalf(testName+": Cannot add multiple synonyms: %s", err)
		}

		tasks = append(tasks, res.TaskID)
	}

	t.Log(testName + ": Add one synonym")
	{
		synonym := NewSynonym("tesla", []string{"tesla", "tesla motors"})
		synonyms = append(synonyms, synonym)

		res, err := i.SaveSynonym(synonym, true)
		if err != nil {
			t.Fatalf(testName+": Cannot add one synonym: %s", err)
		}

		tasks = append(tasks, res.TaskID)
	}

	t.Log(testName + ": Wait for all the previous tasks to complete")
	waitTasksAsync(t, i, tasks)

	return synonyms
}

// synonymsAreEqual returns `true` if the two synonyms are the same.
func synonymsAreEqual(s1, s2 Synonym) bool {
	return s1.ObjectID == s2.ObjectID &&
		s1.Type == s2.Type &&
		s1.Word == s2.Word &&
		s1.Input == s2.Input &&
		s1.Placeholder == s2.Placeholder &&
		stringSlicesAreEqual(s1.Corrections, s2.Corrections) &&
		stringSlicesAreEqual(s1.Synonyms, s2.Synonyms) &&
		stringSlicesAreEqual(s1.Replacements, s2.Replacements)
}

// synonymSlicesAreEqual returns `true` if the two slices contains the exact
// same synonyms. Slices don't need to be sorted.
func synonymSlicesAreEqual(synonyms1, synonyms2 []Synonym) bool {
	count := 0

	if len(synonyms1) != len(synonyms2) {
		return false
	}

	for _, s1 := range synonyms1 {
		for _, s2 := range synonyms2 {
			if synonymsAreEqual(s1, s2) {
				count++
				break
			}
		}
	}

	return count == len(synonyms1)
}

// rulesAreEqual returns `true` if the two rules are the same.
func rulesAreEqual(r1, r2 Rule) bool {
	return r1.ObjectID == r2.ObjectID &&
		r1.Description == r2.Description &&
		reflect.DeepEqual(r1.Condition, r2.Condition)
}

// ruleSlicesAreEqual returns `true` if the two slices contains the exact
// same rules. Slices don't need to be sorted.
func ruleSlicesAreEqual(rules1, rules2 []Rule) bool {
	count := 0

	if len(rules1) != len(rules2) {
		return false
	}

	for _, s1 := range rules1 {
		for _, s2 := range rules2 {
			if rulesAreEqual(s1, s2) {
				count++
				break
			}
		}
	}

	return count == len(rules1)
}

// addRules populates the given Index with several query rules. This helper is
// neede as we need to test both Rules methods and the RuleIterator in two
// different tests.
func addRules(t *testing.T, i Index, testName string) []Rule {
	var tasks []int
	var allRules []Rule

	t.Log(testName + ": Add single rule with SaveRule")
	{
		rule := Rule{
			ObjectID:  "brand_tagging",
			Condition: NewSimpleRuleCondition(Contains, "{facet:brand}"),
			Consequence: RuleConsequence{
				Params: Map{
					"automaticFacetFilters": []string{"brand"},
				},
			},
			Description: "Automatic tagging of apple queries with apple brand",
		}
		res, err := i.SaveRule(rule, false)
		if err != nil {
			t.Fatalf(testName+": Cannot perform SaveRule: %s", err)
		}
		tasks = append(tasks, res.TaskID)
		allRules = append(allRules, rule)
	}

	t.Log(testName + ": Add multiple rules with BatchRules")
	{
		rules := []Rule{
			{
				ObjectID:  "remove_js",
				Condition: NewSimpleRuleCondition(Contains, "js"),
				Consequence: RuleConsequence{
					Params: Map{
						"query": QueryIncrementalEdit{Remove: []string{"js"}},
					},
				},
				Description: "Remove `js` from every query",
			},
			{
				ObjectID:  "substitute_coffee_with_tea",
				Condition: NewSimpleRuleCondition(Contains, "coffee"),
				Consequence: RuleConsequence{
					Params: Map{"query": "tea"},
				},
				Description: "substitute `coffee` with `tea`",
			},
		}
		res, err := i.BatchRules(rules, false, false)
		if err != nil {
			t.Fatalf(testName+": Cannot perform BatchRules: %s", err)
		}
		tasks = append(tasks, res.TaskID)
		allRules = append(allRules, rules...)
	}

	t.Log(testName + ": Wait for all the previous tasks to complete")
	waitTasksAsync(t, i, tasks)

	return allRules
}
