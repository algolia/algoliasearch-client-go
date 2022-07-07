package com.algolia.model.predict;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/** AllParams */
public class AllParams {

  @JsonProperty("modelsToRetrieve")
  private List<ModelsToRetrieveEnum> modelsToRetrieve = new ArrayList<>();

  @JsonProperty("typesToRetrieve")
  private List<TypesToRetrieveEnum> typesToRetrieve = new ArrayList<>();

  public AllParams setModelsToRetrieve(List<ModelsToRetrieveEnum> modelsToRetrieve) {
    this.modelsToRetrieve = modelsToRetrieve;
    return this;
  }

  public AllParams addModelsToRetrieve(ModelsToRetrieveEnum modelsToRetrieveItem) {
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

  public AllParams setTypesToRetrieve(List<TypesToRetrieveEnum> typesToRetrieve) {
    this.typesToRetrieve = typesToRetrieve;
    return this;
  }

  public AllParams addTypesToRetrieve(TypesToRetrieveEnum typesToRetrieveItem) {
    this.typesToRetrieve.add(typesToRetrieveItem);
    return this;
  }

  /**
   * Get typesToRetrieve
   *
   * @return typesToRetrieve
   */
  @javax.annotation.Nonnull
  public List<TypesToRetrieveEnum> getTypesToRetrieve() {
    return typesToRetrieve;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    AllParams allParams = (AllParams) o;
    return (
      Objects.equals(this.modelsToRetrieve, allParams.modelsToRetrieve) && Objects.equals(this.typesToRetrieve, allParams.typesToRetrieve)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(modelsToRetrieve, typesToRetrieve);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class AllParams {\n");
    sb.append("    modelsToRetrieve: ").append(toIndentedString(modelsToRetrieve)).append("\n");
    sb.append("    typesToRetrieve: ").append(toIndentedString(typesToRetrieve)).append("\n");
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
