package search

import (
	iopt "github.com/algolia/algoliasearch-client-go/algolia/internal/opt"
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

type deleteByReq struct {
	AroundLatLng      *opt.AroundLatLngOption      `json:"aroundLatLng,omitempty"`
	AroundRadius      *opt.AroundRadiusOption      `json:"aroundRadius,omitempty"`
	FacetFilters      *opt.FacetFiltersOption      `json:"facetFilters,omitempty"`
	Filters           *opt.FiltersOption           `json:"filters,omitempty"`
	InsideBoundingBox *opt.InsideBoundingBoxOption `json:"insideBoundingBox,omitempty"`
	InsidePolygon     *opt.InsidePolygonOption     `json:"insidePolygon,omitempty"`
	NumericFilters    *opt.NumericFiltersOption    `json:"numericFilters,omitempty"`
}

type getObjectsReq struct {
	IndexName            string                          `json:"indexName"`
	ObjectID             string                          `json:"objectID"`
	AttributesToRetrieve *opt.AttributesToRetrieveOption `json:"attributesToRetrieve,omitempty"`
}

type batchReq struct {
	Requests []BatchOperation `json:"requests,omitempty"`
}

type searchReq struct {
	Params string `json:"params"`
}

type QueryParams struct {
	// Attributes
	AttributesToRetrieve         *opt.AttributesToRetrieveOption         `json:"attributesToRetrieve,omitempty"`
	RestrictSearchableAttributes *opt.RestrictSearchableAttributesOption `json:"restrictSearchableAttributes,omitempty"`

	// Filtering
	Filters            *opt.FiltersOption            `json:"filters,omitempty"`
	FacetFilters       *opt.FacetFiltersOption       `json:"facetFilters,omitempty"`
	OptionalFilters    *opt.OptionalFiltersOption    `json:"optionalFilters,omitempty"`
	NumericFilters     *opt.NumericFiltersOption     `json:"numericFilters,omitempty"`
	TagFilters         *opt.TagFiltersOption         `json:"tagFilters,omitempty"`
	SumOrFiltersScores *opt.SumOrFiltersScoresOption `json:"SumOrFiltersScores,omitempty"`

	// Faceting
	Facets                *opt.FacetsOption                `json:"facets,omitempty"`
	MaxValuesPerFacet     *opt.MaxValuesPerFacetOption     `json:"maxValuesPerFacet,omitempty"`
	FacetingAfterDistinct *opt.FacetingAfterDistinctOption `json:"facetingAfterDistinct,omitempty"`
	SortFacetValuesBy     *opt.SortFacetValuesByOption     `json:"sortFacetValuesBy,omitempty"`

	// Highlighting - snippeting
	AttributesToHighlight             *opt.AttributesToHighlightOption             `json:"attributesToHighlight,omitempty"`
	AttributesToSnippet               *opt.AttributesToSnippetOption               `json:"attributesToSnippet,omitempty"`
	HighlightPreTag                   *opt.HighlightPreTagOption                   `json:"highlightPreTag,omitempty"`
	HighlightPostTag                  *opt.HighlightPostTagOption                  `json:"highlightPostTag,omitempty"`
	SnippetEllipsisText               *opt.SnippetEllipsisTextOption               `json:"snippetEllipsisText,omitempty"`
	RestrictHighlightAndSnippetArrays *opt.RestrictHighlightAndSnippetArraysOption `json:"restrictHighlightAndSnippetArrays,omitempty"`

	// Pagination
	Page        *opt.PageOption        `json:"page,omitempty"`
	HitsPerPage *opt.HitsPerPageOption `json:"hitsPerPage,omitempty"`
	Offset      *opt.OffsetOption      `json:"offset,omitempty"`
	Length      *opt.LengthOption      `json:"length,omitempty"`

	// Typos
	MinWordSizeFor1Typo              *opt.MinWordSizeFor1TypoOption              `json:"minWordSizeFor1Typo,omitempty"`
	MinWordSizeFor2Typos             *opt.MinWordSizeFor2TyposOption             `json:"minWordSizeFor2Typos,omitempty"`
	TypoTolerance                    *opt.TypoToleranceOption                    `json:"typoTolerance,omitempty"`
	AllowTyposOnNumericTokens        *opt.AllowTyposOnNumericTokensOption        `json:"allowTyposOnNumericTokens,omitempty"`
	DisableTypoToleranceOnAttributes *opt.DisableTypoToleranceOnAttributesOption `json:"disableTypoToleranceOnAttributes,omitempty"`

	// Geo-search
	AroundLatLng        *opt.AroundLatLngOption        `json:"aroundLatLng,omitempty"`
	AroundLatLngViaIP   *opt.AroundLatLngViaIPOption   `json:"aroundLatLngViaIP,omitempty"`
	AroundRadius        *opt.AroundRadiusOption        `json:"aroundRadius,omitempty"`
	AroundPrecision     *opt.AroundPrecisionOption     `json:"aroundPrecision,omitempty"`
	MinimumAroundRadius *opt.MinimumAroundRadiusOption `json:"minimumAroundRadius,omitempty"`
	InsideBoundingBox   *opt.InsideBoundingBoxOption   `json:"insideBoundingBox,omitempty"`
	InsidePolygon       *opt.InsidePolygonOption       `json:"insidePolygon,omitempty"`

	// Languages
	IgnorePlurals   *opt.IgnorePluralsOption   `json:"ignorePlurals,omitempty"`
	RemoveStopWords *opt.RemoveStopWordsOption `json:"removeStopWords,omitempty"`
	QueryLanguages  *opt.QueryLanguagesOption  `json:"queryLanguages,omitempty"`

	// Query strategy
	QueryType                *opt.QueryTypeOption                `json:"queryType,omitempty"`
	RemoveWordsIfNoResults   *opt.RemoveWordsIfNoResultsOption   `json:"removeWordsIfNoResults,omitempty"`
	AdvancedSyntax           *opt.AdvancedSyntaxOption           `json:"advancedSyntax,omitempty"`
	OptionalWords            *opt.OptionalWordsOption            `json:"optionalWords,omitempty"`
	DisableExactOnAttributes *opt.DisableExactOnAttributesOption `json:"disableExactOnAttributes,omitempty"`
	ExactOnSingleWordQuery   *opt.ExactOnSingleWordQueryOption   `json:"exactOnSingleWordQuery,omitempty"`
	AlternativesAsExact      *opt.AlternativesAsExactOption      `json:"alternativesAsExact,omitempty"`

	// Query rules
	EnableRules  *opt.EnableRulesOption  `json:"enableRules,omitempty"`
	RuleContexts *opt.RuleContextsOption `json:"ruleContexts,omitempty"`

	// Personalization
	EnablePersonalization *opt.EnablePersonalizationOption `json:"enablePersonalization,omitempty"`

	// Advanced
	Distinct                   *opt.DistinctOption                   `json:"distinct,omitempty"`
	GetRankingInfo             *opt.GetRankingInfoOption             `json:"getRankingInfo,omitempty"`
	ClickAnalytics             *opt.ClickAnalyticsOption             `json:"clickAnalytics,omitempty"`
	Analytics                  *opt.AnalyticsOption                  `json:"analytics,omitempty"`
	AnalyticsTags              *opt.AnalyticsTagsOption              `json:"analyticsTags,omitempty"`
	Synonyms                   *opt.SynonymsOption                   `json:"synonyms,omitempty"`
	ReplaceSynonymsInHighlight *opt.ReplaceSynonymsInHighlightOption `json:"replaceSynonymsInHighlight,omitempty"`
	MinProximity               *opt.MinProximityOption               `json:"minProximity,omitempty"`
	ResponseFields             *opt.ResponseFieldsOption             `json:"responseFields,omitempty"`
	MaxFacetHits               *opt.MaxFacetHitsOption               `json:"maxFacetHits,omitempty"`
	PercentileComputation      *opt.PercentileComputationOption      `json:"percentileComputation,omitempty"`
}

func newQueryParams(opts ...interface{}) QueryParams {
	return QueryParams{
		// Attributes
		AttributesToRetrieve:         iopt.ExtractAttributesToRetrieve(opts...),
		RestrictSearchableAttributes: iopt.ExtractRestrictSearchableAttributes(opts...),

		// Filtering
		Filters:            iopt.ExtractFilters(opts...),
		FacetFilters:       iopt.ExtractFacetFilters(opts...),
		OptionalFilters:    iopt.ExtractOptionalFilters(opts...),
		NumericFilters:     iopt.ExtractNumericFilters(opts...),
		TagFilters:         iopt.ExtractTagFilters(opts...),
		SumOrFiltersScores: iopt.ExtractSumOrFiltersScores(opts...),

		// Faceting
		Facets:                iopt.ExtractFacets(opts...),
		MaxValuesPerFacet:     iopt.ExtractMaxValuesPerFacet(opts...),
		FacetingAfterDistinct: iopt.ExtractFacetingAfterDistinct(opts...),
		SortFacetValuesBy:     iopt.ExtractSortFacetValuesBy(opts...),

		// Highlighting - snippeting
		AttributesToHighlight:             iopt.ExtractAttributesToHighlight(opts...),
		AttributesToSnippet:               iopt.ExtractAttributesToSnippet(opts...),
		HighlightPreTag:                   iopt.ExtractHighlightPreTag(opts...),
		HighlightPostTag:                  iopt.ExtractHighlightPostTag(opts...),
		SnippetEllipsisText:               iopt.ExtractSnippetEllipsisText(opts...),
		RestrictHighlightAndSnippetArrays: iopt.ExtractRestrictHighlightAndSnippetArrays(opts...),

		// Pagination
		Page:        iopt.ExtractPage(opts...),
		HitsPerPage: iopt.ExtractHitsPerPage(opts...),
		Offset:      iopt.ExtractOffset(opts...),
		Length:      iopt.ExtractLength(opts...),

		// Typos
		MinWordSizeFor1Typo:              iopt.ExtractMinWordSizeFor1Typo(opts...),
		MinWordSizeFor2Typos:             iopt.ExtractMinWordSizeFor2Typos(opts...),
		TypoTolerance:                    iopt.ExtractTypoTolerance(opts...),
		AllowTyposOnNumericTokens:        iopt.ExtractAllowTyposOnNumericTokens(opts...),
		DisableTypoToleranceOnAttributes: iopt.ExtractDisableTypoToleranceOnAttributes(opts...),

		// Geo-search
		AroundLatLng:        iopt.ExtractAroundLatLng(opts...),
		AroundLatLngViaIP:   iopt.ExtractAroundLatLngViaIP(opts...),
		AroundRadius:        iopt.ExtractAroundRadius(opts...),
		AroundPrecision:     iopt.ExtractAroundPrecision(opts...),
		MinimumAroundRadius: iopt.ExtractMinimumAroundRadius(opts...),
		InsideBoundingBox:   iopt.ExtractInsideBoundingBox(opts...),
		InsidePolygon:       iopt.ExtractInsidePolygon(opts...),

		// Languages
		IgnorePlurals:   iopt.ExtractIgnorePlurals(opts...),
		RemoveStopWords: iopt.ExtractRemoveStopWords(opts...),
		QueryLanguages:  iopt.ExtractQueryLanguages(opts...),

		// Query strategy
		QueryType:                iopt.ExtractQueryType(opts...),
		RemoveWordsIfNoResults:   iopt.ExtractRemoveWordsIfNoResults(opts...),
		AdvancedSyntax:           iopt.ExtractAdvancedSyntax(opts...),
		OptionalWords:            iopt.ExtractOptionalWords(opts...),
		DisableExactOnAttributes: iopt.ExtractDisableExactOnAttributes(opts...),
		ExactOnSingleWordQuery:   iopt.ExtractExactOnSingleWordQuery(opts...),
		AlternativesAsExact:      iopt.ExtractAlternativesAsExact(opts...),

		// Query rules
		EnableRules:  iopt.ExtractEnableRules(opts...),
		RuleContexts: iopt.ExtractRuleContexts(opts...),

		// Personalization
		EnablePersonalization: iopt.ExtractEnablePersonalization(opts...),

		// Advanced
		Distinct:                   iopt.ExtractDistinct(opts...),
		GetRankingInfo:             iopt.ExtractGetRankingInfo(opts...),
		ClickAnalytics:             iopt.ExtractClickAnalytics(opts...),
		Analytics:                  iopt.ExtractAnalytics(opts...),
		AnalyticsTags:              iopt.ExtractAnalyticsTags(opts...),
		Synonyms:                   iopt.ExtractSynonyms(opts...),
		ReplaceSynonymsInHighlight: iopt.ExtractReplaceSynonymsInHighlight(opts...),
		MinProximity:               iopt.ExtractMinProximity(opts...),
		ResponseFields:             iopt.ExtractResponseFields(opts...),
		MaxFacetHits:               iopt.ExtractMaxFacetHits(opts...),
		PercentileComputation:      iopt.ExtractPercentileComputation(opts...),
	}
}

type searchParams struct {
	Query string `json:"query"`
	QueryParams
}

func newSearchParams(query string, opts ...interface{}) searchParams {
	return searchParams{
		Query:       query,
		QueryParams: newQueryParams(opts...),
	}
}

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
