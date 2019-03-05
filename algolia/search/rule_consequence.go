package search

type RuleConsequence struct {
	Params   *RuleParams      `json:"params,omitempty"`
	Promote  []PromotedObject `json:"promote,omitempty"`
	Hide     []HiddenObject   `json:"hide,omitempty"`
	UserData interface{}      `json:"userData,omitempty"`
}

type PromotedObject struct {
	ObjectID string `json:"objectID"`
	Position int    `json:"position"`
}

type HiddenObject struct {
	ObjectID string `json:"objectID"`
}
