package opt

import (
	"encoding/json"
)

type RestrictSearchableAttributesOption struct {
	attributes []string
}

func RestrictSearchableAttributes(attributes ...string) RestrictSearchableAttributesOption {
	return RestrictSearchableAttributesOption{attributes}
}

func (o RestrictSearchableAttributesOption) Get() []string {
	return o.attributes
}

func (o RestrictSearchableAttributesOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.attributes)
}

func (o *RestrictSearchableAttributesOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.attributes)
}
