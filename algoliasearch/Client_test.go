package algoliasearch

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
  index := client.InitIndex("àlgol?à-go")
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

func shouldContainString(json interface{}, attr string, value string, msg string, t *testing.T) {
  array := json.([]interface{})
  for i := range array {
    val, ok := array[i].(map[string]interface{})[attr]
    if ok && value == val.(string) {
      return
    }
  }
  t.Fatalf(msg + ", expected: " + value + " in the array.")
}

func shouldNotContainString(json interface{}, attr string, value string, msg string, t *testing.T) {
  array := json.([]interface{})
  for i := range array {
    val, ok := array[i].(map[string]interface{})[attr]
    if ok && value == val.(string) {
      t.Fatalf(msg + ", expected: " + value + " in the array.")
    }
  }
}

func TestClear(t *testing.T) {
  _, index := initTest(t)
  object := make(map[string]interface{})
  object["name"] = "John Snow"
  resp, err := index.AddObject(object)
  if err != nil {
    t.Fatalf(err.Error())
  }
  index.WaitTask(resp)
  resp, err = index.AddObject(object)
  if err != nil {
    t.Fatalf(err.Error())
  }
  time.Sleep(time.Duration(100) * time.Millisecond) 
  resp, err = index.Clear()
  if err != nil {
    t.Fatalf(err.Error())
  }
  index.WaitTask(resp)
  results, err := index.Query(make(map[string]interface{}))
  if err != nil {
    t.Fatalf(err.Error())
  }
  shouldFloat(results, "nbHits", 0, "Unable to clear the index", t)
  index.Delete()
}

func TestAddObject(t *testing.T) {
  _, index := initTest(t)
  object := make(map[string]interface{})
  object["name"] = "John Snow"
  _, err := index.AddObject(object)
  if err != nil {
    t.Fatalf(err.Error())
  }
  object["name"] = "John Snow"
  object["objectID"] = "àlgol?à"
  _, err = index.AddObject(object)
  if err != nil {
    t.Fatalf(err.Error())
  }
  results, err := index.Query(make(map[string]interface{}))
  if err != nil {
    t.Fatalf(err.Error())
  }
  shouldFloat(results, "nbHits", 2, "Unable to clear the index", t)
  index.Delete()
}

func TestUpdateObject(t *testing.T) {
  _, index := initTest(t)
  object := make(map[string]interface{})
  object["name"] = "John Snow"
  object["objectID"] = "àlgol?à"
  _, err := index.AddObject(object)
  if err != nil {
    t.Fatalf(err.Error())
  }
  object["name"] = "Roger"
  _, err = index.UpdateObject(object)
  if err != nil {
    t.Fatalf(err.Error())
  }
  results, err := index.Query(make(map[string]interface{}))
  if err != nil {
    t.Fatalf(err.Error())
  }
  hits := results.(map[string]interface{})["hits"]
  shouldStr(hits.([]interface{})[0], "name", "Roger", "Unable to update an object", t)
  shouldNotHave(hits.([]interface{})[0], "job", "Unable to update an object", t)
  index.Delete()
}

func TestPartialUpdateObject(t *testing.T) {
  _, index := initTest(t)
  object := make(map[string]interface{})
  object["name"] = "John Snow"
  object["job"] = "Knight"
  object["objectID"] = "àlgol?à"
  _, err := index.AddObject(object)
  if err != nil {
    t.Fatalf(err.Error())
  }
  delete(object, "job")
  object["name"] = "Roger"
  _, err = index.PartialUpdateObject(object)
  if err != nil {
    t.Fatalf(err.Error())
  }
  results, err := index.Query(make(map[string]interface{}))
  if err != nil {
    t.Fatalf(err.Error())
  }
  hits := results.(map[string]interface{})["hits"]
  shouldStr(hits.([]interface{})[0], "name", "Roger", "Unable to update an object", t)
  index.Delete()
}


func TestGetObject(t *testing.T) {
  _, index := initTest(t)
  object := make(map[string]interface{})
  object["name"] = "John Snow"
  object["objectID"] = "àlgol?à"
  resp, err := index.AddObject(object)
  if err != nil {
    t.Fatalf(err.Error())
  }
  _, err = index.WaitTask(resp)
  if err != nil {
    t.Fatalf(err.Error())
  }
  resp, err = index.AddObject(object)
  if err != nil {
    t.Fatalf(err.Error())
  }
  obj, err := index.GetObject("àlgol?à")
  if err != nil {
    t.Fatalf(err.Error())
  }
  shouldStr(obj, "name", "John Snow", "Unable to update an object", t)
  index.Delete()
}

