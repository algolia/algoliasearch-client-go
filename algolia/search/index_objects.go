package search

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/algolia/algoliasearch-client-go/algolia/call"
	"github.com/algolia/algoliasearch-client-go/algolia/errs"
	"github.com/algolia/algoliasearch-client-go/algolia/iterator"
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func (i *Index) GetObject(objectID string, object interface{}, opts ...interface{}) error {
	attributesToRetrieve := opt.ExtractAttributesToRetrieve(opts...)
	if attributesToRetrieve != nil {
		attrs, err := json.Marshal(attributesToRetrieve)
		if err != nil {
			return fmt.Errorf("cannot serialize attributesToRetrieve %v: %v", attrs, err)
		}
		opts = append(opts, opt.ExtraURLParams(map[string]string{"attributesToRetrieve": string(attrs)}))
	}

	path := i.path("/" + url.QueryEscape(objectID))
	return i.transport.Request(&object, http.MethodGet, path, nil, call.Read, opts...)
}

func (i *Index) GetObjects(objectIDs []string, objects interface{}, opts ...interface{}) error {
	var (
		attributesToRetrieve = opt.ExtractAttributesToRetrieve(opts...)
		requests             = make([]map[string]interface{}, len(objectIDs))
		body                 = map[string]interface{}{"requests": requests}
		res                  = getObjectsRes{objects}
	)

	for j, objectID := range objectIDs {
		requests[j] = map[string]interface{}{
			"indexName": i.name,
			"objectID":  url.QueryEscape(objectID),
		}
		if attributesToRetrieve != nil {
			requests[j]["attributesToRetrieve"] = attributesToRetrieve
		}
	}

	path := "/1/indexes/*/objects"
	return i.transport.Request(&res, http.MethodPost, path, body, call.Read, opts...)
}

type getObjectsRes struct {
	Results interface{} `json:"results"`
}

func (i *Index) SaveObject(object interface{}, opts ...interface{}) (res SaveObjectRes, err error) {
	path := i.path("")
	err = i.transport.Request(&res, http.MethodPost, path, object, call.Write, opts...)
	res.wait = i.waitTask
	return
}

func (i *Index) SaveObjects(objects interface{}, opts ...interface{}) (res MultipleBatchRes, err error) {
	return i.Batch(objects, AddObject, opts...)
}

func (i *Index) PartialUpdateObject(object interface{}, opts ...interface{}) (res UpdateTaskRes, err error) {
	objectID, ok := getObjectID(object)
	if !ok {
		err = errs.ErrMissingObjectID
		res.wait = noWait
		return
	}

	createIfNotExists := opt.ExtractCreateIfNotExists(opts...)

	path := i.path("/" + url.QueryEscape(objectID) + "/partial")
	if !createIfNotExists {
		path += "?createIfNotExists=false"
	}

	err = i.transport.Request(&res, "POST", path, object, call.Write, opts...)
	res.wait = i.waitTask
	return
}

func (i *Index) PartialUpdateObjects(objects interface{}, opts ...interface{}) (res MultipleBatchRes, err error) {
	var (
		action            BatchAction
		createIfNotExists = opt.ExtractCreateIfNotExists(opts...)
	)

	if createIfNotExists {
		action = PartialUpdateObject
	} else {
		action = PartialUpdateObjectNoCreate
	}

	return i.Batch(objects, action, opts...)
}

func (i *Index) Batch(objects interface{}, action BatchAction, opts ...interface{}) (MultipleBatchRes, error) {
	var (
		batch      []interface{}
		operations []BatchOperation
		response   BatchRes
		res        MultipleBatchRes
	)

	it := iterator.New(objects)
	autoGenerateObjectIDIfNotExist := opt.ExtractAutoGenerateObjectIDIfNotExist(opts...)

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
			response, err = i.batch(operations, opts...)
			if err != nil {
				return res, fmt.Errorf("could not send intermediate batch: %v", err)
			}
			res.responses = append(res.responses, response)
		} else {
			batch = append(batch, object)
		}

		if object == nil {
			break
		}
	}

	return res, nil
}

func (i *Index) batch(operations []BatchOperation, opts ...interface{}) (res BatchRes, err error) {
	path := i.path("/batch")
	body := map[string][]BatchOperation{
		"requests": operations,
	}
	err = i.transport.Request(&res, http.MethodPost, path, body, call.Write, opts...)
	res.wait = i.waitTask
	return
}
