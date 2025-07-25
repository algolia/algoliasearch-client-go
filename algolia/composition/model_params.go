// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package composition

import (
	"encoding/json"
	"fmt"

	"github.com/algolia/algoliasearch-client-go/v4/algolia/utils"
)

// Params struct for Params.
type Params struct {
	// Search query.
	Query *string `json:"query,omitempty"`
	// Filter expression to only include items that match the filter criteria in the response.  You can use these filter expressions:  - **Numeric filters.** `<facet> <op> <number>`, where `<op>` is one of `<`, `<=`, `=`, `!=`, `>`, `>=`. - **Ranges.** `<facet>:<lower> TO <upper>` where `<lower>` and `<upper>` are the lower and upper limits of the range (inclusive). - **Facet filters.** `<facet>:<value>` where `<facet>` is a facet attribute (case-sensitive) and `<value>` a facet value. - **Tag filters.** `_tags:<value>` or just `<value>` (case-sensitive). - **Boolean filters.** `<facet>: true | false`.  You can combine filters with `AND`, `OR`, and `NOT` operators with the following restrictions:  - You can only combine filters of the same type with `OR`.   **Not supported:** `facet:value OR num > 3`. - You can't use `NOT` with combinations of filters.   **Not supported:** `NOT(facet:value OR facet:value)` - You can't combine conjunctions (`AND`) with `OR`.   **Not supported:** `facet:value OR (facet:value AND facet:value)`  Use quotes around your filters, if the facet attribute name or facet value has spaces, keywords (`OR`, `AND`, `NOT`), or quotes. If a facet attribute is an array, the filter matches if it matches at least one element of the array.  For more information, see [Filters](https://www.algolia.com/doc/guides/managing-results/refine-results/filtering/).
	Filters *string `json:"filters,omitempty"`
	// Page of search results to retrieve.
	Page *int32 `json:"page,omitempty"`
	// Whether the run response should include detailed ranking information.
	GetRankingInfo      *bool            `json:"getRankingInfo,omitempty"`
	RelevancyStrictness *int32           `json:"relevancyStrictness,omitempty"`
	FacetFilters        *FacetFilters    `json:"facetFilters,omitempty"`
	OptionalFilters     *OptionalFilters `json:"optionalFilters,omitempty"`
	NumericFilters      *NumericFilters  `json:"numericFilters,omitempty"`
	// Number of hits per page.
	HitsPerPage *int32 `json:"hitsPerPage,omitempty"`
	// Coordinates for the center of a circle, expressed as a comma-separated string of latitude and longitude.  Only records included within a circle around this central location are included in the results. The radius of the circle is determined by the `aroundRadius` and `minimumAroundRadius` settings. This parameter is ignored if you also specify `insidePolygon` or `insideBoundingBox`.
	AroundLatLng *string `json:"aroundLatLng,omitempty"`
	// Whether to obtain the coordinates from the request's IP address.
	AroundLatLngViaIP *bool            `json:"aroundLatLngViaIP,omitempty"`
	AroundRadius      *AroundRadius    `json:"aroundRadius,omitempty"`
	AroundPrecision   *AroundPrecision `json:"aroundPrecision,omitempty"`
	// Minimum radius (in meters) for a search around a location when `aroundRadius` isn't set.
	MinimumAroundRadius *int32                            `json:"minimumAroundRadius,omitempty"`
	InsideBoundingBox   utils.Nullable[InsideBoundingBox] `json:"insideBoundingBox,omitempty"`
	// Coordinates of a polygon in which to search.  Polygons are defined by 3 to 10,000 points. Each point is represented by its latitude and longitude. Provide multiple polygons as nested arrays. For more information, see [filtering inside polygons](https://www.algolia.com/doc/guides/managing-results/refine-results/geolocation/#filtering-inside-rectangular-or-polygonal-areas). This parameter is ignored if you also specify `insideBoundingBox`.
	InsidePolygon [][]float64 `json:"insidePolygon,omitempty"`
	// Languages for language-specific query processing steps such as plurals, stop-word removal, and word-detection dictionaries This setting sets a default list of languages used by the `removeStopWords` and `ignorePlurals` settings. This setting also sets a dictionary for word detection in the logogram-based [CJK](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/normalization/#normalization-for-logogram-based-languages-cjk) languages. To support this, you must place the CJK language **first** **You should always specify a query language.** If you don't specify an indexing language, the search engine uses all [supported languages](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/supported-languages/), or the languages you specified with the `ignorePlurals` or `removeStopWords` parameters. This can lead to unexpected search results. For more information, see [Language-specific configuration](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/language-specific-configurations/).
	QueryLanguages []SupportedLanguage `json:"queryLanguages,omitempty"`
	// ISO language codes that adjust settings that are useful for processing natural language queries (as opposed to keyword searches) - Sets `removeStopWords` and `ignorePlurals` to the list of provided languages. - Sets `removeWordsIfNoResults` to `allOptional`. - Adds a `natural_language` attribute to `ruleContexts` and `analyticsTags`.
	NaturalLanguages []SupportedLanguage `json:"naturalLanguages,omitempty"`
	// Whether to enable composition rules.
	EnableRules *bool `json:"enableRules,omitempty"`
	// Assigns a rule context to the run query [Rule contexts](https://www.algolia.com/doc/guides/managing-results/rules/rules-overview/how-to/customize-search-results-by-platform/#whats-a-context) are strings that you can use to trigger matching rules.
	RuleContexts []string `json:"ruleContexts,omitempty"`
	// Unique pseudonymous or anonymous user identifier.  This helps with analytics and click and conversion events. For more information, see [user token](https://www.algolia.com/doc/guides/sending-events/concepts/usertoken/).
	UserToken *string `json:"userToken,omitempty"`
	// Whether to include a `queryID` attribute in the response The query ID is a unique identifier for a search query and is required for tracking [click and conversion events](https://www.algolia.com/guides/sending-events/getting-started/).
	ClickAnalytics *bool `json:"clickAnalytics,omitempty"`
	// Whether this search will be included in Analytics.
	Analytics *bool `json:"analytics,omitempty"`
	// Tags to apply to the query for [segmenting analytics data](https://www.algolia.com/doc/guides/search-analytics/guides/segments/).
	AnalyticsTags []string `json:"analyticsTags,omitempty"`
	// Whether to enable index level A/B testing for this run request. If the composition mixes multiple indices, the A/B test is ignored.
	EnableABTest *bool `json:"enableABTest,omitempty"`
	// Whether this search will use [Dynamic Re-Ranking](https://www.algolia.com/doc/guides/algolia-ai/re-ranking/) This setting only has an effect if you activated Dynamic Re-Ranking for this index in the Algolia dashboard.
	EnableReRanking *bool `json:"enableReRanking,omitempty"`
}

