package cts

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"os/user"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/analytics"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/compression"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/insights"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/transport"
)

func InitSearchClient1AndIndex(t *testing.T) (*search.Client, *search.Index, string) {
	c := InitSearchClient1(t)
	canonicalName := GenerateCanonicalPrefixName()
	indexName := canonicalName + "_" + t.Name()
	indexName = strings.Replace(indexName, "/", "_", -1)
	indexName = strings.Replace(indexName, " ", "_", -1)
	i := c.InitIndex(indexName)
	return c, i, indexName
}

func InitSearchClient1(t *testing.T) *search.Client {
	return initSearchClientWith(t, "ALGOLIA_APPLICATION_ID_1", "ALGOLIA_ADMIN_KEY_1")
}

func InitSearchClient2(t *testing.T) *search.Client {
	return initSearchClientWith(t, "ALGOLIA_APPLICATION_ID_2", "ALGOLIA_ADMIN_KEY_2")
}

func InitSearchClientMCM(t *testing.T) *search.Client {
	return initSearchClientWith(t, "ALGOLIA_APPLICATION_ID_MCM", "ALGOLIA_ADMIN_KEY_MCM")
}

func InitAnalyticsClient1(t *testing.T) *analytics.Client {
	return initAnalyticsClientWith(t, "ALGOLIA_APPLICATION_ID_1", "ALGOLIA_ADMIN_KEY_1")
}

func InitInsightsClient(t *testing.T) *insights.Client {
	return initInsightsClientWith(t, "ALGOLIA_APPLICATION_ID_1", "ALGOLIA_ADMIN_KEY_1")
}

func initInsightsClientWith(t *testing.T, appIDEnvVar, apiKeyEnvVar string) *insights.Client {
	appID, key := GetTestingCredentials(t, appIDEnvVar, apiKeyEnvVar)
	c := insights.NewClient(appID, key)
	return c
}

func initSearchClientWith(t *testing.T, appIDEnvVar, apiKeyEnvVar string) *search.Client {
	appID, key := GetTestingCredentials(t, appIDEnvVar, apiKeyEnvVar)
	var c *search.Client
	if appIDEnvVar == "ALGOLIA_APPLICATION_ID_2" { // TODO: temporary workaround while GZIP is not enabled on this cluster
		c = search.NewClientWithConfig(search.Configuration{AppID: appID, APIKey: key, Compression: compression.None})
	} else {
		c = search.NewClient(appID, key)
	}
	deleteOldIndices(c)
	return c
}

func initAnalyticsClientWith(t *testing.T, appIDEnvVar, apiKeyEnvVar string) *analytics.Client {
	appID, key := GetTestingCredentials(t, appIDEnvVar, apiKeyEnvVar)
	c := analytics.NewClient(appID, key)
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

func deleteOldIndices(c *search.Client) {
	today := "go_" + TodayDate()
	indices, _ := c.ListIndices()
	for _, index := range indices.Items {
		if strings.HasPrefix(index.Name, "go_") && !strings.HasPrefix(index.Name, today) {
			_, _ = c.InitIndex(index.Name).Delete()
		}
	}
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

func GenerateSecuredAPIKeyWithArbitraryParameters(
	t *testing.T,
	apiKey string,
	params map[string]interface{},
) string {
	h := hmac.New(sha256.New, []byte(apiKey))

	message := transport.URLEncode(params)
	_, err := h.Write([]byte(message))
	require.NoError(t, err)

	checksum := hex.EncodeToString(h.Sum(nil))
	return base64.StdEncoding.EncodeToString([]byte(checksum + message))
}
