// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.predict;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/** GetAvailableModelTypesResponseInner */
public class GetAvailableModelTypesResponseInner {

  @JsonProperty("name")
  private String name;

  @JsonProperty("type")
  private String type;

  @JsonProperty("compatibleSources")
  private List<CompatibleSources> compatibleSources = new ArrayList<>();

  @JsonProperty("dataRequirements")
  private GetAvailableModelTypesResponseInnerDataRequirements dataRequirements;

  public GetAvailableModelTypesResponseInner setName(String name) {
    this.name = name;
    return this;
  }

  /**
   * Name of the model.
   *
   * @return name
   */
  @javax.annotation.Nonnull
  public String getName() {
    return name;
  }

  public GetAvailableModelTypesResponseInner setType(String type) {
    this.type = type;
    return this;
  }

  /**
   * Description of the model.
   *
   * @return type
   */
  @javax.annotation.Nonnull
  public String getType() {
    return type;
  }

  public GetAvailableModelTypesResponseInner setCompatibleSources(List<CompatibleSources> compatibleSources) {
    this.compatibleSources = compatibleSources;
    return this;
  }

  public GetAvailableModelTypesResponseInner addCompatibleSources(CompatibleSources compatibleSourcesItem) {
    this.compatibleSources.add(compatibleSourcesItem);
    return this;
  }

  /**
   * Get compatibleSources
   *
   * @return compatibleSources
   */
  @javax.annotation.Nonnull
  public List<CompatibleSources> getCompatibleSources() {
    return compatibleSources;
  }

  public GetAvailableModelTypesResponseInner setDataRequirements(GetAvailableModelTypesResponseInnerDataRequirements dataRequirements) {
    this.dataRequirements = dataRequirements;
    return this;
  }

  /**
   * Get dataRequirements
   *
   * @return dataRequirements
   */
  @javax.annotation.Nonnull
  public GetAvailableModelTypesResponseInnerDataRequirements getDataRequirements() {
    return dataRequirements;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    GetAvailableModelTypesResponseInner getAvailableModelTypesResponseInner = (GetAvailableModelTypesResponseInner) o;
    return (
      Objects.equals(this.name, getAvailableModelTypesResponseInner.name) &&
      Objects.equals(this.type, getAvailableModelTypesResponseInner.type) &&
      Objects.equals(this.compatibleSources, getAvailableModelTypesResponseInner.compatibleSources) &&
      Objects.equals(this.dataRequirements, getAvailableModelTypesResponseInner.dataRequirements)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, type, compatibleSources, dataRequirements);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class GetAvailableModelTypesResponseInner {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    compatibleSources: ").append(toIndentedString(compatibleSources)).append("\n");
    sb.append("    dataRequirements: ").append(toIndentedString(dataRequirements)).append("\n");
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
