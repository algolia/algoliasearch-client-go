// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.querysuggestions;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/** SourceIndex */
public class SourceIndex {

  @JsonProperty("indexName")
  private String indexName;

  @JsonProperty("analyticsTags")
  private List<String> analyticsTags;

  @JsonProperty("facets")
  private List<Object> facets;

  @JsonProperty("minHits")
  private Integer minHits;

  @JsonProperty("minLetters")
  private Integer minLetters;

  @JsonProperty("generate")
  private List<List<String>> generate;

  @JsonProperty("external")
  private List<SourceIndexExternal> external;

  public SourceIndex setIndexName(String indexName) {
    this.indexName = indexName;
    return this;
  }

  /**
   * Source index name.
   *
   * @return indexName
   */
  @javax.annotation.Nonnull
  public String getIndexName() {
    return indexName;
  }

  public SourceIndex setAnalyticsTags(List<String> analyticsTags) {
    this.analyticsTags = analyticsTags;
    return this;
  }

  public SourceIndex addAnalyticsTags(String analyticsTagsItem) {
    if (this.analyticsTags == null) {
      this.analyticsTags = new ArrayList<>();
    }
    this.analyticsTags.add(analyticsTagsItem);
    return this;
  }

  /**
   * List of analytics tags to filter the popular searches per tag.
   *
   * @return analyticsTags
   */
  @javax.annotation.Nullable
  public List<String> getAnalyticsTags() {
    return analyticsTags;
  }

  public SourceIndex setFacets(List<Object> facets) {
    this.facets = facets;
    return this;
  }

  public SourceIndex addFacets(Object facetsItem) {
    if (this.facets == null) {
      this.facets = new ArrayList<>();
    }
    this.facets.add(facetsItem);
    return this;
  }

  /**
   * List of facets to define as categories for the query suggestions.
   *
   * @return facets
   */
  @javax.annotation.Nullable
  public List<Object> getFacets() {
    return facets;
  }

  public SourceIndex setMinHits(Integer minHits) {
    this.minHits = minHits;
    return this;
  }

  /**
   * Minimum number of hits (e.g., matching records in the source index) to generate a suggestions.
   *
   * @return minHits
   */
  @javax.annotation.Nullable
  public Integer getMinHits() {
    return minHits;
  }

  public SourceIndex setMinLetters(Integer minLetters) {
    this.minLetters = minLetters;
    return this;
  }

  /**
   * Minimum number of required letters for a suggestion to remain.
   *
   * @return minLetters
   */
  @javax.annotation.Nullable
  public Integer getMinLetters() {
    return minLetters;
  }

  public SourceIndex setGenerate(List<List<String>> generate) {
    this.generate = generate;
    return this;
  }

  public SourceIndex addGenerate(List<String> generateItem) {
    if (this.generate == null) {
      this.generate = new ArrayList<>();
    }
    this.generate.add(generateItem);
    return this;
  }

  /**
   * List of facet attributes used to generate Query Suggestions. The resulting suggestions are
   * every combination of the facets in the nested list (e.g., (facetA and facetB) and facetC).
   *
   * @return generate
   */
  @javax.annotation.Nullable
  public List<List<String>> getGenerate() {
    return generate;
  }

  public SourceIndex setExternal(List<SourceIndexExternal> external) {
    this.external = external;
    return this;
  }

  public SourceIndex addExternal(SourceIndexExternal externalItem) {
    if (this.external == null) {
      this.external = new ArrayList<>();
    }
    this.external.add(externalItem);
    return this;
  }

  /**
   * List of external indices to use to generate custom Query Suggestions.
   *
   * @return external
   */
  @javax.annotation.Nullable
  public List<SourceIndexExternal> getExternal() {
    return external;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SourceIndex sourceIndex = (SourceIndex) o;
    return (
      Objects.equals(this.indexName, sourceIndex.indexName) &&
      Objects.equals(this.analyticsTags, sourceIndex.analyticsTags) &&
      Objects.equals(this.facets, sourceIndex.facets) &&
      Objects.equals(this.minHits, sourceIndex.minHits) &&
      Objects.equals(this.minLetters, sourceIndex.minLetters) &&
      Objects.equals(this.generate, sourceIndex.generate) &&
      Objects.equals(this.external, sourceIndex.external)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(indexName, analyticsTags, facets, minHits, minLetters, generate, external);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SourceIndex {\n");
    sb.append("    indexName: ").append(toIndentedString(indexName)).append("\n");
    sb.append("    analyticsTags: ").append(toIndentedString(analyticsTags)).append("\n");
    sb.append("    facets: ").append(toIndentedString(facets)).append("\n");
    sb.append("    minHits: ").append(toIndentedString(minHits)).append("\n");
    sb.append("    minLetters: ").append(toIndentedString(minLetters)).append("\n");
    sb.append("    generate: ").append(toIndentedString(generate)).append("\n");
    sb.append("    external: ").append(toIndentedString(external)).append("\n");
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