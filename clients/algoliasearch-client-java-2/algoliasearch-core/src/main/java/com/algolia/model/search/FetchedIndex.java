// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.search;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/** FetchedIndex */
public class FetchedIndex {

  @JsonProperty("name")
  private String name;

  @JsonProperty("createdAt")
  private String createdAt;

  @JsonProperty("updatedAt")
  private String updatedAt;

  @JsonProperty("entries")
  private Integer entries;

  @JsonProperty("dataSize")
  private Integer dataSize;

  @JsonProperty("fileSize")
  private Integer fileSize;

  @JsonProperty("lastBuildTimeS")
  private Integer lastBuildTimeS;

  @JsonProperty("numberOfPendingTask")
  private Integer numberOfPendingTask;

  @JsonProperty("pendingTask")
  private Boolean pendingTask;

  @JsonProperty("primary")
  private String primary;

  @JsonProperty("replicas")
  private List<String> replicas;

  public FetchedIndex setName(String name) {
    this.name = name;
    return this;
  }

  /**
   * Index name.
   *
   * @return name
   */
  @javax.annotation.Nonnull
  public String getName() {
    return name;
  }

  public FetchedIndex setCreatedAt(String createdAt) {
    this.createdAt = createdAt;
    return this;
  }

  /**
   * Index creation date. An empty string means that the index has no records.
   *
   * @return createdAt
   */
  @javax.annotation.Nonnull
  public String getCreatedAt() {
    return createdAt;
  }

  public FetchedIndex setUpdatedAt(String updatedAt) {
    this.updatedAt = updatedAt;
    return this;
  }

  /**
   * Date of last update (ISO-8601 format).
   *
   * @return updatedAt
   */
  @javax.annotation.Nonnull
  public String getUpdatedAt() {
    return updatedAt;
  }

  public FetchedIndex setEntries(Integer entries) {
    this.entries = entries;
    return this;
  }

  /**
   * Number of records contained in the index.
   *
   * @return entries
   */
  @javax.annotation.Nonnull
  public Integer getEntries() {
    return entries;
  }

  public FetchedIndex setDataSize(Integer dataSize) {
    this.dataSize = dataSize;
    return this;
  }

  /**
   * Number of bytes of the index in minified format.
   *
   * @return dataSize
   */
  @javax.annotation.Nonnull
  public Integer getDataSize() {
    return dataSize;
  }

  public FetchedIndex setFileSize(Integer fileSize) {
    this.fileSize = fileSize;
    return this;
  }

  /**
   * Number of bytes of the index binary file.
   *
   * @return fileSize
   */
  @javax.annotation.Nonnull
  public Integer getFileSize() {
    return fileSize;
  }

  public FetchedIndex setLastBuildTimeS(Integer lastBuildTimeS) {
    this.lastBuildTimeS = lastBuildTimeS;
    return this;
  }

  /**
   * Last build time.
   *
   * @return lastBuildTimeS
   */
  @javax.annotation.Nonnull
  public Integer getLastBuildTimeS() {
    return lastBuildTimeS;
  }

  public FetchedIndex setNumberOfPendingTask(Integer numberOfPendingTask) {
    this.numberOfPendingTask = numberOfPendingTask;
    return this;
  }

  /**
   * Number of pending indexing operations. This value is deprecated and should not be used.
   *
   * @return numberOfPendingTask
   */
  @javax.annotation.Nullable
  public Integer getNumberOfPendingTask() {
    return numberOfPendingTask;
  }

  public FetchedIndex setPendingTask(Boolean pendingTask) {
    this.pendingTask = pendingTask;
    return this;
  }

  /**
   * A boolean which says whether the index has pending tasks. This value is deprecated and should
   * not be used.
   *
   * @return pendingTask
   */
  @javax.annotation.Nonnull
  public Boolean getPendingTask() {
    return pendingTask;
  }

  public FetchedIndex setPrimary(String primary) {
    this.primary = primary;
    return this;
  }

  /**
   * Only present if the index is a replica. Contains the name of the related primary index.
   *
   * @return primary
   */
  @javax.annotation.Nullable
  public String getPrimary() {
    return primary;
  }

  public FetchedIndex setReplicas(List<String> replicas) {
    this.replicas = replicas;
    return this;
  }

  public FetchedIndex addReplicas(String replicasItem) {
    if (this.replicas == null) {
      this.replicas = new ArrayList<>();
    }
    this.replicas.add(replicasItem);
    return this;
  }

  /**
   * Only present if the index is a primary index with replicas. Contains the names of all linked
   * replicas.
   *
   * @return replicas
   */
  @javax.annotation.Nullable
  public List<String> getReplicas() {
    return replicas;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    FetchedIndex fetchedIndex = (FetchedIndex) o;
    return (
      Objects.equals(this.name, fetchedIndex.name) &&
      Objects.equals(this.createdAt, fetchedIndex.createdAt) &&
      Objects.equals(this.updatedAt, fetchedIndex.updatedAt) &&
      Objects.equals(this.entries, fetchedIndex.entries) &&
      Objects.equals(this.dataSize, fetchedIndex.dataSize) &&
      Objects.equals(this.fileSize, fetchedIndex.fileSize) &&
      Objects.equals(this.lastBuildTimeS, fetchedIndex.lastBuildTimeS) &&
      Objects.equals(this.numberOfPendingTask, fetchedIndex.numberOfPendingTask) &&
      Objects.equals(this.pendingTask, fetchedIndex.pendingTask) &&
      Objects.equals(this.primary, fetchedIndex.primary) &&
      Objects.equals(this.replicas, fetchedIndex.replicas)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(
      name,
      createdAt,
      updatedAt,
      entries,
      dataSize,
      fileSize,
      lastBuildTimeS,
      numberOfPendingTask,
      pendingTask,
      primary,
      replicas
    );
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class FetchedIndex {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    createdAt: ").append(toIndentedString(createdAt)).append("\n");
    sb.append("    updatedAt: ").append(toIndentedString(updatedAt)).append("\n");
    sb.append("    entries: ").append(toIndentedString(entries)).append("\n");
    sb.append("    dataSize: ").append(toIndentedString(dataSize)).append("\n");
    sb.append("    fileSize: ").append(toIndentedString(fileSize)).append("\n");
    sb.append("    lastBuildTimeS: ").append(toIndentedString(lastBuildTimeS)).append("\n");
    sb.append("    numberOfPendingTask: ").append(toIndentedString(numberOfPendingTask)).append("\n");
    sb.append("    pendingTask: ").append(toIndentedString(pendingTask)).append("\n");
    sb.append("    primary: ").append(toIndentedString(primary)).append("\n");
    sb.append("    replicas: ").append(toIndentedString(replicas)).append("\n");
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
