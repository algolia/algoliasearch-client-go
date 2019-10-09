package opt

import "encoding/json"

// AroundPrecisionOption is a wrapper for an AroundPrecision option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type AroundPrecisionOption struct {
	value []AroundPrecisionRange
}

type AroundPrecisionRange struct {
	From  int `json:"from"`
	Value int `json:"value"`
}

func AroundPrecision(ranges ...AroundPrecisionRange) *AroundPrecisionOption {
	if len(ranges) == 0 {
		return &AroundPrecisionOption{
			value: []AroundPrecisionRange{{From: 0, Value: 1}},
		}
	}
	return &AroundPrecisionOption{
		value: ranges,
	}
}

// Get retrieves the actual value of the option parameter.
func (o *AroundPrecisionOption) Get() (int, []AroundPrecisionRange) {
	if o == nil || len(o.value) == 0 {
		return 1, nil
	}
	if len(o.value) == 1 && o.value[0].From == 0 {
		return o.value[0].Value, nil
	}
	return 0, o.value
}

// MarshalJSON implements the json.Marshaler interface for
// AroundPrecisionOption.
func (o AroundPrecisionOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

// UnmarshalJSON implements the json.Unmarshaler interface for
// AroundPrecisionOption.
func (o *AroundPrecisionOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = []AroundPrecisionRange{{From: 0, Value: 1}}
		return nil
	}

	var i int
	if err := json.Unmarshal(data, &i); err == nil {
		o.value = []AroundPrecisionRange{{From: 0, Value: i}}
		return nil
	}

	return json.Unmarshal(data, &o.value)
}

// Equal returns true if the given option is equal to the instance one. In case
// the given option is nil, we checked the instance one is set to the default
// value of the option.
func (o *AroundPrecisionOption) Equal(o2 *AroundPrecisionOption) bool {
	if o == nil {
		return o2 == nil || len(o2.value) == 1 &&
			o2.value[0].From == 0 &&
			o2.value[0].Value == 1
	}
	if o2 == nil {
		return o == nil || len(o.value) == 1 &&
			o.value[0].From == 0 &&
			o.value[0].Value == 1
	}

	if len(o.value) != len(o2.value) {
		return false
	}

	var found int
	for _, r1 := range o.value {
		for _, r2 := range o2.value {
			if r1.From == r2.From && r1.Value == r2.Value {
				found++
				break
			}
		}
	}
	return found == len(o.value)
}

// AroundPrecisionEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func AroundPrecisionEqual(o1, o2 *AroundPrecisionOption) bool {
	return o1.Equal(o2)
}
