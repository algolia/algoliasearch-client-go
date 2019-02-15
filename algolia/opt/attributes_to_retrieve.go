package opt

import "encoding/json"

type AttributesToRetrieveOption struct {
	value []string
}

func AttributesToRetrieve(v []string) AttributesToRetrieveOption {
	return AttributesToRetrieveOption{v}
}

func (o AttributesToRetrieveOption) Get() []string {
	return o.value
}

func (o AttributesToRetrieveOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *AttributesToRetrieveOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = nil
		return nil
	}
	return json.Unmarshal(data, &o.value)
}