package algoliasearch

import (
	"errors"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// Index is the structure used to manipulate an Algolia index.
type Index struct {
	client      *Client
	name        string
	nameEncoded string
}

// NewIndex instantiates a new Index. The `name` parameter corresponds to the
// Algolia index's name while the `client` is used to connect to the Algolia
// API.
func NewIndex(name string, client *Client) *Index {
	return &Index{
		client:      client,
		name:        name,
		nameEncoded: client.transport.urlEncode(name),
	}
}

// Delete deletes the Algolia index.
func (i *Index) Delete() (interface{}, error) {
	return i.client.transport.request("DELETE", "/1/indexes/"+i.nameEncoded, nil, write)
}

// Clear removes every record from the Algolia index.
func (i *Index) Clear() (interface{}, error) {
	return i.client.transport.request("POST", "/1/indexes/"+i.nameEncoded+"/clear", nil, write)
}

// GetObject retrieves the object as an interface representing the JSON-encoded
// object. The `objectID` is used to uniquely identify the object in the index
// while the `attribute` (optional) is a string containing comma-separated
// attributes that you want to retrieve. If this parameter is omitted, all the
// attributes are returned.
func (i *Index) GetObject(objectID string, attribute ...string) (interface{}, error) {
	v := url.Values{}
	if len(attribute) > 1 {
		return nil, errors.New("Too many parameters")
	} else if len(attribute) > 0 {
		v.Add("attribute", attribute[0])
	}

	return i.client.transport.request("GET", "/1/indexes/"+i.nameEncoded+"/"+i.client.transport.urlEncode(objectID)+"?"+v.Encode(), nil, read)
}

// GetObjects retrieves the objects identified by the given `objectIDs`.
func (i *Index) GetObjects(objectIDs ...string) (interface{}, error) {
	requests := make([]interface{}, len(objectIDs))
	for j, id := range objectIDs {
		requests[j] = map[string]interface{}{
			"indexName": i.name,
			"objectID":  id,
		}
	}

	body := map[string]interface{}{
		"requests": requests,
	}

	return i.client.transport.request("POST", "/1/indexes/*/objects", body, read)
}

// DeleteObject deletes an object from the index that is uniquely identified by
// its `objectID`.
func (i *Index) DeleteObject(objectID string) (interface{}, error) {
	return i.client.transport.request("DELETE", "/1/indexes/"+i.nameEncoded+"/"+i.client.transport.urlEncode(objectID), nil, write)
}

// GetSettings retrieves the index settings.
func (i *Index) GetSettings() (interface{}, error) {
	return i.client.transport.request("GET", "/1/indexes/"+i.nameEncoded+"/settings?getVersion=2", nil, read)
}

// SetSettings changes the index settings.
func (i *Index) SetSettings(settings interface{}) (interface{}, error) {
	return i.client.transport.request("PUT", "/1/indexes/"+i.nameEncoded+"/settings", settings, write)
}

// WaitTask waits for the given task to be completed. The interface given is
// typically the returned value of a call to `AddObject`.
func (i *Index) WaitTask(task interface{}) (interface{}, error) {
	if reflect.TypeOf(task).Name() == "float64" {
		return i.WaitTaskWithInit(task.(float64), 100)
	}
	return i.WaitTaskWithInit(task.(map[string]interface{})["taskID"].(float64), 100)
}

// WaitTaskWithInit waits for the task with the `taskID` to be completed. The
// `timeToWait` parameter controls the first duration, in ms, to use between
// each retry (it will be exponentiated up to 10s).
func (i *Index) WaitTaskWithInit(taskID float64, timeToWait float64) (interface{}, error) {
	var err error
	var status interface{}

	for {
		if status, err = i.getStatus(taskID); err != nil {
			return nil, err
		}

		if status.(map[string]interface{})["status"] == "published" {
			return status, nil
		}

		time.Sleep(time.Duration(timeToWait) * time.Millisecond)
		timeToWait = timeToWait * 2
		if timeToWait > 10000 {
			timeToWait = 10000
		}
	}

	return nil, errors.New("Code not reachable")
}

// ListKeys lists all the keys that can access the index.
func (i *Index) ListKeys() (interface{}, error) {
	return i.client.transport.request("GET", "/1/indexes/"+i.nameEncoded+"/keys", nil, read)
}

// GetKey returns the ACL and the validity of the given `key` API key for the
// current index.
func (i *Index) GetKey(key string) (interface{}, error) {
	return i.client.transport.request("GET", "/1/indexes/"+i.nameEncoded+"/keys/"+key, nil, read)
}

// DeleteKey deletes the `key` API key of the current index.
func (i *Index) DeleteKey(key string) (interface{}, error) {
	return i.client.transport.request("DELETE", "/1/indexes/"+i.nameEncoded+"/keys/"+key, nil, write)
}

// AddObject adds a new object to the index.
func (i *Index) AddObject(object interface{}) (interface{}, error) {
	return i.client.transport.request("POST", "/1/indexes/"+i.nameEncoded, object, write)
}

// UpdateObject modifies the record in the Algolia index matching the one given
// in parameter, according to its `objectID` value.
func (i *Index) UpdateObject(object interface{}) (interface{}, error) {
	id := object.(map[string]interface{})["objectID"]
	path := "/1/indexes/" + i.nameEncoded + "/" + i.client.transport.urlEncode(id.(string))
	return i.client.transport.request("PUT", path, object, write)
}

// PartialUpdateObject modifies the record in the Algolia index matching the
// one given in parameter, according to its `objectID` value. However, the
// record is only partially updated i.e. only the specified attributes will be
// updated.
func (i *Index) PartialUpdateObject(object interface{}) (interface{}, error) {
	id := object.(map[string]interface{})["objectID"]
	path := "/1/indexes/" + i.nameEncoded + "/" + i.client.transport.urlEncode(id.(string)) + "/partial"
	return i.client.transport.request("POST", path, object, write)
}

// AddObject adds several objects to the index.
func (i *Index) AddObjects(objects interface{}) (interface{}, error) {
	return i.sameBatch(objects, "addObject")
}

// UpdateObjects adds or updates several objects at the same time, according to
// their respective `objectID` attribute.
func (i *Index) UpdateObjects(objects interface{}) (interface{}, error) {
	return i.sameBatch(objects, "updateObject")
}

// PartialUpdateObjects partially updates several objects at the same time,
// according to their respective `objectID` attribute.
func (i *Index) PartialUpdateObjects(objects interface{}) (interface{}, error) {
	return i.sameBatch(objects, "partialUpdateObject")
}

// DeleteObjects deletes several objects at the same time, according to their
// respective `objectID` attribute.
func (i *Index) DeleteObjects(objectIDs []string) (interface{}, error) {
	objects := make([]interface{}, len(objectIDs))
	for j, id := range objectIDs {
		objects[j] = map[string]interface{}{
			"objectID": id,
		}
	}

	return i.sameBatch(objects, "deleteObject")
}

// DeleteByQuery deletes all the records that are found after performing the
// `query` search query, following the `params` parameters.
func (i *Index) DeleteByQuery(query string, params map[string]interface{}) (interface{}, error) {
	if params == nil {
		params = make(map[string]interface{})
	}
	params["attributesToRetrieve"] = "[\"objectID\"]"
	params["hitsPerPage"] = 1000
	params["distinct"] = false

	var results, task interface{}
	var err error

	if results, err = i.Search(query, params); err != nil {
		return results, err
	}

	for results.(map[string]interface{})["nbHits"].(float64) != 0 {
		objectIDs := make([]string, len(results.(map[string]interface{})["hits"].([]interface{})))

		for i := range results.(map[string]interface{})["hits"].([]interface{}) {
			hits := results.(map[string]interface{})["hits"].([]interface{})[i].(map[string]interface{})
			objectIDs[i] = hits["objectID"].(string)
		}

		if task, err = i.DeleteObjects(objectIDs); err != nil {
			return task, err
		}

		if _, err = i.WaitTask(task); err != nil {
			return nil, err
		}

		if results, err = i.Search(query, params); err != nil {
			return results, err
		}
	}

	return nil, nil
}

// Batch performs each action contained in the `actions` parameter to their
// respective object from the `objects` parameter.
func (i *Index) Batch(objects interface{}, actions []string) (interface{}, error) {
	array := objects.([]interface{})
	queries := make([]map[string]interface{}, len(array))

	for i := range array {
		queries[i] = map[string]interface{}{
			"action": actions[i],
			"body":   array[i],
		}
	}

	return i.CustomBatch(queries)
}

// CustomBatch actually performs the batch request of all `queries`.
func (i *Index) CustomBatch(queries interface{}) (interface{}, error) {
	request := map[string]interface{}{
		"requests": queries,
	}

	return i.client.transport.request("POST", "/1/indexes/"+i.nameEncoded+"/batch", request, write)
}

// Browse returns `hitsPerPage` results from the `page` page.
// Deprecated: Use `BrowseFrom` or `BrowseAll` instead.
func (i *Index) Browse(page, hitsPerPage int) (interface{}, error) {
	return i.client.transport.request("GET", "/1/indexes/"+i.nameEncoded+"/browse?page="+strconv.Itoa(page)+"&hitsPerPage="+strconv.Itoa(hitsPerPage), nil, read)
}

// BrowseFrom browses the results according to the given `params` parameters at
// the position defined by the `cursor` parameter.
func (i *Index) BrowseFrom(params interface{}, cursor string) (interface{}, error) {
	if len(cursor) != 0 {
		cursor = "&cursor=" + i.client.transport.urlEncode(cursor)
	} else {
		cursor = ""
	}
	return i.client.transport.request("GET", "/1/indexes/"+i.nameEncoded+"/browse?"+i.client.transport.EncodeParams(params)+cursor, nil, read)
}

// BrowseAll browses the results according to the given `params` parameter
// starting at the first results. It returns an `IndexIterator` that is used to
// iterate over the results.
func (i *Index) BrowseAll(params interface{}) (*IndexIterator, error) {
	return NewIndexIterator(i, params, "")
}

// Search performs a search query according to the `query` search query and the
// given `params` parameters.
func (i *Index) Search(query string, params interface{}) (interface{}, error) {
	if params == nil {
		params = make(map[string]interface{})
	}

	params.(map[string]interface{})["query"] = query
	body := map[string]interface{}{
		"params": i.client.transport.EncodeParams(params),
	}

	return i.client.transport.request("POST", "/1/indexes/"+i.nameEncoded+"/query", body, search)
}

// Copy copies the index into a new one called `name`.
func (i *Index) Copy(name string) (interface{}, error) {
	return i.operation(name, "copy")
}

// Move renames the index into `name`.
func (i *Index) Move(name string) (interface{}, error) {
	return i.operation(name, "move")
}

// AddKey registers a new API key for the index. The `acl` parameter controls
// which permissions are given, `validity` is the validity duration in seconds
// (0 for unlimited), `maxQueriesPerIPPerHour` is the maximum number of calls
// authorized per hour and `maxHitsPerQuery` controls the number of results
// that each query could return at most.
func (i *Index) AddKey(acl []string, validity int, maxQueriesPerIPPerHour int, maxHitsPerQuery int) (interface{}, error) {
	body := map[string]interface{}{
		"acl":                    acl,
		"maxHitsPerQuery":        maxHitsPerQuery,
		"maxQueriesPerIPPerHour": maxQueriesPerIPPerHour,
		"validity":               validity,
	}

	return i.AddKeyWithParam(body)
}

// AddKeyWithParam registers a new API for the index. The `params` parameter is
// a `map[string]interface{}` of all the parameters given to the `AddKey`
// function.
func (i *Index) AddKeyWithParam(params interface{}) (interface{}, error) {
	return i.client.transport.request("POST", "/1/indexes/"+i.nameEncoded+"/keys", params, write)
}

// UpdateKey updates the `key` API key according to the other given parameters.
func (i *Index) UpdateKey(key string, acl []string, validity int, maxQueriesPerIPPerHour int, maxHitsPerQuery int) (interface{}, error) {
	body := map[string]interface{}{
		"acl":                    acl,
		"maxHitsPerQuery":        maxHitsPerQuery,
		"maxQueriesPerIPPerHour": maxQueriesPerIPPerHour,
		"validity":               validity,
	}

	return i.UpdateKeyWithParam(key, body)
}

// UpdateKeyWithParam updates the `key` API key according to the `params`
// parameters which is a `map[string]interface{}` of all the parameters given
// to the `UpdateKey` function.
func (i *Index) UpdateKeyWithParam(key string, params interface{}) (interface{}, error) {
	return i.client.transport.request("PUT", "/1/indexes/"+i.nameEncoded+"/keys/"+key, params, write)
}

// getStatus returns the status of a task given its ID `taskID`. The returned
// interface is the JSON-encoded answered from the API server. The error is
// non-nil if the REST API has returned an error.
func (i *Index) getStatus(taskID float64) (interface{}, error) {
	return i.client.transport.request("GET", "/1/indexes/"+i.nameEncoded+"/task/"+strconv.FormatFloat(taskID, 'f', -1, 64), nil, read)
}

// sameBatch performs the `action` command on all the objects specified in the
// `objects` parameter.
func (i *Index) sameBatch(objects interface{}, action string) (interface{}, error) {
	method := make([]string, len(objects.([]interface{})))
	for i := range method {
		method[i] = action
	}

	return i.Batch(objects, method)
}

// operation performs the `op` operation on the underlying index and names the
// resulting new index `name`. The `op` operation can be either `copy` or
// `move`.
func (i *Index) operation(name, op string) (interface{}, error) {
	body := map[string]interface{}{
		"destination": name,
		"operation":   op,
	}

	return i.client.transport.request("POST", "/1/indexes/"+i.nameEncoded+"/operation", body, write)
}

// SearchSynonyms returns the synonyms matching `query` whose types match
// `types`. To retrieve the first page, `page` should be set to 0. `hitsPerPage`
// specifies how much synonym sets will be returned by page.
// An error is returned if the underlying HTTP call does not yield a 200
// status code.
func (i *Index) SearchSynonyms(query string, types []string, page, hitsPerPage int) (interface{}, error) {
	body := map[string]interface{}{
		"query":       query,
		"type":        strings.Join(types, ","),
		"page":        page,
		"hitsPerPage": hitsPerPage,
	}

	return i.client.transport.request("POST", "/1/indexes/"+i.nameEncoded+"/synonyms/search", body, search)
}

// GetSynonym retrieves the synonym identified by `objectID`.
// An error is returned if the underlying HTTP call does not yield a 200
// status code.
func (i *Index) GetSynonym(objectID string) (interface{}, error) {
	encodedID := i.client.transport.urlEncode(objectID)
	return i.client.transport.request("GET", "/1/indexes/"+i.nameEncoded+"/synonyms/"+encodedID, nil, read)
}

// DeleteSynonym removes the synonym identified by `objectID`.
// The deletion can be forwarded to the index slaves of the index
// with `forwardToSlaves`.
// An error is returned if the underlying HTTP call does not yield a 200
// status code.
func (i *Index) DeleteSynonym(objectID string, forwardToSlaves bool) (interface{}, error) {
	encodedID := i.client.transport.urlEncode(objectID)
	params := i.client.EncodeParams(map[string]interface{}{
		"forwardToSlaves": strconv.FormatBool(forwardToSlaves),
	})

	return i.client.transport.request("DELETE", "/1/indexes/"+i.nameEncoded+"/synonyms/"+encodedID+"?"+params, nil, write)
}

// ClearSynonyms removes all synonyms from the index. The clear operation can
// be forwarded to the index slaves of the index using `forwardToSlaves`.
// An error is returned if the underlying HTTP call does not yield a 200
// status code.
func (i *Index) ClearSynonyms(forwardToSlaves bool) (interface{}, error) {
	params := i.client.EncodeParams(map[string]interface{}{
		"forwardToSlaves": strconv.FormatBool(forwardToSlaves),
	})

	return i.client.transport.request("POST", "/1/indexes/"+i.nameEncoded+"/synonyms/clear?"+params, nil, write)
}

// BatchSynonyms adds all `synonyms` to the index. The index can be cleared
// before by setting `replaceExistingSynonyms` to `true`. The optional clear
// operation and the additions can be forwarded to the index slaves
// with `forwardToSlaves`
// An error is returned if the underlying HTTP call does not yield a 200
// status code.
func (i *Index) BatchSynonyms(synonyms []interface{}, replaceExistingSynonyms, forwardToSlaves bool) (interface{}, error) {
	params := i.client.EncodeParams(map[string]interface{}{
		"forwardToSlaves":         strconv.FormatBool(forwardToSlaves),
		"replaceExistingSynonyms": strconv.FormatBool(replaceExistingSynonyms),
	})

	return i.client.transport.request("POST", "/1/indexes/"+i.nameEncoded+"/synonyms/batch?"+params, synonyms, write)
}

// SaveSynonym adds the given `synonym` to be identified `objectID`.
// This addition can be forwarded to the index slaves using `forwardToSlaves`.
// An error is returned if the underlying HTTP call does not yield a 200
// status code.
func (i *Index) SaveSynonym(objectID string, synonym interface{}, forwardToSlaves bool) (interface{}, error) {
	encodedID := i.client.EncodeParams(objectID)
	params := i.client.EncodeParams(map[string]interface{}{
		"forwardToSlaves": strconv.FormatBool(forwardToSlaves),
	})

	return i.client.transport.request("PUT", "/1/indexes/"+i.nameEncoded+"/synonyms/"+encodedID+"?"+params, synonym, write)
}
