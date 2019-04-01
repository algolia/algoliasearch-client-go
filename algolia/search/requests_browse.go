package search

type browseReq struct {
	Cursor string `json:"cursor"`
	Params string `json:"params"`
}
