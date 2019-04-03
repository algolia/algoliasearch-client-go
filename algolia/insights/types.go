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

type Event struct {
	EventType string    `json:"eventType"`
	EventName string    `json:"eventName"`
	Index     string    `json:"index"`
	UserToken string    `json:"userToken"`
	Timestamp time.Time `json:"-"`
	ObjectIDs []string  `json:"objectIDs,omitempty"`
	Positions []int     `json:"position,omitempty"`
	QueryID   string    `json:"queryID,omitempty"`
	Filters   []string  `json:"filters,omitempty"`
}

func (e Event) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		struct {
			Timestamp int64 `json:"timestamp,omitempty"`
			Event
		}{
			Timestamp: e.Timestamp.Unix(),
			Event:     e,
		},
	)
}
