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

type ObjectData struct {
	Discount interface{} `json:"discount,omitempty"`
	Price    interface{} `json:"price,omitempty"`
	Quantity int32       `json:"quantity,omitempty"`
	QueryID  *string     `json:"queryID,omitempty"`
}

type Event struct {
	EventName              string       `json:"eventName"`
	EventType              string       `json:"eventType"`
	EventSubtype           string       `json:"eventSubtype,omitempty"`
	Index                  string       `json:"index"`
	ObjectIDs              []string     `json:"objectIDs,omitempty"`
	Positions              []int        `json:"positions,omitempty"`
	QueryID                string       `json:"queryID,omitempty"`
	UserToken              string       `json:"userToken"`
	AuthenticatedUserToken *string      `json:"authenticatedUserToken"`
	Currency               *string      `json:"currency,omitempty"`
	ObjectData             []ObjectData `json:"objectData,omitempty"`
	Timestamp              time.Time    `json:"-"`
	Filters                []string     `json:"filters,omitempty"`
	Value                  interface{}  `json:"value,omitempty"`
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
