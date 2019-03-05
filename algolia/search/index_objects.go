package search

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/algolia/algoliasearch-client-go/algolia/call"
	"github.com/algolia/algoliasearch-client-go/algolia/errs"
	iopt "github.com/algolia/algoliasearch-client-go/algolia/internal/opt"
	"github.com/algolia/algoliasearch-client-go/algolia/iterator"
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/algolia/transport"
)

func (i *Index) GetObject(objectID string, object interface{}, opts ...interface{}) error {
	if attrs := iopt.ExtractAttributesToRetrieve(opts...); attrs != nil {
		attributesToRetrieve := attrs.Get()
		if len(attributesToRetrieve) > 0 {
			opts = append(opts, opt.ExtraURLParams(map[string]string{
				"attributesToRetrieve": strings.Join(attributesToRetrieve, ","),
			}))
		}
	}

	path := i.path("/" + url.QueryEscape(objectID))
	return i.transport.Request(&object, http.MethodGet, path, nil, call.Read, opts...)
}

func (i *Index) SaveObject(object interface{}, opts ...interface{}) (res SaveObjectRes, err error) {
	path := i.path("")
	err = i.transport.Request(&res, http.MethodPost, path, object, call.Write, opts...)
	res.wait = i.waitTask
	return
}

func (i *Index) PartialUpdateObject(object interface{}, opts ...interface{}) (res UpdateTaskRes, err error) {
	objectID, ok := getObjectID(object)
	if !ok {
		err = errs.ErrMissingObjectID
		res.wait = noWait
		return
	}

	createIfNotExists := true
	if opt := iopt.ExtractCreateIfNotExists(opts...); opt != nil {
		createIfNotExists = opt.Get()
	}

	path := i.path("/" + url.QueryEscape(objectID) + "/partial")
	if !createIfNotExists {
		path += "?createIfNotExists=false"
	}

	err = i.transport.Request(&res, "POST", path, object, call.Write, opts...)
	res.wait = i.waitTask
	return
}

func (i *Index) DeleteObject(objectID string, opts ...interface{}) (res DeleteTaskRes, err error) {
	if objectID == "" {
		err = errs.ErrMissingObjectID
		res.wait = noWait
		return
	}

	path := i.path("/" + url.QueryEscape(objectID))
	err = i.transport.Request(&res, http.MethodDelete, path, nil, call.Write, opts...)
	res.wait = i.waitTask
	return
}

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
			ObjectID:             url.QueryEscape(objectID),
			AttributesToRetrieve: attributesToRetrieve,
		}
	}

	path := "/1/indexes/*/objects"
	return i.transport.Request(&res, http.MethodPost, path, body, call.Read, opts...)
}

func (i *Index) SaveObjects(objects interface{}, opts ...interface{}) (res MultipleBatchRes, err error) {
	return i.batch(objects, AddObject, opts...)
}

func (i *Index) PartialUpdateObjects(objects interface{}, opts ...interface{}) (res MultipleBatchRes, err error) {
	var (
		action BatchAction
	)

	createIfNotExists := true
	if opt := iopt.ExtractCreateIfNotExists(opts...); opt != nil {
		createIfNotExists = opt.Get()
	}

	if createIfNotExists {
		action = PartialUpdateObject
	} else {
		action = PartialUpdateObjectNoCreate
	}

	return i.batch(objects, action, opts...)
}

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

func (i *Index) batch(objects interface{}, action BatchAction, opts ...interface{}) (MultipleBatchRes, error) {
	var (
		batch      []interface{}
		operations []BatchOperation
		response   BatchRes
		res        MultipleBatchRes
	)

	autoGenerateObjectIDIfNotExist := false
	if opt := iopt.ExtractAutoGenerateObjectIDIfNotExist(opts...); opt != nil {
		autoGenerateObjectIDIfNotExist = opt.Get()
	}

	it := iterator.New(objects)

	for {
		object, err := it.Next()
		if err != nil {
			return res, fmt.Errorf("iteration failed unexpectedly: %v", err)
		}

		if !autoGenerateObjectIDIfNotExist && object != nil && !hasObjectID(object) {
			return res, fmt.Errorf("missing objectID in object %#v", object)
		}

		if len(batch) >= i.maxBatchSize || object == nil {
			operations, err = newOperationBatch(batch, action)
			if err != nil {
				return res, fmt.Errorf("could not generate intermediate batch: %v", err)
			}
			response, err = i.Batch(operations, opts...)
			if err != nil {
				return res, fmt.Errorf("could not send intermediate batch: %v", err)
			}
			res.Responses = append(res.Responses, response)
		} else {
			batch = append(batch, object)
		}

		if object == nil {
			break
		}
	}

	return res, nil
}

func (i *Index) Batch(operations []BatchOperation, opts ...interface{}) (res BatchRes, err error) {
	body := batchReq{Requests: operations}
	path := i.path("/batch")
	err = i.transport.Request(&res, http.MethodPost, path, body, call.Write, opts...)
	res.wait = i.waitTask
	return
}

func (i *Index) DeleteBy(opts ...interface{}) (res UpdateTaskRes, err error) {
	body := deleteByReq{
		AroundLatLng:      iopt.ExtractAroundLatLng(opts...),
		AroundRadius:      iopt.ExtractAroundRadius(opts...),
		FacetFilters:      iopt.ExtractFacetFilters(opts...),
		Filters:           iopt.ExtractFilters(opts...),
		InsideBoundingBox: iopt.ExtractInsideBoundingBox(opts...),
		InsidePolygon:     iopt.ExtractInsidePolygon(opts...),
		NumericFilters:    iopt.ExtractNumericFilters(opts...),
	}
	path := i.path("/deleteByQuery")
	err = i.transport.Request(&res, http.MethodPost, path, body, call.Write, opts...)
	res.wait = i.waitTask
	return
}

func (i *Index) Search(query string, opts ...interface{}) (res SearchRes, err error) {
	body := searchReq{Params: transport.URLEncode(newSearchParams(query, opts...))}
	path := i.path("/query")
	err = i.transport.Request(&res, http.MethodPost, path, body, call.Read, opts...)
	return
}
