package main

import "time"
import "strconv"
import "net/url"
import "reflect"

type Index struct {
  name string
  nameEncoded string
  client *Client
}

func NewIndex(name string, client *Client) *Index {
  index := new(Index)
  index.name = name
  index.client = client
  index.nameEncoded = client.transport.urlEncode(name)
  return index
}

func (i *Index) Delete() (interface{}, error) {
  return i.client.transport.request("DELETE", "/1/indexes/" + i.nameEncoded, nil)
}

func (i *Index) Clear() (interface{}, error) {
  return i.client.transport.request("POST", "/1/indexes/" + i.nameEncoded + "/clear", nil)
}

func (i *Index) GetObject(objectID string) (interface{}, error) {
  return i.client.transport.request("GET", "/1/indexes/" + i.nameEncoded + "/" + i.client.transport.urlEncode(objectID), nil)
}

func (i *Index) DeleteObject(objectID string) (interface{}, error) {
  return i.client.transport.request("DELETE", "/1/indexes/" + i.nameEncoded + "/" +  i.client.transport.urlEncode(objectID), nil)
}

func (i *Index) GetSettings() (interface{}, error) {
  return i.client.transport.request("GET", "/1/indexes/" + i.nameEncoded + "/settings", nil)
}

func (i *Index) SetSettings(settings interface{}) (interface{}, error) {
  return i.client.transport.request("PUT", "/1/indexes/" + i.nameEncoded + "/settings", settings)
}

func (i *Index) getStatus(taskID float64) (interface{}, error) {
  return i.client.transport.request("GET", "/1/indexes/" + i.nameEncoded + "/task/" + strconv.FormatFloat(taskID, 'f', -1, 64), nil)
}

func (i *Index) WaitTask(task interface{}) (interface{}, error) {
  for true {
    status, err := i.getStatus(task.(map[string]interface{})["taskID"].(float64))
    if err != nil {
      return nil, err
    }
    if status.(map[string]interface{})["status"] == "published" {
      break
    }
    time.Sleep(time.Duration(100) * time.Millisecond) 
  }
  return task, nil
}

func (i *Index) ListKeys() (interface{}, error) {
  return i.client.transport.request("GET", "/1/indexes/" + i.nameEncoded + "/keys", nil)
}

func (i *Index) GetKey(key string) (interface{}, error) {
  return i.client.transport.request("GET", "/1/indexes/" + i.nameEncoded + "/keys/" + key , nil)
}

func (i *Index) DeleteKey(key string) (interface{}, error) {
  return i.client.transport.request("DELETE", "/1/indexes/" + i.nameEncoded + "/keys/" + key , nil)
}

func (i *Index) AddObject(object interface{}) (interface{}, error) {
  method := "POST"
  path := "/1/indexes/" + i.nameEncoded
  if id, ok := object.(map[string]interface{})["objectID"]; ok {
    method = "PUT"
    path = path + "/" + i.client.transport.urlEncode(id.(string))
  }
  return i.client.transport.request(method, path, object)
}

func (i *Index) UpdateObject(object interface{}) (interface{}, error) {
  id := object.(map[string]interface{})["objectID"]
  path := "/1/indexes/" + i.nameEncoded + "/" + i.client.transport.urlEncode(id.(string))
  return i.client.transport.request("PUT", path, object)
}

func (i *Index) PartialUpdateObject(object interface{}) (interface{}, error) {
  id := object.(map[string]interface{})["objectID"]
  path := "/1/indexes/" + i.nameEncoded + "/" + i.client.transport.urlEncode(id.(string)) + "/partial"
  return i.client.transport.request("POST", path, object)
}

func (i *Index) AddObjects(objects interface{}) (interface{}, error) {
  return i.sameBatch(objects, "addObject")
}

func (i *Index) UpdateObjects(objects interface{}) (interface{}, error) {
  return i.sameBatch(objects, "updateObject")
}

func (i *Index) PartialUpdateObjects(objects interface{}) (interface{}, error) {
  return i.sameBatch(objects, "partialUpdateObject")
}

func (i *Index) sameBatch(objects interface{}, action string) (interface{}, error) {
  length := len(objects.([]interface{}))
  method := make([]string, length)
  for i := range method {
    method[i] = action
  }
  return i.Batch(objects, method)
}

func (i *Index) Batch(objects interface{}, actions []string) (interface{}, error) {
  array := objects.([]interface{})
  queries := make([]map[string]interface{}, len(array))
  for i := range array {
    queries[i] = make(map[string]interface{})
    queries[i]["action"] = actions[i]
    queries[i]["body"] = array[i]
  }
  return i.customBatch(queries)
}

func (i *Index) customBatch(queries interface{}) (interface{}, error) {
  request :=  make(map[string]interface{})
  request["requests"] = queries
  return i.client.transport.request("POST", "/1/indexes/" + i.nameEncoded + "/batch", request)
}

func (i *Index) Browse(page, hitsPerPage int) (interface{}, error) {
  return i.client.transport.request("GET", "/1/indexes/" + i.nameEncoded + "/browse?page=" + strconv.Itoa(page) + "&hitsPerPage=" + strconv.Itoa(hitsPerPage) , nil)
}

func (i *Index) Query(query interface{}) (interface{}, error) {
  v := url.Values{}
  for key, value := range query.(map[string]interface{}) {
    if reflect.TypeOf(value).Name() == "string" {
      v.Add(key, value.(string))
    } else if reflect.TypeOf(value).Name() == "float64" {
      v.Add(key, strconv.FormatFloat(value.(float64), 'f', -1, 64))
    } else {
      v.Add(key, strconv.Itoa(value.(int)))
    }
  }
  body := make(map[string]interface{})
  body["params"] = v.Encode()
  return i.client.transport.request("POST", "/1/indexes/" + i.nameEncoded + "/query", body)
}

func (i *Index) operation(name, op string) (interface{}, error) {
  body := make(map[string]interface{})
  body["operation"] = op
  body["destination"] = name
  return i.client.transport.request("POST", "/1/indexes/" + i.nameEncoded + "/operation", body)
}

func (i *Index) Copy(name string) (interface{}, error) {
  return i.operation(name, "copy")
}

func (i *Index) Move(name string) (interface{}, error) {
  return i.operation(name, "move")
}

func (i *Index) AddKey(acl []string, validity int, maxQueriesPerIPPerHour int, maxHitsPerQuery int) (interface{}, error) {
  body := make(map[string]interface{})
  body["acl"] = acl
  body["maxHitsPerQuery"] = maxHitsPerQuery
  body["maxQueriesPerIPPerHour"] = maxQueriesPerIPPerHour
  body["validity"] = validity
  return i.client.transport.request("POST", "/1/indexes/" + i.nameEncoded + "/keys", body)
}
