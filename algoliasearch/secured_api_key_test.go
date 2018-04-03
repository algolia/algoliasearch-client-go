package algoliasearch

import (
	"os"
	"testing"
)

func TestGenerateSecuredAPIKey_Generation(t *testing.T) {
	t.Parallel()
	apiKey := os.Getenv("ALGOLIA_API_KEY")
	if apiKey == "" {
		t.Fatal("TestGenerateSecuredAPIKey_Generation: Missing ALGOLIA_API_KEY")
	}

	t.Log("TestGenerateSecuredAPIKey_Generation: Tests invalid key generations")
	{
		cases := []struct {
			params      Map
			expectedErr error
		}{
			{Map{"userToken": 42}, invalidType("userToken", "string")},
			{Map{"restrictIndices": 42}, invalidType("restrictIndices", "string")},
			{Map{"restrictSources": 42}, invalidType("restrictSources", "string")},
			{Map{"validUntil": "NaN"}, invalidType("validUntil", "int")},
		}

		for _, c := range cases {
			_, err := GenerateSecuredAPIKey(apiKey, c.params)
			if err == nil || err.Error() != c.expectedErr.Error() {
				t.Errorf("TestGenerateSecuredAPIKey_Generation: Calling GenerateSecuredAPIKey(%#v)\nexpected error:  \"%s\"\nbut got instead: \"%s\"",
					c.params,
					c.expectedErr,
					err)
			}
		}
	}

	t.Log("TestGenerateSecuredAPIKey_Generation: Tests valid key generations")
	{
		cases := []struct {
			params      Map
			expectedKey string
		}{
			{Map{"userToken": "user42"}, "NTc5YjBkMTgwYjdkMzBlYzllYzY5MmY3OGRmOGQzMWU3ZWU0ZTI2ZmY4MGQ0ZTZhMWZlNzJiMzllMjg5YzhmZnVzZXJUb2tlbj11c2VyNDI="},
			{Map{"restrictIndices": "myIndex"}, "ZGQ5NjRjYTdmOWRkYzYwMjBkNWQ4ZGQ3MmZlY2RkOTYyZjIxM2FjNDBhMjBhYzhhNDFiYWI4NDE4ZGJiOTgxYXJlc3RyaWN0SW5kaWNlcz1teUluZGV4"},
			{Map{"restrictSources": ""}, "NmNjYzQ0MzI0MmU3OTg1NDJiZDYyNTIwZjE2OWMwYjU1MjQ0ZmFhNDdmNzdjNDg1MGYxYmY1YWJjNWZkOTU2OHJlc3RyaWN0U291cmNlcz0="},
			{Map{"validUntil": 1481901339150}, "NDZiMWNlNDMyMzEzNTRkZjFiYmMyYjE2ZWFmNzVjOWE5MjkzNTllMTgxZjM3NDI1OWNiZjAyOGZjMTc0NzU3MXZhbGlkVW50aWw9MTQ4MTkwMTMzOTE1MA=="},
		}

		for _, c := range cases {
			generatedKey, err := GenerateSecuredAPIKey(apiKey, c.params)
			if err != nil {
				t.Errorf("TestGenerateSecuredAPIKey_Generation: Key with params %#v should have been generated withouth error but got: %s", c.params, err)
			}

			if generatedKey != c.expectedKey {
				t.Errorf("TestGenerateSecuredAPIKey_Generation: Key was not generated correctly for params %#v:\nexpected key:  \"%s\"\nbut got instead: \"%s\"",
					c.params,
					c.expectedKey,
					generatedKey)
			}
		}
	}
}

func TestGenerateSecuredAPIKey_Usage(t *testing.T) {
	t.Parallel()

	t.Log("TestGenerateSecuredAPIKey_Usage: Obtain application ID/API key from the environment")
	appID := os.Getenv("ALGOLIA_APPLICATION_ID")
	if appID == "" {
		t.Fatalf("TestGenerateSecuredAPIKey_Usage: Cannot retrieve application ID from environment")
	}
	apiKey := os.Getenv("ALGOLIA_SEARCH_KEY")
	if apiKey == "" {
		t.Fatal("TestGenerateSecuredAPIKey_Usage: Missing ALGOLIA_SEARCH_KEY")
	}

	t.Log("TestGenerateSecuredAPIKey_Usage: Generate secured API key")
	params := Map{
		"restrictIndices": "TestGenerateSecuredAPIKey_Usage_1,TestGenerateSecuredAPIKey_Usage_2",
	}
	key, err := GenerateSecuredAPIKey(apiKey, params)
	if err != nil {
		t.Fatalf("TestGenerateSecuredAPIKey_Usage: should generate key without error")
	}

	t.Log("TestGenerateSecuredAPIKey_Usage: Check if key correctly restrict accesse to indices")
	client := initClient(t)
	restrictedClient := NewClient(appID, key)

	for _, c := range []struct {
		indexName    string
		isAuthorized bool
	}{
		{"TestGenerateSecuredAPIKey_Usage_1", true},
		{"TestGenerateSecuredAPIKey_Usage_2", true},
		{"TestGenerateSecuredAPIKey_Usage_3", false},
	} {
		i := initIndex(t, client, c.indexName)
		addOneObject(t, i)

		i = restrictedClient.InitIndex(c.indexName)
		_, err := i.Search("", nil)
		if c.isAuthorized && err != nil {
			t.Fatalf("TestGenerateSecuredAPIKey_Usage: index %s should be searchable with restricted API key", c.indexName)
		}
		if !c.isAuthorized && err == nil {
			t.Fatalf("TestGenerateSecuredAPIKey_Usage: index %s should not be searchable with restricted API key", c.indexName)
		}
	}
}
