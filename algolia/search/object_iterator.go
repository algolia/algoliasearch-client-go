package search

type ObjectIterator struct {
	browser func(string) (browseRes, error)
	page    browseRes
	pos     int
}

func newObjectIterator(browser func(string) (browseRes, error)) (it *ObjectIterator, err error) {
	it = &ObjectIterator{browser: browser}
	err = it.loadNextPage()
	return
}

func (it *ObjectIterator) loadNextPage() (err error) {
	if it.page, err = it.browser(it.page.Cursor); err != nil {
		return
	}

	// Return an error if the newly loaded pages contains no results
	if len(it.page.Hits) == 0 {
		err = NoMoreHitsErr
		return
	}

	it.pos = 0
	return
}

func (it *ObjectIterator) Next() (res interface{}, err error) {
	// Abort if the user call `Next()` on a IndexIterator that has been
	// initialized without being able to load the first page.
	if len(it.page.Hits) == 0 {
		err = NoMoreHitsErr
		return
	}

	// If the last element of the page has been reached, the next one is loaded
	// or returned an error if the last element of the last page has already
	// been returned.
	if it.pos == len(it.page.Hits) {
		if it.page.Cursor == "" {
			err = NoMoreHitsErr
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
