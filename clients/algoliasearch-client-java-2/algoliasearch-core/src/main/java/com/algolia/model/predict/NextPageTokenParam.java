// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.predict;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** NextPageTokenParam */
public class NextPageTokenParam {

  @JsonProperty("nextPageToken")
  private String nextPageToken;

  public NextPageTokenParam setNextPageToken(String nextPageToken) {
    this.nextPageToken = nextPageToken;
    return this;
  }

  /**
   * The token is used to navigate forward in the user list. To navigate from the current user list
   * to the next page, the API generates the next page token and it sends the token in the response,
   * beside the current user list. NOTE: This body param cannot be used with `previousPageToken` at
   * the same time.
   *
   * @return nextPageToken
   */
  @javax.annotation.Nullable
  public String getNextPageToken() {
    return nextPageToken;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    NextPageTokenParam nextPageTokenParam = (NextPageTokenParam) o;
    return Objects.equals(this.nextPageToken, nextPageTokenParam.nextPageToken);
  }

  @Override
  public int hashCode() {
    return Objects.hash(nextPageToken);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class NextPageTokenParam {\n");
    sb.append("    nextPageToken: ").append(toIndentedString(nextPageToken)).append("\n");
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
