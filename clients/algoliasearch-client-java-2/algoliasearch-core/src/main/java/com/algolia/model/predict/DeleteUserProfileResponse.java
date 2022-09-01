// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.predict;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** DeleteUserProfileResponse */
public class DeleteUserProfileResponse {

  @JsonProperty("user")
  private String user;

  @JsonProperty("deletedUntil")
  private String deletedUntil;

  public DeleteUserProfileResponse setUser(String user) {
    this.user = user;
    return this;
  }

  /**
   * The ID of the user that was deleted.
   *
   * @return user
   */
  @javax.annotation.Nonnull
  public String getUser() {
    return user;
  }

  public DeleteUserProfileResponse setDeletedUntil(String deletedUntil) {
    this.deletedUntil = deletedUntil;
    return this;
  }

  /**
   * The time the same user ID will be imported again when the data is ingested.
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
    DeleteUserProfileResponse deleteUserProfileResponse = (DeleteUserProfileResponse) o;
    return (
      Objects.equals(this.user, deleteUserProfileResponse.user) && Objects.equals(this.deletedUntil, deleteUserProfileResponse.deletedUntil)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(user, deletedUntil);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class DeleteUserProfileResponse {\n");
    sb.append("    user: ").append(toIndentedString(user)).append("\n");
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