func TestDeleteObject(t *testing.T) {
  _, index := initTest(t)
  object := make(map[string]interface{})
  object["name"] = "John Snow"
  object["objectID"] = "àlgol?à"
  resp, err := index.AddObject(object)
  if err != nil {
    t.Fatalf(err.Error())
  }
  _, err = index.WaitTask(resp)
  if err != nil {
    t.Fatalf(err.Error())
  }
  resp, err = index.DeleteObject("àlgol?à")
  if err != nil {
    t.Fatalf(err.Error())
  }
  _, err = index.WaitTask(resp)
  if err != nil {
    t.Fatalf(err.Error())
  }
  results, err := index.Query(make(map[string]interface{}))
  if err != nil {
    t.Fatalf(err.Error())
  }
  shouldFloat(results, "nbHits", 0, "Unable to clear the index", t)
  index.Delete()
}

func TestSetSettings(t *testing.T) {
  _, index := initTest(t)
  settings := make(map[string]interface{})
  settings["hitsPerPage"] = 30
  resp, err := index.SetSettings(settings)
  if err != nil {
    t.Fatalf(err.Error())
  }
  _, err = index.WaitTask(resp)
  if err != nil {
    t.Fatalf(err.Error())
  }
  settingsChanged, err := index.GetSettings()
  if err != nil {
    t.Fatalf(err.Error())
  }
  shouldFloat(settingsChanged, "hitsPerPage", 30, "Unable to change setting", t)
  index.Delete()
}

func TestGetLogs(t *testing.T) {
  client, _ := initTest(t)
  logs, err := client.GetLogs(0, 100, false)
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
  resp, err := index.AddObject(object)
  if err != nil {
    t.Fatalf(err.Error())
  }
  _, err = index.WaitTask(resp)
  if err != nil {
    t.Fatalf(err.Error())
  }
  items, err := index.Browse(1, 1)
  if err != nil {
    t.Fatalf(err.Error())
  }
  shouldHave(items, "hits", "Unable to browse index", t)
  index.Delete()
}

func TestQuery(t *testing.T) {
  _, index := initTest(t)
  object := make(map[string]interface{})
  object["name"] = "John Snow"
  object["objectID"] = "àlgol?à"
  _, err := index.AddObject(object)
  if err != nil {
    t.Fatalf(err.Error())
  }
  query := make(map[string]interface{})
  query["query"] = ""
  query["attributesToRetrieve"] = "*"
  query["getRankingInfo"] = 1
  results, err := index.Query(query)
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
  _, err := index.AddObject(object)
  if err != nil {
    t.Fatalf(err.Error())
  }
  resp, err := index.Copy("àlgo?à2-go")
  if err != nil {
    t.Fatalf(err.Error())
  }
  _, err = index.WaitTask(resp)
  if err != nil {
    t.Fatalf(err.Error())
  }
  indexCopy := client.InitIndex("àlgo?à2-go")
  results, err := indexCopy.Query(make(map[string]interface{}))
  if err != nil {
    t.Fatalf(err.Error())
  }
  shouldFloat(results, "nbHits", 1, "Unable to copy an index", t)
  index.Delete()
  indexCopy.Delete()
}

func TestMove(t *testing.T) {
  client, index := initTest(t)
  object := make(map[string]interface{})
  object["name"] = "John Snow"
  object["objectID"] = "àlgol?à"
  _, err := index.AddObject(object)
  if err != nil {
    t.Fatalf(err.Error())
  }
  resp, err := index.Move("àlgo?à2-go")
  if err != nil {
    t.Fatalf(err.Error())
  }
  _, err = index.WaitTask(resp)
  if err != nil {
    t.Fatalf(err.Error())
  }
  indexMove := client.InitIndex("àlgo?à2-go")
  results, err := indexMove.Query(make(map[string]interface{}))
  if err != nil {
    t.Fatalf(err.Error())
  }
  shouldFloat(results, "nbHits", 1, "Unable to move an index", t)
  indexMove.Delete()
}

