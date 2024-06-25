package insights

type UserClient struct {
	UserToken string
	Client
}

func (c *UserClient) ClickedObjectIDs(
	eventName string,
	indexName string,
	objectIDs []string,
	opts ...interface{},
) (StatusMessageRes, error) {
	return c.Client.SendEvent(Event{
		UserToken: c.UserToken,
		EventType: EventTypeClick,
		EventName: eventName,
		Index:     indexName,
		ObjectIDs: objectIDs,
	}, opts...)
}

func (c *UserClient) ClickedObjectIDsAfterSearch(
	eventName string,
	indexName string,
	objectIDs []string,
	positions []int,
	queryID string,
	opts ...interface{},
) (StatusMessageRes, error) {
	return c.Client.SendEvent(Event{
		UserToken: c.UserToken,
		EventType: EventTypeClick,
		EventName: eventName,
		Index:     indexName,
		ObjectIDs: objectIDs,
		Positions: positions,
		QueryID:   queryID,
	}, opts...)
}

func (c *UserClient) ClickedFilters(
	eventName string,
	indexName string,
	filters []string,
	opts ...interface{},
) (res StatusMessageRes, err error) {
	return c.Client.SendEvent(Event{
		UserToken: c.UserToken,
		EventType: EventTypeClick,
		EventName: eventName,
		Index:     indexName,
		Filters:   filters,
	}, opts...)
}

func (c *UserClient) ConvertedObjectIDs(
	eventName string,
	indexName string,
	objectIDs []string,
	opts ...interface{},
) (StatusMessageRes, error) {
	return c.Client.SendEvent(Event{
		UserToken: c.UserToken,
		EventType: EventTypeConversion,
		EventName: eventName,
		Index:     indexName,
		ObjectIDs: objectIDs,
	}, opts...)
}

func (c *UserClient) ConvertedObjectIDsPurchase(
	eventName string,
	indexName string,
	objectIDs []string,
	opts ...interface{},
) (StatusMessageRes, error) {
	return c.Client.SendEvent(Event{
		UserToken:    c.UserToken,
		EventType:    EventTypeConversion,
		EventSubtype: EventSubtypePurchase,
		EventName:    eventName,
		Index:        indexName,
		ObjectIDs:    objectIDs,
	}, opts...)
}

func (c *UserClient) ConvertedObjectIDsAddToCart(
	eventName string,
	indexName string,
	objectIDs []string,
	opts ...interface{},
) (StatusMessageRes, error) {
	return c.Client.SendEvent(Event{
		UserToken:    c.UserToken,
		EventType:    EventTypeConversion,
		EventSubtype: EventSubtypeAddToCart,
		EventName:    eventName,
		Index:        indexName,
		ObjectIDs:    objectIDs,
	}, opts...)
}

func (c *UserClient) ConvertedObjectIDsAfterSearch(
	eventName string,
	indexName string,
	objectIDs []string,
	queryID string,
	opts ...interface{},
) (StatusMessageRes, error) {
	return c.Client.SendEvent(Event{
		UserToken: c.UserToken,
		EventType: EventTypeConversion,
		EventName: eventName,
		Index:     indexName,
		ObjectIDs: objectIDs,
		QueryID:   queryID,
	}, opts...)
}

func (c *UserClient) ConvertedObjectIDsAfterSearchPurchase(
	eventName string,
	indexName string,
	objectIDs []string,
	queryID string,
	opts ...interface{},
) (StatusMessageRes, error) {
	return c.Client.SendEvent(Event{
		UserToken:    c.UserToken,
		EventType:    EventTypeConversion,
		EventSubtype: EventSubtypePurchase,
		EventName:    eventName,
		Index:        indexName,
		ObjectIDs:    objectIDs,
		QueryID:      queryID,
	}, opts...)
}

func (c *UserClient) ConvertedObjectIDsAfterSearchAddToCart(
	eventName string,
	indexName string,
	objectIDs []string,
	queryID string,
	opts ...interface{},
) (StatusMessageRes, error) {
	return c.Client.SendEvent(Event{
		UserToken:    c.UserToken,
		EventType:    EventTypeConversion,
		EventSubtype: EventSubtypeAddToCart,
		EventName:    eventName,
		Index:        indexName,
		ObjectIDs:    objectIDs,
		QueryID:      queryID,
	}, opts...)
}

func (c *UserClient) ConvertedFilters(
	eventName string,
	indexName string,
	filters []string,
	opts ...interface{},
) (StatusMessageRes, error) {
	return c.Client.SendEvent(Event{
		UserToken: c.UserToken,
		EventType: EventTypeConversion,
		EventName: eventName,
		Index:     indexName,
		Filters:   filters,
	}, opts...)
}

func (c *UserClient) ConvertedFiltersPurchase(
	eventName string,
	indexName string,
	filters []string,
	opts ...interface{},
) (StatusMessageRes, error) {
	return c.Client.SendEvent(Event{
		UserToken:    c.UserToken,
		EventType:    EventTypeConversion,
		EventSubtype: EventSubtypePurchase,
		EventName:    eventName,
		Index:        indexName,
		Filters:      filters,
	}, opts...)
}

func (c *UserClient) ConvertedFiltersAddToCart(
	eventName string,
	indexName string,
	filters []string,
	opts ...interface{},
) (StatusMessageRes, error) {
	return c.Client.SendEvent(Event{
		UserToken:    c.UserToken,
		EventType:    EventTypeConversion,
		EventSubtype: EventSubtypeAddToCart,
		EventName:    eventName,
		Index:        indexName,
		Filters:      filters,
	}, opts...)
}

func (c *UserClient) ViewedObjectIDs(
	eventName string,
	indexName string,
	objectIDs []string,
	opts ...interface{},
) (StatusMessageRes, error) {
	return c.Client.SendEvent(Event{
		UserToken: c.UserToken,
		EventType: EventTypeView,
		EventName: eventName,
		Index:     indexName,
		ObjectIDs: objectIDs,
	}, opts...)
}

func (c *UserClient) ViewedFilters(
	eventName string,
	indexName string,
	filters []string,
	opts ...interface{},
) (StatusMessageRes, error) {
	return c.Client.SendEvent(Event{
		UserToken: c.UserToken,
		EventType: EventTypeView,
		EventName: eventName,
		Index:     indexName,
		Filters:   filters,
	}, opts...)
}
