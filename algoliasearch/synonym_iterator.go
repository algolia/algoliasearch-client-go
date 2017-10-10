package algoliasearch

// SynonymIterator is the exposed structure to iterate over all the synonyms of
// an index.
type SynonymIterator struct {
	index       Index
	synonyms    []Synonym
	hitsPerPage int
	page        int
	pos         int
}

// NewSynonymIterator returns a new SynonymIterator that will iterate over all
// the synonyms of the declared index.
func NewSynonymIterator(index Index) *SynonymIterator {
	return &SynonymIterator{
		index:       index,
		synonyms:    nil,
		hitsPerPage: 1000,
		page:        -1,
		pos:         -1,
	}
}

// Next returns iterate to the next synonym of the underlying index. Every call
// to Next should yield a different synonym with a nil error until the
// algoliasearch.NoMoreSynonymsErr is returned which means that all the
// synonyms have been retrieved. If the error is of a different type, it means
// that the iteration could not have been done correctly.
func (it *SynonymIterator) Next() (*Synonym, error) {
	if it.synonyms == nil || it.pos >= len(it.synonyms) {
		if err := it.loadNextPage(); err != nil {
			it.reset()
			return nil, err
		}
	}

	it.pos++
	if it.pos >= len(it.synonyms) {
		return nil, NoMoreSynonymsErr
	}

	synonym := it.synonyms[it.pos]
	synonym.HighlightResult = nil
	return &synonym, nil
}

func (it *SynonymIterator) loadNextPage() error {
	it.pos = -1
	it.page++

	synonyms, err := it.index.SearchSynonyms("", nil, it.page, it.hitsPerPage)
	if err != nil {
		return err
	}

	it.synonyms = synonyms
	return nil
}

func (it *SynonymIterator) reset() {
	it.synonyms = nil
	it.page = -1
	it.pos = 0
}
