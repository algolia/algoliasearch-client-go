// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.search;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** SearchForHitsOptions */
public class SearchForHitsOptions {

  @JsonProperty("indexName")
  private String indexName;

  @JsonProperty("type")
  private SearchTypeDefault type;

  public SearchForHitsOptions setIndexName(String indexName) {
    this.indexName = indexName;
    return this;
  }

  /**
   * The Algolia index name.
   *
   * @return indexName
   */
  @javax.annotation.Nonnull
  public String getIndexName() {
    return indexName;
  }

  public SearchForHitsOptions setType(SearchTypeDefault type) {
    this.type = type;
    return this;
  }

  /**
   * Get type
   *
   * @return type
   */
  @javax.annotation.Nullable
  public SearchTypeDefault getType() {
    return type;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SearchForHitsOptions searchForHitsOptions = (SearchForHitsOptions) o;
    return Objects.equals(this.indexName, searchForHitsOptions.indexName) && Objects.equals(this.type, searchForHitsOptions.type);
  }

  @Override
  public int hashCode() {
    return Objects.hash(indexName, type);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SearchForHitsOptions {\n");
    sb.append("    indexName: ").append(toIndentedString(indexName)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
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
