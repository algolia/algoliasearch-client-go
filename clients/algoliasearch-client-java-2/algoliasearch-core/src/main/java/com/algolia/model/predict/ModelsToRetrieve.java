// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.predict;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/** ModelsToRetrieve */
public class ModelsToRetrieve {

  @JsonProperty("modelsToRetrieve")
  private List<ModelsToRetrieveEnum> modelsToRetrieve = new ArrayList<>();

  public ModelsToRetrieve setModelsToRetrieve(List<ModelsToRetrieveEnum> modelsToRetrieve) {
    this.modelsToRetrieve = modelsToRetrieve;
    return this;
  }

  public ModelsToRetrieve addModelsToRetrieve(ModelsToRetrieveEnum modelsToRetrieveItem) {
    this.modelsToRetrieve.add(modelsToRetrieveItem);
    return this;
  }

  /**
   * Get modelsToRetrieve
   *
   * @return modelsToRetrieve
   */
  @javax.annotation.Nonnull
  public List<ModelsToRetrieveEnum> getModelsToRetrieve() {
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
    ModelsToRetrieve modelsToRetrieve = (ModelsToRetrieve) o;
    return Objects.equals(this.modelsToRetrieve, modelsToRetrieve.modelsToRetrieve);
  }

  @Override
  public int hashCode() {
    return Objects.hash(modelsToRetrieve);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ModelsToRetrieve {\n");
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
