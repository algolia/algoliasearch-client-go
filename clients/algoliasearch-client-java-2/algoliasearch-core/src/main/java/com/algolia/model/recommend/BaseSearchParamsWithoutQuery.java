package com.algolia.model.recommend;

import com.google.gson.annotations.SerializedName;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/** BaseSearchParamsWithoutQuery */
public class BaseSearchParamsWithoutQuery {

  @SerializedName("similarQuery")
  private String similarQuery;

  @SerializedName("filters")
  private String filters;

  @SerializedName("facetFilters")
  private FacetFilters facetFilters;

  @SerializedName("optionalFilters")
  private OptionalFilters optionalFilters;

  @SerializedName("numericFilters")
  private NumericFilters numericFilters;

  @SerializedName("tagFilters")
  private TagFilters tagFilters;

  @SerializedName("sumOrFiltersScores")
  private Boolean sumOrFiltersScores;

  @SerializedName("facets")
  private List<String> facets;

  @SerializedName("maxValuesPerFacet")
  private Integer maxValuesPerFacet;

  @SerializedName("facetingAfterDistinct")
  private Boolean facetingAfterDistinct;

  @SerializedName("sortFacetValuesBy")
  private String sortFacetValuesBy;

  @SerializedName("page")
  private Integer page;

  @SerializedName("offset")
  private Integer offset;

  @SerializedName("length")
  private Integer length;

  @SerializedName("aroundLatLng")
  private String aroundLatLng;

  @SerializedName("aroundLatLngViaIP")
  private Boolean aroundLatLngViaIP;

  @SerializedName("aroundRadius")
  private AroundRadius aroundRadius;

  @SerializedName("aroundPrecision")
  private Integer aroundPrecision;

  @SerializedName("minimumAroundRadius")
  private Integer minimumAroundRadius;

  @SerializedName("insideBoundingBox")
  private List<Double> insideBoundingBox;

  @SerializedName("insidePolygon")
  private List<Double> insidePolygon;

  @SerializedName("naturalLanguages")
  private List<String> naturalLanguages;

  @SerializedName("ruleContexts")
  private List<String> ruleContexts;

  @SerializedName("personalizationImpact")
  private Integer personalizationImpact;

  @SerializedName("userToken")
  private String userToken;

  @SerializedName("getRankingInfo")
  private Boolean getRankingInfo;

  @SerializedName("clickAnalytics")
  private Boolean clickAnalytics;

  @SerializedName("analytics")
  private Boolean analytics;

  @SerializedName("analyticsTags")
  private List<String> analyticsTags;

  @SerializedName("percentileComputation")
  private Boolean percentileComputation;

  @SerializedName("enableABTest")
  private Boolean enableABTest;

  @SerializedName("enableReRanking")
  private Boolean enableReRanking;

  @SerializedName("reRankingApplyFilter")
  private ReRankingApplyFilter reRankingApplyFilter;

  public BaseSearchParamsWithoutQuery setSimilarQuery(String similarQuery) {
    this.similarQuery = similarQuery;
    return this;
  }

  /**
   * Overrides the query parameter and performs a more generic search that can be used to find
   * \"similar\" results.
   *
   * @return similarQuery
   */
  @javax.annotation.Nullable
  public String getSimilarQuery() {
    return similarQuery;
  }

  public BaseSearchParamsWithoutQuery setFilters(String filters) {
    this.filters = filters;
    return this;
  }

  /**
   * Filter the query with numeric, facet and/or tag filters.
   *
   * @return filters
   */
  @javax.annotation.Nullable
  public String getFilters() {
    return filters;
  }

  public BaseSearchParamsWithoutQuery setFacetFilters(FacetFilters facetFilters) {
    this.facetFilters = facetFilters;
    return this;
  }

  /**
   * Get facetFilters
   *
   * @return facetFilters
   */
  @javax.annotation.Nullable
  public FacetFilters getFacetFilters() {
    return facetFilters;
  }

  public BaseSearchParamsWithoutQuery setOptionalFilters(OptionalFilters optionalFilters) {
    this.optionalFilters = optionalFilters;
    return this;
  }

