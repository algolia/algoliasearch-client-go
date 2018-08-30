package it

import (
	"fmt"
	"os"
	"os/user"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/stretchr/testify/require"
)

func InitSearchClient1AndIndex(t *testing.T) (algoliasearch.Client, algoliasearch.Index, string) {
	c := InitSearchClient1(t)
	canonicalName := GenerateCanonicalPrefixName()
	indexName := canonicalName + "_" + t.Name()
	indexName = strings.Replace(indexName, "/", "_", -1)
	indexName = strings.Replace(indexName, " ", "_", -1)
	i := c.InitIndex(indexName)
	return c, i, indexName
}

func InitSearchClient1(t *testing.T) algoliasearch.Client {
	return initSearchClientWith(t, "ALGOLIA_APPLICATION_ID_1", "ALGOLIA_ADMIN_KEY_1")
}

func InitSearchClient2(t *testing.T) algoliasearch.Client {
	return initSearchClientWith(t, "ALGOLIA_APPLICATION_ID_2", "ALGOLIA_ADMIN_KEY_2")
}

func InitSearchClientMCM(t *testing.T) algoliasearch.Client {
	return initSearchClientWith(t, "ALGOLIA_APPLICATION_ID_MCM", "ALGOLIA_ADMIN_KEY_MCM")
}

func initSearchClientWith(t *testing.T, appIDEnvVar, apiKeyEnvVar string) algoliasearch.Client {
	appID, key := GetTestingCredentials(t, appIDEnvVar, apiKeyEnvVar)
	c := algoliasearch.NewClient(appID, key)
	go deleteOldIndices(c)
	return c
}

func GetTestingCredentials(t *testing.T, appIDEnvVar, apiKeyEnvVar string) (string, string) {
	appID := os.Getenv(appIDEnvVar)
	key := os.Getenv(apiKeyEnvVar)
	require.NotEmpty(t, appID)
	require.NotEmpty(t, key)
	return appID, key
}

func GenerateCanonicalPrefixName() string {
	var instanceName string

	if build := os.Getenv("TRAVIS_JOB_NUMBER"); build != "" {
		instanceName = build
	} else {
		if user, err := user.Current(); err == nil {
			instanceName = user.Username
		} else {
			instanceName = "unknown"
		}
	}

	return fmt.Sprintf(
		"go_%s_%s",
		TodayDateTime(),
		instanceName,
	)
}

func deleteOldIndices(c algoliasearch.Client) {
	today := "go_" + TodayDate()
	indices, _ := c.ListIndexes()
	for _, index := range indices {
		if strings.HasPrefix(index.Name, "go_") && !strings.HasPrefix(index.Name, today) {
			_, _ = c.InitIndex(index.Name).Delete()
		}
	}
}

func WaitTasks(t *testing.T, i algoliasearch.Index, tasks ...int) {
	var wg sync.WaitGroup
	for _, taskID := range tasks {
		wg.Add(1)
		go func(wg *sync.WaitGroup, taskID int) {
			defer wg.Done()
			err := i.WaitTask(taskID)
			require.NoError(t, err)
		}(&wg, taskID)
	}
	wg.Wait()
}

func Retry(shouldStopFunc func() bool) {
	for !shouldStopFunc() {
		time.Sleep(100 * time.Millisecond)
	}
}

func TodayDate() string {
	return time.Now().Format("2006-01-02")
}

func TodayDateTime() string {
	return time.Now().Format("2006-01-02_15:04:05")
}
