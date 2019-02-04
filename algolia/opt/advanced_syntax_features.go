package opt

import (
	"encoding/json"
)

type AdvancedSyntaxFeaturesOption struct {
	features []string
}

func AdvancedSyntaxFeatures(features ...string) AdvancedSyntaxFeaturesOption {
	return AdvancedSyntaxFeaturesOption{features}
}

func (o AdvancedSyntaxFeaturesOption) Get() []string {
	return o.features
}

func (o AdvancedSyntaxFeaturesOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.features)
}

func (o *AdvancedSyntaxFeaturesOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.features)
}
