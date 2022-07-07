package com.algolia.model.search;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** SearchForFacetsOptions */
public class SearchForFacetsOptions {

  @JsonProperty("facet")
  private String facet;

  @JsonProperty("indexName")
  private String indexName;

  @JsonProperty("facetQuery")
  private String facetQuery;

  @JsonProperty("maxFacetHits")
  private Integer maxFacetHits;

  @JsonProperty("type")
  private SearchTypeFacet type;

  public SearchForFacetsOptions setFacet(String facet) {
    this.facet = facet;
    return this;
  }

  /**
   * The `facet` name.
   *
   * @return facet
   */
  @javax.annotation.Nonnull
  public String getFacet() {
    return facet;
  }

  public SearchForFacetsOptions setIndexName(String indexName) {
    this.indexName = indexName;
    return this;
  }

  /**
   * The Algolia index name.
   *
   * @return indexName
   */
  @javax.annotation.Nonnull
  public String getIndexName() {
    return indexName;
  }

  public SearchForFacetsOptions setFacetQuery(String facetQuery) {
    this.facetQuery = facetQuery;
    return this;
  }

  /**
   * Text to search inside the facet's values.
   *
   * @return facetQuery
   */
  @javax.annotation.Nullable
  public String getFacetQuery() {
    return facetQuery;
  }

  public SearchForFacetsOptions setMaxFacetHits(Integer maxFacetHits) {
    this.maxFacetHits = maxFacetHits;
    return this;
  }

  /**
   * Maximum number of facet hits to return during a search for facet values. For performance
   * reasons, the maximum allowed number of returned values is 100. maximum: 100
   *
   * @return maxFacetHits
   */
  @javax.annotation.Nullable
  public Integer getMaxFacetHits() {
    return maxFacetHits;
  }

  public SearchForFacetsOptions setType(SearchTypeFacet type) {
    this.type = type;
    return this;
  }

  /**
   * Get type
   *
   * @return type
   */
  @javax.annotation.Nonnull
  public SearchTypeFacet getType() {
    return type;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SearchForFacetsOptions searchForFacetsOptions = (SearchForFacetsOptions) o;
    return (
      Objects.equals(this.facet, searchForFacetsOptions.facet) &&
      Objects.equals(this.indexName, searchForFacetsOptions.indexName) &&
      Objects.equals(this.facetQuery, searchForFacetsOptions.facetQuery) &&
      Objects.equals(this.maxFacetHits, searchForFacetsOptions.maxFacetHits) &&
      Objects.equals(this.type, searchForFacetsOptions.type)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(facet, indexName, facetQuery, maxFacetHits, type);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SearchForFacetsOptions {\n");
    sb.append("    facet: ").append(toIndentedString(facet)).append("\n");
    sb.append("    indexName: ").append(toIndentedString(indexName)).append("\n");
    sb.append("    facetQuery: ").append(toIndentedString(facetQuery)).append("\n");
    sb.append("    maxFacetHits: ").append(toIndentedString(maxFacetHits)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
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
