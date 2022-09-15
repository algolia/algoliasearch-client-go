// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.predict;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** DeleteSegmentResponse */
public class DeleteSegmentResponse {

  @JsonProperty("segmentID")
  private String segmentID;

  @JsonProperty("deletedUntil")
  private String deletedUntil;

  public DeleteSegmentResponse setSegmentID(String segmentID) {
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

  public DeleteSegmentResponse setDeletedUntil(String deletedUntil) {
    this.deletedUntil = deletedUntil;
    return this;
  }

  /**
   * The date and time at which the segment will be re-ingested.
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
    DeleteSegmentResponse deleteSegmentResponse = (DeleteSegmentResponse) o;
    return (
      Objects.equals(this.segmentID, deleteSegmentResponse.segmentID) &&
      Objects.equals(this.deletedUntil, deleteSegmentResponse.deletedUntil)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(segmentID, deletedUntil);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class DeleteSegmentResponse {\n");
    sb.append("    segmentID: ").append(toIndentedString(segmentID)).append("\n");
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
