package algoliasearch

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

// Index is the structure used to manipulate an Algolia index.
type Index struct {
	client *Client
	name   string
	route  string
}

// NewIndex instantiates a new Index. The `name` parameter corresponds to the
// Algolia index's name while the `client` is used to connect to the Algolia
// API.
func NewIndex(name string, client *Client) *Index {
	return &Index{
		client: client,
		name:   name,
		route:  "/1/indexes/" + url.QueryEscape(name),
	}
}

// Delete deletes the Algolia index.
func (i *Index) Delete() (res DeleteTaskRes, err error) {
	path := i.route
	err = i.client.request(&res, "DELETE", path, nil, write)
	return
}

// Clear removes every record from the Algolia index.
func (i *Index) Clear() (res UpdateTaskRes, err error) {
	path := i.route + "/clear"
	err = i.client.request(&res, "POST", path, nil, write)
	return
}

// GetObject retrieves the object as an interface representing the JSON-encoded
// object. The `objectID` is used to uniquely identify the object in the index
// while `attributes` contains the list of attributes to retrieve.
func (i *Index) GetObject(objectID string, attributes []string) (object Object, err error) {
	params := Params{
		"attributes": strings.Join(attributes, ","),
	}

	path := i.route + "/" + url.QueryEscape(objectID) + "?" + encodeParams(params)
	err = i.client.request(&object, "GET", path, nil, read)
	return
}

// GetObjects retrieves the objects identified according to their `objectIDs`.
func (i *Index) GetObjects(objectIDs []string) (objects []Object, err error) {
	requests := make([]map[string]string, len(objectIDs))
	for j, id := range objectIDs {
		requests[j] = map[string]string{
			"indexName": i.name,
			"objectID":  url.QueryEscape(id),
		}
	}

	body := map[string]interface{}{
		"requests": requests,
	}

	path := i.route + "/*/objects"
	err = i.client.request(&objects, "POST", path, body, read)
	return
}

// DeleteObject deletes an object from the index that is uniquely identified by
// its `objectID`.
func (i *Index) DeleteObject(objectID string) (res DeleteTaskRes, err error) {
	path := i.route + "/" + url.QueryEscape(objectID)
	err = i.client.request(&res, "DELETE", path, nil, write)
	return
}

// GetSettings retrieves the index' settings.
func (i *Index) GetSettings() (settings Settings, err error) {
	path := i.route + "/settings?getVersion=2"
	err = i.client.request(&settings, "GET", path, nil, read)
	return
}

// SetSettings changes the index settings.
func (i *Index) SetSettings(settings map[string]interface{}) (res UpdateTaskRes, err error) {
	if err = checkSettings(settings); err != nil {
		return
	}

	path := i.route + "/settings"
	err = i.client.request(&res, "PUT", path, settings, write)
	return
}

// WaitTask stops the current execution until the task identified by its
// `taskID` is finished. The waiting time between each check starts at 1s and
// is increased by a factor of 2 at each retry (but is bounded at around
// 20min).
func (i *Index) WaitTask(taskID int) error {
	var res TaskStatusRes
	var err error

	var maxDuration time.Duration = time.Second
	var sleepDuration time.Duration

	for {
		if res, err = i.GetStatus(taskID); err != nil {
			return err
		}

		if res.Status == "published" {
			return nil
		}

		sleepDuration = randDuration(maxDuration)
		time.Sleep(sleepDuration)

		// Increase the upper boundary used to generate the sleep
		// duration
		if maxDuration < 10*time.Minute {
			maxDuration *= 2
		}
	}

	return nil
}

// ListKeys lists all the keys that can access the index.
func (i *Index) ListKeys() (keys []Key, err error) {
	var res listKeysRes

	path := i.route + "/keys"
	if err = i.client.request(&res, "GET", path, nil, read); err != nil {
		return
	}

	keys = res.Keys
	return
}

// AddKey adds the `k` key to the list of authorized keys of the index.
func (i *Index) AddKey(k Key) (res AddKeyRes, err error) {
	path := i.route + "/keys"
	err = i.client.request(&res, "POST", path, k, read)
	return
}

// UpdateKey updates the key identified by the value `k.Value` by replacing all
// the originla key's fields by the ones of `k`.
func (i *Index) UpdateKey(value string, k Key) (res UpdateKeyRes, err error) {
	path := i.route + "/keys/" + value
	err = i.client.request(&res, "PUT", path, k, read)
	return
}

// GetKey retrieves the key identified by its `value` string.
func (i *Index) GetKey(value string) (key Key, err error) {
	path := i.route + "/keys/" + value
	err = i.client.request(&key, "GET", path, nil, read)
	return
}

// DeleteKey deletes the key identified by its `value` string.
func (i *Index) DeleteKey(value string) (res DeleteRes, err error) {
	path := i.route + "/keys/" + value
	err = i.client.request(&res, "DELETE", path, nil, write)
	return
}