type ParamsOption func(f *Params)

func WithParamsQuery(val string) ParamsOption {
	return func(f *Params) {
		f.Query = &val
	}
}

func WithParamsFilters(val string) ParamsOption {
	return func(f *Params) {
		f.Filters = &val
	}
}

func WithParamsPage(val int32) ParamsOption {
	return func(f *Params) {
		f.Page = &val
	}
}

func WithParamsGetRankingInfo(val bool) ParamsOption {
	return func(f *Params) {
		f.GetRankingInfo = &val
	}
}

func WithParamsRelevancyStrictness(val int32) ParamsOption {
	return func(f *Params) {
		f.RelevancyStrictness = &val
	}
}

func WithParamsFacetFilters(val FacetFilters) ParamsOption {
	return func(f *Params) {
		f.FacetFilters = &val
	}
}

func WithParamsOptionalFilters(val OptionalFilters) ParamsOption {
	return func(f *Params) {
		f.OptionalFilters = &val
	}
}

func WithParamsNumericFilters(val NumericFilters) ParamsOption {
	return func(f *Params) {
		f.NumericFilters = &val
	}
}

func WithParamsHitsPerPage(val int32) ParamsOption {
	return func(f *Params) {
		f.HitsPerPage = &val
	}
}

func WithParamsAroundLatLng(val string) ParamsOption {
	return func(f *Params) {
		f.AroundLatLng = &val
	}
}

func WithParamsAroundLatLngViaIP(val bool) ParamsOption {
	return func(f *Params) {
		f.AroundLatLngViaIP = &val
	}
}

func WithParamsAroundRadius(val AroundRadius) ParamsOption {
	return func(f *Params) {
		f.AroundRadius = &val
	}
}

func WithParamsAroundPrecision(val AroundPrecision) ParamsOption {
	return func(f *Params) {
		f.AroundPrecision = &val
	}
}

func WithParamsMinimumAroundRadius(val int32) ParamsOption {
	return func(f *Params) {
		f.MinimumAroundRadius = &val
	}
}

func WithParamsInsideBoundingBox(val utils.Nullable[InsideBoundingBox]) ParamsOption {
	return func(f *Params) {
		f.InsideBoundingBox = val
	}
}

func WithParamsInsidePolygon(val [][]float64) ParamsOption {
	return func(f *Params) {
		f.InsidePolygon = val
	}
}

func WithParamsQueryLanguages(val []SupportedLanguage) ParamsOption {
	return func(f *Params) {
		f.QueryLanguages = val
	}
}

func WithParamsNaturalLanguages(val []SupportedLanguage) ParamsOption {
	return func(f *Params) {
		f.NaturalLanguages = val
	}
}

func WithParamsEnableRules(val bool) ParamsOption {
	return func(f *Params) {
		f.EnableRules = &val
	}
}

func WithParamsRuleContexts(val []string) ParamsOption {
	return func(f *Params) {
		f.RuleContexts = val
	}
}

