package algoliasearch

import (
	"encoding/json"
	"fmt"
	"time"
)

type ABTest struct {
	Name     string
	Variants []Variant
	EndAt    time.Time
}

const ISO8601 = "2006-01-02T15:04:05Z"

func (abTest ABTest) MarshalJSON() ([]byte, error) {
	jsonVariants, err := json.Marshal(abTest.Variants)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal variants: %s", err)
	}
	jsonAbTest := fmt.Sprintf(
		`{"name":"%s","variants":%s,"endAt":"%s"}`,
		abTest.Name,
		jsonVariants,
		abTest.EndAt.In(time.UTC).Format(ISO8601),
	)
	return []byte(jsonAbTest), nil
}

type Variant struct {
	Index             string `json:"index"`
	TrafficPercentage int    `json:"trafficPercentage"`
	Description       string `json:"description,omitempty"`
}

type ABTestTaskRes struct {
	ABTestID int    `json:"abTestID"`
	Index    string `json:"index"`
	TaskID   int    `json:"taskID"`
}

type GetABTestsRes struct {
	ABTests []ABTestResponse `json:"abtests"`
	Count   int              `json:"count"`
	Total   int              `json:"total"`
}

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

type VariantResponse struct {
	AverageClickPosition int     `json:"averageClickPosition"`
	ClickCount           int     `json:"clickCount"`
	ClickThroughRate     float64 `json:"clickThroughRate"`
	ConversionCount      int     `json:"conversionCount"`
	ConversionRate       float64 `json:"conversionRate"`
	Description          string  `json:"description"`
	Index                string  `json:"index"`
	NoResultCount        int     `json:"noResultCount"`
	SearchCount          int     `json:"searchCount"`
	TrafficPercentage    int     `json:"trafficPercentage"`
	UserCount            int     `json:"userCount"`
}
