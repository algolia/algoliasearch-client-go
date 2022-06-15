package com.algolia.model.search;

import com.google.gson.annotations.SerializedName;
import java.util.Objects;

/** UserHit */
public class UserHit {

  @SerializedName("userID")
  private String userID;

  @SerializedName("clusterName")
  private String clusterName;

  @SerializedName("nbRecords")
  private Integer nbRecords;

  @SerializedName("dataSize")
  private Integer dataSize;

  @SerializedName("objectID")
  private String objectID;

  @SerializedName("_highlightResult")
  private UserHighlightResult highlightResult;

  public UserHit setUserID(String userID) {
    this.userID = userID;
    return this;
  }

  /**
   * userID of the user.
   *
   * @return userID
   */
  @javax.annotation.Nonnull
  public String getUserID() {
    return userID;
  }

  public UserHit setClusterName(String clusterName) {
    this.clusterName = clusterName;
    return this;
  }

  /**
   * Name of the cluster.
   *
   * @return clusterName
   */
  @javax.annotation.Nonnull
  public String getClusterName() {
    return clusterName;
  }

  public UserHit setNbRecords(Integer nbRecords) {
    this.nbRecords = nbRecords;
    return this;
  }

  /**
   * Number of records in the cluster.
   *
   * @return nbRecords
   */
  @javax.annotation.Nonnull
  public Integer getNbRecords() {
    return nbRecords;
  }

  public UserHit setDataSize(Integer dataSize) {
    this.dataSize = dataSize;
    return this;
  }

  /**
   * Data size taken by all the users assigned to the cluster.
   *
   * @return dataSize
   */
  @javax.annotation.Nonnull
  public Integer getDataSize() {
    return dataSize;
  }

  public UserHit setObjectID(String objectID) {
    this.objectID = objectID;
    return this;
  }

  /**
   * userID of the requested user. Same as userID.
   *
   * @return objectID
   */
  @javax.annotation.Nonnull
  public String getObjectID() {
    return objectID;
  }

  public UserHit setHighlightResult(UserHighlightResult highlightResult) {
    this.highlightResult = highlightResult;
    return this;
  }

  /**
   * Get highlightResult
   *
   * @return highlightResult
   */
  @javax.annotation.Nonnull
  public UserHighlightResult getHighlightResult() {
    return highlightResult;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    UserHit userHit = (UserHit) o;
    return (
      Objects.equals(this.userID, userHit.userID) &&
      Objects.equals(this.clusterName, userHit.clusterName) &&
      Objects.equals(this.nbRecords, userHit.nbRecords) &&
      Objects.equals(this.dataSize, userHit.dataSize) &&
      Objects.equals(this.objectID, userHit.objectID) &&
      Objects.equals(this.highlightResult, userHit.highlightResult)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(userID, clusterName, nbRecords, dataSize, objectID, highlightResult);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class UserHit {\n");
    sb.append("    userID: ").append(toIndentedString(userID)).append("\n");
    sb.append("    clusterName: ").append(toIndentedString(clusterName)).append("\n");
    sb.append("    nbRecords: ").append(toIndentedString(nbRecords)).append("\n");
    sb.append("    dataSize: ").append(toIndentedString(dataSize)).append("\n");
    sb.append("    objectID: ").append(toIndentedString(objectID)).append("\n");
    sb.append("    highlightResult: ").append(toIndentedString(highlightResult)).append("\n");
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
