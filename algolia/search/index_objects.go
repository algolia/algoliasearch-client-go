package search

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/call"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/errs"
	iopt "github.com/algolia/algoliasearch-client-go/v3/algolia/internal/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/iterator"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/transport"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/wait"
)

// GetObject retrieves the record identified by the given objectID and
// deserializes it into the object parameter.
func (i *Index) GetObject(objectID string, object interface{}, opts ...interface{}) error {
	if objectID == "" {
		return errs.ErrMissingObjectID
	}

	opts = opt.InsertExtraURLParam(
		opts,
		"attributesToRetrieve",
		strings.Join(iopt.ExtractAttributesToRetrieve(opts...).Get(), ","),
	)

	path := i.path("/%s", url.QueryEscape(objectID))
	return i.transport.Request(&object, http.MethodGet, path, nil, call.Read, opts...)
}

// SaveObject saves the given object to the index.
func (i *Index) SaveObject(object interface{}, opts ...interface{}) (res SaveObjectRes, err error) {
	path := i.path("")
	err = i.transport.Request(&res, http.MethodPost, path, object, call.Write, opts...)
	res.wait = i.WaitTask
	return
}

// PartialUpdateObject replaces the object content of the given object according
// to its respective objectID field.
func (i *Index) PartialUpdateObject(object interface{}, opts ...interface{}) (res UpdateTaskRes, err error) {
	objectID, ok := getObjectID(object)
	if !ok {
		err = errs.ErrMissingObjectID
		res.wait = noWait
		return
	}

	opts = opt.InsertExtraURLParam(
		opts,
		"createIfNotExists",
		iopt.ExtractCreateIfNotExists(opts...).Get(),
	)

	path := i.path("/%s/partial", url.QueryEscape(objectID))
	err = i.transport.Request(&res, http.MethodPost, path, object, call.Write, opts...)
	res.wait = i.WaitTask
	return
}

// DeleteObject removes the record identified by the given objectID.
func (i *Index) DeleteObject(objectID string, opts ...interface{}) (res DeleteTaskRes, err error) {
	if objectID == "" {
		err = errs.ErrMissingObjectID
		res.wait = noWait
		return
	}

	path := i.path("/%s", url.QueryEscape(objectID))
	err = i.transport.Request(&res, http.MethodDelete, path, nil, call.Write, opts...)
	res.wait = i.WaitTask
	return
}

// GetObjects retrieves the records identified by the given objectIDs and
// deserializes them into the objects parameter.
func (i *Index) GetObjects(objectIDs []string, objects interface{}, opts ...interface{}) error {
	var (
		attributesToRetrieve = iopt.ExtractAttributesToRetrieve(opts...)
		requests             = make([]getObjectsReq, len(objectIDs))
		body                 = map[string]interface{}{"requests": requests}
		res                  = getObjectsRes{objects}
	)

	for j, objectID := range objectIDs {
		requests[j] = getObjectsReq{
			IndexName:            i.name,
			ObjectID:             objectID,
			AttributesToRetrieve: attributesToRetrieve,
		}
	}

	path := "/1/indexes/*/objects"
	return i.transport.Request(&res, http.MethodPost, path, body, call.Read, opts...)
}

// SaveObjects saves the given objects to the index.
//
// Objects can either be a slice, an array or an object implementing the
// iterator.Iterator interface. In the last case, SaveObjects will call the Next
// method of the iterator to retrieve the objects sent one by one.
//
// This method does not send all the provided objects in a single call. Objects
// are sent by batches whose size is controlled by Configuration.MaxBatchSize
// (defaults to search.DefaultMaxBatchSize).
func (i *Index) SaveObjects(objects interface{}, opts ...interface{}) (res GroupBatchRes, err error) {
	return i.batch(objects, AddObject, opts...)
}

// PartialUpdateObjects replaces object content of all the given objects
// according to their respective objectID field.
//
// If opt.CreateIfNotExists(true) is passed, non-existing objects will be
// created, otherwise, the call will fail.
//
// This method does not send all the provided objects in a single call. Objects
// are sent by batches whose size is controlled by Configuration.MaxBatchSize
// (defaults to search.DefaultMaxBatchSize).
func (i *Index) PartialUpdateObjects(objects interface{}, opts ...interface{}) (res GroupBatchRes, err error) {
	var action BatchAction

	if iopt.ExtractCreateIfNotExists(opts...).Get() {
		action = PartialUpdateObject
	} else {
		action = PartialUpdateObjectNoCreate
	}

	return i.batch(objects, action, opts...)
}