// AddObject adds a new record to the index.
func (i *Index) AddObject(object Object) (res CreateObjectRes, err error) {
	path := i.route
	err = i.client.request(&res, "POST", path, object, write)
	return
}

// UpdateObject replaces the record in the index matching the one given in
// parameter, according to its `objectID`.
func (i *Index) UpdateObject(object Object) (res UpdateObjectRes, err error) {
	objectID, err := object.ObjectID()
	if err != nil {
		return
	}

	path := i.route + "/" + url.QueryEscape(objectID)
	err = i.client.request(&res, "PUT", path, object, write)
	return
}

// PartialUpdateObject modifies the record in the index matching the one given
// in parameter, according to its `objectID`. However, the record is only
// partially updated i.e. only the specified attributes will be
// updated, the original record won't be replaced.
func (i *Index) PartialUpdateObject(object Object) (res UpdateTaskRes, err error) {
	objectID, err := object.ObjectID()
	if err != nil {
		return
	}

	path := i.route + "/" + url.QueryEscape(objectID) + "/partial"
	err = i.client.request(&res, "POST", path, object, write)
	return
}

// AddObjects adds several objects to the index.
func (i *Index) AddObjects(objects []Object) (BatchRes, error) {
	operations := newBatchOperations(objects, "addObject")
	return i.Batch(operations)
}

// UpdateObjects adds or replaces several objects at the same time, according
// to their respective `objectID` attribute.
func (i *Index) UpdateObjects(objects []Object) (BatchRes, error) {
	operations := newBatchOperations(objects, "updateObject")
	return i.Batch(operations)
}

// PartialUpdateObjects partially updates several objects at the same time,
// according to their respective `objectID` attribute.
func (i *Index) PartialUpdateObjects(objects []Object) (BatchRes, error) {
	operations := newBatchOperations(objects, "partialUpdateObject")
	return i.Batch(operations)
}

// DeleteObjects deletes several objects at the same time, according to their
// respective `objectID` attribute.
func (i *Index) DeleteObjects(objectIDs []string) (BatchRes, error) {
	objects := make([]Object, len(objectIDs))

	for j, id := range objectIDs {
		objects[j] = Object{
			"objectID": url.QueryEscape(id),
		}
	}

	operations := newBatchOperations(objects, "deleteObject")
	return i.Batch(operations)
}

// Batch processes all the specified `operations` in a batch manner. The
// operations's actions could be one of the following:
//   - addObject
//   - updateObject
//   - partialUpdateObject
//   - partialUpdateObjectNoCreate
//   - deleteObject
//   - clear
// For more informations, please refer to the official REST API documentation
// available here: https://www.algolia.com/doc/rest#batch-write-operations.
func (i *Index) Batch(operations []BatchOperation) (res BatchRes, err error) {
	body := map[string][]BatchOperation{
		"requests": operations,
	}

	path := i.route + "/batch"
	err = i.client.request(&res, "POST", path, body, write)
	return
}

// Copy copies the index into a new one called `name`.
func (i *Index) Copy(name string) (UpdateTaskRes, error) {
	return i.operation(name, "copy")
}

// Move renames the index into `name`.
func (i *Index) Move(name string) (UpdateTaskRes, error) {
	return i.operation(name, "move")
}

// operation performs the `op` operation on the underlying index and names the
// resulting new index `name`. The `op` operation can be either `copy` or
// `move`.
func (i *Index) operation(dst, op string) (res UpdateTaskRes, err error) {
	o := IndexOperation{
		Destination: dst,
		Operation:   op,
	}

	path := i.route + "/operation"
	err = i.client.request(&res, "POST", path, o, write)
	return
}

// GetStatus returns the status of a task given its ID `taskID`.
func (i *Index) GetStatus(taskID int) (res TaskStatusRes, err error) {
	path := i.route + fmt.Sprintf("/task/%d", taskID)
	err = i.client.request(&res, "GET", path, nil, read)
	return
}

// SearchSynonyms returns the synonyms matching `query` whose types match
// `types`. To retrieve the first page, `page` should be set to 0.
// `hitsPerPage` specifies how many synonym sets will be returned per page.
func (i *Index) SearchSynonyms(query string, types []string, page, hitsPerPage int) (synonyms Synonyms, err error) {
	body := map[string]interface{}{
		"query":       query,
		"type":        strings.Join(types, ","),
		"page":        page,
		"hitsPerPage": hitsPerPage,
	}

	path := i.route + "/synonyms/search"
	var rawSynonyms map[string]interface{}
	err = i.client.request(&rawSynonyms, "POST", path, body, search)

	if hits, ok := rawSynonyms["hits"]; ok {
		synonyms, err = generateSynonyms(hits)
	} else {
		err = fmt.Errorf("Cannot retrieve the `hits` field")
	}

	return
}

// GetSynonym retrieves the synonym identified by `objectID`.
func (i *Index) GetSynonym(objectID string) (s Synonym, err error) {
	var rawSynonym map[string]interface{}

	path := i.route + "/synonyms/" + url.QueryEscape(objectID)
	i.client.request(&rawSynonym, "GET", path, nil, read)

	if err != nil {
		return
	}

	s, err = generateSynonym(rawSynonym)
	return
}

