// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.search;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/** The `batch` parameters. */
public class BatchWriteParams {

  @JsonProperty("requests")
  private List<BatchOperation> requests;

  public BatchWriteParams setRequests(List<BatchOperation> requests) {
    this.requests = requests;
    return this;
  }

  public BatchWriteParams addRequests(BatchOperation requestsItem) {
    if (this.requests == null) {
      this.requests = new ArrayList<>();
    }
    this.requests.add(requestsItem);
    return this;
  }

  /**
   * Get requests
   *
   * @return requests
   */
  @javax.annotation.Nullable
  public List<BatchOperation> getRequests() {
    return requests;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    BatchWriteParams batchWriteParams = (BatchWriteParams) o;
    return Objects.equals(this.requests, batchWriteParams.requests);
  }

  @Override
  public int hashCode() {
    return Objects.hash(requests);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class BatchWriteParams {\n");
    sb.append("    requests: ").append(toIndentedString(requests)).append("\n");
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
