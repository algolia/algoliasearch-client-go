// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.predict;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** Affinity */
public class Affinity {

  @JsonProperty("name")
  private String name;

  @JsonProperty("value")
  private String value;

  @JsonProperty("probability")
  private Double probability;

  public Affinity setName(String name) {
    this.name = name;
    return this;
  }

  /**
   * Get name
   *
   * @return name
   */
  @javax.annotation.Nonnull
  public String getName() {
    return name;
  }

  public Affinity setValue(String value) {
    this.value = value;
    return this;
  }

  /**
   * Get value
   *
   * @return value
   */
  @javax.annotation.Nonnull
  public String getValue() {
    return value;
  }

  public Affinity setProbability(Double probability) {
    this.probability = probability;
    return this;
  }

  /**
   * Get probability minimum: 0 maximum: 1
   *
   * @return probability
   */
  @javax.annotation.Nonnull
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
    Affinity affinity = (Affinity) o;
    return (
      Objects.equals(this.name, affinity.name) &&
      Objects.equals(this.value, affinity.value) &&
      Objects.equals(this.probability, affinity.probability)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, value, probability);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Affinity {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    value: ").append(toIndentedString(value)).append("\n");
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
