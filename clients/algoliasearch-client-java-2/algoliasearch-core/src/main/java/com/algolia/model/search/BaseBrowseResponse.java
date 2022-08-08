// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.search;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** BaseBrowseResponse */
public class BaseBrowseResponse {

  @JsonProperty("cursor")
  private String cursor;

  public BaseBrowseResponse setCursor(String cursor) {
    this.cursor = cursor;
    return this;
  }

  /**
   * Cursor indicating the location to resume browsing from. Must match the value returned by the
   * previous call.
   *
   * @return cursor
   */
  @javax.annotation.Nullable
  public String getCursor() {
    return cursor;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    BaseBrowseResponse baseBrowseResponse = (BaseBrowseResponse) o;
    return Objects.equals(this.cursor, baseBrowseResponse.cursor);
  }

  @Override
  public int hashCode() {
    return Objects.hash(cursor);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class BaseBrowseResponse {\n");
    sb.append("    cursor: ").append(toIndentedString(cursor)).append("\n");
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