func WithParamsUserToken(val string) ParamsOption {
	return func(f *Params) {
		f.UserToken = &val
	}
}

func WithParamsClickAnalytics(val bool) ParamsOption {
	return func(f *Params) {
		f.ClickAnalytics = &val
	}
}

func WithParamsAnalytics(val bool) ParamsOption {
	return func(f *Params) {
		f.Analytics = &val
	}
}

func WithParamsAnalyticsTags(val []string) ParamsOption {
	return func(f *Params) {
		f.AnalyticsTags = val
	}
}

func WithParamsEnableABTest(val bool) ParamsOption {
	return func(f *Params) {
		f.EnableABTest = &val
	}
}

func WithParamsEnableReRanking(val bool) ParamsOption {
	return func(f *Params) {
		f.EnableReRanking = &val
	}
}

// NewParams instantiates a new Params object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParams(opts ...ParamsOption) *Params {
	this := &Params{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyParams return a pointer to an empty Params object.
func NewEmptyParams() *Params {
	return &Params{}
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *Params) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Params) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *Params) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *Params) SetQuery(v string) *Params {
	o.Query = &v
	return o
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *Params) GetFilters() string {
	if o == nil || o.Filters == nil {
		var ret string
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Params) GetFiltersOk() (*string, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *Params) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given string and assigns it to the Filters field.
func (o *Params) SetFilters(v string) *Params {
	o.Filters = &v
	return o
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *Params) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Params) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *Params) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *Params) SetPage(v int32) *Params {
	o.Page = &v
	return o
}

// GetGetRankingInfo returns the GetRankingInfo field value if set, zero value otherwise.
func (o *Params) GetGetRankingInfo() bool {
	if o == nil || o.GetRankingInfo == nil {
		var ret bool
		return ret
	}
	return *o.GetRankingInfo
}

// GetGetRankingInfoOk returns a tuple with the GetRankingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Params) GetGetRankingInfoOk() (*bool, bool) {
	if o == nil || o.GetRankingInfo == nil {
		return nil, false
	}
	return o.GetRankingInfo, true
}

// HasGetRankingInfo returns a boolean if a field has been set.
func (o *Params) HasGetRankingInfo() bool {
	if o != nil && o.GetRankingInfo != nil {
		return true
	}

	return false
}

// SetGetRankingInfo gets a reference to the given bool and assigns it to the GetRankingInfo field.
func (o *Params) SetGetRankingInfo(v bool) *Params {
	o.GetRankingInfo = &v
	return o
}

// GetRelevancyStrictness returns the RelevancyStrictness field value if set, zero value otherwise.
func (o *Params) GetRelevancyStrictness() int32 {
	if o == nil || o.RelevancyStrictness == nil {
		var ret int32
		return ret
	}
	return *o.RelevancyStrictness
}

// GetRelevancyStrictnessOk returns a tuple with the RelevancyStrictness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Params) GetRelevancyStrictnessOk() (*int32, bool) {
	if o == nil || o.RelevancyStrictness == nil {
		return nil, false
	}
	return o.RelevancyStrictness, true
}

// HasRelevancyStrictness returns a boolean if a field has been set.
func (o *Params) HasRelevancyStrictness() bool {
	if o != nil && o.RelevancyStrictness != nil {
		return true
	}

	return false
}

// SetRelevancyStrictness gets a reference to the given int32 and assigns it to the RelevancyStrictness field.
func (o *Params) SetRelevancyStrictness(v int32) *Params {
	o.RelevancyStrictness = &v
	return o
}

// GetFacetFilters returns the FacetFilters field value if set, zero value otherwise.
func (o *Params) GetFacetFilters() FacetFilters {
	if o == nil || o.FacetFilters == nil {
		var ret FacetFilters
		return ret
	}
	return *o.FacetFilters
}

// GetFacetFiltersOk returns a tuple with the FacetFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Params) GetFacetFiltersOk() (*FacetFilters, bool) {
	if o == nil || o.FacetFilters == nil {
		return nil, false
	}
	return o.FacetFilters, true
}

// HasFacetFilters returns a boolean if a field has been set.
func (o *Params) HasFacetFilters() bool {
	if o != nil && o.FacetFilters != nil {
		return true
	}

	return false
}

// SetFacetFilters gets a reference to the given FacetFilters and assigns it to the FacetFilters field.
func (o *Params) SetFacetFilters(v *FacetFilters) *Params {
	o.FacetFilters = v
	return o
}

// GetOptionalFilters returns the OptionalFilters field value if set, zero value otherwise.
func (o *Params) GetOptionalFilters() OptionalFilters {
	if o == nil || o.OptionalFilters == nil {
		var ret OptionalFilters
		return ret
	}
	return *o.OptionalFilters
}

