package search

import "github.com/algolia/algoliasearch-client-go/v3/algolia/opt"

type RuleConsequence struct {
	Params         *RuleParams               `json:"params,omitempty"`
	Promote        []PromotedObject          `json:"promote,omitempty"`
	FilterPromotes *opt.FilterPromotesOption `json:"filterPromotes,omitempty"`
	Hide           []HiddenObject            `json:"hide,omitempty"`
	UserData       interface{}               `json:"userData,omitempty"`
	Redirect       *RuleRedirect             `json:"redirect,omitempty"`
}

type RuleRedirect struct {
	IndexName string `json:"indexName,omitempty"`
}

type PromotedObject struct {
	ObjectID  string   `json:"objectID,omitempty"`
	ObjectIDs []string `json:"objectIDs,omitempty"`
	Position  int      `json:"position"`
}

type HiddenObject struct {
	ObjectID string `json:"objectID"`
}
