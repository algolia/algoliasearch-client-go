package opt

// AttributesToHighlightEmptyList is a workaround function to set empty array into a AttributesToHighlightOption.
// since passing empty params to `opt.AttributesToHighlight()` will result in `null`
func AttributesToHighlightEmptyList() *AttributesToHighlightOption {
	return &AttributesToHighlightOption{[]string{}}
}
