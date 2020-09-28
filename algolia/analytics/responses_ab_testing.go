package analytics

import (
	"time"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

// ABTestTaskRes represents a generic task submitted to the Analytics API for an
// AB test.
type ABTestTaskRes struct {
	ABTestID int    `json:"abTestID"`
	Index    string `json:"index"`
	TaskID   int64  `json:"taskID"`
	wait     func(index string, taskID int64) error
}

// Wait blocks until the AB test task completes or if there is an error while
// waiting for its completion.
func (r ABTestTaskRes) Wait() error {
	return r.wait(r.Index, r.TaskID)
}

// GetABTestsRes represents a response payload from the Analytics API when
// performing a GetABTests call.
type GetABTestsRes struct {
	ABTests []ABTestResponse `json:"abtests"`
	Count   int              `json:"count"`
	Total   int              `json:"total"`
}

// ABTestResponse represents an AB test as returned by the Analytics API when
// retrieved as part of a response.
type ABTestResponse struct {
	ABTestID               int               `json:"abTestID"`
	ClickSignificance      int               `json:"clickSignificance"`
	ConversionSignificance float64           `json:"conversionSignificance"`
	CreatedAt              time.Time         `json:"createdAt"`
	EndAt                  time.Time         `json:"endAt"`
	Name                   string            `json:"name"`
	Status                 string            `json:"status"`
	Variants               []VariantResponse `json:"variants"`
}

// VariantResponse represents an AB test's variant as returned by the Analytics
// API when retrieved as part of a response.
type VariantResponse struct {
	AverageClickPosition   int                 `json:"averageClickPosition"`
	ClickCount             int                 `json:"clickCount"`
	ClickThroughRate       float64             `json:"clickThroughRate"`
	ConversionCount        int                 `json:"conversionCount"`
	ConversionRate         float64             `json:"conversionRate"`
	Description            string              `json:"description"`
	Index                  string              `json:"index"`
	NoResultCount          int                 `json:"noResultCount"`
	SearchCount            int                 `json:"searchCount"`
	TrackedSearchCount     int                 `json:"trackedSearchCount"`
	TrafficPercentage      int                 `json:"trafficPercentage"`
	UserCount              int                 `json:"userCount"`
	CustomSearchParameters *search.QueryParams `json:"customSearchParameters"`
}
