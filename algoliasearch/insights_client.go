package algoliasearch

type insights struct {
	client *client
	path   string
}

func NewInsights(client *client) Insights {
	return &insights{
		client: client,
		path:   "/1/events",
	}
}

func (i *insights) User(userToken string) InsightsWithUser {
	return &insightsWithUser{i, userToken}
}

type insightsWithUser struct {
	*insights
	userToken string
}

func (i *insightsWithUser) ClickedFilters(eventName, indexName string, filters []string) (res InsightsResponse, err error) {
	return i.ClickedFiltersWithRequestOptions(eventName, indexName, filters, nil)
}

func (i *insightsWithUser) ClickedFiltersWithRequestOptions(eventName, indexName string, filters []string, opts *RequestOptions) (res InsightsResponse, err error) {
	req := clickInsightsRequest(i.userToken, eventName, indexName)
	req.Filters = filters
	return i.SendEventWithRequestOptions(req, opts)
}

func (i *insightsWithUser) ClickedObjectIDs(eventName, indexName string, objectIDs []string) (res InsightsResponse, err error) {
	return i.ClickedObjectIDsWithRequestOptions(eventName, indexName, objectIDs, nil)
}

func (i *insightsWithUser) ClickedObjectIDsWithRequestOptions(eventName, indexName string, objectIDs []string, opts *RequestOptions) (res InsightsResponse, err error) {
	req := clickInsightsRequest(i.userToken, eventName, indexName)
	req.ObjectIDs = objectIDs
	return i.SendEventWithRequestOptions(req, opts)
}

func (i *insightsWithUser) ClickedObjectIDsAfterSearch(eventName, indexName string, objectIDs []string, positions []int, queryID string) (res InsightsResponse, err error) {
	return i.ClickedObjectIDsAfterSearchWithRequestOptions(eventName, indexName, objectIDs, positions, queryID, nil)
}

func (i *insightsWithUser) ClickedObjectIDsAfterSearchWithRequestOptions(eventName, indexName string, objectIDs []string, positions []int, queryID string, opts *RequestOptions) (res InsightsResponse, err error) {
	req := clickInsightsRequest(i.userToken, eventName, indexName)
	req.ObjectIDs = objectIDs
	req.Positions = positions
	req.QueryID = queryID
	return i.SendEventWithRequestOptions(req, opts)
}

func (i *insightsWithUser) ConvertedObjectIDs(eventName, indexName string, objectIDs []string) (res InsightsResponse, err error) {
	return i.ConvertedObjectIDsWithRequestOptions(eventName, indexName, objectIDs, nil)
}

func (i *insightsWithUser) ConvertedObjectIDsWithRequestOptions(eventName, indexName string, objectIDs []string, opts *RequestOptions) (res InsightsResponse, err error) {
	req := conversionInsightsRequest(i.userToken, eventName, indexName)
	req.ObjectIDs = objectIDs
	return i.SendEventWithRequestOptions(req, opts)
}

func (i *insightsWithUser) ConvertedFilters(eventName, indexName string, filters []string) (res InsightsResponse, err error) {
	return i.ConvertedFiltersWithRequestOptions(eventName, indexName, filters, nil)
}

func (i *insightsWithUser) ConvertedFiltersWithRequestOptions(eventName, indexName string, filters []string, opts *RequestOptions) (res InsightsResponse, err error) {
	req := conversionInsightsRequest(i.userToken, eventName, indexName)
	req.Filters = filters
	return i.SendEventWithRequestOptions(req, opts)
}

func (i *insightsWithUser) ConvertedObjectIDsAfterSearch(eventName, indexName string, objectIDs []string, queryID string) (res InsightsResponse, err error) {
	return i.ConvertedObjectIDsAfterSearchWithRequestOptions(eventName, indexName, objectIDs, queryID, nil)
}

func (i *insightsWithUser) ConvertedObjectIDsAfterSearchWithRequestOptions(eventName, indexName string, objectIDs []string, queryID string, opts *RequestOptions) (res InsightsResponse, err error) {
	req := conversionInsightsRequest(i.userToken, eventName, indexName)
	req.ObjectIDs = objectIDs
	req.QueryID = queryID
	return i.SendEventWithRequestOptions(req, opts)
}

func (i *insightsWithUser) ViewedFilters(eventName, indexName string, filters []string) (res InsightsResponse, err error) {
	return i.ViewedFiltersWithRequestOptions(eventName, indexName, filters, nil)
}

func (i *insightsWithUser) ViewedFiltersWithRequestOptions(eventName, indexName string, filters []string, opts *RequestOptions) (res InsightsResponse, err error) {
	req := viewInsightsRequest(i.userToken, eventName, indexName)
	req.Filters = filters
	return i.SendEventWithRequestOptions(req, opts)
}

func (i *insightsWithUser) ViewedObjectIDs(eventName, indexName string, objectIDs []string) (res InsightsResponse, err error) {
	return i.ViewedObjectIDsWithRequestOptions(eventName, indexName, objectIDs, nil)
}

func (i *insightsWithUser) ViewedObjectIDsWithRequestOptions(eventName, indexName string, objectIDs []string, opts *RequestOptions) (res InsightsResponse, err error) {
	req := viewInsightsRequest(i.userToken, eventName, indexName)
	req.ObjectIDs = objectIDs
	return i.SendEventWithRequestOptions(req, opts)
}

func (i *insightsWithUser) SendEvent(req InsightsRequest) (res InsightsResponse, err error) {
	return i.SendEventWithRequestOptions(req, nil)
}

func (i *insightsWithUser) SendEventWithRequestOptions(req InsightsRequest, opts *RequestOptions) (res InsightsResponse, err error) {
	return i.SendEventsWithRequestOptions([]InsightsRequest{req}, opts)
}

func (i *insightsWithUser) SendEvents(req []InsightsRequest) (res InsightsResponse, err error) {
	return i.SendEventsWithRequestOptions(req, nil)
}

func (i *insightsWithUser) SendEventsWithRequestOptions(req []InsightsRequest, opts *RequestOptions) (res InsightsResponse, err error) {
	body := map[string]interface{}{
		"events": req,
	}
	err = i.client.request(&res, "POST", i.path, body, insightsCall, opts)
	return
}
