// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.predict;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** CreateSegmentParams */
public class CreateSegmentParams {

  @JsonProperty("name")
  private String name;

  @JsonProperty("conditions")
  private String conditions;

  public CreateSegmentParams setName(String name) {
    this.name = name;
    return this;
  }

  /**
   * The name or description of the segment.
   *
   * @return name
   */
  @javax.annotation.Nonnull
  public String getName() {
    return name;
  }

  public CreateSegmentParams setConditions(String conditions) {
    this.conditions = conditions;
    return this;
  }

  /**
   * The filters that define the segment, defined in the same way as filters for Rules.
   *
   * @return conditions
   */
  @javax.annotation.Nonnull
  public String getConditions() {
    return conditions;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    CreateSegmentParams createSegmentParams = (CreateSegmentParams) o;
    return Objects.equals(this.name, createSegmentParams.name) && Objects.equals(this.conditions, createSegmentParams.conditions);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, conditions);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class CreateSegmentParams {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    conditions: ").append(toIndentedString(conditions)).append("\n");
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
