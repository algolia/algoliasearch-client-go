package search

import (
	"io"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/debug"
)

type RuleIterator struct {
	rules   []Rule
	browser func(int) (SearchRulesRes, error)
	nbPages int
	page    int
	pos     int
}

func newRuleIterator(browser func(page int) (SearchRulesRes, error)) (it *RuleIterator, err error) {
	it = &RuleIterator{browser: browser}
	err = it.loadNextPage()
	return
}

func (it *RuleIterator) Next(opts ...interface{}) (*Rule, error) {
	if it.pos >= len(it.rules) {
		if it.page >= it.nbPages {
			return nil, io.EOF
		}
		err := it.loadNextPage()
		if err != nil {
			return nil, err
		}
	}

	rule := it.rules[it.pos]
	it.pos++
	return &rule, nil
}

func (it *RuleIterator) loadNextPage() error {
	res, err := it.browser(it.page)
	if err != nil {
		return err
	}

	it.pos = 0
	it.page++
	it.nbPages = res.NbPages
	it.rules, err = res.Rules()
	debug.Printf("RuleIterator: new page %d/%d loaded (%d rules)\n", it.page, it.nbPages, len(it.rules))

	return err
}
