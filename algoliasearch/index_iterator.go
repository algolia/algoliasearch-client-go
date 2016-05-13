package algoliasearch

import "errors"

// IndexIterator is used by the BrowseAll functions to iterate over all the
// records of an index (or a subset according to what the query was).
type IndexIterator struct {
	cursor string
	index  *Index
	page   BrowseRes
	params Map
	pos    int
}

// NewIndexIterator instantiates a IndexIterator on the `index` and according
// to the given `params`. It is also trying to load the first page of results
// and return an error if something goes wrong.
func NewIndexIterator(index *Index, params Map) (it IndexIterator, err error) {
	it = IndexIterator{
		index:  index,
		params: duplicateMap(params),
		pos:    0,
	}
	err = it.loadNextPage()
	return
}

// Next returns the next record each time is is called. Subsequent pages of
// results are automatically loaded and an error is returned if a problem
// arises. When the last element has been reached, an error is returned with
// the following message: "No more hits".
func (it *IndexIterator) Next() (res Map, err error) {
	// Abort if the user call `Next()` on a IndexIterator that has been
	// initialized without being able to load the first page.
	if len(it.page.Hits) == 0 {
		err = errors.New("No more hits")
		return
	}

	// If the last element of the page has been reached, the next one is loaded
	// or returned an error if the last element of the last page has already
	// been returned.
	if it.pos == len(it.page.Hits) {
		if it.cursor == "" {
			err = errors.New("No more hits")
		} else {
			err = it.loadNextPage()
		}

		if err != nil {
			return
		}
	}

	res = it.page.Hits[it.pos]
	it.pos++

	return
}

// loadNextPage is used internally to load the next page of results, using the
// underlying Browse cursor.
func (it *IndexIterator) loadNextPage() (err error) {
	// Update the cursor for each new page except for the first one
	if it.cursor != "" {
		it.params["cursor"] = it.cursor
	}

	if it.page, err = it.index.Browse(it.params); err != nil {
		return
	}

	// Return an error if the newly loaded pages contains no results
	if len(it.page.Hits) == 0 {
		err = errors.New("No more hits")
		return
	}

	it.cursor = it.page.Cursor
	it.pos = 0
	return
}
