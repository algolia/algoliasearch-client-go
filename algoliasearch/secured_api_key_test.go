package algoliasearch

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestGenerateSecuredAPIKey_Generation(t *testing.T) {
	t.Parallel()

	apiKey := os.Getenv("ALGOLIA_SEARCH_KEY_1")
	if apiKey == "" {
		t.Fatal("TestGenerateSecuredAPIKey_Generation: Missing ALGOLIA_SEARCH_KEY_1")
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
			require.Equal(t, c.expectedErr, err, "wrong error when generating secured API key for %#v", c.params)
		}
	}

	t.Log("TestGenerateSecuredAPIKey_Generation: Tests valid key generations")
	{
		cases := []struct {
			params      Map
			expectedKey string
		}{
			{Map{"userToken": "user42"}, "ZDhjZmNmMmMzY2Y4MWY1NjBhMDg0NjNkYzMxYTA2ZTM5YTIyY2JkZDVmMDAwMDNhMDQ2ZGUyYjZkMjc0ZjdmOXVzZXJUb2tlbj11c2VyNDI="},
			{Map{"restrictIndices": "myIndex"}, "OWEwZGMwZGRjYTg3YzY2NzA2MTA2YTQ2YTIyN2Q1MDk4N2NlY2Y5OTE4ZWU4Y2U5NTkyYWQ2MWI1ZTQ3NjMwOHJlc3RyaWN0SW5kaWNlcz1teUluZGV4"},
			{Map{"restrictSources": ""}, "NWMyZjI2ZWNmMTZiY2YyZTNlZmQzZTFjZTg0MGE1OTU2ZjE1NWE2ZGQ2OWI4OTRjMmVhMzJhNDkyNzQ4MjQwN3Jlc3RyaWN0U291cmNlcz0="},
			{Map{"validUntil": 1481901339150}, "NDFiNDFkMzNkNmQ5NDZhZTU2MmIxMzVmODg1YWZkNWI2ODIzMWE5NDIyZGJlNDRlNjM3OWMwOGJhY2RmY2I0NXZhbGlkVW50aWw9MTQ4MTkwMTMzOTE1MA=="},
		}

		for _, c := range cases {
			generatedKey, err := GenerateSecuredAPIKey(apiKey, c.params)
			assert.NoError(t, err, "should not error while generating secured API key for %#v", c.params)
			assert.Equal(t, c.expectedKey, generatedKey, "wrong secured API key was generated for %#v", c.params)
		}
	}
}

func TestGenerateSecuredAPIKey_Usage(t *testing.T) {
	t.Parallel()

	appID := os.Getenv("ALGOLIA_APPLICATION_ID_1")
	require.NotEmpty(t, appID)

	apiKey := os.Getenv("ALGOLIA_SEARCH_KEY_1")
	require.NotEmpty(t, apiKey)

	params := Map{
		"restrictIndices": "TestGenerateSecuredAPIKey_Usage_1,TestGenerateSecuredAPIKey_Usage_2",
	}
	key, err := GenerateSecuredAPIKey(apiKey, params)
	require.NoError(t, err)

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
		if c.isAuthorized {
			assert.NoError(t, err, "should be able to search in index %q", c.indexName)
		} else {
			assert.Error(t, err, "should be able to search in index %q", c.indexName)
		}
	}
}
