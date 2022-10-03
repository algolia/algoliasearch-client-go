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

  @JsonProperty("modelAttributes")
  private List<String> modelAttributes;

  @JsonProperty("contentAttributes")
  private List<String> contentAttributes;

  @JsonProperty("modelStatus")
  private ModelStatus modelStatus;

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

  public UpdateModelParams setModelAttributes(List<String> modelAttributes) {
    this.modelAttributes = modelAttributes;
    return this;
  }

  public UpdateModelParams addModelAttributes(String modelAttributesItem) {
    if (this.modelAttributes == null) {
      this.modelAttributes = new ArrayList<>();
    }
    this.modelAttributes.add(modelAttributesItem);
    return this;
  }

  /**
   * Get modelAttributes
   *
   * @return modelAttributes
   */
  @javax.annotation.Nullable
  public List<String> getModelAttributes() {
    return modelAttributes;
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

  public UpdateModelParams setModelStatus(ModelStatus modelStatus) {
    this.modelStatus = modelStatus;
    return this;
  }

  /**
   * Get modelStatus
   *
   * @return modelStatus
   */
  @javax.annotation.Nullable
  public ModelStatus getModelStatus() {
    return modelStatus;
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
      Objects.equals(this.modelAttributes, updateModelParams.modelAttributes) &&
      Objects.equals(this.contentAttributes, updateModelParams.contentAttributes) &&
      Objects.equals(this.modelStatus, updateModelParams.modelStatus)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, modelAttributes, contentAttributes, modelStatus);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class UpdateModelParams {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    modelAttributes: ").append(toIndentedString(modelAttributes)).append("\n");
    sb.append("    contentAttributes: ").append(toIndentedString(contentAttributes)).append("\n");
    sb.append("    modelStatus: ").append(toIndentedString(modelStatus)).append("\n");
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
