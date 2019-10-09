package opt

import (
	"encoding/json"
	"reflect"
)

// RemoveStopWordsOption is a wrapper for an RemoveStopWords option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type RemoveStopWordsOption struct {
	removeStopWords bool
	languages       []string
}

// RemoveStopWords returns an RemoveStopWordsOption whose value is set to the given boolean.
func RemoveStopWords(v bool) *RemoveStopWordsOption {
	return &RemoveStopWordsOption{removeStopWords: v}
}

// RemoveStopWordsFor returns an RemoveStopWordsOption whose value is set to the given list of
// languages.
func RemoveStopWordsFor(languages ...string) *RemoveStopWordsOption {
	return &RemoveStopWordsOption{languages: languages}
}

// Get retrieves the actual value of the option parameter.
func (o *RemoveStopWordsOption) Get() (bool, []string) {
	if o == nil {
		return false, nil
	}
	return o.removeStopWords, o.languages
}

// MarshalJSON implements the json.Marshaler interface for
// RemoveStopWordsOption.
func (o RemoveStopWordsOption) MarshalJSON() ([]byte, error) {
	if len(o.languages) > 0 {
		return json.Marshal(o.languages)
	}
	return json.Marshal(o.removeStopWords)
}

// UnmarshalJSON implements the json.Unmarshaler interface for
// RemoveStopWordsOption.
func (o *RemoveStopWordsOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	if err := json.Unmarshal(data, &o.languages); err == nil {
		return nil
	}

	return json.Unmarshal(data, &o.removeStopWords)
}

// Equal returns true if the given option is equal to the instance one. In case
// the given option is nil, we checked the instance one is set to the default
// value of the option.
func (o *RemoveStopWordsOption) Equal(o2 *RemoveStopWordsOption) bool {
	if o == nil {
		return o2 == nil || !o2.removeStopWords && len(o2.languages) == 0
	}
	if o2 == nil {
		return o == nil || !o.removeStopWords && len(o.languages) == 0
	}
	return reflect.DeepEqual(o, o2)
}

// RemoveStopWordsEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func RemoveStopWordsEqual(o1, o2 *RemoveStopWordsOption) bool {
	return o1.Equal(o2)
}
