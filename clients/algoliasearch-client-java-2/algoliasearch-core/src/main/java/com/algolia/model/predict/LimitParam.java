// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.predict;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** LimitParam */
public class LimitParam {

  @JsonProperty("limit")
  private Integer limit;

  public LimitParam setLimit(Integer limit) {
    this.limit = limit;
    return this;
  }

  /**
   * Get limit
   *
   * @return limit
   */
  @javax.annotation.Nullable
  public Integer getLimit() {
    return limit;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    LimitParam limitParam = (LimitParam) o;
    return Objects.equals(this.limit, limitParam.limit);
  }

  @Override
  public int hashCode() {
    return Objects.hash(limit);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class LimitParam {\n");
    sb.append("    limit: ").append(toIndentedString(limit)).append("\n");
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