// GetOptionalFiltersOk returns a tuple with the OptionalFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Params) GetOptionalFiltersOk() (*OptionalFilters, bool) {
	if o == nil || o.OptionalFilters == nil {
		return nil, false
	}
	return o.OptionalFilters, true
}

// HasOptionalFilters returns a boolean if a field has been set.
func (o *Params) HasOptionalFilters() bool {
	if o != nil && o.OptionalFilters != nil {
		return true
	}

	return false
}

// SetOptionalFilters gets a reference to the given OptionalFilters and assigns it to the OptionalFilters field.
func (o *Params) SetOptionalFilters(v *OptionalFilters) *Params {
	o.OptionalFilters = v
	return o
}

// GetNumericFilters returns the NumericFilters field value if set, zero value otherwise.
func (o *Params) GetNumericFilters() NumericFilters {
	if o == nil || o.NumericFilters == nil {
		var ret NumericFilters
		return ret
	}
	return *o.NumericFilters
}

// GetNumericFiltersOk returns a tuple with the NumericFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Params) GetNumericFiltersOk() (*NumericFilters, bool) {
	if o == nil || o.NumericFilters == nil {
		return nil, false
	}
	return o.NumericFilters, true
}

// HasNumericFilters returns a boolean if a field has been set.
func (o *Params) HasNumericFilters() bool {
	if o != nil && o.NumericFilters != nil {
		return true
	}

	return false
}

// SetNumericFilters gets a reference to the given NumericFilters and assigns it to the NumericFilters field.
func (o *Params) SetNumericFilters(v *NumericFilters) *Params {
	o.NumericFilters = v
	return o
}

// GetHitsPerPage returns the HitsPerPage field value if set, zero value otherwise.
func (o *Params) GetHitsPerPage() int32 {
	if o == nil || o.HitsPerPage == nil {
		var ret int32
		return ret
	}
	return *o.HitsPerPage
}

// GetHitsPerPageOk returns a tuple with the HitsPerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Params) GetHitsPerPageOk() (*int32, bool) {
	if o == nil || o.HitsPerPage == nil {
		return nil, false
	}
	return o.HitsPerPage, true
}

// HasHitsPerPage returns a boolean if a field has been set.
func (o *Params) HasHitsPerPage() bool {
	if o != nil && o.HitsPerPage != nil {
		return true
	}

	return false
}

// SetHitsPerPage gets a reference to the given int32 and assigns it to the HitsPerPage field.
func (o *Params) SetHitsPerPage(v int32) *Params {
	o.HitsPerPage = &v
	return o
}

// GetAroundLatLng returns the AroundLatLng field value if set, zero value otherwise.
func (o *Params) GetAroundLatLng() string {
	if o == nil || o.AroundLatLng == nil {
		var ret string
		return ret
	}
	return *o.AroundLatLng
}

// GetAroundLatLngOk returns a tuple with the AroundLatLng field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Params) GetAroundLatLngOk() (*string, bool) {
	if o == nil || o.AroundLatLng == nil {
		return nil, false
	}
	return o.AroundLatLng, true
}

// HasAroundLatLng returns a boolean if a field has been set.
func (o *Params) HasAroundLatLng() bool {
	if o != nil && o.AroundLatLng != nil {
		return true
	}

	return false
}

// SetAroundLatLng gets a reference to the given string and assigns it to the AroundLatLng field.
func (o *Params) SetAroundLatLng(v string) *Params {
	o.AroundLatLng = &v
	return o
}

// GetAroundLatLngViaIP returns the AroundLatLngViaIP field value if set, zero value otherwise.
func (o *Params) GetAroundLatLngViaIP() bool {
	if o == nil || o.AroundLatLngViaIP == nil {
		var ret bool
		return ret
	}
	return *o.AroundLatLngViaIP
}

// GetAroundLatLngViaIPOk returns a tuple with the AroundLatLngViaIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Params) GetAroundLatLngViaIPOk() (*bool, bool) {
	if o == nil || o.AroundLatLngViaIP == nil {
		return nil, false
	}
	return o.AroundLatLngViaIP, true
}

// HasAroundLatLngViaIP returns a boolean if a field has been set.
func (o *Params) HasAroundLatLngViaIP() bool {
	if o != nil && o.AroundLatLngViaIP != nil {
		return true
	}

	return false
}

// SetAroundLatLngViaIP gets a reference to the given bool and assigns it to the AroundLatLngViaIP field.
func (o *Params) SetAroundLatLngViaIP(v bool) *Params {
	o.AroundLatLngViaIP = &v
	return o
}

// GetAroundRadius returns the AroundRadius field value if set, zero value otherwise.
func (o *Params) GetAroundRadius() AroundRadius {
	if o == nil || o.AroundRadius == nil {
		var ret AroundRadius
		return ret
	}
	return *o.AroundRadius
}

