// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.predict;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** PreviousPageTokenParam */
public class PreviousPageTokenParam {

  @JsonProperty("previousPageToken")
  private String previousPageToken;

  public PreviousPageTokenParam setPreviousPageToken(String previousPageToken) {
    this.previousPageToken = previousPageToken;
    return this;
  }

  /**
   * The token is used to navigate backward in the user list. To navigate from the current user list
   * to the previous page, the API generates the previous page token and it sends the token in the
   * response, beside the current user list. NOTE: This body param cannot be used with
   * `nextPageToken` at the same time.
   *
   * @return previousPageToken
   */
  @javax.annotation.Nullable
  public String getPreviousPageToken() {
    return previousPageToken;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    PreviousPageTokenParam previousPageTokenParam = (PreviousPageTokenParam) o;
    return Objects.equals(this.previousPageToken, previousPageTokenParam.previousPageToken);
  }

  @Override
  public int hashCode() {
    return Objects.hash(previousPageToken);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class PreviousPageTokenParam {\n");
    sb.append("    previousPageToken: ").append(toIndentedString(previousPageToken)).append("\n");
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
