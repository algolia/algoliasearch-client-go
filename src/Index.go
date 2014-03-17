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

func (i *Index) deleteIndex() (interface{}, error) {
  return i.client.transport.request("DELETE", "/1/indexes/" + i.nameEncoded, nil)
}

func (i *Index) clearIndex() (interface{}, error) {
  return i.client.transport.request("POST", "/1/indexes/" + i.nameEncoded + "/clear", nil)
}

func (i *Index) getObject(objectID string) (interface{}, error) {
  return i.client.transport.request("GET", "/1/indexes/" + i.nameEncoded + "/" + i.client.transport.urlEncode(objectID), nil)
}

func (i *Index) deleteObject(objectID string) (interface{}, error) {
  return i.client.transport.request("DELETE", "/1/indexes/" + i.nameEncoded + "/" +  i.client.transport.urlEncode(objectID), nil)
}

func (i *Index) getSettings() (interface{}, error) {
  return i.client.transport.request("GET", "/1/indexes/" + i.nameEncoded + "/settings", nil)
}

func (i *Index) setSettings(settings interface{}) (interface{}, error) {
  return i.client.transport.request("PUT", "/1/indexes/" + i.nameEncoded + "/settings", settings)
}

func (i *Index) getStatus(taskID float64) (interface{}, error) {
  return i.client.transport.request("GET", "/1/indexes/" + i.nameEncoded + "/task/" + strconv.FormatFloat(taskID, 'f', -1, 64), nil)
}

func (i *Index) waitTask(task interface{}) (interface{}, error) {
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

func (i *Index) listIndexKeys() (interface{}, error) {
  return i.client.transport.request("GET", "/1/indexes/" + i.nameEncoded + "/keys", nil)
}

func (i *Index) getIndexKey(key string) (interface{}, error) {
  return i.client.transport.request("GET", "/1/indexes/" + i.nameEncoded + "/keys/" + key , nil)
}

func (i *Index) deleteIndexKey(key string) (interface{}, error) {
  return i.client.transport.request("DELETE", "/1/indexes/" + i.nameEncoded + "/keys/" + key , nil)
}

func (i *Index) addObject(object interface{}) (interface{}, error) {
  method := "POST"
  path := "/1/indexes/" + i.nameEncoded
  if id, ok := object.(map[string]interface{})["objectID"]; ok {
    method = "PUT"
    path = path + "/" + i.client.transport.urlEncode(id.(string))
  }
  return i.client.transport.request(method, path, object)
}

func (i *Index) updateObject(object interface{}) (interface{}, error) {
  id := object.(map[string]interface{})["objectID"]
  path := "/1/indexes/" + i.nameEncoded + "/" + i.client.transport.urlEncode(id.(string))
  return i.client.transport.request("PUT", path, object)
}

func (i *Index) partialUpdateObject(object interface{}) (interface{}, error) {
  id := object.(map[string]interface{})["objectID"]
  path := "/1/indexes/" + i.nameEncoded + "/" + i.client.transport.urlEncode(id.(string)) + "/partial"
  return i.client.transport.request("POST", path, object)
}

func (i *Index) browse(page, hitsPerPage int) (interface{}, error) {
  return i.client.transport.request("GET", "/1/indexes/" + i.nameEncoded + "/browse?page=" + strconv.Itoa(page) + "&hitsPerPage=" + strconv.Itoa(hitsPerPage) , nil)
}

func (i *Index) query(query interface{}) (interface{}, error) {
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

func (i *Index) copy(name string) (interface{}, error) {
  return i.operation(name, "copy")
}

func (i *Index) move(name string) (interface{}, error) {
  return i.operation(name, "move")
}

func (i *Index) addKey(acl []string, validity int, maxQueriesPerIPPerHour int, maxHitsPerQuery int) (interface{}, error) {
  body := make(map[string]interface{})
  body["acl"] = acl
  body["maxHitsPerQuery"] = maxHitsPerQuery
  body["maxQueriesPerIPPerHour"] = maxQueriesPerIPPerHour
  body["validity"] = validity
  return i.client.transport.request("POST", "/1/indexes/" + i.nameEncoded + "/keys", body)
}
