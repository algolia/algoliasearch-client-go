package algoliasearch

import (
	"fmt"
	"math/rand"
	"syscall"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func fillRandom(random []byte) {
	letters := "abcdefghijklmnopqrstuvwxyzBCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := range random {
		random[i] = letters[rand.Int()%len(letters)]
	}
}

func genIndexName(ignore string) string {
	baseName := "àlgol?à-go"

	random := make([]byte, 8)
	fillRandom(random)
	for ignore == string(random) {
		fillRandom(random)
	}

	return safeName(fmt.Sprintf("%s-%s", string(random), baseName))
}

func safeName(name string) string {
	travis, haveTravis := syscall.Getenv("TRAVIS")
	buildID, haveBuildID := syscall.Getenv("TRAVIS_JOB_NUMBER")
	if !haveTravis || !haveBuildID || travis != "true" {
		return name
	}

	return name + "_travis-" + buildID
}

func initTest(t *testing.T) (*client, *index) {
	appID, haveAppID := syscall.Getenv("ALGOLIA_APPLICATION_ID")
	apiKey, haveAPIKey := syscall.Getenv("ALGOLIA_API_KEY")
	if !haveAPIKey || !haveAppID {
		t.Fatal("Need ALGOLIA_APPLICATION_ID and ALGOLIA_API_KEY")
	}
	hosts := make([]string, 3)
	hosts[0] = appID + "-1.algolia.net"
	hosts[1] = appID + "-2.algolia.net"
	hosts[2] = appID + "-3.algolia.net"
	c := NewClientWithHosts(appID, apiKey, hosts).(*client)
	i := c.InitIndex(genIndexName("")).(*index)

	return c, i
}

