// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.predict;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/** Prediction for the **funnel_stage** model. */
public class FunnelStageSuccess {

  @JsonProperty("value")
  private List<FunnelStage> value = new ArrayList<>();

  @JsonProperty("lastUpdatedAt")
  private String lastUpdatedAt;

  public FunnelStageSuccess setValue(List<FunnelStage> value) {
    this.value = value;
    return this;
  }

  public FunnelStageSuccess addValue(FunnelStage valueItem) {
    this.value.add(valueItem);
    return this;
  }

  /**
   * Get value
   *
   * @return value
   */
  @javax.annotation.Nonnull
  public List<FunnelStage> getValue() {
    return value;
  }

  public FunnelStageSuccess setLastUpdatedAt(String lastUpdatedAt) {
    this.lastUpdatedAt = lastUpdatedAt;
    return this;
  }

  /**
   * Get lastUpdatedAt
   *
   * @return lastUpdatedAt
   */
  @javax.annotation.Nonnull
  public String getLastUpdatedAt() {
    return lastUpdatedAt;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    FunnelStageSuccess funnelStageSuccess = (FunnelStageSuccess) o;
    return Objects.equals(this.value, funnelStageSuccess.value) && Objects.equals(this.lastUpdatedAt, funnelStageSuccess.lastUpdatedAt);
  }

  @Override
  public int hashCode() {
    return Objects.hash(value, lastUpdatedAt);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class FunnelStageSuccess {\n");
    sb.append("    value: ").append(toIndentedString(value)).append("\n");
    sb.append("    lastUpdatedAt: ").append(toIndentedString(lastUpdatedAt)).append("\n");
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
