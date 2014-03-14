package main

import "syscall"
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

func TestClear(t *testing.T) {
  _, index := initTest(t)
  object := make(map[string]interface{})
  _ = index.addObject(object)
  _ = index.clearIndex()
  //TODO
}

func TestAddObject(t *testing.T) {
  _, index := initTest(t)
  object := make(map[string]interface{})
  object["name"] = "John Snow"
  _ = index.addObject(object)
  object["objectID"] = "algolia"
  _ = index.addObject(object)
  //TODO check index
}

func TestUpdateObject(t *testing.T) {
  _, index := initTest(t)
  object := make(map[string]interface{})
  object["name"] = "John Snow"
  object["objectID"] = "algolia"
  _ = index.addObject(object)
  object["name"] = "Roger"
  _ = index.updateObject(object)
  //TODO check return
}

func TestPartialUpdateObject(t *testing.T) {
  _, index := initTest(t)
  object := make(map[string]interface{})
  object["name"] = "John Snow"
  object["job"] = "Knight"
  object["objectID"] = "algolia"
  _ = index.addObject(object)
  delete(object, "job")
  object["name"] = "Roger"
  _ = index.partialUpdateObject(object)
  //TODO check return
}


func TestGetObject(t *testing.T) {
  _, index := initTest(t)
  object := make(map[string]interface{})
  object["name"] = "John Snow"
  object["objectID"] = "algolia"
  _ = index.addObject(object)
  _ = index.getObject("algolia")
  //TODO check return
}

func TestDeleteObject(t *testing.T) {
  _, index := initTest(t)
  object := make(map[string]interface{})
  object["name"] = "John Snow"
  object["objectID"] = "algolia"
  _ = index.addObject(object)
  _ = index.deleteObject("algolia")
  //TODO check return
}

func TestGetSettings(t *testing.T) {
  _, index := initTest(t)
  _ = index.getSettings()
  //TODO check return
}

func TestSetSettings(t *testing.T) {
  _, index := initTest(t)
  settings := make(map[string]interface{})
  _ = index.setSettings(settings)
  //TODO check return
}

func TestGetLogs(t *testing.T) {
  client, _ := initTest(t)
  logs := client.getLogs(0, 100, false)
  if _, ok := logs.(map[string]interface{})["logs"]; !ok {
    t.Fatalf("Unable to get logs")
  }
}

func TestBrowse(t *testing.T) {
  _, index := initTest(t)
  items := index.browse(1, 1)
  if _, ok := items.(map[string]interface{})["hits"]; !ok {
    t.Fatalf("Unable to browse")
  }
}

func TestQuery(t *testing.T) {
  _, index := initTest(t)
  object := make(map[string]interface{})
  object["name"] = "John Snow"
  object["objectID"] = "algolia"
  _ = index.addObject(object)
  query := make(map[string]interface{})
  query["query"] = ""
  _ = index.query(query)
  //TODO
}

func TestCopy(t *testing.T) {
  _, index := initTest(t)
  _ = index.copy("àlgo?à2-go")
  //TODO
}

//func TestMove(t *testing.T) {
//  _, index := initTest(t)
//  _ = index.move("àlgo?à2-go")
  //TODO
//}
