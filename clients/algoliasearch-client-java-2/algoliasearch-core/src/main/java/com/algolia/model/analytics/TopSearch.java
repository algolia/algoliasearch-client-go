// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.analytics;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** TopSearch */
public class TopSearch {

  @JsonProperty("search")
  private String search;

  @JsonProperty("count")
  private Integer count;

  @JsonProperty("nbHits")
  private Integer nbHits;

  public TopSearch setSearch(String search) {
    this.search = search;
    return this;
  }

  /**
   * The search query.
   *
   * @return search
   */
  @javax.annotation.Nonnull
  public String getSearch() {
    return search;
  }

  public TopSearch setCount(Integer count) {
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

  public TopSearch setNbHits(Integer nbHits) {
    this.nbHits = nbHits;
    return this;
  }

  /**
   * Number of hits that the search query matched.
   *
   * @return nbHits
   */
  @javax.annotation.Nonnull
  public Integer getNbHits() {
    return nbHits;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    TopSearch topSearch = (TopSearch) o;
    return (
      Objects.equals(this.search, topSearch.search) &&
      Objects.equals(this.count, topSearch.count) &&
      Objects.equals(this.nbHits, topSearch.nbHits)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(search, count, nbHits);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class TopSearch {\n");
    sb.append("    search: ").append(toIndentedString(search)).append("\n");
    sb.append("    count: ").append(toIndentedString(count)).append("\n");
    sb.append("    nbHits: ").append(toIndentedString(nbHits)).append("\n");
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
