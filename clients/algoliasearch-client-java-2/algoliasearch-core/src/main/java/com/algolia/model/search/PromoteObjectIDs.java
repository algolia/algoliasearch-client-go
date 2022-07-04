package com.algolia.model.search;

import com.google.gson.annotations.SerializedName;
import java.util.List;
import java.util.Objects;

/** Multiple objectIDs to promote as hits. */
public class PromoteObjectIDs {

  @SerializedName("objectIDs")
  private List<String> objectIDs;

  @SerializedName("position")
  private Integer position;

  public PromoteObjectIDs setObjectIDs(List<String> objectIDs) {
    this.objectIDs = objectIDs;
    return this;
  }

  public PromoteObjectIDs addObjectIDs(String objectIDsItem) {
    this.objectIDs.add(objectIDsItem);
    return this;
  }

  /**
   * Array of unique identifiers of the objects to promote.
   *
   * @return objectIDs
   */
  @javax.annotation.Nonnull
  public List<String> getObjectIDs() {
    return objectIDs;
  }

  public PromoteObjectIDs setPosition(Integer position) {
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
    PromoteObjectIDs promoteObjectIDs = (PromoteObjectIDs) o;
    return Objects.equals(this.objectIDs, promoteObjectIDs.objectIDs) && Objects.equals(this.position, promoteObjectIDs.position);
  }

  @Override
  public int hashCode() {
    return Objects.hash(objectIDs, position);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class PromoteObjectIDs {\n");
    sb.append("    objectIDs: ").append(toIndentedString(objectIDs)).append("\n");
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
