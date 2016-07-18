package algoliasearch

import (
	"sync"
	"testing"
	"time"
)

func TestClientOperations(t *testing.T) {
	c, i := initClientAndIndex(t, "TestClientOperations")

	objectID := addOneObject(t, c, i)

	// Test CopyIndex
	{
		res, err := c.CopyIndex("TestClientOperations", "TestClientOperations_copy")
		if err != nil {
			t.Fatalf("TestClientOperations: Cannot copy the index: %s", err)
		}

		waitTask(t, i, res.TaskID)
	}

	// Test MoveIndex
	i = c.InitIndex("TestClientOperations_copy")
	{
		res, err := c.MoveIndex("TestClientOperations_copy", "TestClientOperations_move")
		if err != nil {
			t.Fatalf("TestClientOperations: Cannot move the index: %s", err)
		}

		waitTask(t, i, res.TaskID)
	}

	// Test ClearIndex
	i = c.InitIndex("TestClientOperations_move")
	{
		res, err := c.ClearIndex("TestClientOperations_move")
		if err != nil {
			t.Fatalf("TestClear: Cannot clear the index: %s, err")
		}

		waitTask(t, i, res.TaskID)

		_, err = i.GetObject(objectID, nil)
		if err == nil || err.Error() != "{\"message\":\"ObjectID does not exist\",\"status\":404}\n" {
			t.Fatalf("TestClientOperations: Object %s should be deleted after clear: %s", objectID, err)
		}
	}

	// Test DeleteIndex
	{
		_, err := c.DeleteIndex("TestClientOperations_move")
		if err != nil {
			t.Fatalf("TestClientOperations: Cannot delete the moved index: %s", err)
		}
	}
}

func deleteClientKey(t *testing.T, c Client, key string) {
	_, err := c.DeleteUserKey(key)
	if err != nil {
		t.Fatalf("deleteClientKey: Cannot delete key: %s", err)
	}
}

func waitClientKey(t *testing.T, c Client, keyID string, f func(k Key) bool) {
	retries := 10

	for r := 0; r < retries; r++ {
		key, err := c.GetUserKey(keyID)

		if err == nil && (f == nil || f(key)) {
			return
		}
		time.Sleep(1 * time.Second)
	}

	t.Fatalf("waitClientKey: Key not found or function call failed")
}

func waitClientKeysAsync(t *testing.T, c Client, keyIDs []string, f func(k Key) bool) {
	var wg sync.WaitGroup

	for _, keyID := range keyIDs {
		wg.Add(1)

		go func(keyID string) {
			defer wg.Done()
			waitClientKey(t, c, keyID, f)
		}(keyID)
	}

	wg.Wait()
}

func TestClientKeys(t *testing.T) {
	c := initClient(t)

	// Check that no key was previously existing
	{
		keys, err := c.ListKeys()

		if err != nil {
			t.Fatalf("TestClientKeys: Cannot list the keys: %s", err)
		}

		if len(keys) != 1 || keys[0].Description != "Search-only API Key" {
			t.Fatalf("TestClientKeys: Should return the Search-only API Key instead of %d key(s)", len(keys))
		}
	}

	var searchKey, allRightsKey string

	// Add a search key with parameters
	{
		params := Map{
			"description":            "",
			"maxQueriesPerIPPerHour": 1000,
			"referers":               []string{},
			"queryParameters":        "typoTolerance=strict",
			"validity":               600,
			"maxHitsPerQuery":        1,
		}

		res, err := c.AddUserKey([]string{"search"}, params)
		if err != nil {
			t.Fatalf("TestClientKeys: Cannot create the search key: %s", err)
		}

		searchKey = res.Key
	}
	defer deleteClientKey(t, c, searchKey)

	// Add an all-permissions key
	{
		acl := []string{
			"search",
			"browse",
			"addObject",
			"deleteObject",
			"deleteIndex",
			"settings",
			"editSettings",
			"analytics",
			"listIndexes",
		}

		res, err := c.AddUserKey(acl, nil)
		if err != nil {
			t.Fatalf("TestClientKeys: Cannot create the all-rights key: %s", err)
		}

		allRightsKey = res.Key
	}
	defer deleteClientKey(t, c, allRightsKey)

	waitClientKeysAsync(t, c, []string{searchKey, allRightsKey}, nil)

	// Check that the 2 previous keys were added
	{
		keys, err := c.ListKeys()

		if err != nil {
			t.Fatalf("TestClientKeys: Cannot list the added keys: %s", err)
		}

		if len(keys) != 3 {
			t.Fatalf("TestClientKeys: Should return 3 keys instead of %d", len(keys))
		}
	}

	// Update search key description
	{
		params := Map{"description": "Search-Only Key"}

		_, err := c.UpdateUserKey(searchKey, params)
		if err != nil {
			t.Fatalf("TestClientKeys: Cannot update search only key's description: %s", err)
		}

		waitClientKey(t, c, searchKey, func(k Key) bool { return k.Description == "Search-Only Key" })
	}
}
