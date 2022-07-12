// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.analytics;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** TopHit */
public class TopHit {

  @JsonProperty("hit")
  private String hit;

  @JsonProperty("count")
  private Integer count;

  public TopHit setHit(String hit) {
    this.hit = hit;
    return this;
  }

  /**
   * The hit.
   *
   * @return hit
   */
  @javax.annotation.Nonnull
  public String getHit() {
    return hit;
  }

  public TopHit setCount(Integer count) {
    this.count = count;
    return this;
  }

  /**
   * The number of occurrences.
   *
   * @return count
   */
  @javax.annotation.Nonnull
  public Integer getCount() {
    return count;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    TopHit topHit = (TopHit) o;
    return Objects.equals(this.hit, topHit.hit) && Objects.equals(this.count, topHit.count);
  }

  @Override
  public int hashCode() {
    return Objects.hash(hit, count);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class TopHit {\n");
    sb.append("    hit: ").append(toIndentedString(hit)).append("\n");
    sb.append("    count: ").append(toIndentedString(count)).append("\n");
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
