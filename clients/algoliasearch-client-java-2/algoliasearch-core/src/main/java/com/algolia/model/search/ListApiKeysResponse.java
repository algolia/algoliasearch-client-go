// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.search;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/** ListApiKeysResponse */
public class ListApiKeysResponse {

  @JsonProperty("keys")
  private List<GetApiKeyResponse> keys = new ArrayList<>();

  public ListApiKeysResponse setKeys(List<GetApiKeyResponse> keys) {
    this.keys = keys;
    return this;
  }

  public ListApiKeysResponse addKeys(GetApiKeyResponse keysItem) {
    this.keys.add(keysItem);
    return this;
  }

  /**
   * List of api keys.
   *
   * @return keys
   */
  @javax.annotation.Nonnull
  public List<GetApiKeyResponse> getKeys() {
    return keys;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ListApiKeysResponse listApiKeysResponse = (ListApiKeysResponse) o;
    return Objects.equals(this.keys, listApiKeysResponse.keys);
  }

  @Override
  public int hashCode() {
    return Objects.hash(keys);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ListApiKeysResponse {\n");
    sb.append("    keys: ").append(toIndentedString(keys)).append("\n");
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