func TestAddIndexKey(t *testing.T) {
  _, index := initTest(t)
  acl := []string{"search"}
  newKey, err := index.AddKey(acl, 300, 100, 100)
  if err != nil {
    t.Fatalf(err.Error())
  }
  key, err := index.GetKey(newKey.(map[string]interface{})["key"].(string))
  if err != nil {
    t.Fatalf(err.Error())
  }
  shouldStr(key, "value", newKey.(map[string]interface{})["key"].(string), "Unable to get a key", t)
  list, err := index.ListKeys()
  if err != nil {
    t.Fatalf(err.Error())
  }
  shouldContainString(list.(map[string]interface{})["keys"], "value", newKey.(map[string]interface{})["key"].(string), "Unable to add a key", t)
  _, err = index.DeleteKey(newKey.(map[string]interface{})["key"].(string))
  if err != nil {
    t.Fatalf(err.Error())
  }
  list, err = index.ListKeys()
  if err != nil {
    t.Fatalf(err.Error())
  }
  shouldNotContainString(list.(map[string]interface{})["keys"], "value", newKey.(map[string]interface{})["key"].(string), "Unable to add a key", t)
  index.Delete() 
}

func TestAddKey(t *testing.T) {
  client, index := initTest(t)
  acl := []string{"search"}
  indexes := []string{index.name}
  newKey, err := client.AddKey(acl, indexes, 300, 100, 100)
  if err != nil {
    t.Fatalf(err.Error())
  }
  key, err := client.GetKey(newKey.(map[string]interface{})["key"].(string))
  if err != nil {
    t.Fatalf(err.Error())
  }
  shouldStr(key, "value", newKey.(map[string]interface{})["key"].(string), "Unable to get a key", t)
  list, err := client.ListKeys()
  if err != nil {
    t.Fatalf(err.Error())
  }
  shouldContainString(list.(map[string]interface{})["keys"], "value", newKey.(map[string]interface{})["key"].(string), "Unable to add a key", t)
  _, err = client.DeleteKey(newKey.(map[string]interface{})["key"].(string))
  if err != nil {
    t.Fatalf(err.Error())
  }
  list, err = client.ListKeys()
  if err != nil {
    t.Fatalf(err.Error())
  }
  shouldNotContainString(list.(map[string]interface{})["keys"], "value", newKey.(map[string]interface{})["key"].(string), "Unable to add a key", t)
}

func TestAddObjects(t *testing.T) {
  _, index := initTest(t)
  objects := make([]interface{}, 2)

  object := make(map[string]interface{})
  object["name"] = "John"
  object["city"] = "San Francisco"
  objects[0] = object

  object = make(map[string]interface{})
  object["name"] = "Roger"
  object["city"] = "New York"
  objects[1] = object
  task, err := index.AddObjects(objects)
  if err != nil {
    t.Fatalf(err.Error())
  }
  index.WaitTask(task)
  results, err := index.Query(make(map[string]interface{}))
  if err != nil {
    t.Fatalf(err.Error())
  }
  shouldFloat(results, "nbHits", 2, "Unable to add objects", t)
  index.Delete()
}

func TestUpdateObjects(t *testing.T) {
  _, index := initTest(t)
  objects := make([]interface{}, 2)

  object := make(map[string]interface{})
  object["name"] = "John"
  object["city"] = "San Francisco"
  object["objectID"] = "àlgo?à-1"
  objects[0] = object

  object = make(map[string]interface{})
  object["name"] = "Roger"
  object["city"] = "New York"
  object["objectID"] = "àlgo?à-2"
  objects[1] = object
  task, err := index.UpdateObjects(objects)
  if err != nil {
    t.Fatalf(err.Error())
  }
  index.WaitTask(task)
  results, err := index.Query(make(map[string]interface{}))
  if err != nil {
    t.Fatalf(err.Error())
  }
  shouldFloat(results, "nbHits", 2, "Unable to update objects", t)
  index.Delete()
}

func TestPartialUpdateObjects(t *testing.T) {
  _, index := initTest(t)
  objects := make([]interface{}, 2)

  object := make(map[string]interface{})
  object["name"] = "John"
  object["objectID"] = "àlgo?à-1"
  objects[0] = object

  object = make(map[string]interface{})
  object["name"] = "Roger"
  object["objectID"] = "àlgo?à-2"
  objects[1] = object
  task, err := index.PartialUpdateObjects(objects)
  if err != nil {
    t.Fatalf(err.Error())
  }
  index.WaitTask(task)
  results, err := index.Query(make(map[string]interface{}))
  if err != nil {
    t.Fatalf(err.Error())
  }
  shouldFloat(results, "nbHits", 2, "Unable to partial update objects", t)
  index.Delete()
}

/*
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
}*/
