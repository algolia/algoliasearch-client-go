// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.recommend;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/** Value */
public class Value {

  @JsonProperty("order")
  private List<String> order;

  @JsonProperty("sortRemainingBy")
  private SortRemainingBy sortRemainingBy;

  public Value setOrder(List<String> order) {
    this.order = order;
    return this;
  }

  public Value addOrder(String orderItem) {
    if (this.order == null) {
      this.order = new ArrayList<>();
    }
    this.order.add(orderItem);
    return this;
  }

  /**
   * Pinned order of facet lists.
   *
   * @return order
   */
  @javax.annotation.Nullable
  public List<String> getOrder() {
    return order;
  }

  public Value setSortRemainingBy(SortRemainingBy sortRemainingBy) {
    this.sortRemainingBy = sortRemainingBy;
    return this;
  }

  /**
   * Get sortRemainingBy
   *
   * @return sortRemainingBy
   */
  @javax.annotation.Nullable
  public SortRemainingBy getSortRemainingBy() {
    return sortRemainingBy;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Value value = (Value) o;
    return Objects.equals(this.order, value.order) && Objects.equals(this.sortRemainingBy, value.sortRemainingBy);
  }

  @Override
  public int hashCode() {
    return Objects.hash(order, sortRemainingBy);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Value {\n");
    sb.append("    order: ").append(toIndentedString(order)).append("\n");
    sb.append("    sortRemainingBy: ").append(toIndentedString(sortRemainingBy)).append("\n");
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
