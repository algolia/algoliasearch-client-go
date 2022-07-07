package com.algolia.model.recommend;

import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonValue;

/**
 * How to display the remaining items. - `count`: facet count (descending). - `alpha`: alphabetical
 * (ascending). - `hidden`: show only pinned values.
 */
public enum SortRemainingBy {
  COUNT("count"),

  ALPHA("alpha"),

  HIDDEN("hidden");

  private final String value;

  SortRemainingBy(String value) {
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
  public static SortRemainingBy fromValue(String value) {
    for (SortRemainingBy b : SortRemainingBy.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }
}
