// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.predict;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/** UpdateModelParams */
public class UpdateModelParams {

  @JsonProperty("name")
  private String name;

  @JsonProperty("affinities")
  private List<String> affinities;

  @JsonProperty("contentAttributes")
  private List<String> contentAttributes;

  @JsonProperty("status")
  private Status status;

  public UpdateModelParams setName(String name) {
    this.name = name;
    return this;
  }

  /**
   * The model’s instance name.
   *
   * @return name
   */
  @javax.annotation.Nullable
  public String getName() {
    return name;
  }

  public UpdateModelParams setAffinities(List<String> affinities) {
    this.affinities = affinities;
    return this;
  }

  public UpdateModelParams addAffinities(String affinitiesItem) {
    if (this.affinities == null) {
      this.affinities = new ArrayList<>();
    }
    this.affinities.add(affinitiesItem);
    return this;
  }

  /**
   * Get affinities
   *
   * @return affinities
   */
  @javax.annotation.Nullable
  public List<String> getAffinities() {
    return affinities;
  }

  public UpdateModelParams setContentAttributes(List<String> contentAttributes) {
    this.contentAttributes = contentAttributes;
    return this;
  }

  public UpdateModelParams addContentAttributes(String contentAttributesItem) {
    if (this.contentAttributes == null) {
      this.contentAttributes = new ArrayList<>();
    }
    this.contentAttributes.add(contentAttributesItem);
    return this;
  }

  /**
   * Get contentAttributes
   *
   * @return contentAttributes
   */
  @javax.annotation.Nullable
  public List<String> getContentAttributes() {
    return contentAttributes;
  }

  public UpdateModelParams setStatus(Status status) {
    this.status = status;
    return this;
  }

  /**
   * Get status
   *
   * @return status
   */
  @javax.annotation.Nullable
  public Status getStatus() {
    return status;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    UpdateModelParams updateModelParams = (UpdateModelParams) o;
    return (
      Objects.equals(this.name, updateModelParams.name) &&
      Objects.equals(this.affinities, updateModelParams.affinities) &&
      Objects.equals(this.contentAttributes, updateModelParams.contentAttributes) &&
      Objects.equals(this.status, updateModelParams.status)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, affinities, contentAttributes, status);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class UpdateModelParams {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    affinities: ").append(toIndentedString(affinities)).append("\n");
    sb.append("    contentAttributes: ").append(toIndentedString(contentAttributes)).append("\n");
    sb.append("    status: ").append(toIndentedString(status)).append("\n");
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
