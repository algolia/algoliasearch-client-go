// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.search;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** FacetsStats */
public class FacetsStats {

  @JsonProperty("min")
  private Integer min;

  @JsonProperty("max")
  private Integer max;

  @JsonProperty("avg")
  private Integer avg;

  @JsonProperty("sum")
  private Integer sum;

  public FacetsStats setMin(Integer min) {
    this.min = min;
    return this;
  }

  /**
   * The minimum value in the result set.
   *
   * @return min
   */
  @javax.annotation.Nullable
  public Integer getMin() {
    return min;
  }

  public FacetsStats setMax(Integer max) {
    this.max = max;
    return this;
  }

  /**
   * The maximum value in the result set.
   *
   * @return max
   */
  @javax.annotation.Nullable
  public Integer getMax() {
    return max;
  }

  public FacetsStats setAvg(Integer avg) {
    this.avg = avg;
    return this;
  }

  /**
   * The average facet value in the result set.
   *
   * @return avg
   */
  @javax.annotation.Nullable
  public Integer getAvg() {
    return avg;
  }

  public FacetsStats setSum(Integer sum) {
    this.sum = sum;
    return this;
  }

  /**
   * The sum of all values in the result set.
   *
   * @return sum
   */
  @javax.annotation.Nullable
  public Integer getSum() {
    return sum;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    FacetsStats facetsStats = (FacetsStats) o;
    return (
      Objects.equals(this.min, facetsStats.min) &&
      Objects.equals(this.max, facetsStats.max) &&
      Objects.equals(this.avg, facetsStats.avg) &&
      Objects.equals(this.sum, facetsStats.sum)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(min, max, avg, sum);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class FacetsStats {\n");
    sb.append("    min: ").append(toIndentedString(min)).append("\n");
    sb.append("    max: ").append(toIndentedString(max)).append("\n");
    sb.append("    avg: ").append(toIndentedString(avg)).append("\n");
    sb.append("    sum: ").append(toIndentedString(sum)).append("\n");
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