// GetAroundRadiusOk returns a tuple with the AroundRadius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Params) GetAroundRadiusOk() (*AroundRadius, bool) {
	if o == nil || o.AroundRadius == nil {
		return nil, false
	}
	return o.AroundRadius, true
}

// HasAroundRadius returns a boolean if a field has been set.
func (o *Params) HasAroundRadius() bool {
	if o != nil && o.AroundRadius != nil {
		return true
	}

	return false
}

// SetAroundRadius gets a reference to the given AroundRadius and assigns it to the AroundRadius field.
func (o *Params) SetAroundRadius(v *AroundRadius) *Params {
	o.AroundRadius = v
	return o
}

// GetAroundPrecision returns the AroundPrecision field value if set, zero value otherwise.
func (o *Params) GetAroundPrecision() AroundPrecision {
	if o == nil || o.AroundPrecision == nil {
		var ret AroundPrecision
		return ret
	}
	return *o.AroundPrecision
}

// GetAroundPrecisionOk returns a tuple with the AroundPrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Params) GetAroundPrecisionOk() (*AroundPrecision, bool) {
	if o == nil || o.AroundPrecision == nil {
		return nil, false
	}
	return o.AroundPrecision, true
}

// HasAroundPrecision returns a boolean if a field has been set.
func (o *Params) HasAroundPrecision() bool {
	if o != nil && o.AroundPrecision != nil {
		return true
	}

	return false
}

// SetAroundPrecision gets a reference to the given AroundPrecision and assigns it to the AroundPrecision field.
func (o *Params) SetAroundPrecision(v *AroundPrecision) *Params {
	o.AroundPrecision = v
	return o
}

// GetMinimumAroundRadius returns the MinimumAroundRadius field value if set, zero value otherwise.
func (o *Params) GetMinimumAroundRadius() int32 {
	if o == nil || o.MinimumAroundRadius == nil {
		var ret int32
		return ret
	}
	return *o.MinimumAroundRadius
}

// GetMinimumAroundRadiusOk returns a tuple with the MinimumAroundRadius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Params) GetMinimumAroundRadiusOk() (*int32, bool) {
	if o == nil || o.MinimumAroundRadius == nil {
		return nil, false
	}
	return o.MinimumAroundRadius, true
}

// HasMinimumAroundRadius returns a boolean if a field has been set.
func (o *Params) HasMinimumAroundRadius() bool {
	if o != nil && o.MinimumAroundRadius != nil {
		return true
	}

	return false
}

// SetMinimumAroundRadius gets a reference to the given int32 and assigns it to the MinimumAroundRadius field.
func (o *Params) SetMinimumAroundRadius(v int32) *Params {
	o.MinimumAroundRadius = &v
	return o
}

// GetInsideBoundingBox returns the InsideBoundingBox field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Params) GetInsideBoundingBox() InsideBoundingBox {
	if o == nil || o.InsideBoundingBox.Get() == nil {
		var ret InsideBoundingBox
		return ret
	}
	return *o.InsideBoundingBox.Get()
}

// GetInsideBoundingBoxOk returns a tuple with the InsideBoundingBox field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Params) GetInsideBoundingBoxOk() (*InsideBoundingBox, bool) {
	if o == nil {
		return nil, false
	}
	return o.InsideBoundingBox.Get(), o.InsideBoundingBox.IsSet()
}

// HasInsideBoundingBox returns a boolean if a field has been set.
func (o *Params) HasInsideBoundingBox() bool {
	if o != nil && o.InsideBoundingBox.IsSet() {
		return true
	}

	return false
}

// SetInsideBoundingBox gets a reference to the given utils.Nullable[InsideBoundingBox] and assigns it to the InsideBoundingBox field.
func (o *Params) SetInsideBoundingBox(v *InsideBoundingBox) *Params {
	o.InsideBoundingBox.Set(v)
	return o
}

// SetInsideBoundingBoxNil sets the value for InsideBoundingBox to be an explicit nil.
func (o *Params) SetInsideBoundingBoxNil() {
	o.InsideBoundingBox.Set(nil)
}

// UnsetInsideBoundingBox ensures that no value is present for InsideBoundingBox, not even an explicit nil.
func (o *Params) UnsetInsideBoundingBox() {
	o.InsideBoundingBox.Unset()
}

// GetInsidePolygon returns the InsidePolygon field value if set, zero value otherwise.
func (o *Params) GetInsidePolygon() [][]float64 {
	if o == nil || o.InsidePolygon == nil {
		var ret [][]float64
		return ret
	}
	return o.InsidePolygon
}

// GetInsidePolygonOk returns a tuple with the InsidePolygon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Params) GetInsidePolygonOk() ([][]float64, bool) {
	if o == nil || o.InsidePolygon == nil {
		return nil, false
	}
	return o.InsidePolygon, true
}

// HasInsidePolygon returns a boolean if a field has been set.
func (o *Params) HasInsidePolygon() bool {
	if o != nil && o.InsidePolygon != nil {
		return true
	}

	return false
}

