// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.predict;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** GetAvailableModelTypesResponseInnerDataRequirements */
public class GetAvailableModelTypesResponseInnerDataRequirements {

  @JsonProperty("minUsers")
  private Integer minUsers;

  @JsonProperty("minDays")
  private Integer minDays;

  public GetAvailableModelTypesResponseInnerDataRequirements setMinUsers(Integer minUsers) {
    this.minUsers = minUsers;
    return this;
  }

  /**
   * Minimum number of users required for this model.
   *
   * @return minUsers
   */
  @javax.annotation.Nonnull
  public Integer getMinUsers() {
    return minUsers;
  }

  public GetAvailableModelTypesResponseInnerDataRequirements setMinDays(Integer minDays) {
    this.minDays = minDays;
    return this;
  }

  /**
   * Minimum number of days model needs to run.
   *
   * @return minDays
   */
  @javax.annotation.Nonnull
  public Integer getMinDays() {
    return minDays;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    GetAvailableModelTypesResponseInnerDataRequirements getAvailableModelTypesResponseInnerDataRequirements =
      (GetAvailableModelTypesResponseInnerDataRequirements) o;
    return (
      Objects.equals(this.minUsers, getAvailableModelTypesResponseInnerDataRequirements.minUsers) &&
      Objects.equals(this.minDays, getAvailableModelTypesResponseInnerDataRequirements.minDays)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(minUsers, minDays);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class GetAvailableModelTypesResponseInnerDataRequirements {\n");
    sb.append("    minUsers: ").append(toIndentedString(minUsers)).append("\n");
    sb.append("    minDays: ").append(toIndentedString(minDays)).append("\n");
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
