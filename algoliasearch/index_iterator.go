package algoliasearch

import "fmt"

type IndexIterator struct {
	cursor string
	index  *Index
	page   BrowseRes
	params map[string]interface{}
	pos    int64
}

func NewIndexIterator(index *Index, params map[string]interface{}) (it IndexIterator, err error) {
	it = IndexIterator{
		index:  index,
		params: duplicateMap(params),
		pos:    0,
	}
	err = it.loadNextPage()
	return
}

func (it *IndexIterator) Next() (res map[string]interface{}, err error) {
	// Abort if the user call `Next()` on a IndexIterator that has been
	// initialized without being able to load the first page.
	if it.page.NbHits == 0 {
		err = fmt.Errorf("No more hits")
		return
	}

	// If the last element of the page has been reached, the next one is loaded
	if it.pos == it.page.NbHits {
		if err = it.loadNextPage(); err != nil {
			return
		}
	}

	res = it.page.Hits[it.pos]
	it.pos++

	return
}

func (it *IndexIterator) loadNextPage() (err error) {
	// Update the cursor for each new page except for the first one
	if it.cursor != "" {
		it.params["cursor"] = it.cursor
	}

	if it.page, err = it.index.Browse(it.params); err != nil {
		return
	}

	// Return an error if the newly loaded pages contains no results
	if it.page.NbHits == 0 {
		err = fmt.Errorf("No more hits")
		return
	}

	it.cursor = it.page.Cursor
	it.pos = 0
	return
}