  /**
   * Get optionalFilters
   *
   * @return optionalFilters
   */
  @javax.annotation.Nullable
  public OptionalFilters getOptionalFilters() {
    return optionalFilters;
  }

  public BaseSearchParamsWithoutQuery setNumericFilters(NumericFilters numericFilters) {
    this.numericFilters = numericFilters;
    return this;
  }

  /**
   * Get numericFilters
   *
   * @return numericFilters
   */
  @javax.annotation.Nullable
  public NumericFilters getNumericFilters() {
    return numericFilters;
  }

  public BaseSearchParamsWithoutQuery setTagFilters(TagFilters tagFilters) {
    this.tagFilters = tagFilters;
    return this;
  }

  /**
   * Get tagFilters
   *
   * @return tagFilters
   */
  @javax.annotation.Nullable
  public TagFilters getTagFilters() {
    return tagFilters;
  }

  public BaseSearchParamsWithoutQuery setSumOrFiltersScores(Boolean sumOrFiltersScores) {
    this.sumOrFiltersScores = sumOrFiltersScores;
    return this;
  }

  /**
   * Determines how to calculate the total score for filtering.
   *
   * @return sumOrFiltersScores
   */
  @javax.annotation.Nullable
  public Boolean getSumOrFiltersScores() {
    return sumOrFiltersScores;
  }

  public BaseSearchParamsWithoutQuery setFacets(List<String> facets) {
    this.facets = facets;
    return this;
  }

  public BaseSearchParamsWithoutQuery addFacets(String facetsItem) {
    if (this.facets == null) {
      this.facets = new ArrayList<>();
    }
    this.facets.add(facetsItem);
    return this;
  }

  /**
   * Retrieve facets and their facet values.
   *
   * @return facets
   */
  @javax.annotation.Nullable
  public List<String> getFacets() {
    return facets;
  }

  public BaseSearchParamsWithoutQuery setMaxValuesPerFacet(Integer maxValuesPerFacet) {
    this.maxValuesPerFacet = maxValuesPerFacet;
    return this;
  }

  /**
   * Maximum number of facet values to return for each facet during a regular search.
   *
   * @return maxValuesPerFacet
   */
  @javax.annotation.Nullable
  public Integer getMaxValuesPerFacet() {
    return maxValuesPerFacet;
  }

  public BaseSearchParamsWithoutQuery setFacetingAfterDistinct(Boolean facetingAfterDistinct) {
    this.facetingAfterDistinct = facetingAfterDistinct;
    return this;
  }

  /**
   * Force faceting to be applied after de-duplication (via the Distinct setting).
   *
   * @return facetingAfterDistinct
   */
  @javax.annotation.Nullable
  public Boolean getFacetingAfterDistinct() {
    return facetingAfterDistinct;
  }

  public BaseSearchParamsWithoutQuery setSortFacetValuesBy(String sortFacetValuesBy) {
    this.sortFacetValuesBy = sortFacetValuesBy;
    return this;
  }

  /**
   * Controls how facet values are fetched.
   *
   * @return sortFacetValuesBy
   */
  @javax.annotation.Nullable
  public String getSortFacetValuesBy() {
    return sortFacetValuesBy;
  }

  public BaseSearchParamsWithoutQuery setPage(Integer page) {
    this.page = page;
    return this;
  }

  /**
   * Specify the page to retrieve.
   *
   * @return page
   */
  @javax.annotation.Nullable
  public Integer getPage() {
    return page;
  }

  public BaseSearchParamsWithoutQuery setOffset(Integer offset) {
    this.offset = offset;
    return this;
  }

  /**
   * Specify the offset of the first hit to return.
   *
   * @return offset
   */
  @javax.annotation.Nullable
  public Integer getOffset() {
    return offset;
  }

  public BaseSearchParamsWithoutQuery setLength(Integer length) {
    this.length = length;
    return this;
  }

  /**
   * Set the number of hits to retrieve (used only with offset). minimum: 1 maximum: 1000
   *
   * @return length
   */
  @javax.annotation.Nullable
  public Integer getLength() {
    return length;
  }

  public BaseSearchParamsWithoutQuery setAroundLatLng(String aroundLatLng) {
    this.aroundLatLng = aroundLatLng;
    return this;
  }

