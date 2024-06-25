package insights

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEvent_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		event    Event
		expected string
	}{
		{
			name: "Click Event",
			event: Event{
				EventType: EventTypeClick,
				EventName: "foo",
				Index:     "indexName",
				UserToken: "bar",
				Timestamp: time.Date(2024, 5, 31, 21, 30, 0, 0, time.UTC),
				ObjectIDs: []string{"one", "two"},
			},
			expected: `{"timestamp":1717191000000,"eventType":"click","eventName":"foo","index":"indexName","userToken":"bar","objectIDs":["one","two"]}`,
		},
		{
			name: "Conversion Event without subtype",
			event: Event{
				EventType: EventTypeConversion,
				EventName: "foo",
				Index:     "indexName",
				UserToken: "bar",
				Timestamp: time.Date(2024, 5, 31, 21, 30, 0, 0, time.UTC),
				ObjectIDs: []string{"one", "two"},
			},
			expected: `{"timestamp":1717191000000,"eventType":"conversion","eventName":"foo","index":"indexName","userToken":"bar","objectIDs":["one","two"]}`,
		},
		{
			name: "Conversion Event - Purchase",
			event: Event{
				EventType:    EventTypeConversion,
				EventSubtype: EventSubtypePurchase,
				EventName:    "foo",
				Index:        "indexName",
				UserToken:    "bar",
				Timestamp:    time.Date(2024, 5, 31, 21, 30, 0, 0, time.UTC),
				ObjectIDs:    []string{"one", "two"},
			},
			expected: `{"timestamp":1717191000000,"eventType":"conversion","eventSubtype":"purchase","eventName":"foo","index":"indexName","userToken":"bar","objectIDs":["one","two"]}`,
		},
		{
			name: "Event with zero timestamp",
			event: Event{
				EventType: EventTypeView,
				EventName: "foo",
				Index:     "indexName",
				UserToken: "bar",
			},
			expected: `{"eventType":"view","eventName":"foo","index":"indexName","userToken":"bar"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bytes, err := tt.event.MarshalJSON()
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expected, string(bytes))
		})
	}
}
