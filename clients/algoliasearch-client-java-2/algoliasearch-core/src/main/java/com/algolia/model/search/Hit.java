// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.search;

import com.fasterxml.jackson.annotation.*;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;

/** A single hit. */
public class Hit {

  @JsonProperty("objectID")
  private String objectID;

  @JsonProperty("_highlightResult")
  private Map<String, HighlightResult> highlightResult;

  @JsonProperty("_snippetResult")
  private Map<String, SnippetResult> snippetResult;

  @JsonProperty("_rankingInfo")
  private RankingInfo rankingInfo;

  @JsonProperty("_distinctSeqID")
  private Integer distinctSeqID;

  private Map<String, Object> additionalProperties = new HashMap<>();

  @JsonAnyGetter
  public Map<String, Object> getAdditionalProperties() {
    return this.additionalProperties;
  }

  @JsonAnySetter
  public Hit setAdditionalProperty(String name, Object value) {
    this.additionalProperties.put(name, value);
    return this;
  }

  public Hit setObjectID(String objectID) {
    this.objectID = objectID;
    return this;
  }

  /**
   * Unique identifier of the object.
   *
   * @return objectID
   */
  @javax.annotation.Nonnull
  public String getObjectID() {
    return objectID;
  }

  public Hit setHighlightResult(Map<String, HighlightResult> highlightResult) {
    this.highlightResult = highlightResult;
    return this;
  }

  public Hit putHighlightResult(String key, HighlightResult highlightResultItem) {
    if (this.highlightResult == null) {
      this.highlightResult = new HashMap<>();
    }
    this.highlightResult.put(key, highlightResultItem);
    return this;
  }

  /**
   * Show highlighted section and words matched on a query.
   *
   * @return highlightResult
   */
  @javax.annotation.Nullable
  public Map<String, HighlightResult> getHighlightResult() {
    return highlightResult;
  }

  public Hit setSnippetResult(Map<String, SnippetResult> snippetResult) {
    this.snippetResult = snippetResult;
    return this;
  }

  public Hit putSnippetResult(String key, SnippetResult snippetResultItem) {
    if (this.snippetResult == null) {
      this.snippetResult = new HashMap<>();
    }
    this.snippetResult.put(key, snippetResultItem);
    return this;
  }

  /**
   * Snippeted attributes show parts of the matched attributes. Only returned when
   * attributesToSnippet is non-empty.
   *
   * @return snippetResult
   */
  @javax.annotation.Nullable
  public Map<String, SnippetResult> getSnippetResult() {
    return snippetResult;
  }

  public Hit setRankingInfo(RankingInfo rankingInfo) {
    this.rankingInfo = rankingInfo;
    return this;
  }

  /**
   * Get rankingInfo
   *
   * @return rankingInfo
   */
  @javax.annotation.Nullable
  public RankingInfo getRankingInfo() {
    return rankingInfo;
  }

  public Hit setDistinctSeqID(Integer distinctSeqID) {
    this.distinctSeqID = distinctSeqID;
    return this;
  }

  /**
   * Get distinctSeqID
   *
   * @return distinctSeqID
   */
  @javax.annotation.Nullable
  public Integer getDistinctSeqID() {
    return distinctSeqID;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Hit hit = (Hit) o;
    return (
      Objects.equals(this.objectID, hit.objectID) &&
      Objects.equals(this.highlightResult, hit.highlightResult) &&
      Objects.equals(this.snippetResult, hit.snippetResult) &&
      Objects.equals(this.rankingInfo, hit.rankingInfo) &&
      Objects.equals(this.distinctSeqID, hit.distinctSeqID) &&
      super.equals(o)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(objectID, highlightResult, snippetResult, rankingInfo, distinctSeqID, super.hashCode());
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Hit {\n");
    sb.append("    ").append(toIndentedString(super.toString())).append("\n");
    sb.append("    objectID: ").append(toIndentedString(objectID)).append("\n");
    sb.append("    highlightResult: ").append(toIndentedString(highlightResult)).append("\n");
    sb.append("    snippetResult: ").append(toIndentedString(snippetResult)).append("\n");
    sb.append("    rankingInfo: ").append(toIndentedString(rankingInfo)).append("\n");
    sb.append("    distinctSeqID: ").append(toIndentedString(distinctSeqID)).append("\n");
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
