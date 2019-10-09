package opt

import "encoding/json"

// SnippetEllipsisTextOption is a wrapper for an SnippetEllipsisText option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type SnippetEllipsisTextOption struct {
	value string
}

const (
	oldDefaultSnippetEllipsisTextValue = ""
	newDefaultSnippetEllipsisTextValue = "…"
)

// SnippetEllipsisText wraps the given value into a SnippetEllipsisTextOption.
func SnippetEllipsisText(v string) *SnippetEllipsisTextOption {
	return &SnippetEllipsisTextOption{v}
}

// Get retrieves the actual value of the option parameter.
func (o *SnippetEllipsisTextOption) Get() string {
	if o == nil {
		return newDefaultSnippetEllipsisTextValue
	}
	return o.value
}

// MarshalJSON implements the json.Marshaler interface for
// SnippetEllipsisTextOption.
func (o SnippetEllipsisTextOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

// UnmarshalJSON implements the json.Unmarshaler interface for
// SnippetEllipsisTextOption.
func (o *SnippetEllipsisTextOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = newDefaultSnippetEllipsisTextValue
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

// Equal returns true if the given option is equal to the instance one. In case
// the given option is nil, we checked the instance one is set to the default
// value of the option.
func (o *SnippetEllipsisTextOption) Equal(o2 *SnippetEllipsisTextOption) bool {
	// We cannot generate this setting using `go generate` as the default
	// value can either be "…" or "". From the documentation:
	//   * Defaults to an empty string for all accounts created before
	//     February 10th, 2016.
	//   * Defaults to "…" (U+2026, HORIZONTAL ELLIPSIS) for accounts
	//     created after that date.
	if o == nil {
		return o2 == nil || o2.value == newDefaultSnippetEllipsisTextValue || o2.value == oldDefaultSnippetEllipsisTextValue
	}
	if o2 == nil {
		return o == nil || o.value == newDefaultSnippetEllipsisTextValue || o.value == oldDefaultSnippetEllipsisTextValue
	}
	return o.value == o2.value
}

// SnippetEllipsisTextEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func SnippetEllipsisTextEqual(o1, o2 *SnippetEllipsisTextOption) bool {
	return o1.Equal(o2)
}
