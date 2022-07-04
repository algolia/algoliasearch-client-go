package com.algolia.model.search;

import com.google.gson.annotations.SerializedName;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/** ConsequenceQueryObject */
public class ConsequenceQueryObject {

  @SerializedName("remove")
  private List<String> remove;

  @SerializedName("edits")
  private List<Edit> edits;

  public ConsequenceQueryObject setRemove(List<String> remove) {
    this.remove = remove;
    return this;
  }

  public ConsequenceQueryObject addRemove(String removeItem) {
    if (this.remove == null) {
      this.remove = new ArrayList<>();
    }
    this.remove.add(removeItem);
    return this;
  }

  /**
   * Words to remove.
   *
   * @return remove
   */
  @javax.annotation.Nullable
  public List<String> getRemove() {
    return remove;
  }

  public ConsequenceQueryObject setEdits(List<Edit> edits) {
    this.edits = edits;
    return this;
  }

  public ConsequenceQueryObject addEdits(Edit editsItem) {
    if (this.edits == null) {
      this.edits = new ArrayList<>();
    }
    this.edits.add(editsItem);
    return this;
  }

  /**
   * Edits to apply.
   *
   * @return edits
   */
  @javax.annotation.Nullable
  public List<Edit> getEdits() {
    return edits;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ConsequenceQueryObject consequenceQueryObject = (ConsequenceQueryObject) o;
    return Objects.equals(this.remove, consequenceQueryObject.remove) && Objects.equals(this.edits, consequenceQueryObject.edits);
  }

  @Override
  public int hashCode() {
    return Objects.hash(remove, edits);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ConsequenceQueryObject {\n");
    sb.append("    remove: ").append(toIndentedString(remove)).append("\n");
    sb.append("    edits: ").append(toIndentedString(edits)).append("\n");
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
