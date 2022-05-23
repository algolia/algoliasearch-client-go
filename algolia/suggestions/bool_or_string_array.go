package suggestions

import "encoding/json"

// BoolOrStringArray can hold either a bool or a []string.
// IsBool indicates the valid value: Bool or StringArray.
type BoolOrStringArray struct {
	IsBool      bool
	Bool        bool
	StringArray []string
}

func NewBool(boolValue bool) BoolOrStringArray {
	return BoolOrStringArray{
		IsBool:      true,
		Bool:        boolValue,
		StringArray: []string{},
	}
}

func NewStringArray(array []string) BoolOrStringArray {
	return BoolOrStringArray{
		IsBool:      false,
		Bool:        false,
		StringArray: array,
	}
}

// UnmarshalJSON decodes a JSON bool or string array into a BoolOrStringArray.
func (b *BoolOrStringArray) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &b.Bool)
	if err == nil {
		b.IsBool = true
		b.StringArray = []string{}
		return nil
	}
	err = json.Unmarshal(data, &b.StringArray)
	if err == nil {
		b.IsBool = false
		b.Bool = false
		return nil
	}
	return err
}

// MarshalJSON encodes a BoolOrStringArray to a JSON bool or string array.
func (b BoolOrStringArray) MarshalJSON() ([]byte, error) {
	if b.IsBool {
		return json.Marshal(b.Bool)
	}
	return json.Marshal(b.StringArray)
}
