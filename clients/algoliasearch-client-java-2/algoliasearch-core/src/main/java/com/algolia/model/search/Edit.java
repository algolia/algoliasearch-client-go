package com.algolia.model.search;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** Edit */
public class Edit {

  @JsonProperty("type")
  private EditType type;

  @JsonProperty("delete")
  private String delete;

  @JsonProperty("insert")
  private String insert;

  public Edit setType(EditType type) {
    this.type = type;
    return this;
  }

  /**
   * Get type
   *
   * @return type
   */
  @javax.annotation.Nullable
  public EditType getType() {
    return type;
  }

  public Edit setDelete(String delete) {
    this.delete = delete;
    return this;
  }

  /**
   * Text or patterns to remove from the query string.
   *
   * @return delete
   */
  @javax.annotation.Nullable
  public String getDelete() {
    return delete;
  }

  public Edit setInsert(String insert) {
    this.insert = insert;
    return this;
  }

  /**
   * Text that should be inserted in place of the removed text inside the query string.
   *
   * @return insert
   */
  @javax.annotation.Nullable
  public String getInsert() {
    return insert;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Edit edit = (Edit) o;
    return Objects.equals(this.type, edit.type) && Objects.equals(this.delete, edit.delete) && Objects.equals(this.insert, edit.insert);
  }

  @Override
  public int hashCode() {
    return Objects.hash(type, delete, insert);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Edit {\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    delete: ").append(toIndentedString(delete)).append("\n");
    sb.append("    insert: ").append(toIndentedString(insert)).append("\n");
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
