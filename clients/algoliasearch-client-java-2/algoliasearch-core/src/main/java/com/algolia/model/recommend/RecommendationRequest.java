// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.recommend;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** RecommendationRequest */
public class RecommendationRequest {

  @JsonProperty("model")
  private RecommendationModels model;

  @JsonProperty("objectID")
  private String objectID;

  @JsonProperty("indexName")
  private String indexName;

  @JsonProperty("threshold")
  private Integer threshold;

  @JsonProperty("maxRecommendations")
  private Integer maxRecommendations;

  @JsonProperty("queryParameters")
  private SearchParamsObject queryParameters;

  @JsonProperty("fallbackParameters")
  private SearchParamsObject fallbackParameters;

  public RecommendationRequest setModel(RecommendationModels model) {
    this.model = model;
    return this;
  }

  /**
   * Get model
   *
   * @return model
   */
  @javax.annotation.Nonnull
  public RecommendationModels getModel() {
    return model;
  }

  public RecommendationRequest setObjectID(String objectID) {
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

  public RecommendationRequest setIndexName(String indexName) {
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

  public RecommendationRequest setThreshold(Integer threshold) {
    this.threshold = threshold;
    return this;
  }

  /**
   * The threshold to use when filtering recommendations by their score. minimum: 0 maximum: 100
   *
   * @return threshold
   */
  @javax.annotation.Nonnull
  public Integer getThreshold() {
    return threshold;
  }

  public RecommendationRequest setMaxRecommendations(Integer maxRecommendations) {
    this.maxRecommendations = maxRecommendations;
    return this;
  }

  /**
   * The max number of recommendations to retrieve. If it's set to 0, all the recommendations of the
   * objectID may be returned.
   *
   * @return maxRecommendations
   */
  @javax.annotation.Nullable
  public Integer getMaxRecommendations() {
    return maxRecommendations;
  }

  public RecommendationRequest setQueryParameters(SearchParamsObject queryParameters) {
    this.queryParameters = queryParameters;
    return this;
  }

  /**
   * Get queryParameters
   *
   * @return queryParameters
   */
  @javax.annotation.Nullable
  public SearchParamsObject getQueryParameters() {
    return queryParameters;
  }

  public RecommendationRequest setFallbackParameters(SearchParamsObject fallbackParameters) {
    this.fallbackParameters = fallbackParameters;
    return this;
  }

  /**
   * Get fallbackParameters
   *
   * @return fallbackParameters
   */
  @javax.annotation.Nullable
  public SearchParamsObject getFallbackParameters() {
    return fallbackParameters;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    RecommendationRequest recommendationRequest = (RecommendationRequest) o;
    return (
      Objects.equals(this.model, recommendationRequest.model) &&
      Objects.equals(this.objectID, recommendationRequest.objectID) &&
      Objects.equals(this.indexName, recommendationRequest.indexName) &&
      Objects.equals(this.threshold, recommendationRequest.threshold) &&
      Objects.equals(this.maxRecommendations, recommendationRequest.maxRecommendations) &&
      Objects.equals(this.queryParameters, recommendationRequest.queryParameters) &&
      Objects.equals(this.fallbackParameters, recommendationRequest.fallbackParameters)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(model, objectID, indexName, threshold, maxRecommendations, queryParameters, fallbackParameters);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class RecommendationRequest {\n");
    sb.append("    model: ").append(toIndentedString(model)).append("\n");
    sb.append("    objectID: ").append(toIndentedString(objectID)).append("\n");
    sb.append("    indexName: ").append(toIndentedString(indexName)).append("\n");
    sb.append("    threshold: ").append(toIndentedString(threshold)).append("\n");
    sb.append("    maxRecommendations: ").append(toIndentedString(maxRecommendations)).append("\n");
    sb.append("    queryParameters: ").append(toIndentedString(queryParameters)).append("\n");
    sb.append("    fallbackParameters: ").append(toIndentedString(fallbackParameters)).append("\n");
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