// DeleteObjects removes the records identified by the given objectIDs.
func (i *Index) DeleteObjects(objectIDs []string, opts ...interface{}) (res BatchRes, err error) {
	objects := make([]interface{}, len(objectIDs))

	for j, id := range objectIDs {
		objects[j] = map[string]string{"objectID": id}
	}

	var operations []BatchOperation
	if operations, err = newOperationBatch(objects, DeleteObject); err == nil {
		res, err = i.Batch(operations, opts...)
	} else {
		res.wait = noWait
	}
	return
}

func (i *Index) batch(objects interface{}, action BatchAction, opts ...interface{}) (GroupBatchRes, error) {
	var (
		batch       []interface{}
		operations  []BatchOperation
		singleRes   BatchRes
		multipleRes GroupBatchRes
	)

	autoGenerateObjectIDIfNotExist := iopt.ExtractAutoGenerateObjectIDIfNotExist(opts...).Get()
	it := iterator.New(objects)

	for {
		object, err := it.Next()
		if err != nil {
			return multipleRes, fmt.Errorf("iteration failed unexpectedly: %v", err)
		}

		if !autoGenerateObjectIDIfNotExist && object != nil && !hasObjectID(object) {
			return multipleRes, fmt.Errorf("missing objectID in object %#v", object)
		}

		if shouldSendBatch(i.maxBatchSize, batch, object) {
			operations, err = newOperationBatch(batch, action)
			if err != nil {
				return multipleRes, fmt.Errorf("could not generate intermediate batch: %v", err)
			}
			singleRes, err = i.Batch(operations, opts...)
			if err != nil {
				return multipleRes, fmt.Errorf("could not send intermediate batch: %v", err)
			}
			multipleRes.Responses = append(multipleRes.Responses, singleRes)
			batch = nil
		}

		if object == nil {
			break
		} else {
			batch = append(batch, object)
		}
	}

	return multipleRes, nil
}

func shouldSendBatch(maxBatchSize int, batch []interface{}, object interface{}) bool {
	isMaxBatchSizePositive := maxBatchSize > 0
	isBatchBigEnough := len(batch) >= maxBatchSize
	isLastBatch := object == nil
	return isMaxBatchSizePositive && (isBatchBigEnough || isLastBatch)
}

// Batch sends all the given indexing operations with a single call.
func (i *Index) Batch(operations []BatchOperation, opts ...interface{}) (res BatchRes, err error) {
	body := batchReq{Requests: operations}
	path := i.path("/batch")
	err = i.transport.Request(&res, http.MethodPost, path, body, call.Write, opts...)
	res.wait = i.WaitTask
	return
}

// DeleteBy removes all the records that match the given query parameters.
//
// Not all query parameters are supported, please refer to the official
// documentation for an exhaustive list:
// https://www.algolia.com/doc/api-reference/api-methods/delete-by/#method-param-filterparameters
func (i *Index) DeleteBy(opts ...interface{}) (res UpdateTaskRes, err error) {
	body := newDeleteByReq(opts...)
	path := i.path("/deleteByQuery")
	err = i.transport.Request(&res, http.MethodPost, path, body, call.Write, opts...)
	res.wait = i.WaitTask
	return
}

// Search performs a search query according to the given query string and any
// given query parameter among all the index records.
func (i *Index) Search(query string, opts ...interface{}) (res QueryRes, err error) {
	body := searchReq{Params: transport.URLEncode(newSearchParams(query, opts...))}
	path := i.path("/query")
	err = i.transport.Request(&res, http.MethodPost, path, body, call.Read, opts...)
	return
}

// FindObject searches iteratively through the search response `Hits`
// field to find the first response hit that would match against the given
// `filterFunc` function.
//
// If no object has been found within the first result set, the function
// will perform a new search operation on the next page of results, if any,
// until a matching object is found or the end of results, whichever
// happens first.
//
// To prevent the iteration through pages of results, `doNotPaginate`
// parameter can be set to true. This will stop the function at the end of
// the first page of search results even if no object does match.
func (i *Index) FindObject(
	filterFunc func(object map[string]interface{}) bool,
	query string,
	paginate bool,
	opts ...interface{},
) (*ObjectWithPosition, error) {
	res, err := i.Search(query, opts...)
	if err != nil {
		return nil, err
	}

	var hits []map[string]interface{}
	err = res.UnmarshalHits(&hits)
	if err != nil {
		return nil, err
	}

	for pos, hit := range hits {
		if filterFunc(hit) {
			return &ObjectWithPosition{
				Object:   hit,
				Position: pos,
				Page:     res.Page,
			}, nil
		}
	}

	hasNextPage := res.Page+1 < res.NbPages
	if !paginate || !hasNextPage {
		return nil, errs.ErrObjectNotFound
	}

	return i.FindObject(
		filterFunc,
		query,
		paginate,
		opt.InsertOrReplaceOption(opts, opt.Page(res.Page+1))...,
	)
}

