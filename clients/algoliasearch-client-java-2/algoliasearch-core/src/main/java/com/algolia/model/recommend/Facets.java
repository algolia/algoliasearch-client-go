package com.algolia.model.recommend;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/** The ordering of facets (widgets). */
public class Facets {

  @JsonProperty("order")
  private List<String> order;

  public Facets setOrder(List<String> order) {
    this.order = order;
    return this;
  }

  public Facets addOrder(String orderItem) {
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

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Facets facets = (Facets) o;
    return Objects.equals(this.order, facets.order);
  }

  @Override
  public int hashCode() {
    return Objects.hash(order);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Facets {\n");
    sb.append("    order: ").append(toIndentedString(order)).append("\n");
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
