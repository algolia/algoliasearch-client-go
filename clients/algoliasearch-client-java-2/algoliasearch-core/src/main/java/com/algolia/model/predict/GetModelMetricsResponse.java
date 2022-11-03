// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.predict;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/** GetModelMetricsResponse */
public class GetModelMetricsResponse {

  @JsonProperty("modelID")
  private String modelID;

  @JsonProperty("metrics")
  private List<ModelMetrics> metrics = new ArrayList<>();

  public GetModelMetricsResponse setModelID(String modelID) {
    this.modelID = modelID;
    return this;
  }

  /**
   * The ID of the model.
   *
   * @return modelID
   */
  @javax.annotation.Nonnull
  public String getModelID() {
    return modelID;
  }

  public GetModelMetricsResponse setMetrics(List<ModelMetrics> metrics) {
    this.metrics = metrics;
    return this;
  }

  public GetModelMetricsResponse addMetrics(ModelMetrics metricsItem) {
    this.metrics.add(metricsItem);
    return this;
  }

  /**
   * Get metrics
   *
   * @return metrics
   */
  @javax.annotation.Nonnull
  public List<ModelMetrics> getMetrics() {
    return metrics;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    GetModelMetricsResponse getModelMetricsResponse = (GetModelMetricsResponse) o;
    return Objects.equals(this.modelID, getModelMetricsResponse.modelID) && Objects.equals(this.metrics, getModelMetricsResponse.metrics);
  }

  @Override
  public int hashCode() {
    return Objects.hash(modelID, metrics);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class GetModelMetricsResponse {\n");
    sb.append("    modelID: ").append(toIndentedString(modelID)).append("\n");
    sb.append("    metrics: ").append(toIndentedString(metrics)).append("\n");
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