  /**
   * Search for entries around a central geolocation, enabling a geo search within a circular area.
   *
   * @return aroundLatLng
   */
  @javax.annotation.Nullable
  public String getAroundLatLng() {
    return aroundLatLng;
  }

  public BaseSearchParamsWithoutQuery setAroundLatLngViaIP(Boolean aroundLatLngViaIP) {
    this.aroundLatLngViaIP = aroundLatLngViaIP;
    return this;
  }

  /**
   * Search for entries around a given location automatically computed from the requester's IP
   * address.
   *
   * @return aroundLatLngViaIP
   */
  @javax.annotation.Nullable
  public Boolean getAroundLatLngViaIP() {
    return aroundLatLngViaIP;
  }

  public BaseSearchParamsWithoutQuery setAroundRadius(AroundRadius aroundRadius) {
    this.aroundRadius = aroundRadius;
    return this;
  }

  /**
   * Get aroundRadius
   *
   * @return aroundRadius
   */
  @javax.annotation.Nullable
  public AroundRadius getAroundRadius() {
    return aroundRadius;
  }

  public BaseSearchParamsWithoutQuery setAroundPrecision(Integer aroundPrecision) {
    this.aroundPrecision = aroundPrecision;
    return this;
  }

  /**
   * Precision of geo search (in meters), to add grouping by geo location to the ranking formula.
   *
   * @return aroundPrecision
   */
  @javax.annotation.Nullable
  public Integer getAroundPrecision() {
    return aroundPrecision;
  }

  public BaseSearchParamsWithoutQuery setMinimumAroundRadius(Integer minimumAroundRadius) {
    this.minimumAroundRadius = minimumAroundRadius;
    return this;
  }

  /**
   * Minimum radius (in meters) used for a geo search when aroundRadius is not set. minimum: 1
   *
   * @return minimumAroundRadius
   */
  @javax.annotation.Nullable
  public Integer getMinimumAroundRadius() {
    return minimumAroundRadius;
  }

  public BaseSearchParamsWithoutQuery setInsideBoundingBox(List<Double> insideBoundingBox) {
    this.insideBoundingBox = insideBoundingBox;
    return this;
  }

  public BaseSearchParamsWithoutQuery addInsideBoundingBox(Double insideBoundingBoxItem) {
    if (this.insideBoundingBox == null) {
      this.insideBoundingBox = new ArrayList<>();
    }
    this.insideBoundingBox.add(insideBoundingBoxItem);
    return this;
  }

  /**
   * Search inside a rectangular area (in geo coordinates).
   *
   * @return insideBoundingBox
   */
  @javax.annotation.Nullable
  public List<Double> getInsideBoundingBox() {
    return insideBoundingBox;
  }

  public BaseSearchParamsWithoutQuery setInsidePolygon(List<Double> insidePolygon) {
    this.insidePolygon = insidePolygon;
    return this;
  }

  public BaseSearchParamsWithoutQuery addInsidePolygon(Double insidePolygonItem) {
    if (this.insidePolygon == null) {
      this.insidePolygon = new ArrayList<>();
    }
    this.insidePolygon.add(insidePolygonItem);
    return this;
  }

  /**
   * Search inside a polygon (in geo coordinates).
   *
   * @return insidePolygon
   */
  @javax.annotation.Nullable
  public List<Double> getInsidePolygon() {
    return insidePolygon;
  }

  public BaseSearchParamsWithoutQuery setNaturalLanguages(List<String> naturalLanguages) {
    this.naturalLanguages = naturalLanguages;
    return this;
  }

  public BaseSearchParamsWithoutQuery addNaturalLanguages(String naturalLanguagesItem) {
    if (this.naturalLanguages == null) {
      this.naturalLanguages = new ArrayList<>();
    }
    this.naturalLanguages.add(naturalLanguagesItem);
    return this;
  }

  /**
   * This parameter changes the default values of certain parameters and settings that work best for
   * a natural language query, such as ignorePlurals, removeStopWords, removeWordsIfNoResults,
   * analyticsTags and ruleContexts. These parameters and settings work well together when the query
   * is formatted in natural language instead of keywords, for example when your user performs a
   * voice search.
   *
   * @return naturalLanguages
   */
  @javax.annotation.Nullable
  public List<String> getNaturalLanguages() {
    return naturalLanguages;
  }

