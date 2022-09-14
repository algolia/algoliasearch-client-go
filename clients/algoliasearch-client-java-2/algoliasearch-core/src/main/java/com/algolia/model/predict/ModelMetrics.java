// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.predict;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** ModelMetrics */
public class ModelMetrics {

  @JsonProperty("precision")
  private Double precision;

  @JsonProperty("recall")
  private Double recall;

  @JsonProperty("mrr")
  private Double mrr;

  @JsonProperty("coverage")
  private Double coverage;

  @JsonProperty("f1_score")
  private Double f1Score;

  @JsonProperty("updatedAt")
  private String updatedAt;

  public ModelMetrics setPrecision(Double precision) {
    this.precision = precision;
    return this;
  }

  /**
   * Get precision
   *
   * @return precision
   */
  @javax.annotation.Nullable
  public Double getPrecision() {
    return precision;
  }

  public ModelMetrics setRecall(Double recall) {
    this.recall = recall;
    return this;
  }

  /**
   * Get recall
   *
   * @return recall
   */
  @javax.annotation.Nullable
  public Double getRecall() {
    return recall;
  }

  public ModelMetrics setMrr(Double mrr) {
    this.mrr = mrr;
    return this;
  }

  /**
   * Get mrr
   *
   * @return mrr
   */
  @javax.annotation.Nullable
  public Double getMrr() {
    return mrr;
  }

  public ModelMetrics setCoverage(Double coverage) {
    this.coverage = coverage;
    return this;
  }

  /**
   * Get coverage
   *
   * @return coverage
   */
  @javax.annotation.Nullable
  public Double getCoverage() {
    return coverage;
  }

  public ModelMetrics setF1Score(Double f1Score) {
    this.f1Score = f1Score;
    return this;
  }

  /**
   * Get f1Score
   *
   * @return f1Score
   */
  @javax.annotation.Nullable
  public Double getF1Score() {
    return f1Score;
  }

  public ModelMetrics setUpdatedAt(String updatedAt) {
    this.updatedAt = updatedAt;
    return this;
  }

  /**
   * Date of last update (ISO-8601 format).
   *
   * @return updatedAt
   */
  @javax.annotation.Nullable
  public String getUpdatedAt() {
    return updatedAt;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ModelMetrics modelMetrics = (ModelMetrics) o;
    return (
      Objects.equals(this.precision, modelMetrics.precision) &&
      Objects.equals(this.recall, modelMetrics.recall) &&
      Objects.equals(this.mrr, modelMetrics.mrr) &&
      Objects.equals(this.coverage, modelMetrics.coverage) &&
      Objects.equals(this.f1Score, modelMetrics.f1Score) &&
      Objects.equals(this.updatedAt, modelMetrics.updatedAt)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(precision, recall, mrr, coverage, f1Score, updatedAt);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ModelMetrics {\n");
    sb.append("    precision: ").append(toIndentedString(precision)).append("\n");
    sb.append("    recall: ").append(toIndentedString(recall)).append("\n");
    sb.append("    mrr: ").append(toIndentedString(mrr)).append("\n");
    sb.append("    coverage: ").append(toIndentedString(coverage)).append("\n");
    sb.append("    f1Score: ").append(toIndentedString(f1Score)).append("\n");
    sb.append("    updatedAt: ").append(toIndentedString(updatedAt)).append("\n");
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
