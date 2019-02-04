package opt

import (
	"encoding/json"
)

type AnalyticsTagsOption struct {
	tags []string
}

func AnalyticsTags(tags ...string) AnalyticsTagsOption {
	return AnalyticsTagsOption{tags}
}

func (o AnalyticsTagsOption) Get() []string {
	return o.tags
}

func (o AnalyticsTagsOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.tags)
}

func (o *AnalyticsTagsOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.tags)
}
