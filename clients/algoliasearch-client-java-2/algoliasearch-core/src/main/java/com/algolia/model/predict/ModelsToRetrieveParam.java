// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.predict;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/** ModelsToRetrieveParam */
public class ModelsToRetrieveParam {

  @JsonProperty("modelsToRetrieve")
  private List<ModelsToRetrieve> modelsToRetrieve;

  public ModelsToRetrieveParam setModelsToRetrieve(List<ModelsToRetrieve> modelsToRetrieve) {
    this.modelsToRetrieve = modelsToRetrieve;
    return this;
  }

  public ModelsToRetrieveParam addModelsToRetrieve(ModelsToRetrieve modelsToRetrieveItem) {
    if (this.modelsToRetrieve == null) {
      this.modelsToRetrieve = new ArrayList<>();
    }
    this.modelsToRetrieve.add(modelsToRetrieveItem);
    return this;
  }

  /**
   * Get modelsToRetrieve
   *
   * @return modelsToRetrieve
   */
  @javax.annotation.Nullable
  public List<ModelsToRetrieve> getModelsToRetrieve() {
    return modelsToRetrieve;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ModelsToRetrieveParam modelsToRetrieveParam = (ModelsToRetrieveParam) o;
    return Objects.equals(this.modelsToRetrieve, modelsToRetrieveParam.modelsToRetrieve);
  }

  @Override
  public int hashCode() {
    return Objects.hash(modelsToRetrieve);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ModelsToRetrieveParam {\n");
    sb.append("    modelsToRetrieve: ").append(toIndentedString(modelsToRetrieve)).append("\n");
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
