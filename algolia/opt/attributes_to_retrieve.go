package opt

import (
	"encoding/json"
)

type AttributesToRetrieveOption struct {
	attributes []string
}

func AttributesToRetrieve(attributes ...string) AttributesToRetrieveOption {
	return AttributesToRetrieveOption{attributes}
}

func (o AttributesToRetrieveOption) Get() []string {
	return o.attributes
}

func (o AttributesToRetrieveOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.attributes)
}

func (o *AttributesToRetrieveOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.attributes)
}
