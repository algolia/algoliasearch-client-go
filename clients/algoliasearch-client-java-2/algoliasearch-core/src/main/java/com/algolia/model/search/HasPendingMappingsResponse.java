package com.algolia.model.search;

import com.fasterxml.jackson.annotation.*;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;

/** HasPendingMappingsResponse */
public class HasPendingMappingsResponse {

  @JsonProperty("pending")
  private Boolean pending;

  @JsonProperty("clusters")
  private Map<String, List<String>> clusters;

  public HasPendingMappingsResponse setPending(Boolean pending) {
    this.pending = pending;
    return this;
  }

  /**
   * If there is any clusters with pending mapping state.
   *
   * @return pending
   */
  @javax.annotation.Nonnull
  public Boolean getPending() {
    return pending;
  }

  public HasPendingMappingsResponse setClusters(Map<String, List<String>> clusters) {
    this.clusters = clusters;
    return this;
  }

  public HasPendingMappingsResponse putClusters(String key, List<String> clustersItem) {
    if (this.clusters == null) {
      this.clusters = new HashMap<>();
    }
    this.clusters.put(key, clustersItem);
    return this;
  }

  /**
   * Describe cluster pending (migrating, creating, deleting) mapping state.
   *
   * @return clusters
   */
  @javax.annotation.Nullable
  public Map<String, List<String>> getClusters() {
    return clusters;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    HasPendingMappingsResponse hasPendingMappingsResponse = (HasPendingMappingsResponse) o;
    return (
      Objects.equals(this.pending, hasPendingMappingsResponse.pending) && Objects.equals(this.clusters, hasPendingMappingsResponse.clusters)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(pending, clusters);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class HasPendingMappingsResponse {\n");
    sb.append("    pending: ").append(toIndentedString(pending)).append("\n");
    sb.append("    clusters: ").append(toIndentedString(clusters)).append("\n");
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
