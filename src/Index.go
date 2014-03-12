package main

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
  index.nameEncoded = url.QueryEscape(name)
  return index
}

func (i *Index) deleteIndex() interface{} {
  return i.client.transport.request("DELETE", "/1/indexes/" + i.nameEncoded, "")
}

func (i *Index) clearIndex() interface{} {
  return i.client.transport.request("POST", "/1/indexes/" + i.nameEncoded + "/clear", "")
}

func (i *Index) getObject(objectID string) interface{} {
  return i.client.transport.request("GET", "/1/indexes/" + i.nameEncoded + "/" + url.QueryEscape(objectID), "")
}

func (i *Index) deleteObject(objectID string) interface{} {
  return i.client.transport.request("DELETE", "/1/indexes/" + i.nameEncoded + "/" +  url.QueryEscape(objectID), "")
}

func (i *Index) getSettings() interface{} {
  return i.client.transport.request("GET", "/1/indexes/" + i.nameEncoded + "/settings", "")
}

func (i *Index) setSettings(settings interface{}) interface{} {
  return i.client.transport.request("PUT", "/1/indexes/" + i.nameEncoded + "/settings", settings)
}

func (i *Index) getStatus(taskID string) interface{} {
  return i.client.transport.request("GET", "/1/indexes/" + i.nameEncoded + "/task/" + taskID, "")
}

func (i *Index) listIndexKeys() interface{} {
  return i.client.transport.request("GET", "/1/indexes/" + i.nameEncoded + "/keys", "")
}

func (i *Index) getIndexKey(key string) interface{} {
  return i.client.transport.request("GET", "/1/indexes/" + i.nameEncoded + "/keys/" + key , "")
}

func (i *Index) deleteIndexKey(key string) interface{} {
  return i.client.transport.request("DELETE", "/1/indexes/" + i.nameEncoded + "/keys/" + key , "")
}

func (i *Index) addObject(object interface{}) interface{} {
  method := "POST"
  path := "/1/indexes/" + i.nameEncoded
  if id, ok := object.(map[string]interface{})["objectID"]; ok {
    method = "PUT"
    path = path + "/" + url.QueryEscape(id.(string))
  }
  return i.client.transport.request(method, path, object)
}

func (i *Index) updateObject(object interface{}) interface{} {
  id := object.(map[string]interface{})["objectID"]
  path := "/1/indexes/" + i.nameEncoded + "/" + url.QueryEscape(id.(string))
  return i.client.transport.request("PUT", path, object)
}

func (i *Index) partialUpdateObject(object interface{}) interface{} {
  id := object.(map[string]interface{})["objectID"]
  path := "/1/indexes/" + i.nameEncoded + "/" + url.QueryEscape(id.(string)) + "/partial"
  return i.client.transport.request("POST", path, object)
}

func (i *Index) browse(page, hitsPerPage int) interface{} {
  body := make(map[string]interface{})
  body["page"] = page
  body["hitsPerPage"] = hitsPerPage
  return i.client.transport.request("GET", "/1/indexes/" + i.nameEncoded + "/browse", body)
}

func (i *Index) query(query interface{}) interface{} {
  return i.client.transport.request("GET", "/1/indexes/" + i.nameEncoded, query)
}

func (i *Index) queryPost(query interface{}) interface{} {
  return i.client.transport.request("POST", "/1/indexes/" + i.nameEncoded + "/query", query)
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
