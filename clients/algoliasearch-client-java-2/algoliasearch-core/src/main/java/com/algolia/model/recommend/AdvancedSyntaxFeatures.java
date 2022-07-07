package com.algolia.model.recommend;

import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonValue;

/** Gets or Sets advancedSyntaxFeatures */
public enum AdvancedSyntaxFeatures {
  EXACT_PHRASE("exactPhrase"),

  EXCLUDE_WORDS("excludeWords");

  private final String value;

  AdvancedSyntaxFeatures(String value) {
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
  public static AdvancedSyntaxFeatures fromValue(String value) {
    for (AdvancedSyntaxFeatures b : AdvancedSyntaxFeatures.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }
}
