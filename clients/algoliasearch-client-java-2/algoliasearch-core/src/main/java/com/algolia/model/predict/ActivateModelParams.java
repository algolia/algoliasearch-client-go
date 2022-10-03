// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.predict;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/** ActivateModelParams */
public class ActivateModelParams {

  @JsonProperty("type")
  private ModelsToRetrieve type;

  @JsonProperty("name")
  private String name;

  @JsonProperty("sourceID")
  private String sourceID;

  @JsonProperty("index")
  private String index;

  @JsonProperty("modelAttributes")
  private List<String> modelAttributes;

  @JsonProperty("contentAttributes")
  private List<String> contentAttributes;

  public ActivateModelParams setType(ModelsToRetrieve type) {
    this.type = type;
    return this;
  }

  /**
   * Get type
   *
   * @return type
   */
  @javax.annotation.Nonnull
  public ModelsToRetrieve getType() {
    return type;
  }

  public ActivateModelParams setName(String name) {
    this.name = name;
    return this;
  }

  /**
   * The model’s instance name.
   *
   * @return name
   */
  @javax.annotation.Nonnull
  public String getName() {
    return name;
  }

  public ActivateModelParams setSourceID(String sourceID) {
    this.sourceID = sourceID;
    return this;
  }

  /**
   * The data source ID, as returned by the (external) sources API.
   *
   * @return sourceID
   */
  @javax.annotation.Nonnull
  public String getSourceID() {
    return sourceID;
  }

  public ActivateModelParams setIndex(String index) {
    this.index = index;
    return this;
  }

  /**
   * The index name.
   *
   * @return index
   */
  @javax.annotation.Nonnull
  public String getIndex() {
    return index;
  }

  public ActivateModelParams setModelAttributes(List<String> modelAttributes) {
    this.modelAttributes = modelAttributes;
    return this;
  }

  public ActivateModelParams addModelAttributes(String modelAttributesItem) {
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

  public ActivateModelParams setContentAttributes(List<String> contentAttributes) {
    this.contentAttributes = contentAttributes;
    return this;
  }

  public ActivateModelParams addContentAttributes(String contentAttributesItem) {
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

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ActivateModelParams activateModelParams = (ActivateModelParams) o;
    return (
      Objects.equals(this.type, activateModelParams.type) &&
      Objects.equals(this.name, activateModelParams.name) &&
      Objects.equals(this.sourceID, activateModelParams.sourceID) &&
      Objects.equals(this.index, activateModelParams.index) &&
      Objects.equals(this.modelAttributes, activateModelParams.modelAttributes) &&
      Objects.equals(this.contentAttributes, activateModelParams.contentAttributes)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(type, name, sourceID, index, modelAttributes, contentAttributes);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ActivateModelParams {\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    sourceID: ").append(toIndentedString(sourceID)).append("\n");
    sb.append("    index: ").append(toIndentedString(index)).append("\n");
    sb.append("    modelAttributes: ").append(toIndentedString(modelAttributes)).append("\n");
    sb.append("    contentAttributes: ").append(toIndentedString(contentAttributes)).append("\n");
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
