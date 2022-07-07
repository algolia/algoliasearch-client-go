package com.algolia.model.search;

import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonValue;

/** Selects a strategy to remove words from the query when it doesn't match any hits. */
public enum RemoveWordsIfNoResults {
  NONE("none"),

  LAST_WORDS("lastWords"),

  FIRST_WORDS("firstWords"),

  ALL_OPTIONAL("allOptional");

  private final String value;

  RemoveWordsIfNoResults(String value) {
    this.value = value;
  }

  @JsonValue
  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  @JsonCreator
  public static RemoveWordsIfNoResults fromValue(String value) {
    for (RemoveWordsIfNoResults b : RemoveWordsIfNoResults.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }
}