  public BaseSearchParamsWithoutQuery setRuleContexts(List<String> ruleContexts) {
    this.ruleContexts = ruleContexts;
    return this;
  }

  public BaseSearchParamsWithoutQuery addRuleContexts(String ruleContextsItem) {
    if (this.ruleContexts == null) {
      this.ruleContexts = new ArrayList<>();
    }
    this.ruleContexts.add(ruleContextsItem);
    return this;
  }

  /**
   * Enables contextual rules.
   *
   * @return ruleContexts
   */
  @javax.annotation.Nullable
  public List<String> getRuleContexts() {
    return ruleContexts;
  }

  public BaseSearchParamsWithoutQuery setPersonalizationImpact(Integer personalizationImpact) {
    this.personalizationImpact = personalizationImpact;
    return this;
  }

  /**
   * Define the impact of the Personalization feature.
   *
   * @return personalizationImpact
   */
  @javax.annotation.Nullable
  public Integer getPersonalizationImpact() {
    return personalizationImpact;
  }

  public BaseSearchParamsWithoutQuery setUserToken(String userToken) {
    this.userToken = userToken;
    return this;
  }

  /**
   * Associates a certain user token with the current search.
   *
   * @return userToken
   */
  @javax.annotation.Nullable
  public String getUserToken() {
    return userToken;
  }

  public BaseSearchParamsWithoutQuery setGetRankingInfo(Boolean getRankingInfo) {
    this.getRankingInfo = getRankingInfo;
    return this;
  }

  /**
   * Retrieve detailed ranking information.
   *
   * @return getRankingInfo
   */
  @javax.annotation.Nullable
  public Boolean getGetRankingInfo() {
    return getRankingInfo;
  }

  public BaseSearchParamsWithoutQuery setClickAnalytics(Boolean clickAnalytics) {
    this.clickAnalytics = clickAnalytics;
    return this;
  }

  /**
   * Enable the Click Analytics feature.
   *
   * @return clickAnalytics
   */
  @javax.annotation.Nullable
  public Boolean getClickAnalytics() {
    return clickAnalytics;
  }

  public BaseSearchParamsWithoutQuery setAnalytics(Boolean analytics) {
    this.analytics = analytics;
    return this;
  }

  /**
   * Whether the current query will be taken into account in the Analytics.
   *
   * @return analytics
   */
  @javax.annotation.Nullable
  public Boolean getAnalytics() {
    return analytics;
  }

  public BaseSearchParamsWithoutQuery setAnalyticsTags(List<String> analyticsTags) {
    this.analyticsTags = analyticsTags;
    return this;
  }

  public BaseSearchParamsWithoutQuery addAnalyticsTags(String analyticsTagsItem) {
    if (this.analyticsTags == null) {
      this.analyticsTags = new ArrayList<>();
    }
    this.analyticsTags.add(analyticsTagsItem);
    return this;
  }

  /**
   * List of tags to apply to the query for analytics purposes.
   *
   * @return analyticsTags
   */
  @javax.annotation.Nullable
  public List<String> getAnalyticsTags() {
    return analyticsTags;
  }

  public BaseSearchParamsWithoutQuery setPercentileComputation(Boolean percentileComputation) {
    this.percentileComputation = percentileComputation;
    return this;
  }

  /**
   * Whether to include or exclude a query from the processing-time percentile computation.
   *
   * @return percentileComputation
   */
  @javax.annotation.Nullable
  public Boolean getPercentileComputation() {
    return percentileComputation;
  }

  public BaseSearchParamsWithoutQuery setEnableABTest(Boolean enableABTest) {
    this.enableABTest = enableABTest;
    return this;
  }

  /**
   * Whether this search should participate in running AB tests.
   *
   * @return enableABTest
   */
  @javax.annotation.Nullable
  public Boolean getEnableABTest() {
    return enableABTest;
  }

  public BaseSearchParamsWithoutQuery setEnableReRanking(Boolean enableReRanking) {
    this.enableReRanking = enableReRanking;
    return this;
  }

