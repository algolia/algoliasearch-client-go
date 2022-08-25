// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.search;

import com.fasterxml.jackson.annotation.*;
import java.util.Objects;

/** Disable the builtin Algolia entries for a type of dictionary per language. */
public class DictionarySettingsParams {

  @JsonProperty("disableStandardEntries")
  private StandardEntries disableStandardEntries;

  public DictionarySettingsParams setDisableStandardEntries(StandardEntries disableStandardEntries) {
    this.disableStandardEntries = disableStandardEntries;
    return this;
  }

  /**
   * Get disableStandardEntries
   *
   * @return disableStandardEntries
   */
  @javax.annotation.Nonnull
  public StandardEntries getDisableStandardEntries() {
    return disableStandardEntries;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    DictionarySettingsParams dictionarySettingsParams = (DictionarySettingsParams) o;
    return Objects.equals(this.disableStandardEntries, dictionarySettingsParams.disableStandardEntries);
  }

  @Override
  public int hashCode() {
    return Objects.hash(disableStandardEntries);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class DictionarySettingsParams {\n");
    sb.append("    disableStandardEntries: ").append(toIndentedString(disableStandardEntries)).append("\n");
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