// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.predict;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** ActivateModelInstanceResponse */
public class ActivateModelInstanceResponse {

  @JsonProperty("modelID")
  private String modelID;

  @JsonProperty("updatedAt")
  private String updatedAt;

  public ActivateModelInstanceResponse setModelID(String modelID) {
    this.modelID = modelID;
    return this;
  }

  /**
   * The ID of the model.
   *
   * @return modelID
   */
  @javax.annotation.Nonnull
  public String getModelID() {
    return modelID;
  }

  public ActivateModelInstanceResponse setUpdatedAt(String updatedAt) {
    this.updatedAt = updatedAt;
    return this;
  }

  /**
   * Date of last update (ISO-8601 format).
   *
   * @return updatedAt
   */
  @javax.annotation.Nonnull
  public String getUpdatedAt() {
    return updatedAt;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ActivateModelInstanceResponse activateModelInstanceResponse = (ActivateModelInstanceResponse) o;
    return (
      Objects.equals(this.modelID, activateModelInstanceResponse.modelID) &&
      Objects.equals(this.updatedAt, activateModelInstanceResponse.updatedAt)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(modelID, updatedAt);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ActivateModelInstanceResponse {\n");
    sb.append("    modelID: ").append(toIndentedString(modelID)).append("\n");
    sb.append("    updatedAt: ").append(toIndentedString(updatedAt)).append("\n");
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
