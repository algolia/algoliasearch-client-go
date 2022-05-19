package com.algolia.model.search;

import com.google.gson.annotations.SerializedName;
import java.util.Objects;

/** SearchQueries */
public class SearchQueries {

  @SerializedName("indexName")
  private String indexName;

  @SerializedName("query")
  private String query = "";

  @SerializedName("type")
  private SearchType type = SearchType.DEFAULT;

  @SerializedName("facet")
  private String facet;

  @SerializedName("params")
  private SearchParams params;

  public SearchQueries setIndexName(String indexName) {
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

  public SearchQueries setQuery(String query) {
    this.query = query;
    return this;
  }

  /**
   * The text to search in the index.
   *
   * @return query
   */
  @javax.annotation.Nullable
  public String getQuery() {
    return query;
  }

  public SearchQueries setType(SearchType type) {
    this.type = type;
    return this;
  }

  /**
   * Get type
   *
   * @return type
   */
  @javax.annotation.Nullable
  public SearchType getType() {
    return type;
  }

  public SearchQueries setFacet(String facet) {
    this.facet = facet;
    return this;
  }

  /**
   * The `facet` name.
   *
   * @return facet
   */
  @javax.annotation.Nullable
  public String getFacet() {
    return facet;
  }

  public SearchQueries setParams(SearchParams params) {
    this.params = params;
    return this;
  }

  /**
   * Get params
   *
   * @return params
   */
  @javax.annotation.Nullable
  public SearchParams getParams() {
    return params;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SearchQueries searchQueries = (SearchQueries) o;
    return (
      Objects.equals(this.indexName, searchQueries.indexName) &&
      Objects.equals(this.query, searchQueries.query) &&
      Objects.equals(this.type, searchQueries.type) &&
      Objects.equals(this.facet, searchQueries.facet) &&
      Objects.equals(this.params, searchQueries.params)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(indexName, query, type, facet, params);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SearchQueries {\n");
    sb
      .append("    indexName: ")
      .append(toIndentedString(indexName))
      .append("\n");
    sb.append("    query: ").append(toIndentedString(query)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    facet: ").append(toIndentedString(facet)).append("\n");
    sb.append("    params: ").append(toIndentedString(params)).append("\n");
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
