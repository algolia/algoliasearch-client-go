package search

import (
	iopt "github.com/algolia/algoliasearch-client-go/v3/algolia/internal/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
)

type searchRulesParams struct {
	Query       string                  `json:"query"`
	Anchoring   *opt.AnchoringOption    `json:"anchoring,omitempty"`
	Context     *opt.RuleContextsOption `json:"context,omitempty"`
	Page        *opt.PageOption         `json:"page,omitempty"`
	HitsPerPage *opt.HitsPerPageOption  `json:"hitsPerPage,omitempty"`
	Enabled     *opt.EnableRulesOption  `json:"enabled,omitempty"`
}

func newSearchRulesParams(query string, opts ...interface{}) searchRulesParams {
	return searchRulesParams{
		Query:       query,
		Anchoring:   iopt.ExtractAnchoring(opts...),
		Context:     iopt.ExtractRuleContexts(opts...),
		Page:        iopt.ExtractPage(opts...),
		HitsPerPage: iopt.ExtractHitsPerPage(opts...),
		Enabled:     iopt.ExtractEnableRules(opts...),
	}
}
