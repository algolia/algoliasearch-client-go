// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.predict;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** AllUpdateSegmentParams */
public class AllUpdateSegmentParams {

  @JsonProperty("name")
  private String name;

  @JsonProperty("conditions")
  private String conditions;

  public AllUpdateSegmentParams setName(String name) {
    this.name = name;
    return this;
  }

  /**
   * The name or description of the segment.
   *
   * @return name
   */
  @javax.annotation.Nullable
  public String getName() {
    return name;
  }

  public AllUpdateSegmentParams setConditions(String conditions) {
    this.conditions = conditions;
    return this;
  }

  /**
   * The filters that define the segment, defined in the same way as filters for Rules.
   *
   * @return conditions
   */
  @javax.annotation.Nullable
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
    AllUpdateSegmentParams allUpdateSegmentParams = (AllUpdateSegmentParams) o;
    return Objects.equals(this.name, allUpdateSegmentParams.name) && Objects.equals(this.conditions, allUpdateSegmentParams.conditions);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, conditions);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class AllUpdateSegmentParams {\n");
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
