package com.algolia.model.search;

import com.google.gson.annotations.SerializedName;
import java.util.Objects;

/** Single objectID to promote as hits. */
public class PromoteObjectID {

  @SerializedName("objectID")
  private String objectID;

  @SerializedName("position")
  private Integer position;

  public PromoteObjectID setObjectID(String objectID) {
    this.objectID = objectID;
    return this;
  }

  /**
   * Unique identifier of the object to promote.
   *
   * @return objectID
   */
  @javax.annotation.Nonnull
  public String getObjectID() {
    return objectID;
  }

  public PromoteObjectID setPosition(Integer position) {
    this.position = position;
    return this;
  }

  /**
   * The position to promote the objects to (zero-based). If you pass objectIDs, the objects are
   * placed at this position as a group. For example, if you pass four objectIDs to position 0, the
   * objects take the first four positions.
   *
   * @return position
   */
  @javax.annotation.Nonnull
  public Integer getPosition() {
    return position;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    PromoteObjectID promoteObjectID = (PromoteObjectID) o;
    return Objects.equals(this.objectID, promoteObjectID.objectID) && Objects.equals(this.position, promoteObjectID.position);
  }

  @Override
  public int hashCode() {
    return Objects.hash(objectID, position);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class PromoteObjectID {\n");
    sb.append("    objectID: ").append(toIndentedString(objectID)).append("\n");
    sb.append("    position: ").append(toIndentedString(position)).append("\n");
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
