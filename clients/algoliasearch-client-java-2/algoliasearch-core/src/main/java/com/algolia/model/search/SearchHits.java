// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.search;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/** SearchHits */
public class SearchHits<T> {

  @JsonProperty("hits")
  private List<T> hits = new ArrayList<>();

  public SearchHits setHits(List<T> hits) {
    this.hits = hits;
    return this;
  }

  public SearchHits addHits(T hitsItem) {
    this.hits.add(hitsItem);
    return this;
  }

  /**
   * Get hits
   *
   * @return hits
   */
  @javax.annotation.Nonnull
  public List<T> getHits() {
    return hits;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SearchHits searchHits = (SearchHits) o;
    return Objects.equals(this.hits, searchHits.hits);
  }

  @Override
  public int hashCode() {
    return Objects.hash(hits);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SearchHits {\n");
    sb.append("    hits: ").append(toIndentedString(hits)).append("\n");
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
