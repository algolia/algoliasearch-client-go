package algoliasearch

// RuleIterator is the exposed structure to iterate over all the rules of
// an index.
type RuleIterator struct {
	index       Index
	rules       []Rule
	hitsPerPage int
	page        int
	pos         int
}

// NewRuleIterator returns a new RuleIterator that will iterate over all
// the rules of the declared index.
func NewRuleIterator(index Index) *RuleIterator {
	return &RuleIterator{
		index:       index,
		rules:       nil,
		hitsPerPage: 1000,
		page:        -1,
		pos:         -1,
	}
}

// Next returns iterate to the next rule of the underlying index. Every call
// to Next should yield a different rule with a nil error until the
// algoliasearch.NoMoreRulesErr is returned which means that all the
// rules have been retrieved. If the error is of a different type, it means
// that the iteration could not have been done correctly.
func (it *RuleIterator) Next() (*Rule, error) {
	if it.rules == nil || it.pos >= len(it.rules) {
		if err := it.loadNextPage(); err != nil {
			it.reset()
			return nil, err
		}
	}

	it.pos++
	if it.pos >= len(it.rules) {
		return nil, NoMoreRulesErr
	}

	rule := it.rules[it.pos]
	rule.HighlightResult = nil
	return &rule, nil
}

func (it *RuleIterator) loadNextPage() error {
	it.pos = -1
	it.page++

	res, err := it.index.SearchRules(Map{"query": ""})
	if err != nil {
		return err
	}

	it.rules = res.Hits
	return nil
}

func (it *RuleIterator) reset() {
	it.rules = nil
	it.page = -1
	it.pos = 0
}
