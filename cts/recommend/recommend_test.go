package recommend

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/recommend"
	"github.com/stretchr/testify/require"
)

type FakeRequester struct {
	Do func(http.Request) (*http.Response, error)
}

func (f *FakeRequester) Request(req *http.Request) (*http.Response, error) {
	return f.Do(*req)
}

type multipleOptions struct {
	Requests []recommend.RecommendationsOptions `json:"requests"`
}

func TestGetRecommendations(t *testing.T) {

	config := recommend.Configuration{AppID: "testAppID", APIKey: "testApiKey", Requester: &FakeRequester{Do: func(r http.Request) (*http.Response, error) {
		require.Equal(t, r.Method, "POST")
		require.Equal(t, r.URL.String(), "https://testAppID-dsn.algolia.net/1/indexes/*/recommendations")
		require.Equal(t, r.Header.Get("X-Algolia-Application-Id"), "testAppID")
		require.Equal(t, r.Header.Get("X-Algolia-API-Key"), "testApiKey")

		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		require.Nil(t, err)

		var options multipleOptions

		err = json.Unmarshal(body, &options)
		require.Nil(t, err)

		var option = options.Requests[len(options.Requests)-1]
		require.Equal(t, option.IndexName, "IndexName")
		require.Equal(t, option.ObjectID, "B018APC4LE")
		require.Equal(t, option.Model, recommend.BoughtTogether)
		return &http.Response{Status: "200"}, nil
	}}}

	client := recommend.NewClientWithConfig(config)

	options := recommend.RecommendationsOptions{IndexName: "IndexName", ObjectID: "B018APC4LE", Model: recommend.BoughtTogether}

	_, _ = client.GetRecommendations([]recommend.RecommendationsOptions{options})

}

func TestGetRelatedProducts(t *testing.T) {

	config := recommend.Configuration{AppID: "testAppID", APIKey: "testApiKey", Requester: &FakeRequester{Do: func(r http.Request) (*http.Response, error) {
		require.Equal(t, r.Method, "POST")
		require.Equal(t, r.URL.String(), "https://testAppID-dsn.algolia.net/1/indexes/*/recommendations")
		require.Equal(t, r.Header.Get("X-Algolia-Application-Id"), "testAppID")
		require.Equal(t, r.Header.Get("X-Algolia-API-Key"), "testApiKey")

		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		require.Nil(t, err)

		var options multipleOptions

		err = json.Unmarshal(body, &options)
		require.Nil(t, err)

		var option = options.Requests[len(options.Requests)-1]
		require.Equal(t, option.IndexName, "IndexName")
		require.Equal(t, option.ObjectID, "B018APC4LE")
		require.Equal(t, option.Model, recommend.RelatedProducts)
		return &http.Response{Status: "200"}, nil
	}}}

	client := recommend.NewClientWithConfig(config)

	options := recommend.NewRelatedProductsOptions("IndexName", "B018APC4LE", 0, nil, nil, nil)

	_, _ = client.GetRelatedProducts([]recommend.RelatedProductsOptions{options})

}

func TestGetBoughtTogether(t *testing.T) {

	config := recommend.Configuration{AppID: "testAppID", APIKey: "testApiKey", Requester: &FakeRequester{Do: func(r http.Request) (*http.Response, error) {
		require.Equal(t, r.Method, "POST")
		require.Equal(t, r.URL.String(), "https://testAppID-dsn.algolia.net/1/indexes/*/recommendations")
		require.Equal(t, r.Header.Get("X-Algolia-Application-Id"), "testAppID")
		require.Equal(t, r.Header.Get("X-Algolia-API-Key"), "testApiKey")

		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		require.Nil(t, err)

		var options multipleOptions

		err = json.Unmarshal(body, &options)
		require.Nil(t, err)

		var option = options.Requests[len(options.Requests)-1]
		require.Equal(t, option.IndexName, "IndexName")
		require.Equal(t, option.ObjectID, "B018APC4LE")
		require.Equal(t, option.Model, recommend.BoughtTogether)
		return &http.Response{Status: "200"}, nil
	}}}

	client := recommend.NewClientWithConfig(config)

	options := recommend.NewBoughtTogetherOptions("IndexName", "B018APC4LE", 0, nil, nil)

	_, _ = client.GetFrequentlyBoughtTogether([]recommend.FrequentlyBoughtTogetherOptions{options})

}
