// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.predict;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** Prediction for the **order_value** model. */
public class OrderValueSuccess {

  @JsonProperty("value")
  private Double value;

  @JsonProperty("lastUpdatedAt")
  private String lastUpdatedAt;

  public OrderValueSuccess setValue(Double value) {
    this.value = value;
    return this;
  }

  /**
   * Get value minimum: 0
   *
   * @return value
   */
  @javax.annotation.Nonnull
  public Double getValue() {
    return value;
  }

  public OrderValueSuccess setLastUpdatedAt(String lastUpdatedAt) {
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
    OrderValueSuccess orderValueSuccess = (OrderValueSuccess) o;
    return Objects.equals(this.value, orderValueSuccess.value) && Objects.equals(this.lastUpdatedAt, orderValueSuccess.lastUpdatedAt);
  }

  @Override
  public int hashCode() {
    return Objects.hash(value, lastUpdatedAt);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class OrderValueSuccess {\n");
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
