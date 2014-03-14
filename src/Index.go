package main

import "time"
import "strconv"
import "net/url"

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

func (i *Index) deleteIndex() interface{} {
  return i.client.transport.request("DELETE", "/1/indexes/" + i.nameEncoded, nil)
}

func (i *Index) clearIndex() interface{} {
  return i.client.transport.request("POST", "/1/indexes/" + i.nameEncoded + "/clear", nil)
}

func (i *Index) getObject(objectID string) interface{} {
  return i.client.transport.request("GET", "/1/indexes/" + i.nameEncoded + "/" + i.client.transport.urlEncode(objectID), nil)
}

func (i *Index) deleteObject(objectID string) interface{} {
  return i.client.transport.request("DELETE", "/1/indexes/" + i.nameEncoded + "/" +  i.client.transport.urlEncode(objectID), nil)
}

func (i *Index) getSettings() interface{} {
  return i.client.transport.request("GET", "/1/indexes/" + i.nameEncoded + "/settings", nil)
}

func (i *Index) setSettings(settings interface{}) interface{} {
  return i.client.transport.request("PUT", "/1/indexes/" + i.nameEncoded + "/settings", settings)
}

func (i *Index) getStatus(taskID float64) interface{} {
  return i.client.transport.request("GET", "/1/indexes/" + i.nameEncoded + "/task/" + strconv.FormatFloat(taskID, 'f', -1, 64), nil)
}

func (i *Index) waitTask(task interface{}) interface{} {
  for true {
    status := i.getStatus(task.(map[string]interface{})["taskID"].(float64))
    if status.(map[string]interface{})["status"] == "published" {
      break
    }
    time.Sleep(time.Duration(100) * time.Millisecond) 
  }
  return task
}

func (i *Index) listIndexKeys() interface{} {
  return i.client.transport.request("GET", "/1/indexes/" + i.nameEncoded + "/keys", nil)
}

func (i *Index) getIndexKey(key string) interface{} {
  return i.client.transport.request("GET", "/1/indexes/" + i.nameEncoded + "/keys/" + key , nil)
}

func (i *Index) deleteIndexKey(key string) interface{} {
  return i.client.transport.request("DELETE", "/1/indexes/" + i.nameEncoded + "/keys/" + key , nil)
}

func (i *Index) addObject(object interface{}) interface{} {
  method := "POST"
  path := "/1/indexes/" + i.nameEncoded
  if id, ok := object.(map[string]interface{})["objectID"]; ok {
    method = "PUT"
    path = path + "/" + i.client.transport.urlEncode(id.(string))
  }
  return i.client.transport.request(method, path, object)
}

func (i *Index) updateObject(object interface{}) interface{} {
  id := object.(map[string]interface{})["objectID"]
  path := "/1/indexes/" + i.nameEncoded + "/" + i.client.transport.urlEncode(id.(string))
  return i.client.transport.request("PUT", path, object)
}

func (i *Index) partialUpdateObject(object interface{}) interface{} {
  id := object.(map[string]interface{})["objectID"]
  path := "/1/indexes/" + i.nameEncoded + "/" + i.client.transport.urlEncode(id.(string)) + "/partial"
  return i.client.transport.request("POST", path, object)
}

func (i *Index) browse(page, hitsPerPage int) interface{} {
  body := make(map[string]interface{})
  body["page"] = page
  body["hitsPerPage"] = hitsPerPage
  return i.client.transport.request("GET", "/1/indexes/" + i.nameEncoded + "/browse", body)
}

func (i *Index) query(query interface{}) interface{} {
  v := url.Values{}
  for key, value := range query.(map[string]interface{}) {
    v.Add(key, value.(string))
  }
  body := make(map[string]interface{})
  body["params"] = v.Encode()
  return i.client.transport.request("POST", "/1/indexes/" + i.nameEncoded + "/query", body)
}

func (i *Index) operation(name, op string) interface{} {
  body := make(map[string]interface{})
  body["operation"] = op
  body["destination"] = name
  return i.client.transport.request("POST", "/1/indexes/" + i.nameEncoded + "/operation", body)
}

func (i *Index) copy(name string) interface{} {
  return i.operation(name, "copy")
}

func (i *Index) move(name string) interface{} {
  return i.operation(name, "move")
}

func (i *Index) addKey(acl []string, validity int, maxQueriesPerIPPerHour int, maxHitsPerQuery int) interface{} {
  body := make(map[string]interface{})
  body["acl"] = acl
  body["maxHitsPerQuery"] = maxHitsPerQuery
  body["maxQueriesPerIPPerHour"] = maxQueriesPerIPPerHour
  body["validity"] = validity
  return i.client.transport.request("POST", "/1/indexes/" + i.nameEncoded + "/keys", body)
}
