// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.search;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/** SearchResponses */
public class SearchResponses<T> {

  @JsonProperty("results")
  private List<SearchResponse<T>> results = new ArrayList<>();

  public SearchResponses setResults(List<SearchResponse<T>> results) {
    this.results = results;
    return this;
  }

  public SearchResponses addResults(SearchResponse<T> resultsItem) {
    this.results.add(resultsItem);
    return this;
  }

  /**
   * Get results
   *
   * @return results
   */
  @javax.annotation.Nonnull
  public List<SearchResponse<T>> getResults() {
    return results;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SearchResponses searchResponses = (SearchResponses) o;
    return Objects.equals(this.results, searchResponses.results);
  }

  @Override
  public int hashCode() {
    return Objects.hash(results);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SearchResponses {\n");
    sb.append("    results: ").append(toIndentedString(results)).append("\n");
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
