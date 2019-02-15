package algoliasearch

import (
	"encoding/json"
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

// NewIndex instantiates a new `Index`. The `name` parameter corresponds to the
// Algolia index name while the `client` is used to connect to the Algolia API.
func NewIndex(name string, client *client) Index {
	return &index{
		client: client,
		name:   name,
		route:  "/1/indexes/" + url.QueryEscape(name),
	}
}

func (i *index) GetAppID() string {
	return i.client.GetAppID()
}

func (i *index) Delete() (res DeleteTaskRes, err error) {
	return i.DeleteWithRequestOptions(nil)
}

func (i *index) DeleteWithRequestOptions(opts *RequestOptions) (res DeleteTaskRes, err error) {
	path := i.route
	err = i.client.request(&res, "DELETE", path, nil, write, opts)
	return
}

func (i *index) Clear() (res UpdateTaskRes, err error) {
	return i.ClearWithRequestOptions(nil)
}

func (i *index) ClearWithRequestOptions(opts *RequestOptions) (res UpdateTaskRes, err error) {
	path := i.route + "/clear"
	err = i.client.request(&res, "POST", path, nil, write, opts)
	return
}

func (i *index) GetObject(objectID string, attributes []string) (object Object, err error) {
	return i.GetObjectWithRequestOptions(objectID, attributes, nil)
}

func (i *index) GetObjectWithRequestOptions(objectID string, attributes []string, opts *RequestOptions) (object Object, err error) {
	var params Map
	if attributes != nil {
		var attrBytes []byte
		attrBytes, err = json.Marshal(attributes)
		if err != nil {
			return
		}
		params = Map{
			"attributes": string(attrBytes),
		}
	}

	path := i.route + "/" + url.QueryEscape(objectID) + "?" + encodeMap(params)
	err = i.client.request(&object, "GET", path, nil, read, opts)
	return
}

func (i *index) getObjects(objectIDs, attributesToRetrieve []string, opts *RequestOptions) (objs []Object, err error) {
	attrs := strings.Join(attributesToRetrieve, ",")

	requests := make([]map[string]string, len(objectIDs))
	for j, id := range objectIDs {
		requests[j] = map[string]string{
			"indexName": i.name,
			"objectID":  url.QueryEscape(id),
		}
		if attributesToRetrieve != nil {
			requests[j]["attributesToRetrieve"] = attrs
		}
	}

	body := Map{
		"requests": requests,
	}

	var res objects
	path := "/1/indexes/*/objects"
	err = i.client.request(&res, "POST", path, body, read, opts)
	objs = res.Results
	return
}

func (i *index) GetObjects(objectIDs []string) (objs []Object, err error) {
	return i.GetObjectsWithRequestOptions(objectIDs, nil)
}

func (i *index) GetObjectsWithRequestOptions(objectIDs []string, opts *RequestOptions) (objs []Object, err error) {
	return i.getObjects(objectIDs, nil, opts)
}

func (i *index) GetObjectsAttrs(objectIDs, attrs []string) (objs []Object, err error) {
	return i.GetObjectsAttrsWithRequestOptions(objectIDs, attrs, nil)
}

func (i *index) GetObjectsAttrsWithRequestOptions(objectIDs, attrs []string, opts *RequestOptions) (objs []Object, err error) {
	return i.getObjects(objectIDs, attrs, opts)
}

func (i *index) DeleteObject(objectID string) (res DeleteTaskRes, err error) {
	return i.DeleteObjectWithRequestOptions(objectID, nil)
}

func (i *index) DeleteObjectWithRequestOptions(objectID string, opts *RequestOptions) (res DeleteTaskRes, err error) {
	if objectID == "" {
		err = fmt.Errorf("objectID cannot be empty")
		return
	}

	path := i.route + "/" + url.QueryEscape(objectID)
	err = i.client.request(&res, "DELETE", path, nil, write, opts)
	return
}

func (i *index) GetSettings() (settings Settings, err error) {
	return i.GetSettingsWithRequestOptions(nil)
}

func (i *index) GetSettingsWithRequestOptions(opts *RequestOptions) (settings Settings, err error) {
	path := i.route + "/settings?getVersion=2"
	err = i.client.request(&settings, "GET", path, nil, read, opts)
	settings.clean()
	return
}

func (i *index) SetSettings(settings Map) (res UpdateTaskRes, err error) {
	return i.SetSettingsWithRequestOptions(settings, nil)
}

func (i *index) SetSettingsWithRequestOptions(settings Map, opts *RequestOptions) (res UpdateTaskRes, err error) {
	if err = checkSettings(settings); err != nil {
		return
	}

	// Handle forwardToReplicas separately
	forwardToReplicas, ok := settings["forwardToReplicas"]
	if !ok {
		forwardToReplicas = false
	}
	delete(settings, "forwardToReplicas")

	path := i.route + "/settings?forwardToReplicas=" + fmt.Sprintf("%t", forwardToReplicas)
	err = i.client.request(&res, "PUT", path, settings, write, opts)
	return
}

func (i *index) WaitTask(taskID int) error {
	return i.WaitTaskWithRequestOptions(taskID, nil)
}

func (i *index) WaitTaskWithRequestOptions(taskID int, opts *RequestOptions) error {
	return i.client.WaitTaskWithRequestOptions(i.name, taskID, opts)
}

func (i *index) ListKeys() (keys []Key, err error) {
	return i.ListKeysWithRequestOptions(nil)
}

func (i *index) ListKeysWithRequestOptions(opts *RequestOptions) (keys []Key, err error) {
	var res listAPIKeysRes

	path := i.route + "/keys"
	err = i.client.request(&res, "GET", path, nil, read, opts)
	keys = res.Keys
	return
}

func (i *index) AddUserKey(ACL []string, params Map) (AddKeyRes, error) {
	return i.AddAPIKey(ACL, params)
}

func (i *index) AddAPIKey(ACL []string, params Map) (res AddKeyRes, err error) {
	return i.AddAPIKeyWithRequestOptions(ACL, params, nil)
}

func (i *index) AddAPIKeyWithRequestOptions(ACL []string, params Map, opts *RequestOptions) (res AddKeyRes, err error) {
	req := duplicateMap(params)
	req["acl"] = ACL

	if err = checkKey(req); err != nil {
		return
	}

	path := i.route + "/keys"
	err = i.client.request(&res, "POST", path, req, write, opts)
	return
}

func (i *index) UpdateUserKey(key string, params Map) (UpdateKeyRes, error) {
	return i.UpdateAPIKey(key, params)
}

func (i *index) UpdateAPIKey(key string, params Map) (res UpdateKeyRes, err error) {
	return i.UpdateAPIKeyWithRequestOptions(key, params, nil)
}

func (i *index) UpdateAPIKeyWithRequestOptions(key string, params Map, opts *RequestOptions) (res UpdateKeyRes, err error) {
	if err = checkKey(params); err != nil {
		return
	}

	path := i.route + "/keys/" + url.QueryEscape(key)
	err = i.client.request(&res, "PUT", path, params, read, opts)
	return
}

func (i *index) GetUserKey(value string) (Key, error) {
	return i.GetAPIKey(value)
}

func (i *index) GetAPIKey(value string) (key Key, err error) {
	return i.GetAPIKeyWithRequestOptions(value, nil)
}

func (i *index) GetAPIKeyWithRequestOptions(value string, opts *RequestOptions) (key Key, err error) {
	path := i.route + "/keys/" + url.QueryEscape(value)
	err = i.client.request(&key, "GET", path, nil, read, opts)
	return
}

func (i *index) DeleteUserKey(value string) (DeleteRes, error) {
	return i.DeleteAPIKey(value)
}

func (i *index) DeleteAPIKey(value string) (res DeleteRes, err error) {
	return i.DeleteAPIKeyWithRequestOptions(value, nil)
}

func (i *index) DeleteAPIKeyWithRequestOptions(value string, opts *RequestOptions) (res DeleteRes, err error) {
	path := i.route + "/keys/" + value
	err = i.client.request(&res, "DELETE", path, nil, write, opts)
	return
}

func (i *index) AddObject(object Object) (res CreateObjectRes, err error) {
	return i.AddObjectWithRequestOptions(object, nil)
}

func (i *index) AddObjectWithRequestOptions(object Object, opts *RequestOptions) (res CreateObjectRes, err error) {
	path := i.route
	err = i.client.request(&res, "POST", path, object, write, opts)
	return
}

func (i *index) UpdateObject(object Object) (res UpdateObjectRes, err error) {
	return i.UpdateObjectWithRequestOptions(object, nil)
}

func (i *index) UpdateObjectWithRequestOptions(object Object, opts *RequestOptions) (res UpdateObjectRes, err error) {
	objectID, err := object.ObjectID()
	if err != nil {
		return
	}

	path := i.route + "/" + url.QueryEscape(objectID)
	err = i.client.request(&res, "PUT", path, object, write, opts)
	return
}

func (i *index) partialUpdateObject(object Object, createIfNotExists bool, opts *RequestOptions) (res UpdateTaskRes, err error) {
	objectID, err := object.ObjectID()
	if err != nil {
		return
	}

	path := i.route + "/" + url.QueryEscape(objectID) + "/partial"
	if !createIfNotExists {
		path += "?createIfNotExists=false"
	}
	err = i.client.request(&res, "POST", path, object, write, opts)
	return
}

func (i *index) PartialUpdateObject(object Object) (res UpdateTaskRes, err error) {
	return i.PartialUpdateObjectWithRequestOptions(object, nil)
}

func (i *index) PartialUpdateObjectWithRequestOptions(object Object, opts *RequestOptions) (res UpdateTaskRes, err error) {
	return i.partialUpdateObject(object, true, opts)
}

func (i *index) PartialUpdateObjectNoCreate(object Object) (res UpdateTaskRes, err error) {
	return i.PartialUpdateObjectNoCreateWithRequestOptions(object, nil)
}

func (i *index) PartialUpdateObjectNoCreateWithRequestOptions(object Object, opts *RequestOptions) (res UpdateTaskRes, err error) {
	return i.partialUpdateObject(object, false, opts)
}

func (i *index) AddObjects(objects []Object) (res BatchRes, err error) {
	return i.AddObjectsWithRequestOptions(objects, nil)
}

func (i *index) AddObjectsWithRequestOptions(objects []Object, opts *RequestOptions) (res BatchRes, err error) {
	var operations []BatchOperation

	if operations, err = newBatchOperations(objects, "addObject"); err == nil {
		res, err = i.BatchWithRequestOptions(operations, opts)
	}

	return
}

func (i *index) UpdateObjects(objects []Object) (res BatchRes, err error) {
	return i.UpdateObjectsWithRequestOptions(objects, nil)
}

func (i *index) UpdateObjectsWithRequestOptions(objects []Object, opts *RequestOptions) (res BatchRes, err error) {
	var operations []BatchOperation

	if operations, err = newBatchOperations(objects, "updateObject"); err == nil {
		res, err = i.BatchWithRequestOptions(operations, opts)
	}

	return
}

func (i *index) partialUpdateObjects(objects []Object, action string, opts *RequestOptions) (res BatchRes, err error) {
	var operations []BatchOperation

	if operations, err = newBatchOperations(objects, action); err == nil {
		res, err = i.BatchWithRequestOptions(operations, opts)
	}

	return
}

func (i *index) PartialUpdateObjects(objects []Object) (res BatchRes, err error) {
	return i.PartialUpdateObjectsWithRequestOptions(objects, nil)
}

func (i *index) PartialUpdateObjectsWithRequestOptions(objects []Object, opts *RequestOptions) (res BatchRes, err error) {
	return i.partialUpdateObjects(objects, "partialUpdateObject", opts)
}

func (i *index) PartialUpdateObjectsNoCreate(objects []Object) (res BatchRes, err error) {
	return i.PartialUpdateObjectsNoCreateWithRequestOptions(objects, nil)
}

func (i *index) PartialUpdateObjectsNoCreateWithRequestOptions(objects []Object, opts *RequestOptions) (res BatchRes, err error) {
	return i.partialUpdateObjects(objects, "partialUpdateObjectNoCreate", opts)
}

func (i *index) DeleteObjects(objectIDs []string) (res BatchRes, err error) {
	return i.DeleteObjectsWithRequestOptions(objectIDs, nil)
}

func (i *index) DeleteObjectsWithRequestOptions(objectIDs []string, opts *RequestOptions) (res BatchRes, err error) {
	objects := make([]Object, len(objectIDs))

	for j, id := range objectIDs {
		objects[j] = Object{
			"objectID": id,
		}
	}

	var operations []BatchOperation
	if operations, err = newBatchOperations(objects, "deleteObject"); err == nil {
		res, err = i.BatchWithRequestOptions(operations, opts)
	}

	return
}

func (i *index) Batch(operations []BatchOperation) (res BatchRes, err error) {
	return i.BatchWithRequestOptions(operations, nil)
}

func (i *index) BatchWithRequestOptions(operations []BatchOperation, opts *RequestOptions) (res BatchRes, err error) {
	body := map[string][]BatchOperation{
		"requests": operations,
	}

	path := i.route + "/batch"
	err = i.client.request(&res, "POST", path, body, write, opts)
	return
}

func (i *index) Copy(name string) (UpdateTaskRes, error) {
	return i.CopyWithRequestOptions(name, nil)
}

func (i *index) CopyWithRequestOptions(name string, opts *RequestOptions) (UpdateTaskRes, error) {
	return i.ScopedCopyWithRequestOptions(name, nil, opts)
}

func (i *index) ScopedCopy(name string, scopes []string) (UpdateTaskRes, error) {
	return i.ScopedCopyWithRequestOptions(name, scopes, nil)
}

func (i *index) ScopedCopyWithRequestOptions(name string, scopes []string, opts *RequestOptions) (UpdateTaskRes, error) {
	return i.client.ScopedCopyIndexWithRequestOptions(i.name, name, scopes, opts)
}

func (i *index) Move(name string) (UpdateTaskRes, error) {
	return i.MoveTo(name)
}

func (i *index) MoveWithRequestOptions(name string, opts *RequestOptions) (UpdateTaskRes, error) {
	return i.MoveToWithRequestOptions(name, opts)
}

func (i *index) MoveTo(name string) (UpdateTaskRes, error) {
	return i.MoveToWithRequestOptions(name, nil)
}

func (i *index) MoveToWithRequestOptions(name string, opts *RequestOptions) (UpdateTaskRes, error) {
	return i.client.MoveIndexWithRequestOptions(i.name, name, opts)
}

func (i *index) GetStatus(taskID int) (res TaskStatusRes, err error) {
	return i.GetStatusWithRequestOptions(taskID, nil)
}

func (i *index) GetStatusWithRequestOptions(taskID int, opts *RequestOptions) (res TaskStatusRes, err error) {
	res, err = i.client.GetStatusWithRequestOptions(i.name, taskID, opts)
	return
}

func (i *index) SearchSynonyms(query string, types []string, page, hitsPerPage int) (synonyms []Synonym, err error) {
	return i.SearchSynonymsWithRequestOptions(query, types, page, hitsPerPage, nil)
}

func (i *index) SearchSynonymsWithRequestOptions(query string, types []string, page, hitsPerPage int, opts *RequestOptions) (synonyms []Synonym, err error) {
	body := Map{
		"query":       query,
		"type":        strings.Join(types, ","),
		"page":        page,
		"hitsPerPage": hitsPerPage,
	}

	path := i.route + "/synonyms/search"
	var res SearchSynonymsRes
	err = i.client.request(&res, "POST", path, body, search, opts)

	if err == nil {
		synonyms = res.Hits
	}

	return
}

func (i *index) GetSynonym(objectID string) (s Synonym, err error) {
	return i.GetSynonymWithRequestOptions(objectID, nil)
}

func (i *index) GetSynonymWithRequestOptions(objectID string, opts *RequestOptions) (s Synonym, err error) {
	path := i.route + "/synonyms/" + url.QueryEscape(objectID)
	err = i.client.request(&s, "GET", path, nil, read, opts)
	return
}

func (i *index) AddSynonym(synonym Synonym, forwardToReplicas bool) (res UpdateTaskWithIDRes, err error) {
	return i.SaveSynonym(synonym, forwardToReplicas)
}

func (i *index) AddSynonymWithRequestOptions(synonym Synonym, forwardToReplicas bool, opts *RequestOptions) (res UpdateTaskWithIDRes, err error) {
	return i.SaveSynonymWithRequestOptions(synonym, forwardToReplicas, opts)
}

func (i *index) SaveSynonym(synonym Synonym, forwardToReplicas bool) (res UpdateTaskWithIDRes, err error) {
	return i.SaveSynonymWithRequestOptions(synonym, forwardToReplicas, nil)
}

func (i *index) SaveSynonymWithRequestOptions(synonym Synonym, forwardToReplicas bool, opts *RequestOptions) (res UpdateTaskWithIDRes, err error) {
	params := Map{
		"forwardToReplicas": forwardToReplicas,
	}

	path := i.route + "/synonyms/" + url.QueryEscape(synonym.ObjectID) + "?" + encodeMap(params)
	err = i.client.request(&res, "PUT", path, synonym, write, opts)
	return
}

func (i *index) DeleteSynonym(objectID string, forwardToReplicas bool) (res DeleteTaskRes, err error) {
	return i.DeleteSynonymWithRequestOptions(objectID, forwardToReplicas, nil)
}

func (i *index) DeleteSynonymWithRequestOptions(objectID string, forwardToReplicas bool, opts *RequestOptions) (res DeleteTaskRes, err error) {
	params := Map{
		"forwardToReplicas": forwardToReplicas,
	}

	path := i.route + "/synonyms/" + url.QueryEscape(objectID) + "?" + encodeMap(params)
	err = i.client.request(&res, "DELETE", path, nil, write, opts)
	return
}

func (i *index) ClearSynonyms(forwardToReplicas bool) (res UpdateTaskRes, err error) {
	return i.ClearSynonymsWithRequestOptions(forwardToReplicas, nil)
}

func (i *index) ClearSynonymsWithRequestOptions(forwardToReplicas bool, opts *RequestOptions) (res UpdateTaskRes, err error) {
	params := Map{
		"forwardToReplicas": forwardToReplicas,
	}

	path := i.route + "/synonyms/clear?" + encodeMap(params)
	err = i.client.request(&res, "POST", path, nil, write, opts)
	return
}

func (i *index) BatchSynonyms(synonyms []Synonym, replaceExistingSynonyms, forwardToReplicas bool) (res UpdateTaskRes, err error) {
	return i.BatchSynonymsWithRequestOptions(synonyms, replaceExistingSynonyms, forwardToReplicas, nil)
}

func (i *index) BatchSynonymsWithRequestOptions(synonyms []Synonym, replaceExistingSynonyms, forwardToReplicas bool, opts *RequestOptions) (res UpdateTaskRes, err error) {
	params := Map{
		"replaceExistingSynonyms": replaceExistingSynonyms,
		"forwardToReplicas":       forwardToReplicas,
	}

	path := i.route + "/synonyms/batch?" + encodeMap(params)
	err = i.client.request(&res, "POST", path, synonyms, write, opts)
	return
}

func (i *index) Browse(params Map, cursor string) (res BrowseRes, err error) {
	return i.BrowseWithRequestOptions(params, cursor, nil)
}

func (i *index) BrowseWithRequestOptions(params Map, cursor string, opts *RequestOptions) (res BrowseRes, err error) {
	copy := duplicateMap(params)
	if err = checkQuery(copy); err != nil {
		return
	}

	if cursor != "" {
		copy["cursor"] = cursor
	}

	req := Map{
		"params": encodeMap(copy),
	}

	path := i.route + "/browse"
	err = i.client.request(&res, "POST", path, req, read, opts)
	return
}

func (i *index) BrowseAll(params Map) (it IndexIterator, err error) {
	return i.BrowseAllWithRequestOptions(params, nil)
}

func (i *index) BrowseAllWithRequestOptions(params Map, opts *RequestOptions) (it IndexIterator, err error) {
	if err = checkQuery(params); err != nil {
		return
	}

	it, err = newIndexIterator(i, params, opts)
	return
}

func (i *index) Search(query string, params Map) (res QueryRes, err error) {
	return i.SearchWithRequestOptions(query, params, nil)
}

func (i *index) SearchWithRequestOptions(query string, params Map, opts *RequestOptions) (res QueryRes, err error) {
	copy := duplicateMap(params)
	copy["query"] = query

	if err = checkQuery(copy); err != nil {
		return
	}

	req := Map{
		"params": encodeMap(copy),
	}

	path := i.route + "/query"
	err = i.client.request(&res, "POST", path, req, search, opts)
	return
}

func (i *index) DeleteBy(params Map) (res UpdateTaskRes, err error) {
	return i.DeleteByWithRequestOptions(params, nil)
}

func (i *index) DeleteByWithRequestOptions(params Map, opts *RequestOptions) (res UpdateTaskRes, err error) {
	if err = checkQuery(params); err != nil {
		return
	}

	req := Map{
		"params": encodeMap(params),
	}

	path := i.route + "/deleteByQuery"
	err = i.client.request(&res, "POST", path, req, write, opts)
	return
}

func (i *index) DeleteByQuery(query string, params Map) (err error) {
	return i.DeleteByQueryWithRequestOptions(query, params, nil)
}

func (i *index) DeleteByQueryWithRequestOptions(query string, params Map, opts *RequestOptions) (err error) {
	copy := duplicateMap(params)
	copy["attributesToRetrieve"] = []string{"objectID"}
	copy["hitsPerPage"] = 1000
	copy["query"] = query
	copy["distinct"] = 0

	var browseRes BrowseRes
	var batchRes BatchRes
	var objectIDs []string
	var cursor string

	for {
		// Start browsing the content by cursor
		if browseRes, err = i.BrowseWithRequestOptions(copy, cursor, opts); err != nil {
			return
		}

		// Collect all objectIDs
		for _, hit := range browseRes.Hits {
			objectIDs = append(objectIDs, hit["objectID"].(string))
		}

		// Set the new cursor from response
		cursor = browseRes.Cursor

		// Break if there's no more matching records
		if cursor == "" {
			break
		}
	}

	// Delete all the objects
	if batchRes, err = i.DeleteObjectsWithRequestOptions(objectIDs, opts); err != nil {
		return
	}

	// Wait until DeleteObjects completion
	err = i.WaitTaskWithRequestOptions(batchRes.TaskID, opts)
	return
}

func (i *index) SearchFacet(facet, query string, params Map) (res SearchFacetRes, err error) {
	return i.SearchForFacetValues(facet, query, params)
}

func (i *index) SearchForFacetValues(facet, query string, params Map) (res SearchFacetRes, err error) {
	return i.SearchForFacetValuesWithRequestOptions(facet, query, params, nil)
}

func (i *index) SearchForFacetValuesWithRequestOptions(facet, query string, params Map, opts *RequestOptions) (res SearchFacetRes, err error) {
	copy := duplicateMap(params)
	if err = checkQuery(copy); err != nil {
		return
	}

	copy["facetQuery"] = query

	req := Map{
		"params": encodeMap(copy),
	}

	path := i.route + "/facets/" + facet + "/query"
	err = i.client.request(&res, "POST", path, req, search, opts)
	return
}

func (i *index) SaveRule(rule Rule, forwardToReplicas bool) (res SaveRuleRes, err error) {
	return i.SaveRuleWithRequestOptions(rule, forwardToReplicas, nil)
}

func (i *index) SaveRuleWithRequestOptions(rule Rule, forwardToReplicas bool, opts *RequestOptions) (res SaveRuleRes, err error) {
	if err = checkRule(rule); err != nil {
		return
	}

	rule.enableImplicitly()

	params := Map{"forwardToReplicas": forwardToReplicas}
	path := i.route + "/rules/" + rule.ObjectID + "?" + encodeMap(params)
	err = i.client.request(&res, "PUT", path, rule, write, opts)
	return
}

func (i *index) BatchRules(rules []Rule, forwardToReplicas, clearExistingRules bool) (res BatchRulesRes, err error) {
	return i.BatchRulesWithRequestOptions(rules, forwardToReplicas, clearExistingRules, nil)
}

func (i *index) BatchRulesWithRequestOptions(rules []Rule, forwardToReplicas, clearExistingRules bool, opts *RequestOptions) (res BatchRulesRes, err error) {
	if err = checkRules(rules); err != nil {
		return
	}

	for i, _ := range rules {
		rules[i].enableImplicitly()
	}

	params := Map{
		"forwardToReplicas":  forwardToReplicas,
		"clearExistingRules": clearExistingRules,
	}
	path := i.route + "/rules/batch?" + encodeMap(params)
	err = i.client.request(&res, "POST", path, rules, write, opts)
	return
}

func (i *index) GetRule(objectID string) (rule *Rule, err error) {
	return i.GetRuleWithRequestOptions(objectID, nil)
}

func (i *index) GetRuleWithRequestOptions(objectID string, opts *RequestOptions) (rule *Rule, err error) {
	path := i.route + "/rules/" + objectID
	err = i.client.request(&rule, "GET", path, nil, read, opts)
	return
}

func (i *index) DeleteRule(objectID string, forwardToReplicas bool) (res DeleteRuleRes, err error) {
	return i.DeleteRuleWithRequestOptions(objectID, forwardToReplicas, nil)
}

func (i *index) DeleteRuleWithRequestOptions(objectID string, forwardToReplicas bool, opts *RequestOptions) (res DeleteRuleRes, err error) {
	params := Map{"forwardToReplicas": forwardToReplicas}
	path := i.route + "/rules/" + objectID + "?" + encodeMap(params)
	err = i.client.request(&res, "DELETE", path, nil, write, opts)
	return
}

func (i *index) ClearRules(forwardToReplicas bool) (res ClearRulesRes, err error) {
	return i.ClearRulesWithRequestOptions(forwardToReplicas, nil)
}

func (i *index) ClearRulesWithRequestOptions(forwardToReplicas bool, opts *RequestOptions) (res ClearRulesRes, err error) {
	params := Map{"forwardToReplicas": forwardToReplicas}
	path := i.route + "/rules/clear?" + encodeMap(params)
	err = i.client.request(&res, "POST", path, nil, write, opts)
	return
}

func (i *index) SearchRules(params Map) (res SearchRulesRes, err error) {
	return i.SearchRulesWithRequestOptions(params, nil)
}

func (i *index) SearchRulesWithRequestOptions(params Map, opts *RequestOptions) (res SearchRulesRes, err error) {
	if err = checkSearchRulesParams(params); err != nil {
		return
	}

	path := i.route + "/rules/search"
	err = i.client.request(&res, "POST", path, params, read, opts)
	return
}

// ReplaceAllSynonyms replace all the synonyms of the current index with the given ones.
func (i *index) ReplaceAllSynonyms(synonyms []Synonym) (res UpdateTaskRes, err error) {
	return i.ReplaceAllSynonymsWithRequestOptions(synonyms, nil)
}

// ReplaceAllSynonymsWithRequestOptions is the same as ReplaceAllSynonyms but it also
// accepts extra RequestOptions.
func (i *index) ReplaceAllSynonymsWithRequestOptions(synonyms []Synonym, opts *RequestOptions) (res UpdateTaskRes, err error) {
	return i.BatchSynonymsWithRequestOptions(synonyms, true, false, opts)
}

// ReplaceAllRules replace all the rules of the current index with the given ones.
func (i *index) ReplaceAllRules(rules []Rule) (res BatchRulesRes, err error) {
	return i.ReplaceAllRulesWithRequestOptions(rules, nil)
}

// ReplaceAllRulesWithRequestOptions is the same as ReplaceAllRules but it also
// accepts extra RequestOptions.
func (i *index) ReplaceAllRulesWithRequestOptions(rules []Rule, opts *RequestOptions) (res BatchRulesRes, err error) {
	return i.BatchRulesWithRequestOptions(rules, false, true, opts)
}

// ReplaceAllObjects replace all the objects of the current index with the given ones.
func (i *index) ReplaceAllObjects(objects []Object) error {
	return i.ReplaceAllObjectsWithRequestOptions(objects, nil)
}

// ReplaceAllObjectsWithRequestOptions is the same as ReplaceAllObjects but it also
// accepts extra RequestOptions.
func (i *index) ReplaceAllObjectsWithRequestOptions(objects []Object, opts *RequestOptions) (err error) {
	tmpIndexName := fmt.Sprintf("%s_tmp_%d", i.name, time.Now().Nanosecond())

	defer func() {
		if err != nil {
			i.client.DeleteIndexWithRequestOptions(tmpIndexName, opts)
		}
	}()

	var taskIDs []int

	// Copy settings/synonyms/rules to the temporary index
	{
		var res UpdateTaskRes
		res, err = i.client.ScopedCopyIndexWithRequestOptions(i.name, tmpIndexName, []string{"settings", "synonyms", "rules"}, opts)
		if err != nil {
			return
		}
		taskIDs = append(taskIDs, res.TaskID)
	}

	tmpIndex := i.client.InitIndex(tmpIndexName)

	// Copy objects to the temporary index
	{
		batchSize := 1000

		for i := 0; i < len(objects); i += batchSize {
			j := i + batchSize
			if j > len(objects) {
				j = len(objects)
			}
			var res BatchRes
			res, err = tmpIndex.AddObjectsWithRequestOptions(objects[i:j], opts)
			if err != nil {
				return
			}
			taskIDs = append(taskIDs, res.TaskID)
		}
	}

	// Wait for all the tasks to finish before performing the move of the temporary index to the final one.
	for _, taskID := range taskIDs {
		err = tmpIndex.WaitTaskWithRequestOptions(taskID, opts)
		if err != nil {
			return
		}
	}

	// Perform the move operation
	{
		var res UpdateTaskRes
		res, err = i.client.MoveIndexWithRequestOptions(tmpIndexName, i.name, opts)
		if err != nil {
			return
		}
		tmpIndex.WaitTaskWithRequestOptions(res.TaskID, opts)
	}

	return
}
