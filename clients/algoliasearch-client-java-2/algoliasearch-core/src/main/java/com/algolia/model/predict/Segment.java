// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.predict;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** Segment */
public class Segment {

  @JsonProperty("segmentID")
  private String segmentID;

  @JsonProperty("name")
  private String name;

  @JsonProperty("conditions")
  private String conditions;

  @JsonProperty("segmentStatus")
  private SegmentStatus segmentStatus;

  @JsonProperty("type")
  private SegmentType type;

  @JsonProperty("errorMessage")
  private String errorMessage;

  public Segment setSegmentID(String segmentID) {
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

  public Segment setName(String name) {
    this.name = name;
    return this;
  }

  /**
   * Get name
   *
   * @return name
   */
  @javax.annotation.Nonnull
  public String getName() {
    return name;
  }

  public Segment setConditions(String conditions) {
    this.conditions = conditions;
    return this;
  }

  /**
   * Get conditions
   *
   * @return conditions
   */
  @javax.annotation.Nonnull
  public String getConditions() {
    return conditions;
  }

  public Segment setSegmentStatus(SegmentStatus segmentStatus) {
    this.segmentStatus = segmentStatus;
    return this;
  }

  /**
   * Get segmentStatus
   *
   * @return segmentStatus
   */
  @javax.annotation.Nonnull
  public SegmentStatus getSegmentStatus() {
    return segmentStatus;
  }

  public Segment setType(SegmentType type) {
    this.type = type;
    return this;
  }

  /**
   * Get type
   *
   * @return type
   */
  @javax.annotation.Nonnull
  public SegmentType getType() {
    return type;
  }

  public Segment setErrorMessage(String errorMessage) {
    this.errorMessage = errorMessage;
    return this;
  }

  /**
   * Get errorMessage
   *
   * @return errorMessage
   */
  @javax.annotation.Nullable
  public String getErrorMessage() {
    return errorMessage;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Segment segment = (Segment) o;
    return (
      Objects.equals(this.segmentID, segment.segmentID) &&
      Objects.equals(this.name, segment.name) &&
      Objects.equals(this.conditions, segment.conditions) &&
      Objects.equals(this.segmentStatus, segment.segmentStatus) &&
      Objects.equals(this.type, segment.type) &&
      Objects.equals(this.errorMessage, segment.errorMessage)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(segmentID, name, conditions, segmentStatus, type, errorMessage);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Segment {\n");
    sb.append("    segmentID: ").append(toIndentedString(segmentID)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    conditions: ").append(toIndentedString(conditions)).append("\n");
    sb.append("    segmentStatus: ").append(toIndentedString(segmentStatus)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    errorMessage: ").append(toIndentedString(errorMessage)).append("\n");
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
