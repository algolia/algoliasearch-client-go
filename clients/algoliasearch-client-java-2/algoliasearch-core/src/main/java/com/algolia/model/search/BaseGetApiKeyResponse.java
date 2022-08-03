// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.search;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** BaseGetApiKeyResponse */
public class BaseGetApiKeyResponse {

  @JsonProperty("value")
  private String value;

  @JsonProperty("createdAt")
  private Long createdAt;

  public BaseGetApiKeyResponse setValue(String value) {
    this.value = value;
    return this;
  }

  /**
   * The API key.
   *
   * @return value
   */
  @javax.annotation.Nullable
  public String getValue() {
    return value;
  }

  public BaseGetApiKeyResponse setCreatedAt(Long createdAt) {
    this.createdAt = createdAt;
    return this;
  }

  /**
   * Time of the event expressed in milliseconds since the Unix epoch.
   *
   * @return createdAt
   */
  @javax.annotation.Nonnull
  public Long getCreatedAt() {
    return createdAt;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    BaseGetApiKeyResponse baseGetApiKeyResponse = (BaseGetApiKeyResponse) o;
    return Objects.equals(this.value, baseGetApiKeyResponse.value) && Objects.equals(this.createdAt, baseGetApiKeyResponse.createdAt);
  }

  @Override
  public int hashCode() {
    return Objects.hash(value, createdAt);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class BaseGetApiKeyResponse {\n");
    sb.append("    value: ").append(toIndentedString(value)).append("\n");
    sb.append("    createdAt: ").append(toIndentedString(createdAt)).append("\n");
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
