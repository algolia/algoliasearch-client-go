package search

type browseRes struct {
	Cursor  string `json:"cursor"`
	Warning string `json:"warning"`
	QueryRes
}
