package com.algolia.model.search;

import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonValue;

/** Controls how the exact ranking criterion is computed when the query contains only one word. */
public enum ExactOnSingleWordQuery {
  ATTRIBUTE("attribute"),

  NONE("none"),

  WORD("word");

  private final String value;

  ExactOnSingleWordQuery(String value) {
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
  public static ExactOnSingleWordQuery fromValue(String value) {
    for (ExactOnSingleWordQuery b : ExactOnSingleWordQuery.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }
}
