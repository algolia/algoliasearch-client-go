package insights

import (
	"encoding/json"
	"time"
)

const (
	EventTypeClick      = "click"
	EventTypeConversion = "conversion"
	EventTypeView       = "view"
)

const (
	EventSubtypePurchase  = "purchase"
	EventSubtypeAddToCart = "addToCart"
)

type Event struct {
	EventType    string    `json:"eventType"`
	EventSubtype string    `json:"eventSubtype,omitempty"`
	EventName    string    `json:"eventName"`
	Index        string    `json:"index"`
	UserToken    string    `json:"userToken"`
	Timestamp    time.Time `json:"-"`
	ObjectIDs    []string  `json:"objectIDs,omitempty"`
	Positions    []int     `json:"positions,omitempty"`
	QueryID      string    `json:"queryID,omitempty"`
	Filters      []string  `json:"filters,omitempty"`
}

func (e Event) MarshalJSON() ([]byte, error) {
	type EventJSON Event

	var timestamp int64
	if !e.Timestamp.IsZero() {
		timestamp = int64(time.Nanosecond) * e.Timestamp.UnixNano() / int64(time.Millisecond)
	}

	return json.Marshal(
		struct {
			Timestamp int64 `json:"timestamp,omitempty"`
			EventJSON
		}{
			Timestamp: timestamp,
			EventJSON: EventJSON(e),
		},
	)
}
