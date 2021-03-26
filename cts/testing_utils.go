package cts

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"os/user"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/analytics"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/compression"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/insights"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/recommendation"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/region"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/transport"
)

func InitSearchClient1AndIndex(t *testing.T) (*search.Client, *search.Index, string) {
	c := InitSearchClient1(t)
	indexName := GenerateIndexName(t)
	i := c.InitIndex(indexName)
	return c, i, indexName
}

func GenerateIndexName(t *testing.T) string {
	indexName := GenerateCanonicalPrefixName() + "_" + t.Name()
	indexName = cleanIndexName(indexName)
	return indexName
}

func cleanIndexName(indexName string) string {
	for _, char := range []string{
		"/",
		"-",
		".",
		":",
		" ",
		"%",
		"+",
		"^",
	} {
		indexName = strings.Replace(indexName, char, "_", -1)
	}
	return indexName
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

func InitRecommendationClient(t *testing.T) *recommendation.Client {
	return initRecommendationClientWith(t, "ALGOLIA_APPLICATION_ID_1", "ALGOLIA_ADMIN_KEY_1")
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
	return c
}

func initAnalyticsClientWith(t *testing.T, appIDEnvVar, apiKeyEnvVar string) *analytics.Client {
	appID, key := GetTestingCredentials(t, appIDEnvVar, apiKeyEnvVar)
	c := analytics.NewClient(appID, key)
	return c
}

func initRecommendationClientWith(t *testing.T, appIDEnvVar, apiKeyEnvVar string) *recommendation.Client {
	appID, key := GetTestingCredentials(t, appIDEnvVar, apiKeyEnvVar)
	c := recommendation.NewClient(appID, key, region.US)
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

	if build := os.Getenv("CIRCLE_BUILD_NUM"); build != "" {
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

func Retry(shouldStopFunc func() bool) {
	for !shouldStopFunc() {
		time.Sleep(100 * time.Millisecond)
	}
}

func TodayDate() string {
	return time.Now().Format("2006_01_02")
}

func TodayDateTime() string {
	return time.Now().Format("2006_01_02_15_04_05")
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

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateRandomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
