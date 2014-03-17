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
  resp, err := index.addObject(object)
  if err != nil {
    t.Fatalf(err.Error())
  }
  index.waitTask(resp)
  resp, err = index.addObject(object)
  if err != nil {
    t.Fatalf(err.Error())
  }
  time.Sleep(time.Duration(100) * time.Millisecond) 
  resp, err = index.clearIndex()
  if err != nil {
    t.Fatalf(err.Error())
  }
  index.waitTask(resp)
  results, err := index.query(make(map[string]interface{}))
  if err != nil {
    t.Fatalf(err.Error())
  }
  shouldFloat(results, "nbHits", 0, "Unable to clear the index", t)
  index.deleteIndex()
}

func TestAddObject(t *testing.T) {
  _, index := initTest(t)
  object := make(map[string]interface{})
  object["name"] = "John Snow"
  _, err := index.addObject(object)
  if err != nil {
    t.Fatalf(err.Error())
  }
  object["name"] = "John Snow"
  object["objectID"] = "àlgol?à"
  _, err = index.addObject(object)
  if err != nil {
    t.Fatalf(err.Error())
  }
  results, err := index.query(make(map[string]interface{}))
  if err != nil {
    t.Fatalf(err.Error())
  }
  shouldFloat(results, "nbHits", 2, "Unable to clear the index", t)
  index.deleteIndex()
}

func TestUpdateObject(t *testing.T) {
  _, index := initTest(t)
  object := make(map[string]interface{})
  object["name"] = "John Snow"
  object["objectID"] = "àlgol?à"
  _, err := index.addObject(object)
  if err != nil {
    t.Fatalf(err.Error())
  }
  object["name"] = "Roger"
  _, err = index.updateObject(object)
  if err != nil {
    t.Fatalf(err.Error())
  }
  results, err := index.query(make(map[string]interface{}))
  if err != nil {
    t.Fatalf(err.Error())
  }
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
  _, err := index.addObject(object)
  if err != nil {
    t.Fatalf(err.Error())
  }
  delete(object, "job")
  object["name"] = "Roger"
  _, err = index.partialUpdateObject(object)
  if err != nil {
    t.Fatalf(err.Error())
  }
  results, err := index.query(make(map[string]interface{}))
  if err != nil {
    t.Fatalf(err.Error())
  }
  hits := results.(map[string]interface{})["hits"]
  shouldStr(hits.([]interface{})[0], "name", "Roger", "Unable to update an object", t)
  index.deleteIndex()
}


func TestGetObject(t *testing.T) {
  _, index := initTest(t)
  object := make(map[string]interface{})
  object["name"] = "John Snow"
  object["objectID"] = "àlgol?à"
  resp, err := index.addObject(object)
  if err != nil {
    t.Fatalf(err.Error())
  }
  _, err = index.waitTask(resp)
  if err != nil {
    t.Fatalf(err.Error())
  }
  resp, err = index.addObject(object)
  if err != nil {
    t.Fatalf(err.Error())
  }
  obj, err := index.getObject("àlgol?à")
  if err != nil {
    t.Fatalf(err.Error())
  }
  shouldStr(obj, "name", "John Snow", "Unable to update an object", t)
  index.deleteIndex()
}

func TestDeleteObject(t *testing.T) {
  _, index := initTest(t)
  object := make(map[string]interface{})
  object["name"] = "John Snow"
  object["objectID"] = "àlgol?à"
  resp, err := index.addObject(object)
  if err != nil {
    t.Fatalf(err.Error())
  }
  _, err = index.waitTask(resp)
  if err != nil {
    t.Fatalf(err.Error())
  }
  resp, err = index.deleteObject("àlgol?à")
  if err != nil {
    t.Fatalf(err.Error())
  }
  _, err = index.waitTask(resp)
  if err != nil {
    t.Fatalf(err.Error())
  }
  results, err := index.query(make(map[string]interface{}))
  if err != nil {
    t.Fatalf(err.Error())
  }
  shouldFloat(results, "nbHits", 0, "Unable to clear the index", t)
  index.deleteIndex()
}

