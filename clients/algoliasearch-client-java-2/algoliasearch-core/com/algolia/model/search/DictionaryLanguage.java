package com.algolia.model.search;

import com.google.gson.annotations.SerializedName;
import java.util.Objects;

/** Custom entries for a dictionary. */
public class DictionaryLanguage {

  @SerializedName("nbCustomEntires")
  private Integer nbCustomEntires;

  public DictionaryLanguage setNbCustomEntires(Integer nbCustomEntires) {
    this.nbCustomEntires = nbCustomEntires;
    return this;
  }

  /**
   * When nbCustomEntries is set to 0, the user didn't customize the dictionary. The dictionary is
   * still supported with standard, Algolia-provided entries.
   *
   * @return nbCustomEntires
   */
  @javax.annotation.Nullable
  public Integer getNbCustomEntires() {
    return nbCustomEntires;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    DictionaryLanguage dictionaryLanguage = (DictionaryLanguage) o;
    return Objects.equals(
      this.nbCustomEntires,
      dictionaryLanguage.nbCustomEntires
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(nbCustomEntires);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class DictionaryLanguage {\n");
    sb
      .append("    nbCustomEntires: ")
      .append(toIndentedString(nbCustomEntires))
      .append("\n");
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
