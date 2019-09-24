package search

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/iterator"
)

// ObjectIterator represents an iterator over records of an index.
//
// ObjectIterator implements the iterator.Iterator interface.
type ObjectIterator struct {
	browser func(string) (browseRes, error)
	page    browseRes
	pos     int
}

var _ iterator.Iterator = &ObjectIterator{}

func newObjectIterator(browser func(string) (browseRes, error)) (it *ObjectIterator, err error) {
	it = &ObjectIterator{browser: browser}
	it.page, err = it.browser("")
	return
}

func (it *ObjectIterator) loadNextPage() (err error) {
	if it.page, err = it.browser(it.page.Cursor); err != nil {
		return
	}

	// Return an error if the newly loaded pages contains no results
	if len(it.page.Hits) == 0 {
		err = io.EOF
		return
	}

	it.pos = 0
	return
}

// Next returns one record from the index. To directly decode the underlying
// object instead of getting it from the returned empty interface, passes the
// object to decode to as a first argument of the method, such as:
//
//     var obj struct { ... }
//     _, err := it.Next(&obj)
//
func (it *ObjectIterator) Next(opts ...interface{}) (interface{}, error) {
	// Abort if the user call `Next()` on a IndexIterator that has been
	// initialized without being able to load the first page.
	if len(it.page.Hits) == 0 {
		return nil, io.EOF
	}

	var err error

	// If the last element of the page has been reached, the next one is loaded
	// or returned an error if the last element of the last page has already
	// been returned.
	if it.pos == len(it.page.Hits) {
		if it.page.Cursor == "" {
			err = io.EOF
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
