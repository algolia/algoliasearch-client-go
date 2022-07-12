// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.predict;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** FunnelStage */
public class FunnelStage {

  @JsonProperty("name")
  private String name;

  @JsonProperty("probability")
  private Double probability;

  public FunnelStage setName(String name) {
    this.name = name;
    return this;
  }

  /**
   * Get name
   *
   * @return name
   */
  @javax.annotation.Nullable
  public String getName() {
    return name;
  }

  public FunnelStage setProbability(Double probability) {
    this.probability = probability;
    return this;
  }

  /**
   * Get probability minimum: 0 maximum: 1
   *
   * @return probability
   */
  @javax.annotation.Nullable
  public Double getProbability() {
    return probability;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    FunnelStage funnelStage = (FunnelStage) o;
    return Objects.equals(this.name, funnelStage.name) && Objects.equals(this.probability, funnelStage.probability);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, probability);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class FunnelStage {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    probability: ").append(toIndentedString(probability)).append("\n");
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