func TestSetSettings(t *testing.T) {
  _, index := initTest(t)
  settings := make(map[string]interface{})
  settings["hitsPerPage"] = 30
  resp, err := index.setSettings(settings)
  if err != nil {
    t.Fatalf(err.Error())
  }
  _, err = index.waitTask(resp)
  if err != nil {
    t.Fatalf(err.Error())
  }
  settingsChanged, err := index.getSettings()
  if err != nil {
    t.Fatalf(err.Error())
  }
  shouldFloat(settingsChanged, "hitsPerPage", 30, "Unable to change setting", t)
  index.deleteIndex()
}

func TestGetLogs(t *testing.T) {
  client, _ := initTest(t)
  logs, err := client.getLogs(0, 100, false)
  if err != nil {
    t.Fatalf(err.Error())
  }
  shouldHave(logs, "logs", "Unable to get logs", t)
}

func TestBrowse(t *testing.T) {
  _, index := initTest(t)
  object := make(map[string]interface{})
  object["name"] = "John Snow"
  object["objectID"] = "àlgol?à"
  resp, err := index.addObject(object)
  if err != nil {
    t.Fatalf(err.Error())
  }
  _, err = index.waitTask(resp)
  if err != nil {
    t.Fatalf(err.Error())
  }
  items, err := index.browse(1, 1)
  if err != nil {
    t.Fatalf(err.Error())
  }
  shouldHave(items, "hits", "Unable to browse index", t)
  index.deleteIndex()
}

func TestQuery(t *testing.T) {
  _, index := initTest(t)
  object := make(map[string]interface{})
  object["name"] = "John Snow"
  object["objectID"] = "àlgol?à"
  _, err := index.addObject(object)
  if err != nil {
    t.Fatalf(err.Error())
  }
  query := make(map[string]interface{})
  query["query"] = ""
  query["attributesToRetrieve"] = "*"
  query["getRankingInfo"] = 1
  results, err := index.query(query)
  if err != nil {
    t.Fatalf(err.Error())
  }
  shouldFloat(results, "nbHits", 1, "Unable to query an index", t)
}

func TestCopy(t *testing.T) {
  client, index := initTest(t)
  object := make(map[string]interface{})
  object["name"] = "John Snow"
  object["objectID"] = "àlgol?à"
  _, err := index.addObject(object)
  if err != nil {
    t.Fatalf(err.Error())
  }
  resp, err := index.copy("àlgo?à2-go")
  if err != nil {
    t.Fatalf(err.Error())
  }
  _, err = index.waitTask(resp)
  if err != nil {
    t.Fatalf(err.Error())
  }
  indexCopy := client.initIndex("àlgo?à2-go")
  results, err := indexCopy.query(make(map[string]interface{}))
  if err != nil {
    t.Fatalf(err.Error())
  }
  shouldFloat(results, "nbHits", 1, "Unable to copy an index", t)
  index.deleteIndex()
  indexCopy.deleteIndex()
}

func TestMove(t *testing.T) {
  client, index := initTest(t)
  object := make(map[string]interface{})
  object["name"] = "John Snow"
  object["objectID"] = "àlgol?à"
  _, err := index.addObject(object)
  if err != nil {
    t.Fatalf(err.Error())
  }
  resp, err := index.move("àlgo?à2-go")
  if err != nil {
    t.Fatalf(err.Error())
  }
  _, err = index.waitTask(resp)
  if err != nil {
    t.Fatalf(err.Error())
  }
  indexMove := client.initIndex("àlgo?à2-go")
  results, err := indexMove.query(make(map[string]interface{}))
  if err != nil {
    t.Fatalf(err.Error())
  }
  shouldFloat(results, "nbHits", 1, "Unable to move an index", t)
  indexMove.deleteIndex()
}

func TestKeepAlive(t *testing.T) {
  _, index := initTest(t)
  object := make(map[string]interface{})
  object["name"] = "John Snow"
  object["objectID"] = "àlgol?à"
  _, err := index.addObject(object)
  if err != nil {
    t.Fatalf(err.Error())
  }
  query := make(map[string]interface{})
  for i := 0; i < 100; i++ {
    index.query(query)
  }
}
