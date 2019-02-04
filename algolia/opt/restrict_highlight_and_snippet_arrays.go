package opt

import "encoding/json"

type RestrictHighlightAndSnippetArraysOption struct {
	restrictHighlightAndSnippetArrays bool
}

func RestrictHighlightAndSnippetArrays(restrictHighlightAndSnippetArrays bool) RestrictHighlightAndSnippetArraysOption {
	return RestrictHighlightAndSnippetArraysOption{restrictHighlightAndSnippetArrays}
}

func (o RestrictHighlightAndSnippetArraysOption) Get() bool {
	return o.restrictHighlightAndSnippetArrays
}

func (o RestrictHighlightAndSnippetArraysOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.restrictHighlightAndSnippetArrays)
}

func (o *RestrictHighlightAndSnippetArraysOption) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.restrictHighlightAndSnippetArrays)
}
