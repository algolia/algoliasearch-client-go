package opt

import (
	"encoding/json"
	"time"
)

type ValidUntilOption struct {
	value time.Time
}

func ValidUntil(v time.Time) *ValidUntilOption {
	return &ValidUntilOption{v}
}

func (o *ValidUntilOption) Get() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.value
}

func (o ValidUntilOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *ValidUntilOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = time.Time{}
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

func (o *ValidUntilOption) Equal(o2 *ValidUntilOption) bool {
	if o2 == nil {
		return o.value.IsZero()
	}
	return o.value.Equal(o2.value)
}

func ValidUntilEqual(o1, o2 *ValidUntilOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