func tearDownTest(t *testing.T, index Index) {
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

func waitIndexKey(index Index, key string, f func(Key) bool) {
	for i := 0; i < 60; i++ {
		k, err := index.GetUserKey(key)
		if err == nil && (f == nil || f(k)) {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func waitMissingIndexKey(index Index, key string) {
	for i := 0; i < 60; i++ {
		_, err := index.GetUserKey(key)
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func waitClientKey(client Client, key string, f func(Key) bool) {
	for i := 0; i < 60; i++ {
		k, err := client.GetUserKey(key)
		if err == nil && (f == nil || f(k)) {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func waitMissingClientKey(client Client, key string) {
	for i := 0; i < 60; i++ {
		_, err := client.GetUserKey(key)
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func checkNbHits(t *testing.T, got, expected int) {
	checkEqual(t, got, expected, "nbHits")
}

func checkEqual(t *testing.T, got, expected interface{}, message string) {
	if got != expected {
		t.Fatalf("%s: %v expected %v", message, got, expected)
	}
}

func TestClear(t *testing.T) {
	_, index := initTest(t)
	defer tearDownTest(t, index)
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

	checkNbHits(t, search.NbHits, 0)
}

func TestAddObject(t *testing.T) {
	_, index := initTest(t)
	defer tearDownTest(t, index)
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

	checkNbHits(t, search.NbHits, 2)
}

func TestUpdateObject(t *testing.T) {
	_, index := initTest(t)
	defer tearDownTest(t, index)
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

	checkEqual(t, search.Hits[0]["name"], object["name"], "name")
}

func TestPartialUpdateObject(t *testing.T) {
	_, index := initTest(t)
	defer tearDownTest(t, index)
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

	checkEqual(t, search.Hits[0]["name"], object["name"], "name")

	_, ok := search.Hits[0]["job"]
	checkEqual(t, ok, true, "job presence")
}

func TestGetObject(t *testing.T) {
	_, index := initTest(t)
	defer tearDownTest(t, index)
	object := Object{"name": "John Snow", "objectID": "àlgol?à"}

	addWait(object, index, t)

	obj, err := index.GetObject("àlgol?à", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	checkEqual(t, obj["name"], object["name"], "name")
}

func TestGetObjects(t *testing.T) {
	_, index := initTest(t)
	defer tearDownTest(t, index)
	la := Object{"name": "Los Angeles", "objectID": "1"}

	addWait(la, index, t)

	sf := Object{"name": "San Francisco", "objectID": "2"}

	addWait(sf, index, t)

	objects, err := index.GetObjects([]string{"1", "2"})
	if err != nil {
		t.Fatal(err.Error())
	}

	checkEqual(t, objects[0]["name"], la["name"], "name")
	checkEqual(t, objects[1]["name"], sf["name"], "name")
}

func TestDeleteObject(t *testing.T) {
	_, index := initTest(t)
	defer tearDownTest(t, index)
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

	checkNbHits(t, search.NbHits, 0)
}

func TestSetSettings(t *testing.T) {
	_, index := initTest(t)
	defer tearDownTest(t, index)

	settings := Map{"hitsPerPage": 30}

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

	checkEqual(t, get.HitsPerPage, settings["hitsPerPage"], "hitsPerPage")
}

func TestBrowse(t *testing.T) {
	_, index := initTest(t)
	defer tearDownTest(t, index)
	object := Object{"name": "John Snow", "objectID": "àlgol?à"}

	addWait(object, index, t)

	_, err := index.Browse(Map{"page": 1, "hitsPerPage": 1}, "")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestBrowseWithCursor(t *testing.T) {
	_, index := initTest(t)
	defer tearDownTest(t, index)
	object := Object{"name": "John Snow", "objectID": "àlgol?à"}

	addWait(object, index, t)

	browse, err := index.BrowseAll(Map{"query": ""})
	if err != nil {
		t.Fatal(err.Error())
	}

	hit, err := browse.Next()
	if err != nil {
		t.Fatal(err.Error())
	}

	checkEqual(t, hit["name"], "John Snow", "name")

	_, err = browse.Next()
	if err == nil {
		t.Fatal("Should contains only one element")
	}
}

func TestQuery(t *testing.T) {
	_, index := initTest(t)
	defer tearDownTest(t, index)
	object := Object{"name": "John Snow", "objectID": "àlgol?à"}

	addWait(object, index, t)

	params := Map{"attributesToRetrieve": []string{"*"}, "getRankingInfo": 1}
	search, err := index.Search("", params)
	if err != nil {
		t.Fatal(err.Error())
	}

	checkNbHits(t, search.NbHits, 1)
}

func TestIndexCopy(t *testing.T) {
	client, index := initTest(t)
	defer tearDownTest(t, index)
	object := Object{"name": "John Snow", "objectID": "àlgol?à"}

	addWait(object, index, t)

	copyName := genIndexName(index.name)
	idx, err := index.Copy(copyName)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = index.WaitTask(idx.TaskID)
	if err != nil {
		t.Fatal(err.Error())
	}

	indexCopy := client.InitIndex(copyName)
	defer tearDownTest(t, indexCopy)

	search, err := indexCopy.Search("", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	checkNbHits(t, search.NbHits, 1)
}

func TestCopy(t *testing.T) {
	client, index := initTest(t)
	defer tearDownTest(t, index)
	object := Object{"name": "John Snow", "objectID": "àlgol?à"}

	addWait(object, index, t)

	copyName := genIndexName(index.name)
	copy, err := client.CopyIndex(index.name, copyName)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = index.WaitTask(copy.TaskID)
	if err != nil {
		t.Fatal(err.Error())
	}

	indexCopy := client.InitIndex(copyName)
	defer tearDownTest(t, indexCopy)

	search, err := indexCopy.Search("", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	checkNbHits(t, search.NbHits, 1)
}

func TestIndexMove(t *testing.T) {
	client, index := initTest(t)
	defer tearDownTest(t, index)
	object := Object{"name": "John Snow", "objectID": "àlgol?à"}

	addWait(object, index, t)

	moveName := genIndexName(index.name)
	move, err := index.Move(moveName)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = index.WaitTask(move.TaskID)
	if err != nil {
		t.Fatal(err.Error())
	}

	indexMove := client.InitIndex(moveName)
	defer tearDownTest(t, indexMove)

	search, err := indexMove.Search("", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	checkNbHits(t, search.NbHits, 1)
}

func TestMove(t *testing.T) {
	client, index := initTest(t)
	defer tearDownTest(t, index)
	object := Object{"name": "John Snow", "objectID": "àlgol?à"}

	addWait(object, index, t)

	moveName := genIndexName(index.name)
	move, err := client.MoveIndex(index.name, moveName)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = index.WaitTask(move.TaskID)
	if err != nil {
		t.Fatal(err.Error())
	}

	indexMove := client.InitIndex(moveName)
	defer tearDownTest(t, indexMove)

	search, err := indexMove.Search("", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	checkNbHits(t, search.NbHits, 1)
}

func TestAddIndexKey(t *testing.T) {
	_, index := initTest(t)
	defer tearDownTest(t, index)

	newKey := Map{
		"validity":               300,
		"maxQueriesPerIPPerHour": 100,
		"maxHitsPerQuery":        100,
	}

	add, err := index.AddUserKey([]string{"search"}, newKey)
	if err != nil {
		t.Fatal(err.Error())
	}

	waitIndexKey(index, add.Key, nil)

	get, err := index.GetUserKey(add.Key)
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

	updated := Map{"acl": []string{"addObject"}}
	_, err = index.UpdateUserKey(add.Key, updated)
	if err != nil {
		t.Fatal(err.Error())
	}

	waitIndexKey(index, add.Key, func(k Key) bool {
		return len(k.ACL) == 1 && k.ACL[0] == "addObject"
	})

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

	_, err = index.DeleteUserKey(add.Key)
	if err != nil {
		t.Fatal(err.Error())
	}

	waitMissingIndexKey(index, add.Key)

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
}

func TestAddKey(t *testing.T) {
	client, index := initTest(t)
	defer tearDownTest(t, index)

	acl := []string{"search"}
	params := Map{
		"validity":               300,
		"maxHitsPerQuery":        100,
		"maxQueriesPerIPPerHour": 100,
		"indexes":                []string{index.name},
	}

	add, err := client.AddUserKey(acl, params)
	if err != nil {
		t.Fatal(err.Error())
	}

	waitClientKey(client, add.Key, nil)

	get, err := client.GetUserKey(add.Key)
	if err != nil {
		t.Fatal(err.Error())
	}

	if get.Value != add.Key {

		t.Fatal("Unable to get a key")
	}

	_, err = client.UpdateUserKey(add.Key, Map{"acl": []string{"addObject"}})
	if err != nil {
		t.Fatal(err.Error())
	}

	waitClientKey(client, add.Key, func(k Key) bool {
		return len(k.ACL) == 1 && k.ACL[0] == "addObject"
	})

	list, err := client.ListKeys()
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

	_, err = client.DeleteUserKey(add.Key)
	if err != nil {
		t.Fatal(err.Error())
	}

	waitMissingClientKey(client, add.Key)

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
}

func TestAddObjects(t *testing.T) {
	_, index := initTest(t)
	defer tearDownTest(t, index)
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

	checkNbHits(t, search.NbHits, 2)
}

func TestUpdateObjects(t *testing.T) {
	_, index := initTest(t)
	defer tearDownTest(t, index)
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

	checkNbHits(t, search.NbHits, 2)
}

func TestPartialUpdateObjects(t *testing.T) {
	_, index := initTest(t)
	defer tearDownTest(t, index)
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

	checkNbHits(t, search.NbHits, 2)
}

func TestDeleteObjects(t *testing.T) {
	_, index := initTest(t)
	defer tearDownTest(t, index)
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

	checkNbHits(t, search.NbHits, 0)
}

func TestDeleteByQuery(t *testing.T) {
	_, index := initTest(t)
	defer tearDownTest(t, index)
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

	err = index.DeleteByQuery("San", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	search, err := index.Search("", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	checkNbHits(t, search.NbHits, 1)
}

func TestGenerateNewSecuredApiKey(t *testing.T) {
	base := "182634d8894831d5dbce3b3185c50881"

	key, err := GenerateSecuredAPIKey(base, Map{"tagFilters": "(public,user1)"})
	if err != nil {
		t.Fatal(err.Error())
	}
	expected := "MDZkNWNjNDY4M2MzMDA0NmUyNmNkZjY5OTMzYjVlNmVlMTk1NTEwMGNmNTVjZmJhMmIwOTIzYjdjMTk2NTFiMnRhZ0ZpbHRlcnM9JTI4cHVibGljJTJDdXNlcjElMjk="
	checkEqual(t, key, expected, "secured key")

	key, err = GenerateSecuredAPIKey(base, Map{"tagFilters": "(public,user1)", "userToken": "42"})
	if err != nil {
		t.Fatal(err.Error())
	}
	expected = "OGYwN2NlNTdlOGM2ZmM4MjA5NGM0ZmYwNTk3MDBkNzMzZjQ0MDI3MWZjNTNjM2Y3YTAzMWM4NTBkMzRiNTM5YnRhZ0ZpbHRlcnM9JTI4cHVibGljJTJDdXNlcjElMjkmdXNlclRva2VuPTQy"
	checkEqual(t, key, expected, "secured key")
}

func TestMultipleQueries(t *testing.T) {
	client, index := initTest(t)
	defer tearDownTest(t, index)
	object := Object{"name": "John Snow"}

	addWait(object, index, t)

	queries := []IndexedQuery{
		{IndexName: index.name, Params: Map{"query": "John"}},
	}

	search, err := client.MultipleQueries(queries, "")
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(search) == 0 {
		t.Fatal("search shouldn't be empty")
	}
	checkNbHits(t, search[0].NbHits, 1)
}

func TestFacets(t *testing.T) {
	_, index := initTest(t)
	defer tearDownTest(t, index)

	settings := Map{"attributesForFacetting": []string{"f", "g"}}
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

	search, err := index.Search("", Map{"facets": "f", "facetFilters": "f:f1"})
	if err != nil {
		t.Fatal(err.Error())
	}

	checkNbHits(t, search.NbHits, 2)
}

func TestSynonyms(t *testing.T) {
	_, index := initTest(t)
	defer tearDownTest(t, index)

	object := Object{"name": "589 Howard St., San Francisco"}
	addWait(object, index, t)

	batch, err := index.BatchSynonyms([]Synonym{
		NewSynonym("city", []string{"San Francisco", "SF"}),
		NewAltCorrectionSynomym("street", []string{"St"}, "Street", AltCorrection1),
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

	checkEqual(t, get.ObjectID, "city", "city")

	search, err := index.Search("Howard Street SF", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	checkNbHits(t, search.NbHits, 1)

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

	checkNbHits(t, len(synonyms), 1)

	clear, err := index.ClearSynonyms(false)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = index.WaitTask(clear.TaskID)
	if err != nil {
		t.Fatal(err.Error())
	}

	synonyms, err = index.SearchSynonyms("", []string{}, 0, 5)
	if err != nil {
		t.Fatal(err.Error())
	}

	checkNbHits(t, len(synonyms), 0)
}

func TestDNSTimeout(t *testing.T) {
	appID, haveAppID := syscall.Getenv("ALGOLIA_APPLICATION_ID")
	apiKey, haveAPIKey := syscall.Getenv("ALGOLIA_API_KEY")
	if !haveAPIKey || !haveAppID {
		t.Fatal("Need ALGOLIA_APPLICATION_ID and ALGOLIA_API_KEY")
	}

	hosts := []string{
		appID + "-dsn.algolia.biz",
		appID + "-dsn.algolia.net",
		appID + "-1.algolianet.com",
		appID + "-2.algolianet.com",
		appID + "-3.algolianet.com",
	}

	c := NewClientWithHosts(appID, apiKey, hosts)

	start := time.Now()
	_, _ = c.ListIndexes()

	if start.Add(3 * time.Second).Before(time.Now()) {
		t.Fatal("DNS resolution should have timeouted")
	}
}
