package algoliasearch

import (
	"syscall"
	"testing"
	"time"
)

func safeName(name string) string {
	travis, haveTravis := syscall.Getenv("TRAVIS")
	buildID, haveBuildID := syscall.Getenv("TRAVIS_JOB_NUMBER")
	if !haveTravis || !haveBuildID || travis != "true" {
		return name
	}

	return name + "_travis-" + buildID
}

func initTest(t *testing.T) (Client, Index) {
	appID, haveAppID := syscall.Getenv("ALGOLIA_APPLICATION_ID")
	apiKey, haveAPIKey := syscall.Getenv("ALGOLIA_API_KEY")
	if !haveAPIKey || !haveAppID {
		t.Fatal("Need ALGOLIA_APPLICATION_ID and ALGOLIA_API_KEY")
	}
	client := NewClient(appID, apiKey)
	client.SetTimeout(1000, 10000)
	hosts := make([]string, 3)
	hosts[0] = appID + "-1.algolia.net"
	hosts[1] = appID + "-2.algolia.net"
	hosts[2] = appID + "-3.algolia.net"
	client = NewClientWithHosts(appID, apiKey, hosts)
	index := client.InitIndex(safeName("àlgol?à-go"))

	return client, index
}

func tearDownTest(index Index, t *testing.T) {
	del, err := index.Delete()
	if err != nil {
		t.Fatal(err.Error())
	}

	err = index.WaitTask(del.TaskID)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func addWait(obj Object, index Index, t *testing.T) {
	add, err := index.AddObject(obj)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = index.WaitTask(add.TaskID)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func hasString(slice []string, elt string) bool {
	for _, e := range slice {
		if e == elt {
			return true
		}
	}
	return false
}

func checkNbHits(got, expected int64, t *testing.T) {
	checkEqual(got, expected, "nbHits", t)
}

func checkEqual(got, expected interface{}, message string, t *testing.T) {
	if got != expected {
		t.Fatalf("%s: %v expected %v", message, got, expected)
	}
}

func TestClear(t *testing.T) {
	_, index := initTest(t)
	object := Object{"name": "John snow"}

	addWait(object, index, t)

	clear, err := index.Clear()
	if err != nil {
		t.Fatal(err.Error())
	}

	err = index.WaitTask(clear.TaskID)
	if err != nil {
		t.Fatal(err.Error())
	}

	search, err := index.Search("", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	checkNbHits(search.NbHits, 0, t)
	tearDownTest(index, t)
}

func TestAddObject(t *testing.T) {
	_, index := initTest(t)
	object := Object{"name": "John Snow"}

	addWait(object, index, t)

	object = Object{"name": "John Snow", "objectID": "àlgol?à"}

	add, err := index.AddObject(object)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = index.WaitTask(add.TaskID)
	if err != nil {
		t.Fatal(err.Error())
	}

	search, err := index.Search("", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	checkNbHits(search.NbHits, 2, t)

	tearDownTest(index, t)
}

func TestUpdateObject(t *testing.T) {
	_, index := initTest(t)
	object := Object{"name": "John Snow", "objectID": "àlgol?à"}

	addWait(object, index, t)

	object["name"] = "Roger"

	update, err := index.UpdateObject(object)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = index.WaitTask(update.TaskID)
	if err != nil {
		t.Fatal(err.Error())
	}

	search, err := index.Search("", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	checkEqual(search.Hits[0]["name"], object["name"], "name", t)

	tearDownTest(index, t)
}

func TestPartialUpdateObject(t *testing.T) {
	_, index := initTest(t)
	object := Object{"name": "John Snow", "objectID": "àlgol?à", "job": "Knight"}

	addWait(object, index, t)

	delete(object, "job")
	object["name"] = "Roger"

	update, err := index.PartialUpdateObject(object)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = index.WaitTask(update.TaskID)
	if err != nil {
		t.Fatal(err.Error())
	}

	search, err := index.Search("", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	checkEqual(search.Hits[0]["name"], object["name"], "name", t)

	_, ok := search.Hits[0]["job"]
	checkEqual(ok, false, "job presence", t)

	tearDownTest(index, t)
}

func TestGetObject(t *testing.T) {
	_, index := initTest(t)
	object := Object{"name": "John Snow", "objectID": "àlgol?à"}

	addWait(object, index, t)

	obj, err := index.GetObject("àlgol?à", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	if obj["name"] != object["name"] {
		t.Fatal("Unable to get a single object")
	}

	tearDownTest(index, t)
}

func TestGetObjects(t *testing.T) {
	_, index := initTest(t)
	la := Object{"name": "Los Angeles", "objectID": "1"}

	addWait(la, index, t)

	sf := Object{"name": "San Francisco", "objectID": "2"}

	addWait(sf, index, t)

	objects, err := index.GetObjects([]string{"1", "2"})
	if err != nil {
		t.Fatal(err.Error())
	}

	checkEqual(objects[0]["name"], la["name"], "name", t)
	checkEqual(objects[1]["name"], sf["name"], "name", t)

	tearDownTest(index, t)
}

func TestDeleteObject(t *testing.T) {
	_, index := initTest(t)
	object := Object{"name": "John Snow", "objectID": "àlgol?à"}

	addWait(object, index, t)

	del, err := index.DeleteObject("àlgol?à")
	if err != nil {
		t.Fatal(err.Error())
	}

	err = index.WaitTask(del.TaskID)
	if err != nil {
		t.Fatal(err.Error())
	}

	search, err := index.Search("", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	checkNbHits(search.NbHits, 0, t)

	tearDownTest(index, t)
}

func TestSetSettings(t *testing.T) {
	_, index := initTest(t)

	settings := map[string]interface{}{"hitsPerPage": 30}

	set, err := index.SetSettings(settings)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = index.WaitTask(set.TaskID)
	if err != nil {
		t.Fatal(err.Error())
	}

	get, err := index.GetSettings()
	if err != nil {
		t.Fatal(err.Error())
	}

	checkEqual(get.HitsPerPage, settings["hitsPerPage"], "hitsPerPage", t)

	tearDownTest(index, t)
}

func TestBrowse(t *testing.T) {
	_, index := initTest(t)
	object := Object{"name": "John Snow", "objectID": "àlgol?à"}

	addWait(object, index, t)

	browse, err := index.Browse(map[string]interface{}{"page": 1, "hitsPerPage": 1})
	if err != nil {
		t.Fatal(err.Error())
	}

	tearDownTest(index, t)
}

func TestBrowseWithCursor(t *testing.T) {
	_, index := initTest(t)
	object := Object{"name": "John Snow", "objectID": "àlgol?à"}

	addWait(object, index, t)

	browse, err := index.BrowseAll(map[string]interface{}{"query": ""})
	if err != nil {
		t.Fatal(err.Error())
	}

	hit, err := browse.Next()
	if err != nil {
		t.Fatal(err.Error())
	}

	checkEqual(hit["name"], "John Snow", "name", t)

	_, err = browse.Next()
	if err == nil {
		t.Fatal("Should contains only one element")
	}

	tearDownTest(index, t)
}

func TestQuery(t *testing.T) {
	_, index := initTest(t)
	object := Object{"name": "John Snow", "objectID": "àlgol?à"}

	addWait(object, index, t)

	params := map[string]interface{}{"attributesToRetrieve": "*", "getRankingInfo": 1}
	search, err := index.Search("", params)
	if err != nil {
		t.Fatal(err.Error())
	}

	checkNbHits(search.NbHits, 1, t)

	tearDownTest(index, t)
}

func TestIndexCopy(t *testing.T) {
	client, index := initTest(t)
	object := Object{"name": "John Snow", "objectID": "àlgol?à"}

	addWait(object, index, t)

	idx, err := index.Copy(safeName("àlgo?à2-go"))
	if err != nil {
		t.Fatal(err.Error())
	}

	err = index.WaitTask(idx.TaskID)
	if err != nil {
		t.Fatal(err.Error())
	}

	indexCopy := client.InitIndex(safeName("àlgo?à2-go"))

	search, err := indexCopy.Search("", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	checkNbHits(search.NbHits, 1, t)

	tearDownTest(index, t)
	tearDownTest(indexCopy, t)
}

func TestCopy(t *testing.T) {
	client, index := initTest(t)
	object := Object{"name": "John Snow", "objectID": "àlgol?à"}

	addWait(object, index, t)

	copy, err := client.CopyIndex(safeName("àlgol?à-go"), safeName("àlgo?à2-go"))
	if err != nil {
		t.Fatal(err.Error())
	}

	err = index.WaitTask(copy.TaskID)
	if err != nil {
		t.Fatal(err.Error())
	}

	indexCopy := client.InitIndex(safeName("àlgo?à2-go"))

	search, err := indexCopy.Search("", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	checkNbHits(search.NbHits, 1, t)

	tearDownTest(index, t)
	tearDownTest(indexCopy, t)
}

func TestIndexMove(t *testing.T) {
	client, index := initTest(t)
	object := Object{"name": "John Snow", "objectID": "àlgol?à"}

	addWait(object, index, t)

	move, err := index.Move(safeName("àlgo?à2-go"))
	if err != nil {
		t.Fatal(err.Error())
	}

	err = index.WaitTask(move.TaskID)
	if err != nil {
		t.Fatal(err.Error())
	}

	indexMove := client.InitIndex(safeName("àlgo?à2-go"))

	search, err := indexMove.Search("", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	checkNbHits(search.NbHits, 1, t)

	tearDownTest(indexMove, t)
}

func TestMove(t *testing.T) {
	client, index := initTest(t)
	object := Object{"name": "John Snow", "objectID": "àlgol?à"}

	addWait(object, index, t)

	move, err := client.MoveIndex(safeName("àlgol?à-go"), safeName("àlgo?à2-go"))
	if err != nil {
		t.Fatal(err.Error())
	}

	err = index.WaitTask(move.TaskID)
	if err != nil {
		t.Fatal(err.Error())
	}

	indexMove := client.InitIndex(safeName("àlgo?à2-go"))

	search, err := indexMove.Search("", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	checkNbHits(search.NbHits, 1, t)

	tearDownTest(indexMove, t)
}

func TestAddIndexKey(t *testing.T) {
	_, index := initTest(t)

	newKey := Key{
		ACL:                    []string{"search"},
		Validity:               300,
		MaxQueriesPerIPPerHour: 100,
		MaxHitsPerQuery:        100,
	}

	acl := []string{"search"}
	add, err := index.AddKey(newKey)
	if err != nil {
		t.Fatal(err.Error())
	}

	// No taskID for AddKey, emulate WaitTask.
	time.Sleep(5 * time.Second)

	get, err := index.GetKey(add.Key)
	if err != nil {
		t.Fatal(err.Error())
	}

	if get.Value != add.Key {
		t.Fatal("Unable to get a key")
	}

	list, err := index.ListKeys()
	if err != nil {
		t.Fatal(err.Error())
	}

	keys := make([]string, len(list))
	for i, k := range list {
		keys[i] = k.Value
	}

	if !hasString(keys, add.Key) {
		t.Fatalf("%s should be present", add.Key)
	}

	updated := Key{ACL: []string{"addObject"}, Value: add.Key}
	_, err = index.UpdateKey(updated)
	if err != nil {
		t.Fatal(err.Error())
	}

	// No taskID for UpdateKey, emulate WaitTask.
	time.Sleep(5 * time.Second)

	list, err = index.ListKeys()
	if err != nil {
		t.Fatal(err.Error())
	}

	keys = make([]string, len(list))
	for i, k := range list {
		keys[i] = k.Value
	}

	if !hasString(keys, add.Key) {
		t.Fatalf("%s should be present", add.Key)
	}

	_, err = index.DeleteKey(add.Key)
	if err != nil {
		t.Fatal(err.Error())
	}

	// No taskID for DeleteKey, emulate WaitTask.
	time.Sleep(5 * time.Second)

	list, err = index.ListKeys()
	if err != nil {
		t.Fatal(err.Error())
	}

	keys = make([]string, len(list))
	for i, k := range list {
		keys[i] = k.Value
	}

	if hasString(keys, add.Key) {
		t.Fatalf("%s should not be present", add.Key)
	}

	tearDownTest(index, t)
}

func TestAddKey(t *testing.T) {
	client, index := initTest(t)

	acl := []string{"search"}
	params := map[string]interface{}{
		"Validity":               300,
		"MaxHitsPerQuery":        100,
		"MaxQueriesPerIPPerHour": 100,
		"Indexes":                []string{index.name},
	}

	add, err := client.AddKey(acl, params)
	if err != nil {
		t.Fatal(err.Error())
	}

	// No taskID for AddKey, emulate WaitTask.
	time.Sleep(5 * time.Second)

	get, err := client.GetKey(add.Key)
	if err != nil {
		t.Fatal(err.Error())
	}

	if get.Value != add.Key {
		t.Fatal("Unable to get a key")
	}

	_, err = client.UpdateKey(add.Key, map[string]interface{}{"acl": "addObject"})
	if err != nil {
		t.Fatal(err.Error())
	}

	// No taskID for UpdateKey, emulate WaitTask.
	time.Sleep(5 * time.Second)

	list, err := client.ListKeys()
	if err != nil {
		t.Fatal(err.Error())
	}

	keys := make([]string, len(list))
	for i, k := range list {
		keys[i] = k.Value
	}

	if !hasString(keys, add.Key) {
		t.Fatal("%s should be present", add.Key)
	}

	_, err = client.DeleteKey(add.Key)
	if err != nil {
		t.Fatal(err.Error())
	}

	// No taskID for DeleteKey, emulate WaitTask.
	time.Sleep(5 * time.Second)

	list, err = client.ListKeys()
	if err != nil {
		t.Fatal(err.Error())
	}

	keys = make([]string, len(list))
	for i, k := range list {
		keys[i] = k.Value
	}

	if hasString(keys, add.Key) {
		t.Fatalf("%s should not be present", add.Key)
	}

	tearDownTest(index, t)
}

func TestAddObjects(t *testing.T) {
	_, index := initTest(t)
	objects := []Object{
		Object{"name": "John", "city": "San Francisco"},
		Object{"name": "Roger", "city": "New York"},
	}

	add, err := index.AddObjects(objects)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = index.WaitTask(add.TaskID)
	if err != nil {
		t.Fatal(err.Error())
	}

	search, err := index.Search("", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	checkNbHits(search.NbHits, 2, t)

	tearDownTest(index, t)
}

func TestUpdateObjects(t *testing.T) {
	_, index := initTest(t)
	objects := []Object{
		Object{"name": "John", "city": "San Francisco", "objectID": "àlgo?à-1"},
		Object{"name": "Roger", "city": "New York", "objectID": "àlgo?à-2"},
	}

	update, err := index.UpdateObjects(objects)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = index.WaitTask(update.TaskID)
	if err != nil {
		t.Fatal(err.Error())
	}

	search, err := index.Search("", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	checkNbHits(search.NbHits, 2, t)

	tearDownTest(index, t)
}

func TestPartialUpdateObjects(t *testing.T) {
	_, index := initTest(t)
	objects := []Object{
		Object{"name": "John", "city": "San Francisco", "objectID": "àlgo?à-1"},
		Object{"name": "Roger", "city": "New York", "objectID": "àlgo?à-2"},
	}

	update, err := index.PartialUpdateObjects(objects)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = index.WaitTask(update.TaskID)
	if err != nil {
		t.Fatal(err.Error())
	}

	search, err := index.Search("", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	checkNbHits(search.NbHits, 2, t)

	tearDownTest(index, t)
}

func TestDeleteObjects(t *testing.T) {
	_, index := initTest(t)
	objects := []Object{
		Object{"name": "John", "city": "San Francisco", "objectID": "àlgo?à-1"},
		Object{"name": "Roger", "city": "New York", "objectID": "àlgo?à-2"},
	}

	update, err := index.PartialUpdateObjects(objects)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = index.WaitTask(update.TaskID)
	if err != nil {
		t.Fatal(err.Error())
	}

	objectIDs := []string{"àlgo?à-1", "àlgo?à-2"}

	del, err := index.DeleteObjects(objectIDs)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = index.WaitTask(del.TaskID)
	if err != nil {
		t.Fatal(err.Error())
	}

	search, err := index.Search("", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	checkNbHits(search.NbHits, 0, t)

	tearDownTest(index, t)
}

func TestDeleteByQuery(t *testing.T) {
	_, index := initTest(t)
	objects := []Object{
		Object{"name": "San Jose"},
		Object{"name": "Washington"},
		Object{"name": "San Francisco"},
	}

	add, err := index.AddObjects(objects)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = index.WaitTask(add.TaskID)
	if err != nil {
		t.Fatal(err.Error())
	}

	del, err := index.DeleteByQuery("San", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = index.WaitTask(del.TaskID)
	if err != nil {
		t.Fatal(err.Error())
	}

	search, err := index.Search("", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	checkNbHits(search.NbHits, 1, t)

	tearDownTest(index, t)
}

func TestGenerateNewSecuredApiKey(t *testing.T) {
	client, index := initTest(t)
	base := "182634d8894831d5dbce3b3185c50881"

	key, _ := client.GenerateSecuredAPIKey(base, map[string]interface{}{"tagFilters": "(public,user1)"})
	expected := "MDZkNWNjNDY4M2MzMDA0NmUyNmNkZjY5OTMzYjVlNmVlMTk1NTEwMGNmNTVjZmJhMmIwOTIzYjdjMTk2NTFiMnRhZ0ZpbHRlcnM9JTI4cHVibGljJTJDdXNlcjElMjk="
	checkEqual(key, expected, "secured key", t)

	key, _ = client.GenerateSecuredAPIKey(base, map[string]interface{}{"tagFilters": "(public,user1)", "userToken": "42"})
	expected = "OGYwN2NlNTdlOGM2ZmM4MjA5NGM0ZmYwNTk3MDBkNzMzZjQ0MDI3MWZjNTNjM2Y3YTAzMWM4NTBkMzRiNTM5YnRhZ0ZpbHRlcnM9JTI4cHVibGljJTJDdXNlcjElMjkmdXNlclRva2VuPTQy"
	checkEqual(key, expected, "secured key", t)

	tearDownTest(index, t)
}

func TestMultipleQueries(t *testing.T) {
	client, index := initTest(t)
	object := Object{"name": "John Snow"}

	addWait(object, index, t)

	queries := []map[string]interface{}{
		map[string]interface{}{"indexName": safeName("àlgol?à-go"), "query": ""},
	}

	search, err := client.MultipleQueries(queries, "", "")
	if err != nil {
		t.Fatal(err.Error())
	}

	checkNbHits(search[0].NbHits, 1, t)

	tearDownTest(index, t)
}

func TestFacets(t *testing.T) {
	_, index := initTest(t)

	settings := map[string]interface{}{"attributesForFacetting": []string{"f", "g"}}
	set, err := index.SetSettings(settings)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = index.WaitTask(set.TaskID)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = index.AddObject(Object{"f": "f1", "g": "g1"})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = index.AddObject(Object{"f": "f1", "g": "g2"})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = index.AddObject(Object{"f": "f2", "g": "g2"})
	if err != nil {
		t.Fatal(err.Error())
	}

	add, err := index.AddObject(Object{"f": "f3", "g": "g2"})
	if err != nil {
		t.Fatal(err.Error())
	}

	err = index.WaitTask(add.TaskID)
	if err != nil {
		t.Fatal(err.Error())
	}

	search, err := index.Search("", map[string]interface{}{"facets": "f", "facetFilters": []string{"f:f1"}})
	if err != nil {
		t.Fatal(err.Error())
	}

	checkNbHits(search.NbHits, 2, t)

	tearDownTest(index, t)
}

func TestSynonyms(t *testing.T) {
	_, index := initTest(t)

	object := Object{"name": "589 Howard St., San Francisco"}
	addWait(object, index, t)

	batch, err := index.BatchSynonyms([]Synonym{
		map[string]interface{}{
			"objectID": "city", "type": "synonym",
			"synonyms": []string{"San Francisco", "SF"}},
		map[string]interface{}{
			"objectID": "street", "type": "altCorrection1",
			"word": "Street", "corrections": []string{"St"}},
	}, false, false)

	if err != nil {
		t.Fatal(err.Error())
	}

	err = index.WaitTask(batch.TaskID)
	if err != nil {
		t.Fatal(err.Error())
	}

	get, err := index.GetSynonym("city")
	if err != nil {
		t.Fatal(err.Error())
	}

	checkEqual(get.(map[string]interface{})["objectID"], "city", "city", t)

	search, err := index.Search("Howard Street SF", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	checkNbHits(search.NbHits, 1, t)

	del, err := index.DeleteSynonym("street", false)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = index.WaitTask(del.TaskID)
	if err != nil {
		t.Fatal(err.Error())
	}

	synonyms, err := index.SearchSynonyms("", []string{"synonym"}, 0, 5)
	if err != nil {
		t.Fatal(err.Error())
	}

	checkNbHits(synonyms.NbHits(), 1, t)

	clear, err := index.ClearSynonyms(false)
	if err != nil {
		t.Fatal(err.Error())
	}
	index.WaitTask(clear.TaskID)

	synonyms, err = index.SearchSynonyms("", []string{}, 0, 5)
	if err != nil {
		t.Fatal(err.Error())
	}

	checkNbHits(synonyms.NbHits(), 0, t)

	tearDownTest(index, t)
}
