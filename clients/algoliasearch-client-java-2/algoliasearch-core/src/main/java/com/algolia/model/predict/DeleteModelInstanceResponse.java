// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.predict;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** DeleteModelInstanceResponse */
public class DeleteModelInstanceResponse {

  @JsonProperty("modelID")
  private String modelID;

  @JsonProperty("deletedUntil")
  private String deletedUntil;

  public DeleteModelInstanceResponse setModelID(String modelID) {
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

  public DeleteModelInstanceResponse setDeletedUntil(String deletedUntil) {
    this.deletedUntil = deletedUntil;
    return this;
  }

  /**
   * The date until which you can safely consider the data as being deleted.
   *
   * @return deletedUntil
   */
  @javax.annotation.Nonnull
  public String getDeletedUntil() {
    return deletedUntil;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    DeleteModelInstanceResponse deleteModelInstanceResponse = (DeleteModelInstanceResponse) o;
    return (
      Objects.equals(this.modelID, deleteModelInstanceResponse.modelID) &&
      Objects.equals(this.deletedUntil, deleteModelInstanceResponse.deletedUntil)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(modelID, deletedUntil);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class DeleteModelInstanceResponse {\n");
    sb.append("    modelID: ").append(toIndentedString(modelID)).append("\n");
    sb.append("    deletedUntil: ").append(toIndentedString(deletedUntil)).append("\n");
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
