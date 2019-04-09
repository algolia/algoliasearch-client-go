package insights

type sendEventsReq struct {
	Events []Event `json:"events"`
}

func newSendEventsReq(events []Event, opts ...interface{}) sendEventsReq {
	return sendEventsReq{Events: events}
}
