// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.predict;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** UpdateSegmentResponse */
public class UpdateSegmentResponse {

  @JsonProperty("segmentID")
  private String segmentID;

  @JsonProperty("updatedAt")
  private String updatedAt;

  public UpdateSegmentResponse setSegmentID(String segmentID) {
    this.segmentID = segmentID;
    return this;
  }

  /**
   * The ID of the segment.
   *
   * @return segmentID
   */
  @javax.annotation.Nonnull
  public String getSegmentID() {
    return segmentID;
  }

  public UpdateSegmentResponse setUpdatedAt(String updatedAt) {
    this.updatedAt = updatedAt;
    return this;
  }

  /**
   * The date and time at which the segment was last updated.
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
    UpdateSegmentResponse updateSegmentResponse = (UpdateSegmentResponse) o;
    return (
      Objects.equals(this.segmentID, updateSegmentResponse.segmentID) && Objects.equals(this.updatedAt, updateSegmentResponse.updatedAt)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(segmentID, updatedAt);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class UpdateSegmentResponse {\n");
    sb.append("    segmentID: ").append(toIndentedString(segmentID)).append("\n");
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