// SetInsidePolygon gets a reference to the given [][]float64 and assigns it to the InsidePolygon field.
func (o *Params) SetInsidePolygon(v [][]float64) *Params {
	o.InsidePolygon = v
	return o
}

// GetQueryLanguages returns the QueryLanguages field value if set, zero value otherwise.
func (o *Params) GetQueryLanguages() []SupportedLanguage {
	if o == nil || o.QueryLanguages == nil {
		var ret []SupportedLanguage
		return ret
	}
	return o.QueryLanguages
}

// GetQueryLanguagesOk returns a tuple with the QueryLanguages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Params) GetQueryLanguagesOk() ([]SupportedLanguage, bool) {
	if o == nil || o.QueryLanguages == nil {
		return nil, false
	}
	return o.QueryLanguages, true
}

// HasQueryLanguages returns a boolean if a field has been set.
func (o *Params) HasQueryLanguages() bool {
	if o != nil && o.QueryLanguages != nil {
		return true
	}

	return false
}

// SetQueryLanguages gets a reference to the given []SupportedLanguage and assigns it to the QueryLanguages field.
func (o *Params) SetQueryLanguages(v []SupportedLanguage) *Params {
	o.QueryLanguages = v
	return o
}

// GetNaturalLanguages returns the NaturalLanguages field value if set, zero value otherwise.
func (o *Params) GetNaturalLanguages() []SupportedLanguage {
	if o == nil || o.NaturalLanguages == nil {
		var ret []SupportedLanguage
		return ret
	}
	return o.NaturalLanguages
}

// GetNaturalLanguagesOk returns a tuple with the NaturalLanguages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Params) GetNaturalLanguagesOk() ([]SupportedLanguage, bool) {
	if o == nil || o.NaturalLanguages == nil {
		return nil, false
	}
	return o.NaturalLanguages, true
}

// HasNaturalLanguages returns a boolean if a field has been set.
func (o *Params) HasNaturalLanguages() bool {
	if o != nil && o.NaturalLanguages != nil {
		return true
	}

	return false
}

// SetNaturalLanguages gets a reference to the given []SupportedLanguage and assigns it to the NaturalLanguages field.
func (o *Params) SetNaturalLanguages(v []SupportedLanguage) *Params {
	o.NaturalLanguages = v
	return o
}

// GetEnableRules returns the EnableRules field value if set, zero value otherwise.
func (o *Params) GetEnableRules() bool {
	if o == nil || o.EnableRules == nil {
		var ret bool
		return ret
	}
	return *o.EnableRules
}

// GetEnableRulesOk returns a tuple with the EnableRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Params) GetEnableRulesOk() (*bool, bool) {
	if o == nil || o.EnableRules == nil {
		return nil, false
	}
	return o.EnableRules, true
}

// HasEnableRules returns a boolean if a field has been set.
func (o *Params) HasEnableRules() bool {
	if o != nil && o.EnableRules != nil {
		return true
	}

	return false
}

// SetEnableRules gets a reference to the given bool and assigns it to the EnableRules field.
func (o *Params) SetEnableRules(v bool) *Params {
	o.EnableRules = &v
	return o
}

// GetRuleContexts returns the RuleContexts field value if set, zero value otherwise.
func (o *Params) GetRuleContexts() []string {
	if o == nil || o.RuleContexts == nil {
		var ret []string
		return ret
	}
	return o.RuleContexts
}

// GetRuleContextsOk returns a tuple with the RuleContexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Params) GetRuleContextsOk() ([]string, bool) {
	if o == nil || o.RuleContexts == nil {
		return nil, false
	}
	return o.RuleContexts, true
}

// HasRuleContexts returns a boolean if a field has been set.
func (o *Params) HasRuleContexts() bool {
	if o != nil && o.RuleContexts != nil {
		return true
	}

	return false
}

// SetRuleContexts gets a reference to the given []string and assigns it to the RuleContexts field.
func (o *Params) SetRuleContexts(v []string) *Params {
	o.RuleContexts = v
	return o
}

// GetUserToken returns the UserToken field value if set, zero value otherwise.
func (o *Params) GetUserToken() string {
	if o == nil || o.UserToken == nil {
		var ret string
		return ret
	}
	return *o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Params) GetUserTokenOk() (*string, bool) {
	if o == nil || o.UserToken == nil {
		return nil, false
	}
	return o.UserToken, true
}

// HasUserToken returns a boolean if a field has been set.
func (o *Params) HasUserToken() bool {
	if o != nil && o.UserToken != nil {
		return true
	}

	return false
}

// SetUserToken gets a reference to the given string and assigns it to the UserToken field.
func (o *Params) SetUserToken(v string) *Params {
	o.UserToken = &v
	return o
}