// AddSynonym adds the given `synonym` to be identified `objectID`. This
// addition can be forwarded to the index slaves using `forwardToSlaves`.
func (i *Index) AddSynonym(objectID string, synonym Synonym, forwardToSlaves bool) (res UpdateTaskRes, err error) {
	if err = checkSynonym(synonym); err != nil {
		return
	}

	params := Params{
		"forwardToSlaves": forwardToSlaves,
	}

	path := i.route + "/synonyms/" + url.QueryEscape(objectID) + "?" + encodeParams(params)
	err = i.client.request(&res, "PUT", path, synonym, write)
	return
}

// DeleteSynonym removes the synonym identified by `objectID`. This deletion
// can be forwarded to the index slaves of the index with `forwardToSlaves`.
func (i *Index) DeleteSynonym(objectID string, forwardToSlaves bool) (res DeleteTaskRes, err error) {
	params := Params{
		"forwardToSlaves": forwardToSlaves,
	}

	path := i.route + "/synonyms/" + url.QueryEscape(objectID) + "?" + encodeParams(params)
	err = i.client.request(&res, "DELETE", path, nil, write)
	return
}

// ClearSynonyms removes all synonyms from the index. The clear operation can
// be forwarded to the index slaves of the index using `forwardToSlaves`.
func (i *Index) ClearSynonyms(forwardToSlaves bool) (res UpdateTaskRes, err error) {
	params := Params{
		"forwardToSlaves": forwardToSlaves,
	}

	path := i.route + "/synonyms/clear?" + encodeParams(params)
	err = i.client.request(&res, "POST", path, nil, write)
	return
}

// BatchSynonyms adds all `synonyms` to the index. The index can be cleared
// before by setting `replaceExistingSynonyms` to `true`. The optional clear
// operation and the additions can be forwarded to the index slaves by setting
// `forwardToSlaves` to `true'.
func (i *Index) BatchSynonyms(synonyms []Synonym, replaceExistingSynonyms, forwardToSlaves bool) (res UpdateTaskRes, err error) {
	if err = checkSynonyms(synonyms); err != nil {
		return
	}

	params := Params{
		"replaceExistingSynonyms": replaceExistingSynonyms,
		"forwardToSlaves":         forwardToSlaves,
	}

	path := i.route + "/synonyms/batch?" + encodeParams(params)
	err = i.client.request(&res, "POST", path, synonyms, write)
	return
}

// Browse returns the hits found according to the given `params`. The result
// also contains the cursor needed to paginate the result. This is a low-level
// function, if you simply want to iterate through all the results, it is
// preferable to use BrowseAll instead. For more informations about the Browse
// endpoint, please refer to the REST API documentation:
// https://www.algolia.com/doc/rest#browse-all-index-content
func (i *Index) Browse(params Params) (res BrowseRes, err error) {
	if err = checkQuery(params); err != nil {
		return
	}

	path := i.route + "/browse?" + encodeParams(params)
	err = i.client.request(&res, "GET", path, nil, read)
	return
}

// BrowseAll returns an iterator pointing to the first result that matches the
// search query given the `params`. Calling `Next()` on the iterator will
// returns all the hits one by one, without the 1000 elements limit of the
// Search function.
func (i *Index) BrowseAll(params Params) (it IndexIterator, err error) {
	if err = checkQuery(params); err != nil {
		return
	}

	it, err = NewIndexIterator(i, params)
	return
}

// Search performs a search query according to the `query` search query and the
// given `params`.
func (i *Index) Search(query string, params Params) (res QueryRes, err error) {
	copy := duplicateMap(params)
	copy["query"] = query

	if err = checkQuery(copy); err != nil {
		return
	}

	req := map[string]interface{}{
		"params": encodeParams(copy),
	}

	path := i.route + "/query"
	err = i.client.request(&res, "POST", path, req, search)
	return
}

// DeleteByQuery finds all the records that match the `query`, according to the
// given 'params` and deletes them.
func (i *Index) DeleteByQuery(query string, params Params) (res BatchRes, err error) {
	copy := duplicateMap(params)
	copy["attributesToRetrieve"] = []string{"objectID"}
	copy["hitsPerPage"] = 1000
	copy["query"] = query

	// Retrieve the iterator to browse the results
	var it IndexIterator
	if it, err = i.BrowseAll(copy); err != nil {
		return
	}

	// Iterate through all the objectIDs
	var hit map[string]interface{}
	var objectIDs []string
	for err == nil {
		if hit, err = it.Next(); err == nil {
			objectIDs = append(objectIDs, hit["objectID"].(string))
		}
	}

	// If it errored for something else than finishing the Browse properly, an
	// error is returned.
	if err.Error() != "No more hits" {
		return
	}

	// Delete all the objects
	res, err = i.DeleteObjects(objectIDs)
	return
}
