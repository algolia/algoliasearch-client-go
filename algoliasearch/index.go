package algoliasearch

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

type index struct {
	client *client
	name   string
	route  string
}

// NewIndex instantiates a new Index. The `name` parameter corresponds to the
// Algolia index's name while the `client` is used to connect to the Algolia
// API.
func NewIndex(name string, client *client) Index {
	return &index{
		client: client,
		name:   name,
		route:  "/1/indexes/" + url.QueryEscape(name),
	}
}

func (i *index) Delete() (res DeleteTaskRes, err error) {
	path := i.route
	err = i.client.request(&res, "DELETE", path, nil, write)
	return
}

func (i *index) Clear() (res UpdateTaskRes, err error) {
	path := i.route + "/clear"
	err = i.client.request(&res, "POST", path, nil, write)
	return
}

func (i *index) GetObject(objectID string, attributes []string) (object Object, err error) {
	var params Map
	if attributes != nil {
		params = Map{
			"attributes": strings.Join(attributes, ","),
		}
	}

	path := i.route + "/" + url.QueryEscape(objectID) + "?" + encodeMap(params)
	err = i.client.request(&object, "GET", path, nil, read)
	return
}

func (i *index) GetObjects(objectIDs []string) (objs []Object, err error) {
	requests := make([]map[string]string, len(objectIDs))
	for j, id := range objectIDs {
		requests[j] = map[string]string{
			"indexName": i.name,
			"objectID":  url.QueryEscape(id),
		}
	}

	body := Map{
		"requests": requests,
	}

	var res objects
	path := "/1/indexes/*/objects"
	err = i.client.request(&res, "POST", path, body, read)
	objs = res.Results
	return
}

func (i *index) DeleteObject(objectID string) (res DeleteTaskRes, err error) {
	path := i.route + "/" + url.QueryEscape(objectID)
	err = i.client.request(&res, "DELETE", path, nil, write)
	return
}

func (i *index) GetSettings() (settings Settings, err error) {
	path := i.route + "/settings?getVersion=2"
	err = i.client.request(&settings, "GET", path, nil, read)
	return
}

func (i *index) SetSettings(settings Map) (res UpdateTaskRes, err error) {
	if err = checkSettings(settings); err != nil {
		return
	}

	path := i.route + "/settings"
	err = i.client.request(&res, "PUT", path, settings, write)
	return
}