// GetClickAnalytics returns the ClickAnalytics field value if set, zero value otherwise.
func (o *Params) GetClickAnalytics() bool {
	if o == nil || o.ClickAnalytics == nil {
		var ret bool
		return ret
	}
	return *o.ClickAnalytics
}

// GetClickAnalyticsOk returns a tuple with the ClickAnalytics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Params) GetClickAnalyticsOk() (*bool, bool) {
	if o == nil || o.ClickAnalytics == nil {
		return nil, false
	}
	return o.ClickAnalytics, true
}

// HasClickAnalytics returns a boolean if a field has been set.
func (o *Params) HasClickAnalytics() bool {
	if o != nil && o.ClickAnalytics != nil {
		return true
	}

	return false
}

// SetClickAnalytics gets a reference to the given bool and assigns it to the ClickAnalytics field.
func (o *Params) SetClickAnalytics(v bool) *Params {
	o.ClickAnalytics = &v
	return o
}

// GetAnalytics returns the Analytics field value if set, zero value otherwise.
func (o *Params) GetAnalytics() bool {
	if o == nil || o.Analytics == nil {
		var ret bool
		return ret
	}
	return *o.Analytics
}

// GetAnalyticsOk returns a tuple with the Analytics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Params) GetAnalyticsOk() (*bool, bool) {
	if o == nil || o.Analytics == nil {
		return nil, false
	}
	return o.Analytics, true
}

// HasAnalytics returns a boolean if a field has been set.
func (o *Params) HasAnalytics() bool {
	if o != nil && o.Analytics != nil {
		return true
	}

	return false
}

// SetAnalytics gets a reference to the given bool and assigns it to the Analytics field.
func (o *Params) SetAnalytics(v bool) *Params {
	o.Analytics = &v
	return o
}

// GetAnalyticsTags returns the AnalyticsTags field value if set, zero value otherwise.
func (o *Params) GetAnalyticsTags() []string {
	if o == nil || o.AnalyticsTags == nil {
		var ret []string
		return ret
	}
	return o.AnalyticsTags
}

// GetAnalyticsTagsOk returns a tuple with the AnalyticsTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Params) GetAnalyticsTagsOk() ([]string, bool) {
	if o == nil || o.AnalyticsTags == nil {
		return nil, false
	}
	return o.AnalyticsTags, true
}

// HasAnalyticsTags returns a boolean if a field has been set.
func (o *Params) HasAnalyticsTags() bool {
	if o != nil && o.AnalyticsTags != nil {
		return true
	}

	return false
}

// SetAnalyticsTags gets a reference to the given []string and assigns it to the AnalyticsTags field.
func (o *Params) SetAnalyticsTags(v []string) *Params {
	o.AnalyticsTags = v
	return o
}

// GetEnableABTest returns the EnableABTest field value if set, zero value otherwise.
func (o *Params) GetEnableABTest() bool {
	if o == nil || o.EnableABTest == nil {
		var ret bool
		return ret
	}
	return *o.EnableABTest
}

// GetEnableABTestOk returns a tuple with the EnableABTest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Params) GetEnableABTestOk() (*bool, bool) {
	if o == nil || o.EnableABTest == nil {
		return nil, false
	}
	return o.EnableABTest, true
}

// HasEnableABTest returns a boolean if a field has been set.
func (o *Params) HasEnableABTest() bool {
	if o != nil && o.EnableABTest != nil {
		return true
	}

	return false
}

// SetEnableABTest gets a reference to the given bool and assigns it to the EnableABTest field.
func (o *Params) SetEnableABTest(v bool) *Params {
	o.EnableABTest = &v
	return o
}

// GetEnableReRanking returns the EnableReRanking field value if set, zero value otherwise.
func (o *Params) GetEnableReRanking() bool {
	if o == nil || o.EnableReRanking == nil {
		var ret bool
		return ret
	}
	return *o.EnableReRanking
}

// GetEnableReRankingOk returns a tuple with the EnableReRanking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Params) GetEnableReRankingOk() (*bool, bool) {
	if o == nil || o.EnableReRanking == nil {
		return nil, false
	}
	return o.EnableReRanking, true
}

// HasEnableReRanking returns a boolean if a field has been set.
func (o *Params) HasEnableReRanking() bool {
	if o != nil && o.EnableReRanking != nil {
		return true
	}

	return false
}

// SetEnableReRanking gets a reference to the given bool and assigns it to the EnableReRanking field.
func (o *Params) SetEnableReRanking(v bool) *Params {
	o.EnableReRanking = &v
	return o
}