// FindFirstObject does the same as FindObject except that it reverse the
// doNotPaginate boolean parameter.
//
// Deprecated: Use FindObject instead.
func (i *Index) FindFirstObject(
	filterFunc func(object map[string]interface{}) bool,
	query string,
	doNotPaginate bool,
	opts ...interface{},
) (*ObjectWithPosition, error) {
	return i.FindObject(filterFunc, query, !doNotPaginate, opts...)
}

// SearchForFacetValues performs a search query according to the given query
// string and any given parameter among the values of the given facet.
func (i *Index) SearchForFacetValues(facet, query string, opts ...interface{}) (res SearchForFacetValuesRes, err error) {
	params := newSearchForFacetValuesParams(query, opts...)
	body := map[string]string{
		"params": transport.URLEncode(params),
	}
	path := i.path("/facets/%s/query", facet)
	err = i.transport.Request(&res, http.MethodPost, path, body, call.Read, opts...)
	return
}

// BrowseObjects returns an iterator which will retrieve records one by one from
// the index according to the given query parameters.
//
// The return ObjectIterator can decode objects by passing the address of the
// object to decode to as a first argument of its Next() method.
func (i *Index) BrowseObjects(opts ...interface{}) (*ObjectIterator, error) {
	query := iopt.ExtractQuery(opts...).Get()
	searchParams := newSearchParams(query, opts...)
	browser := i.browserForObjects(searchParams, opts...)
	return newObjectIterator(browser)
}

func (i *Index) browserForObjects(params searchParams, opts ...interface{}) func(string) (browseRes, error) {
	return func(cursor string) (res browseRes, err error) {
		body := browseReq{
			Cursor: cursor,
			Params: transport.URLEncode(params),
		}
		path := i.path("/browse")
		err = i.transport.Request(&res, http.MethodPost, path, body, call.Read, opts...)
		return
	}
}

// ReplaceAllObjects saves all the given objects into the index by replacing all
// the original objects. Settings, rules and synonyms are preserved.
//
// Because this method is performing all operations in a non-blocking way, it
// may be fast but Wait() needs to be called on the returned wait.Group to
// ensure completion.
//
// In rare cases, this operation may fail in case of network disconnections. To
// prevent issues, one may want to pass the opt.Safe(true) option. Note that
// passing this option will make the method blocking.
func (i *Index) ReplaceAllObjects(objects interface{}, opts ...interface{}) (g *wait.Group, err error) {
	safe := iopt.ExtractSafe(opts...).Get()

	exists, err := i.Exists()
	if err != nil {
		err = fmt.Errorf("cannot check if the index exists: %v", err)
		return
	}

	if !exists {
		resSaveObjects, e := i.SaveObjects(objects, opts...)
		if e != nil {
			err = fmt.Errorf("cannot save objects to the index: %v", e)
			return
		}
		if safe {
			if e := resSaveObjects.Wait(); e != nil {
				err = fmt.Errorf("error while waiting for saving objects to the index: %v", e)
				return
			}
		}
		return
	}

	tmpIndex := i.client.InitIndex(fmt.Sprintf(
		"%s_tmp_%d",
		i.name,
		time.Now().UnixNano()),
	)

	defer func() {
		if err != nil {
			if _, e := tmpIndex.Delete(); e != nil {
				err = fmt.Errorf("temporary index cannot be deleted: %v", e)
			}
		}
	}()

	g = wait.NewGroup()
	optsWithScopes := opt.InsertOrReplaceOption(opts, opt.Scopes("rules", "settings", "synonyms"))

	resCopyIndex, err := i.client.CopyIndex(i.name, tmpIndex.name, optsWithScopes...)
	if err != nil {
		err = fmt.Errorf("cannot copy rules, settings and synonyms to the temporary index: %v", err)
		return
	}
	g.Collect(resCopyIndex)

	resSaveObjects, err := tmpIndex.SaveObjects(objects, opts...)
	if err != nil {
		err = fmt.Errorf("cannot save objects to the temporary index: %v", err)
		return
	}
	g.Collect(resSaveObjects)

	if safe {
		if e := g.Wait(); e != nil {
			err = fmt.Errorf("error while waiting for indexing operations to the temporary index: %v", e)
			return
		}
	}

	res, e := i.client.MoveIndex(tmpIndex.name, i.name, opts...)
	if e != nil {
		err = fmt.Errorf("cannot move temporary index to original index: %v", e)
		return
	}

	if safe {
		if e := res.Wait(); e != nil {
			err = fmt.Errorf("error while waiting for move operation of the temporary index: %v", e)
			return
		}
	}

	return
}
