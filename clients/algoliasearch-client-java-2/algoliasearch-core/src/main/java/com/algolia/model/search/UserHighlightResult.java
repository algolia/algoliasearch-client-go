package com.algolia.model.search;

import com.google.gson.annotations.SerializedName;
import java.util.Objects;

/** UserHighlightResult */
public class UserHighlightResult {

  @SerializedName("userID")
  private HighlightResult userID;

  @SerializedName("clusterName")
  private HighlightResult clusterName;

  public UserHighlightResult setUserID(HighlightResult userID) {
    this.userID = userID;
    return this;
  }

  /**
   * Get userID
   *
   * @return userID
   */
  @javax.annotation.Nonnull
  public HighlightResult getUserID() {
    return userID;
  }

  public UserHighlightResult setClusterName(HighlightResult clusterName) {
    this.clusterName = clusterName;
    return this;
  }

  /**
   * Get clusterName
   *
   * @return clusterName
   */
  @javax.annotation.Nonnull
  public HighlightResult getClusterName() {
    return clusterName;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    UserHighlightResult userHighlightResult = (UserHighlightResult) o;
    return Objects.equals(this.userID, userHighlightResult.userID) && Objects.equals(this.clusterName, userHighlightResult.clusterName);
  }

  @Override
  public int hashCode() {
    return Objects.hash(userID, clusterName);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class UserHighlightResult {\n");
    sb.append("    userID: ").append(toIndentedString(userID)).append("\n");
    sb.append("    clusterName: ").append(toIndentedString(clusterName)).append("\n");
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
