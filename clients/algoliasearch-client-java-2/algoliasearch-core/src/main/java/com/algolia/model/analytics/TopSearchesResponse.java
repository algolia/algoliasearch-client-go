// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.analytics;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/** TopSearchesResponse */
public class TopSearchesResponse {

  @JsonProperty("searches")
  private List<TopSearch> searches = new ArrayList<>();

  public TopSearchesResponse setSearches(List<TopSearch> searches) {
    this.searches = searches;
    return this;
  }

  public TopSearchesResponse addSearches(TopSearch searchesItem) {
    this.searches.add(searchesItem);
    return this;
  }

  /**
   * A list of top searches with their count.
   *
   * @return searches
   */
  @javax.annotation.Nonnull
  public List<TopSearch> getSearches() {
    return searches;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    TopSearchesResponse topSearchesResponse = (TopSearchesResponse) o;
    return Objects.equals(this.searches, topSearchesResponse.searches);
  }

  @Override
  public int hashCode() {
    return Objects.hash(searches);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class TopSearchesResponse {\n");
    sb.append("    searches: ").append(toIndentedString(searches)).append("\n");
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
