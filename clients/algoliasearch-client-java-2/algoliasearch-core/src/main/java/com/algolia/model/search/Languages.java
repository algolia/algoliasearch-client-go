// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.search;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** A dictionary language. */
public class Languages {

  @JsonProperty("plurals")
  private DictionaryLanguage plurals;

  @JsonProperty("stopwords")
  private DictionaryLanguage stopwords;

  @JsonProperty("compounds")
  private DictionaryLanguage compounds;

  public Languages setPlurals(DictionaryLanguage plurals) {
    this.plurals = plurals;
    return this;
  }

  /**
   * Get plurals
   *
   * @return plurals
   */
  @javax.annotation.Nullable
  public DictionaryLanguage getPlurals() {
    return plurals;
  }

  public Languages setStopwords(DictionaryLanguage stopwords) {
    this.stopwords = stopwords;
    return this;
  }

  /**
   * Get stopwords
   *
   * @return stopwords
   */
  @javax.annotation.Nullable
  public DictionaryLanguage getStopwords() {
    return stopwords;
  }

  public Languages setCompounds(DictionaryLanguage compounds) {
    this.compounds = compounds;
    return this;
  }

  /**
   * Get compounds
   *
   * @return compounds
   */
  @javax.annotation.Nullable
  public DictionaryLanguage getCompounds() {
    return compounds;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Languages languages = (Languages) o;
    return (
      Objects.equals(this.plurals, languages.plurals) &&
      Objects.equals(this.stopwords, languages.stopwords) &&
      Objects.equals(this.compounds, languages.compounds)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(plurals, stopwords, compounds);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Languages {\n");
    sb.append("    plurals: ").append(toIndentedString(plurals)).append("\n");
    sb.append("    stopwords: ").append(toIndentedString(stopwords)).append("\n");
    sb.append("    compounds: ").append(toIndentedString(compounds)).append("\n");
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
