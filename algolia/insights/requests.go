package insights

type sendEventsReq struct {
	Events []Event `json:"events"`
}

func newSendEventsReq(events []Event) sendEventsReq {
	return sendEventsReq{Events: events}
}
