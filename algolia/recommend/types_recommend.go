package recommend

import "github.com/algolia/algoliasearch-client-go/v3/algolia/search"

type multipleOptions struct {
	Requests []RecommendationsOptions `json:"requests"`
}

type RecommendationModel string

const (
	RelatedProducts RecommendationModel = "related-products"
	BoughtTogether  RecommendationModel = "bought-together"
)

type RecommendationsOptions struct {
	IndexName          string              `json:"indexName"`
	Model              RecommendationModel `json:"model"`
	ObjectID           string              `json:"objectID"`
	Threshold          int                 `json:"threshold"`
	MaxRecommendations *int                `json:"maxRecommendations,omitempty"`
	QueryParameters    *search.QueryParams `json:"queryParameters,omitempty"`
	FallbackParameters *search.QueryParams `json:"fallbackParameters,omitempty"`
}

type RelatedProductsOptions struct {
	recommendationsOptions RecommendationsOptions
}

func NewRelatedProductsOptions(indexName string, objectID string, threshold int, maxRecommendations *int, queryParameters *search.QueryParams, fallbackParameters *search.QueryParams) RelatedProductsOptions {
	return RelatedProductsOptions{recommendationsOptions: RecommendationsOptions{indexName, RelatedProducts, objectID, threshold, maxRecommendations, queryParameters, fallbackParameters}}
}

type FrequentlyBoughtTogetherOptions struct {
	recommendationsOptions RecommendationsOptions
}

func NewBoughtTogetherOptions(indexName string, objectID string, threshold int, maxRecommendations *int, queryParameters *search.QueryParams) FrequentlyBoughtTogetherOptions {
	return FrequentlyBoughtTogetherOptions{recommendationsOptions: RecommendationsOptions{indexName, BoughtTogether, objectID, threshold, maxRecommendations, queryParameters, nil}}
}
