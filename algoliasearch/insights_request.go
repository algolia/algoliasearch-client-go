package algoliasearch

type InsightsRequest struct {
	EventType string   `json:"eventType"`
	EventName string   `json:"eventName"`
	Index     string   `json:"index"`
	UserToken string   `json:"userToken"`
	Timestamp int64    `json:"timestamp"`
	QueryID   string   `json:"queryID"`
	ObjectIDs []string `json:"objectIDs"`
	Filters   []string `json:"filters"`
	Positions []int    `json:"positions"`
}

func newInsightsRequest(userToken, eventName, indexName, eventType string) InsightsRequest {
	return InsightsRequest{
		EventType: eventType,
		EventName: eventName,
		Index:     indexName,
		UserToken: userToken,
	}
}

func clickInsightsRequest(userToken, eventName, indexName string) InsightsRequest {
	return newInsightsRequest(userToken, eventName, indexName, "click")
}

func conversionInsightsRequest(userToken, eventName, indexName string) InsightsRequest {
	return newInsightsRequest(userToken, eventName, indexName, "conversion")
}

func viewInsightsRequest(userToken, eventName, indexName string) InsightsRequest {
	return newInsightsRequest(userToken, eventName, indexName, "view")
}
