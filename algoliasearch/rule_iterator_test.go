package algoliasearch

import "testing"

func TestRuleIterator(t *testing.T) {
	t.Parallel()
	_, i := initClientAndIndex(t, "TestRuleIterator")

	addOneObject(t, i)

	rules := addRules(t, i, "TestRuleIterator")

	it := NewRuleIterator(i)

	var foundRules []Rule
	var rule *Rule
	var err error

	for {
		rule, err = it.Next()
		if err != nil {
			break
		}
		foundRules = append(foundRules, *rule)
	}

	if err != NoMoreRulesErr {
		t.Fatalf("TestRuleIterator: Should have stopped iterating because of a %s error but got %s instead", NoMoreRulesErr, err)
	}

	if !ruleSlicesAreEqual(rules, foundRules) {
		t.Fatalf("TestRuleIterator: Rule slices are not equal:\n%v\n%v\n", rules, foundRules)
	}
}