func (i *index) WaitTask(taskID int) error {
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

func (i *index) ListKeys() (keys []Key, err error) {
	var res listKeysRes

	path := i.route + "/keys"
	if err = i.client.request(&res, "GET", path, nil, read); err != nil {
		return
	}

	keys = res.Keys
	return
}

func (i *index) AddKey(ACL []string, params Map) (res AddKeyRes, err error) {
	req := duplicateMap(params)
	req["acl"] = ACL

	if err = checkKey(req); err != nil {
		return
	}

	path := i.route + "/keys"
	err = i.client.request(&res, "POST", path, req, read)
	return
}

func (i *index) UpdateKey(value string, k Key) (res UpdateKeyRes, err error) {
	path := i.route + "/keys/" + value
	err = i.client.request(&res, "PUT", path, k, read)
	return
}

func (i *index) GetKey(value string) (key Key, err error) {
	path := i.route + "/keys/" + url.QueryEscape(value)
	err = i.client.request(&key, "GET", path, nil, read)
	return
}

func (i *index) DeleteKey(value string) (res DeleteRes, err error) {
	path := i.route + "/keys/" + value
	err = i.client.request(&res, "DELETE", path, nil, write)
	return
}

func (i *index) AddObject(object Object) (res CreateObjectRes, err error) {
	path := i.route
	err = i.client.request(&res, "POST", path, object, write)
	return
}

func (i *index) UpdateObject(object Object) (res UpdateObjectRes, err error) {
	objectID, err := object.ObjectID()
	if err != nil {
		return
	}

	path := i.route + "/" + url.QueryEscape(objectID)
	err = i.client.request(&res, "PUT", path, object, write)
	return
}

func (i *index) PartialUpdateObject(object Object) (res UpdateTaskRes, err error) {
	objectID, err := object.ObjectID()
	if err != nil {
		return
	}

	path := i.route + "/" + url.QueryEscape(objectID) + "/partial"
	err = i.client.request(&res, "POST", path, object, write)
	return
}

func (i *index) AddObjects(objects []Object) (res BatchRes, err error) {
	var operations []BatchOperation

	if operations, err = newBatchOperations(objects, "addObject"); err == nil {
		res, err = i.Batch(operations)
	}

	return
}

func (i *index) UpdateObjects(objects []Object) (res BatchRes, err error) {
	var operations []BatchOperation

	if operations, err = newBatchOperations(objects, "updateObject"); err == nil {
		res, err = i.Batch(operations)
	}

	return
}

func (i *index) PartialUpdateObjects(objects []Object) (res BatchRes, err error) {
	var operations []BatchOperation

	if operations, err = newBatchOperations(objects, "partialUpdateObject"); err == nil {
		res, err = i.Batch(operations)
	}

	return
}

func (i *index) DeleteObjects(objectIDs []string) (res BatchRes, err error) {
	objects := make([]Object, len(objectIDs))

	for j, id := range objectIDs {
		objects[j] = Object{
			"objectID": id,
		}
	}

	var operations []BatchOperation
	if operations, err = newBatchOperations(objects, "deleteObject"); err == nil {
		res, err = i.Batch(operations)
	}

	return
}

func (i *index) Batch(operations []BatchOperation) (res BatchRes, err error) {
	body := map[string][]BatchOperation{
		"requests": operations,
	}

	path := i.route + "/batch"
	err = i.client.request(&res, "POST", path, body, write)
	return
}

func (i *index) Copy(name string) (UpdateTaskRes, error) {
	return i.operation(name, "copy")
}

func (i *index) Move(name string) (UpdateTaskRes, error) {
	return i.operation(name, "move")
}

// operation performs the `op` operation on the underlying index and names the
// resulting new index `name`. The `op` operation can be either `copy` or
// `move`.
func (i *index) operation(dst, op string) (res UpdateTaskRes, err error) {
	o := IndexOperation{
		Destination: dst,
		Operation:   op,
	}

	path := i.route + "/operation"
	err = i.client.request(&res, "POST", path, o, write)
	return
}

func (i *index) GetStatus(taskID int) (res TaskStatusRes, err error) {
	path := i.route + fmt.Sprintf("/task/%d", taskID)
	err = i.client.request(&res, "GET", path, nil, read)
	return
}

func (i *index) SearchSynonyms(query string, types []string, page, hitsPerPage int) (synonyms []Synonym, err error) {
	body := Map{
		"query":       query,
		"type":        strings.Join(types, ","),
		"page":        page,
		"hitsPerPage": hitsPerPage,
	}

	path := i.route + "/synonyms/search"
	var res SearchSynonymsRes
	err = i.client.request(&res, "POST", path, body, search)

	if err == nil {
		synonyms = res.Hits
	}

	return
}

func (i *index) GetSynonym(objectID string) (s Synonym, err error) {
	path := i.route + "/synonyms/" + url.QueryEscape(objectID)
	err = i.client.request(&s, "GET", path, nil, read)
	return
}

func (i *index) AddSynonym(objectID string, synonym Synonym, forwardToSlaves bool) (res UpdateTaskRes, err error) {
	params := Map{
		"forwardToSlaves": forwardToSlaves,
	}

	path := i.route + "/synonyms/" + url.QueryEscape(objectID) + "?" + encodeMap(params)
	err = i.client.request(&res, "PUT", path, synonym, write)
	return
}

func (i *index) DeleteSynonym(objectID string, forwardToSlaves bool) (res DeleteTaskRes, err error) {
	params := Map{
		"forwardToSlaves": forwardToSlaves,
	}

	path := i.route + "/synonyms/" + url.QueryEscape(objectID) + "?" + encodeMap(params)
	err = i.client.request(&res, "DELETE", path, nil, write)
	return
}

func (i *index) ClearSynonyms(forwardToSlaves bool) (res UpdateTaskRes, err error) {
	params := Map{
		"forwardToSlaves": forwardToSlaves,
	}

	path := i.route + "/synonyms/clear?" + encodeMap(params)
	err = i.client.request(&res, "POST", path, nil, write)
	return
}

func (i *index) BatchSynonyms(synonyms []Synonym, replaceExistingSynonyms, forwardToSlaves bool) (res UpdateTaskRes, err error) {
	params := Map{
		"replaceExistingSynonyms": replaceExistingSynonyms,
		"forwardToSlaves":         forwardToSlaves,
	}

	path := i.route + "/synonyms/batch?" + encodeMap(params)
	err = i.client.request(&res, "POST", path, synonyms, write)
	return
}

func (i *index) Browse(params Map) (res BrowseRes, err error) {
	if err = checkQuery(params); err != nil {
		return
	}

	path := i.route + "/browse?" + encodeMap(params)
	err = i.client.request(&res, "GET", path, nil, read)
	return
}

func (i *index) BrowseAll(params Map) (it IndexIterator, err error) {
	if err = checkQuery(params); err != nil {
		return
	}

	it, err = newIndexIterator(i, params)
	return
}

func (i *index) Search(query string, params Map) (res QueryRes, err error) {
	copy := duplicateMap(params)
	copy["query"] = query

	if err = checkQuery(copy); err != nil {
		return
	}

	req := Map{
		"params": encodeMap(copy),
	}

	path := i.route + "/query"
	err = i.client.request(&res, "POST", path, req, search)
	return
}

func (i *index) DeleteByQuery(query string, params Map) (res BatchRes, err error) {
	copy := duplicateMap(params)
	copy["attributesToRetrieve"] = []string{"objectID"}
	copy["hitsPerPage"] = 1000
	copy["query"] = query
	copy["distinct"] = 0

	// Retrieve the iterator to browse the results
	var it IndexIterator
	if it, err = i.BrowseAll(copy); err != nil {
		return
	}

	// Iterate through all the objectIDs
	var hit Map
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