func (o Params) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.GetRankingInfo != nil {
		toSerialize["getRankingInfo"] = o.GetRankingInfo
	}
	if o.RelevancyStrictness != nil {
		toSerialize["relevancyStrictness"] = o.RelevancyStrictness
	}
	if o.FacetFilters != nil {
		toSerialize["facetFilters"] = o.FacetFilters
	}
	if o.OptionalFilters != nil {
		toSerialize["optionalFilters"] = o.OptionalFilters
	}
	if o.NumericFilters != nil {
		toSerialize["numericFilters"] = o.NumericFilters
	}
	if o.HitsPerPage != nil {
		toSerialize["hitsPerPage"] = o.HitsPerPage
	}
	if o.AroundLatLng != nil {
		toSerialize["aroundLatLng"] = o.AroundLatLng
	}
	if o.AroundLatLngViaIP != nil {
		toSerialize["aroundLatLngViaIP"] = o.AroundLatLngViaIP
	}
	if o.AroundRadius != nil {
		toSerialize["aroundRadius"] = o.AroundRadius
	}
	if o.AroundPrecision != nil {
		toSerialize["aroundPrecision"] = o.AroundPrecision
	}
	if o.MinimumAroundRadius != nil {
		toSerialize["minimumAroundRadius"] = o.MinimumAroundRadius
	}
	if o.InsideBoundingBox.IsSet() {
		toSerialize["insideBoundingBox"] = o.InsideBoundingBox.Get()
	}
	if o.InsidePolygon != nil {
		toSerialize["insidePolygon"] = o.InsidePolygon
	}
	if o.QueryLanguages != nil {
		toSerialize["queryLanguages"] = o.QueryLanguages
	}
	if o.NaturalLanguages != nil {
		toSerialize["naturalLanguages"] = o.NaturalLanguages
	}
	if o.EnableRules != nil {
		toSerialize["enableRules"] = o.EnableRules
	}
	if o.RuleContexts != nil {
		toSerialize["ruleContexts"] = o.RuleContexts
	}
	if o.UserToken != nil {
		toSerialize["userToken"] = o.UserToken
	}
	if o.ClickAnalytics != nil {
		toSerialize["clickAnalytics"] = o.ClickAnalytics
	}
	if o.Analytics != nil {
		toSerialize["analytics"] = o.Analytics
	}
	if o.AnalyticsTags != nil {
		toSerialize["analyticsTags"] = o.AnalyticsTags
	}
	if o.EnableABTest != nil {
		toSerialize["enableABTest"] = o.EnableABTest
	}
	if o.EnableReRanking != nil {
		toSerialize["enableReRanking"] = o.EnableReRanking
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal Params: %w", err)
	}

	return serialized, nil
}

func (o Params) String() string {
	out := ""
	out += fmt.Sprintf("  query=%v\n", o.Query)
	out += fmt.Sprintf("  filters=%v\n", o.Filters)
	out += fmt.Sprintf("  page=%v\n", o.Page)
	out += fmt.Sprintf("  getRankingInfo=%v\n", o.GetRankingInfo)
	out += fmt.Sprintf("  relevancyStrictness=%v\n", o.RelevancyStrictness)
	out += fmt.Sprintf("  facetFilters=%v\n", o.FacetFilters)
	out += fmt.Sprintf("  optionalFilters=%v\n", o.OptionalFilters)
	out += fmt.Sprintf("  numericFilters=%v\n", o.NumericFilters)
	out += fmt.Sprintf("  hitsPerPage=%v\n", o.HitsPerPage)
	out += fmt.Sprintf("  aroundLatLng=%v\n", o.AroundLatLng)
	out += fmt.Sprintf("  aroundLatLngViaIP=%v\n", o.AroundLatLngViaIP)
	out += fmt.Sprintf("  aroundRadius=%v\n", o.AroundRadius)
	out += fmt.Sprintf("  aroundPrecision=%v\n", o.AroundPrecision)
	out += fmt.Sprintf("  minimumAroundRadius=%v\n", o.MinimumAroundRadius)
	out += fmt.Sprintf("  insideBoundingBox=%v\n", o.InsideBoundingBox)
	out += fmt.Sprintf("  insidePolygon=%v\n", o.InsidePolygon)
	out += fmt.Sprintf("  queryLanguages=%v\n", o.QueryLanguages)
	out += fmt.Sprintf("  naturalLanguages=%v\n", o.NaturalLanguages)
	out += fmt.Sprintf("  enableRules=%v\n", o.EnableRules)
	out += fmt.Sprintf("  ruleContexts=%v\n", o.RuleContexts)
	out += fmt.Sprintf("  userToken=%v\n", o.UserToken)
	out += fmt.Sprintf("  clickAnalytics=%v\n", o.ClickAnalytics)
	out += fmt.Sprintf("  analytics=%v\n", o.Analytics)
	out += fmt.Sprintf("  analyticsTags=%v\n", o.AnalyticsTags)
	out += fmt.Sprintf("  enableABTest=%v\n", o.EnableABTest)
	out += fmt.Sprintf("  enableReRanking=%v\n", o.EnableReRanking)
	return fmt.Sprintf("Params {\n%s}", out)
}
