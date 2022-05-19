package com.algolia.model.search;

import com.google.gson.annotations.SerializedName;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/** SearchMethodParams */
public class SearchMethodParams {

  @SerializedName("requests")
  private List<SearchQueries> requests = new ArrayList<>();

  @SerializedName("strategy")
  private SearchStrategy strategy;

  public SearchMethodParams setRequests(List<SearchQueries> requests) {
    this.requests = requests;
    return this;
  }

  public SearchMethodParams addRequests(SearchQueries requestsItem) {
    this.requests.add(requestsItem);
    return this;
  }

  /**
   * Get requests
   *
   * @return requests
   */
  @javax.annotation.Nonnull
  public List<SearchQueries> getRequests() {
    return requests;
  }

  public SearchMethodParams setStrategy(SearchStrategy strategy) {
    this.strategy = strategy;
    return this;
  }

  /**
   * Get strategy
   *
   * @return strategy
   */
  @javax.annotation.Nullable
  public SearchStrategy getStrategy() {
    return strategy;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SearchMethodParams searchMethodParams = (SearchMethodParams) o;
    return (
      Objects.equals(this.requests, searchMethodParams.requests) &&
      Objects.equals(this.strategy, searchMethodParams.strategy)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(requests, strategy);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SearchMethodParams {\n");
    sb.append("    requests: ").append(toIndentedString(requests)).append("\n");
    sb.append("    strategy: ").append(toIndentedString(strategy)).append("\n");
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
