package analytics

import "time"

type ABTestTaskRes struct {
	ABTestID int    `json:"abTestID"`
	Index    string `json:"index"`
	TaskID   int    `json:"taskID"`
	wait     func(index string, taskID int) error
}

func (r ABTestTaskRes) Wait() error {
	return r.wait(r.Index, r.TaskID)
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
