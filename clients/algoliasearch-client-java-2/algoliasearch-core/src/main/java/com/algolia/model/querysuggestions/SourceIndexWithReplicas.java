package com.algolia.model.querysuggestions;

import com.google.gson.annotations.SerializedName;
import java.util.List;
import java.util.Objects;

/** Source index with replicas used to generate a Query Suggestions index. */
public class SourceIndexWithReplicas {

  @SerializedName("replicas")
  private Boolean replicas;

  @SerializedName("indexName")
  private String indexName;

  @SerializedName("analyticsTags")
  private List<String> analyticsTags;

  @SerializedName("facets")
  private List<Object> facets;

  @SerializedName("minHits")
  private Integer minHits;

  @SerializedName("minLetters")
  private Integer minLetters;

  @SerializedName("generate")
  private List<List<String>> generate;

  @SerializedName("external")
  private List<SourceIndexExternal> external;

  public SourceIndexWithReplicas setReplicas(Boolean replicas) {
    this.replicas = replicas;
    return this;
  }

  /**
   * true if the Query Suggestions index is a replicas.
   *
   * @return replicas
   */
  @javax.annotation.Nonnull
  public Boolean getReplicas() {
    return replicas;
  }

  public SourceIndexWithReplicas setIndexName(String indexName) {
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

  public SourceIndexWithReplicas setAnalyticsTags(List<String> analyticsTags) {
    this.analyticsTags = analyticsTags;
    return this;
  }

  public SourceIndexWithReplicas addAnalyticsTags(String analyticsTagsItem) {
    this.analyticsTags.add(analyticsTagsItem);
    return this;
  }

  /**
   * List of analytics tags to filter the popular searches per tag.
   *
   * @return analyticsTags
   */
  @javax.annotation.Nonnull
  public List<String> getAnalyticsTags() {
    return analyticsTags;
  }

  public SourceIndexWithReplicas setFacets(List<Object> facets) {
    this.facets = facets;
    return this;
  }

  public SourceIndexWithReplicas addFacets(Object facetsItem) {
    this.facets.add(facetsItem);
    return this;
  }

  /**
   * List of facets to define as categories for the query suggestions.
   *
   * @return facets
   */
  @javax.annotation.Nonnull
  public List<Object> getFacets() {
    return facets;
  }

  public SourceIndexWithReplicas setMinHits(Integer minHits) {
    this.minHits = minHits;
    return this;
  }

  /**
   * Minimum number of hits (e.g., matching records in the source index) to generate a suggestions.
   *
   * @return minHits
   */
  @javax.annotation.Nonnull
  public Integer getMinHits() {
    return minHits;
  }

  public SourceIndexWithReplicas setMinLetters(Integer minLetters) {
    this.minLetters = minLetters;
    return this;
  }

  /**
   * Minimum number of required letters for a suggestion to remain.
   *
   * @return minLetters
   */
  @javax.annotation.Nonnull
  public Integer getMinLetters() {
    return minLetters;
  }

  public SourceIndexWithReplicas setGenerate(List<List<String>> generate) {
    this.generate = generate;
    return this;
  }

  public SourceIndexWithReplicas addGenerate(List<String> generateItem) {
    this.generate.add(generateItem);
    return this;
  }

  /**
   * List of facet attributes used to generate Query Suggestions. The resulting suggestions are
   * every combination of the facets in the nested list (e.g., (facetA and facetB) and facetC).
   *
   * @return generate
   */
  @javax.annotation.Nonnull
  public List<List<String>> getGenerate() {
    return generate;
  }

  public SourceIndexWithReplicas setExternal(List<SourceIndexExternal> external) {
    this.external = external;
    return this;
  }

  public SourceIndexWithReplicas addExternal(SourceIndexExternal externalItem) {
    this.external.add(externalItem);
    return this;
  }

  /**
   * List of external indices to use to generate custom Query Suggestions.
   *
   * @return external
   */
  @javax.annotation.Nonnull
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
    SourceIndexWithReplicas sourceIndexWithReplicas = (SourceIndexWithReplicas) o;
    return (
      Objects.equals(this.replicas, sourceIndexWithReplicas.replicas) &&
      Objects.equals(this.indexName, sourceIndexWithReplicas.indexName) &&
      Objects.equals(this.analyticsTags, sourceIndexWithReplicas.analyticsTags) &&
      Objects.equals(this.facets, sourceIndexWithReplicas.facets) &&
      Objects.equals(this.minHits, sourceIndexWithReplicas.minHits) &&
      Objects.equals(this.minLetters, sourceIndexWithReplicas.minLetters) &&
      Objects.equals(this.generate, sourceIndexWithReplicas.generate) &&
      Objects.equals(this.external, sourceIndexWithReplicas.external)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(replicas, indexName, analyticsTags, facets, minHits, minLetters, generate, external);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SourceIndexWithReplicas {\n");
    sb.append("    replicas: ").append(toIndentedString(replicas)).append("\n");
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
