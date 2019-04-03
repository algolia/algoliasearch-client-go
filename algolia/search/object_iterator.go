package search

import (
	"encoding/json"
	"fmt"
)

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

func (it *ObjectIterator) Next(opts ...interface{}) (interface{}, error) {
	// Abort if the user call `Next()` on a IndexIterator that has been
	// initialized without being able to load the first page.
	if len(it.page.Hits) == 0 {
		return nil, NoMoreHitsErr
	}

	var err error

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
			return nil, err
		}
	}

	res := it.page.Hits[it.pos]
	it.pos++

	if len(opts) > 0 {
		data, err := json.Marshal(res)
		if err != nil {
			return nil, fmt.Errorf("cannot unmarshal next object: raw object cannot be marshalled: %v", err)
		}
		err = json.Unmarshal(data, opts[0])
		if err != nil {
			return nil, fmt.Errorf("cannot unmarshal next object: raw object cannot be unmarshalled to %#v: %v", opts[0], err)
		}
	}

	return res, nil
}
