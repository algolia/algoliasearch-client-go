// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.analytics;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** TopCountry */
public class TopCountry {

  @JsonProperty("country")
  private String country;

  @JsonProperty("count")
  private Integer count;

  public TopCountry setCountry(String country) {
    this.country = country;
    return this;
  }

  /**
   * The country.
   *
   * @return country
   */
  @javax.annotation.Nonnull
  public String getCountry() {
    return country;
  }

  public TopCountry setCount(Integer count) {
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
    TopCountry topCountry = (TopCountry) o;
    return Objects.equals(this.country, topCountry.country) && Objects.equals(this.count, topCountry.count);
  }

  @Override
  public int hashCode() {
    return Objects.hash(country, count);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class TopCountry {\n");
    sb.append("    country: ").append(toIndentedString(country)).append("\n");
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
