// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.predict;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** SegmentConditionsParam */
public class SegmentConditionsParam {

  @JsonProperty("conditions")
  private String conditions;

  public SegmentConditionsParam setConditions(String conditions) {
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
    SegmentConditionsParam segmentConditionsParam = (SegmentConditionsParam) o;
    return Objects.equals(this.conditions, segmentConditionsParam.conditions);
  }

  @Override
  public int hashCode() {
    return Objects.hash(conditions);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SegmentConditionsParam {\n");
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
