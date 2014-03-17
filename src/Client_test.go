package main

import "syscall"
import "time"
import "strconv"
import "testing"

func initTest(t *testing.T) (*Client, *Index) {
  appID, haveAppID := syscall.Getenv("ALGOLIA_APPLICATION_ID")
  apiKey, haveApiKey := syscall.Getenv("ALGOLIA_API_KEY")
  if !haveApiKey || !haveAppID {
    t.Fatalf("Need ALGOLIA_APPLICATION_ID and ALGOLIA_API_KEY")
  }
  client := NewClient(appID, apiKey)
  index := client.initIndex("àlgol?à-go")
  return client, index
}

func shouldHave(json interface{}, attr, msg string, t *testing.T) {
  _, ok := json.(map[string]interface{})[attr]
  if !ok  {
    t.Fatalf(msg + ", expected attribute: " + attr)
  }
}

func shouldNotHave(json interface{}, attr, msg string, t *testing.T) {
  _, ok := json.(map[string]interface{})[attr]
  if ok  {
    t.Fatalf(msg + ", unexpected attribute: " + attr)
  }
}

func shouldStr(json interface{}, attr, value, msg string, t *testing.T) {
  resp, ok := json.(map[string]interface{})[attr]
  if !ok || value != resp.(string) {
    t.Fatalf(msg + ", expected: " + value + " have: " + resp.(string))
  }
}

func shouldFloat(json interface{}, attr string, value float64, msg string, t *testing.T) {
  resp, ok := json.(map[string]interface{})[attr]
  if !ok || value != resp.(float64) {
    t.Fatalf(msg + ", expected: " + strconv.FormatFloat(value, 'f', -1, 64) + " have: " + strconv.FormatFloat(resp.(float64), 'f', -1, 64))
  }
}

func TestClear(t *testing.T) {
  _, index := initTest(t)
  object := make(map[string]interface{})
  object["name"] = "John Snow"
  index.waitTask(index.addObject(object))
  time.Sleep(time.Duration(100) * time.Millisecond) 
  index.waitTask(index.clearIndex())
  results := index.query(make(map[string]interface{}))
  shouldFloat(results, "nbHits", 0, "Unable to clear the index", t)
  index.deleteIndex()
}

func TestAddObject(t *testing.T) {
  _, index := initTest(t)
  object := make(map[string]interface{})
  object["name"] = "John Snow"
  _ = index.addObject(object)
  object["name"] = "John Snow"
  object["objectID"] = "àlgol?à"
  _ = index.addObject(object)
  results := index.query(make(map[string]interface{}))
  shouldFloat(results, "nbHits", 2, "Unable to clear the index", t)
  index.deleteIndex()
}

func TestUpdateObject(t *testing.T) {
  _, index := initTest(t)
  object := make(map[string]interface{})
  object["name"] = "John Snow"
  object["objectID"] = "àlgol?à"
  _ = index.addObject(object)
  object["name"] = "Roger"
  _ = index.updateObject(object)
  results := index.query(make(map[string]interface{}))
  hits := results.(map[string]interface{})["hits"]
  shouldStr(hits.([]interface{})[0], "name", "Roger", "Unable to update an object", t)
  shouldNotHave(hits.([]interface{})[0], "job", "Unable to update an object", t)
  index.deleteIndex()
}

func TestPartialUpdateObject(t *testing.T) {
  _, index := initTest(t)
  object := make(map[string]interface{})
  object["name"] = "John Snow"
  object["job"] = "Knight"
  object["objectID"] = "àlgol?à"
  _ = index.addObject(object)
  delete(object, "job")
  object["name"] = "Roger"
  _ = index.partialUpdateObject(object)
  results := index.query(make(map[string]interface{}))
  hits := results.(map[string]interface{})["hits"]
  shouldStr(hits.([]interface{})[0], "name", "Roger", "Unable to update an object", t)
  index.deleteIndex()
}


func TestGetObject(t *testing.T) {
  _, index := initTest(t)
  object := make(map[string]interface{})
  object["name"] = "John Snow"
  object["objectID"] = "àlgol?à"
  _ = index.waitTask(index.addObject(object))
  obj := index.getObject("àlgol?à")
  shouldStr(obj, "name", "John Snow", "Unable to update an object", t)
  index.deleteIndex()
}

func TestDeleteObject(t *testing.T) {
  _, index := initTest(t)
  object := make(map[string]interface{})
  object["name"] = "John Snow"
  object["objectID"] = "àlgol?à"
  _ = index.waitTask(index.addObject(object))
  _ = index.waitTask(index.deleteObject("àlgol?à"))
  results := index.query(make(map[string]interface{}))
  shouldFloat(results, "nbHits", 0, "Unable to clear the index", t)
  index.deleteIndex()
}

func TestSetSettings(t *testing.T) {
  _, index := initTest(t)
  settings := make(map[string]interface{})
  settings["hitsPerPage"] = 30
  _ = index.waitTask(index.setSettings(settings))
  settingsChanged := index.getSettings()
  shouldFloat(settingsChanged, "hitsPerPage", 30, "Unable to change setting", t)
  index.deleteIndex()
}

func TestGetLogs(t *testing.T) {
  client, _ := initTest(t)
  logs := client.getLogs(0, 100, false)
  shouldHave(logs, "logs", "Unable to get logs", t)
}

func TestBrowse(t *testing.T) {
  _, index := initTest(t)
  object := make(map[string]interface{})
  object["name"] = "John Snow"
  object["objectID"] = "àlgol?à"
  _ = index.waitTask(index.addObject(object))
  items := index.browse(1, 1)
  shouldHave(items, "hits", "Unable to browse index", t)
  index.deleteIndex()
}

func TestQuery(t *testing.T) {
  _, index := initTest(t)
  object := make(map[string]interface{})
  object["name"] = "John Snow"
  object["objectID"] = "àlgol?à"
  _ = index.addObject(object)
  query := make(map[string]interface{})
  query["query"] = ""
  query["attributesToRetrieve"] = "*"
  query["getRankingInfo"] = 1
  results := index.query(query)
  shouldFloat(results, "nbHits", 1, "Unable to query an index", t)
}

func TestCopy(t *testing.T) {
  client, index := initTest(t)
  object := make(map[string]interface{})
  object["name"] = "John Snow"
  object["objectID"] = "àlgol?à"
  _ = index.addObject(object)
  _ = index.waitTask(index.copy("àlgo?à2-go"))
  indexCopy := client.initIndex("àlgo?à2-go")
  results := indexCopy.query(make(map[string]interface{}))
  shouldFloat(results, "nbHits", 1, "Unable to copy an index", t)
  index.deleteIndex()
  indexCopy.deleteIndex()
}

func TestMove(t *testing.T) {
  client, index := initTest(t)
  object := make(map[string]interface{})
  object["name"] = "John Snow"
  object["objectID"] = "àlgol?à"
  _ = index.addObject(object)
  _ = index.waitTask(index.move("àlgo?à2-go"))
  indexMove := client.initIndex("àlgo?à2-go")
  results := indexMove.query(make(map[string]interface{}))
  shouldFloat(results, "nbHits", 1, "Unable to move an index", t)
  indexMove.deleteIndex()
}

func TestKeepAlive(t *testing.T) {
  _, index := initTest(t)
  object := make(map[string]interface{})
  object["name"] = "John Snow"
  object["objectID"] = "àlgol?à"
  _ = index.addObject(object)
  query := make(map[string]interface{})
  for i := 0; i < 100; i++ {
    index.query(query)
  }
}