  /**
   * Whether this search should use AI Re-Ranking.
   *
   * @return enableReRanking
   */
  @javax.annotation.Nullable
  public Boolean getEnableReRanking() {
    return enableReRanking;
  }

  public BaseSearchParamsWithoutQuery setReRankingApplyFilter(ReRankingApplyFilter reRankingApplyFilter) {
    this.reRankingApplyFilter = reRankingApplyFilter;
    return this;
  }

  /**
   * Get reRankingApplyFilter
   *
   * @return reRankingApplyFilter
   */
  @javax.annotation.Nullable
  public ReRankingApplyFilter getReRankingApplyFilter() {
    return reRankingApplyFilter;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    BaseSearchParamsWithoutQuery baseSearchParamsWithoutQuery = (BaseSearchParamsWithoutQuery) o;
    return (
      Objects.equals(this.similarQuery, baseSearchParamsWithoutQuery.similarQuery) &&
      Objects.equals(this.filters, baseSearchParamsWithoutQuery.filters) &&
      Objects.equals(this.facetFilters, baseSearchParamsWithoutQuery.facetFilters) &&
      Objects.equals(this.optionalFilters, baseSearchParamsWithoutQuery.optionalFilters) &&
      Objects.equals(this.numericFilters, baseSearchParamsWithoutQuery.numericFilters) &&
      Objects.equals(this.tagFilters, baseSearchParamsWithoutQuery.tagFilters) &&
      Objects.equals(this.sumOrFiltersScores, baseSearchParamsWithoutQuery.sumOrFiltersScores) &&
      Objects.equals(this.facets, baseSearchParamsWithoutQuery.facets) &&
      Objects.equals(this.maxValuesPerFacet, baseSearchParamsWithoutQuery.maxValuesPerFacet) &&
      Objects.equals(this.facetingAfterDistinct, baseSearchParamsWithoutQuery.facetingAfterDistinct) &&
      Objects.equals(this.sortFacetValuesBy, baseSearchParamsWithoutQuery.sortFacetValuesBy) &&
      Objects.equals(this.page, baseSearchParamsWithoutQuery.page) &&
      Objects.equals(this.offset, baseSearchParamsWithoutQuery.offset) &&
      Objects.equals(this.length, baseSearchParamsWithoutQuery.length) &&
      Objects.equals(this.aroundLatLng, baseSearchParamsWithoutQuery.aroundLatLng) &&
      Objects.equals(this.aroundLatLngViaIP, baseSearchParamsWithoutQuery.aroundLatLngViaIP) &&
      Objects.equals(this.aroundRadius, baseSearchParamsWithoutQuery.aroundRadius) &&
      Objects.equals(this.aroundPrecision, baseSearchParamsWithoutQuery.aroundPrecision) &&
      Objects.equals(this.minimumAroundRadius, baseSearchParamsWithoutQuery.minimumAroundRadius) &&
      Objects.equals(this.insideBoundingBox, baseSearchParamsWithoutQuery.insideBoundingBox) &&
      Objects.equals(this.insidePolygon, baseSearchParamsWithoutQuery.insidePolygon) &&
      Objects.equals(this.naturalLanguages, baseSearchParamsWithoutQuery.naturalLanguages) &&
      Objects.equals(this.ruleContexts, baseSearchParamsWithoutQuery.ruleContexts) &&
      Objects.equals(this.personalizationImpact, baseSearchParamsWithoutQuery.personalizationImpact) &&
      Objects.equals(this.userToken, baseSearchParamsWithoutQuery.userToken) &&
      Objects.equals(this.getRankingInfo, baseSearchParamsWithoutQuery.getRankingInfo) &&
      Objects.equals(this.clickAnalytics, baseSearchParamsWithoutQuery.clickAnalytics) &&
      Objects.equals(this.analytics, baseSearchParamsWithoutQuery.analytics) &&
      Objects.equals(this.analyticsTags, baseSearchParamsWithoutQuery.analyticsTags) &&
      Objects.equals(this.percentileComputation, baseSearchParamsWithoutQuery.percentileComputation) &&
      Objects.equals(this.enableABTest, baseSearchParamsWithoutQuery.enableABTest) &&
      Objects.equals(this.enableReRanking, baseSearchParamsWithoutQuery.enableReRanking) &&
      Objects.equals(this.reRankingApplyFilter, baseSearchParamsWithoutQuery.reRankingApplyFilter)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(
      similarQuery,
      filters,
      facetFilters,
      optionalFilters,
      numericFilters,
      tagFilters,
      sumOrFiltersScores,
      facets,
      maxValuesPerFacet,
      facetingAfterDistinct,
      sortFacetValuesBy,
      page,
      offset,
      length,
      aroundLatLng,
      aroundLatLngViaIP,
      aroundRadius,
      aroundPrecision,
      minimumAroundRadius,
      insideBoundingBox,
      insidePolygon,
      naturalLanguages,
      ruleContexts,
      personalizationImpact,
      userToken,
      getRankingInfo,
      clickAnalytics,
      analytics,
      analyticsTags,
      percentileComputation,
      enableABTest,
      enableReRanking,
      reRankingApplyFilter
    );
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class BaseSearchParamsWithoutQuery {\n");
    sb.append("    similarQuery: ").append(toIndentedString(similarQuery)).append("\n");
    sb.append("    filters: ").append(toIndentedString(filters)).append("\n");
    sb.append("    facetFilters: ").append(toIndentedString(facetFilters)).append("\n");
    sb.append("    optionalFilters: ").append(toIndentedString(optionalFilters)).append("\n");
    sb.append("    numericFilters: ").append(toIndentedString(numericFilters)).append("\n");
    sb.append("    tagFilters: ").append(toIndentedString(tagFilters)).append("\n");
    sb.append("    sumOrFiltersScores: ").append(toIndentedString(sumOrFiltersScores)).append("\n");
    sb.append("    facets: ").append(toIndentedString(facets)).append("\n");
    sb.append("    maxValuesPerFacet: ").append(toIndentedString(maxValuesPerFacet)).append("\n");
    sb.append("    facetingAfterDistinct: ").append(toIndentedString(facetingAfterDistinct)).append("\n");
    sb.append("    sortFacetValuesBy: ").append(toIndentedString(sortFacetValuesBy)).append("\n");
    sb.append("    page: ").append(toIndentedString(page)).append("\n");
    sb.append("    offset: ").append(toIndentedString(offset)).append("\n");
    sb.append("    length: ").append(toIndentedString(length)).append("\n");
    sb.append("    aroundLatLng: ").append(toIndentedString(aroundLatLng)).append("\n");
    sb.append("    aroundLatLngViaIP: ").append(toIndentedString(aroundLatLngViaIP)).append("\n");
    sb.append("    aroundRadius: ").append(toIndentedString(aroundRadius)).append("\n");
    sb.append("    aroundPrecision: ").append(toIndentedString(aroundPrecision)).append("\n");
    sb.append("    minimumAroundRadius: ").append(toIndentedString(minimumAroundRadius)).append("\n");
    sb.append("    insideBoundingBox: ").append(toIndentedString(insideBoundingBox)).append("\n");
    sb.append("    insidePolygon: ").append(toIndentedString(insidePolygon)).append("\n");
    sb.append("    naturalLanguages: ").append(toIndentedString(naturalLanguages)).append("\n");
    sb.append("    ruleContexts: ").append(toIndentedString(ruleContexts)).append("\n");
    sb.append("    personalizationImpact: ").append(toIndentedString(personalizationImpact)).append("\n");
    sb.append("    userToken: ").append(toIndentedString(userToken)).append("\n");
    sb.append("    getRankingInfo: ").append(toIndentedString(getRankingInfo)).append("\n");
    sb.append("    clickAnalytics: ").append(toIndentedString(clickAnalytics)).append("\n");
    sb.append("    analytics: ").append(toIndentedString(analytics)).append("\n");
    sb.append("    analyticsTags: ").append(toIndentedString(analyticsTags)).append("\n");
    sb.append("    percentileComputation: ").append(toIndentedString(percentileComputation)).append("\n");
    sb.append("    enableABTest: ").append(toIndentedString(enableABTest)).append("\n");
    sb.append("    enableReRanking: ").append(toIndentedString(enableReRanking)).append("\n");
    sb.append("    reRankingApplyFilter: ").append(toIndentedString(reRankingApplyFilter)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }
}
