// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.predict;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/** TypesToRetrieveParam */
public class TypesToRetrieveParam {

  @JsonProperty("typesToRetrieve")
  private List<TypesToRetrieve> typesToRetrieve;

  public TypesToRetrieveParam setTypesToRetrieve(List<TypesToRetrieve> typesToRetrieve) {
    this.typesToRetrieve = typesToRetrieve;
    return this;
  }

  public TypesToRetrieveParam addTypesToRetrieve(TypesToRetrieve typesToRetrieveItem) {
    if (this.typesToRetrieve == null) {
      this.typesToRetrieve = new ArrayList<>();
    }
    this.typesToRetrieve.add(typesToRetrieveItem);
    return this;
  }

  /**
   * Get typesToRetrieve
   *
   * @return typesToRetrieve
   */
  @javax.annotation.Nullable
  public List<TypesToRetrieve> getTypesToRetrieve() {
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
    TypesToRetrieveParam typesToRetrieveParam = (TypesToRetrieveParam) o;
    return Objects.equals(this.typesToRetrieve, typesToRetrieveParam.typesToRetrieve);
  }

  @Override
  public int hashCode() {
    return Objects.hash(typesToRetrieve);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class TypesToRetrieveParam {\n");
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
