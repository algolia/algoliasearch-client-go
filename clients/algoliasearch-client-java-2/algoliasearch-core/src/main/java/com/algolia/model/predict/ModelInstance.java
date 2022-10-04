// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.predict;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/** ModelInstance */
public class ModelInstance {

  @JsonProperty("modelID")
  private String modelID;

  @JsonProperty("name")
  private String name;

  @JsonProperty("type")
  private String type;

  @JsonProperty("sourceID")
  private String sourceID;

  @JsonProperty("index")
  private String index;

  @JsonProperty("modelAttributes")
  private List<ModelAttributes> modelAttributes;

  @JsonProperty("contentAttributes")
  private List<String> contentAttributes;

  @JsonProperty("lastTrained")
  private String lastTrained;

  @JsonProperty("lastInference")
  private String lastInference;

  @JsonProperty("errorMessage")
  private String errorMessage;

  @JsonProperty("modelStatus")
  private GetModelInstanceConfigStatus modelStatus;

  public ModelInstance setModelID(String modelID) {
    this.modelID = modelID;
    return this;
  }

  /**
   * ID of the model.
   *
   * @return modelID
   */
  @javax.annotation.Nonnull
  public String getModelID() {
    return modelID;
  }

  public ModelInstance setName(String name) {
    this.name = name;
    return this;
  }

  /**
   * Name of model instance.
   *
   * @return name
   */
  @javax.annotation.Nonnull
  public String getName() {
    return name;
  }

  public ModelInstance setType(String type) {
    this.type = type;
    return this;
  }

  /**
   * Type of the model.
   *
   * @return type
   */
  @javax.annotation.Nonnull
  public String getType() {
    return type;
  }

  public ModelInstance setSourceID(String sourceID) {
    this.sourceID = sourceID;
    return this;
  }

  /**
   * Get sourceID
   *
   * @return sourceID
   */
  @javax.annotation.Nonnull
  public String getSourceID() {
    return sourceID;
  }

  public ModelInstance setIndex(String index) {
    this.index = index;
    return this;
  }

  /**
   * Get index
   *
   * @return index
   */
  @javax.annotation.Nonnull
  public String getIndex() {
    return index;
  }

  public ModelInstance setModelAttributes(List<ModelAttributes> modelAttributes) {
    this.modelAttributes = modelAttributes;
    return this;
  }

  public ModelInstance addModelAttributes(ModelAttributes modelAttributesItem) {
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
  public List<ModelAttributes> getModelAttributes() {
    return modelAttributes;
  }

  public ModelInstance setContentAttributes(List<String> contentAttributes) {
    this.contentAttributes = contentAttributes;
    return this;
  }

  public ModelInstance addContentAttributes(String contentAttributesItem) {
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

  public ModelInstance setLastTrained(String lastTrained) {
    this.lastTrained = lastTrained;
    return this;
  }

  /**
   * The date and time this model instance was last trained.
   *
   * @return lastTrained
   */
  @javax.annotation.Nonnull
  public String getLastTrained() {
    return lastTrained;
  }

  public ModelInstance setLastInference(String lastInference) {
    this.lastInference = lastInference;
    return this;
  }

  /**
   * The date and time this model instance generated its last inference.
   *
   * @return lastInference
   */
  @javax.annotation.Nonnull
  public String getLastInference() {
    return lastInference;
  }

  public ModelInstance setErrorMessage(String errorMessage) {
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

  public ModelInstance setModelStatus(GetModelInstanceConfigStatus modelStatus) {
    this.modelStatus = modelStatus;
    return this;
  }

  /**
   * Get modelStatus
   *
   * @return modelStatus
   */
  @javax.annotation.Nonnull
  public GetModelInstanceConfigStatus getModelStatus() {
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
    ModelInstance modelInstance = (ModelInstance) o;
    return (
      Objects.equals(this.modelID, modelInstance.modelID) &&
      Objects.equals(this.name, modelInstance.name) &&
      Objects.equals(this.type, modelInstance.type) &&
      Objects.equals(this.sourceID, modelInstance.sourceID) &&
      Objects.equals(this.index, modelInstance.index) &&
      Objects.equals(this.modelAttributes, modelInstance.modelAttributes) &&
      Objects.equals(this.contentAttributes, modelInstance.contentAttributes) &&
      Objects.equals(this.lastTrained, modelInstance.lastTrained) &&
      Objects.equals(this.lastInference, modelInstance.lastInference) &&
      Objects.equals(this.errorMessage, modelInstance.errorMessage) &&
      Objects.equals(this.modelStatus, modelInstance.modelStatus)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(
      modelID,
      name,
      type,
      sourceID,
      index,
      modelAttributes,
      contentAttributes,
      lastTrained,
      lastInference,
      errorMessage,
      modelStatus
    );
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ModelInstance {\n");
    sb.append("    modelID: ").append(toIndentedString(modelID)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    sourceID: ").append(toIndentedString(sourceID)).append("\n");
    sb.append("    index: ").append(toIndentedString(index)).append("\n");
    sb.append("    modelAttributes: ").append(toIndentedString(modelAttributes)).append("\n");
    sb.append("    contentAttributes: ").append(toIndentedString(contentAttributes)).append("\n");
    sb.append("    lastTrained: ").append(toIndentedString(lastTrained)).append("\n");
    sb.append("    lastInference: ").append(toIndentedString(lastInference)).append("\n");
    sb.append("    errorMessage: ").append(toIndentedString(errorMessage)).append("\n");
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
