package algoliasearch

import (
	"errors"
)

type IndexIterator struct {
  answer interface{}
  params string
  pos int
  index *Index
}

func (it *IndexIterator) Next() (interface{}, error) {
	var err error
	for (err == nil) {
		hits := it.answer.(map[string]interface{})["hits"].([]interface{})
		if it.pos < len(hits) {
			it.pos++
			return hits[it.pos - 1], nil
		}
		if cursor, ok := it.GetCursor(); ok && len(cursor) > 0 {
			err = it.loadNextPage()
			continue
		}
		return nil, errors.New("End of the index reached")
	}
	return nil, err
}

func (it *IndexIterator) GetCursor() (string, bool) {
	cursor, ok := it.answer.(map[string]interface{})["cursor"]
	cursorStr := ""
	if (cursor != nil) {
		cursorStr = cursor.(string)
	}
	return cursorStr, ok
}

func (it *IndexIterator) loadNextPage() (error) {
	it.pos = 0
	cursor, ok := it.GetCursor()
	if (ok && len(cursor) != 0) {
		cursor = "&cursor=" + cursor	
	} else {
		cursor = ""
	}
	answer, err := it.index.client.transport.request("GET", "/1/indexes/" + it.index.nameEncoded + "/browse?" + it.params + cursor, nil, read) 
	it.answer = answer
	return err
}