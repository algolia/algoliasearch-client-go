package com.algolia.model.recommend;

import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonValue;

/** Controls if and how query words are interpreted as prefixes. */
public enum QueryType {
  PREFIX_LAST("prefixLast"),

  PREFIX_ALL("prefixAll"),

  PREFIX_NONE("prefixNone");

  private final String value;

  QueryType(String value) {
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
  public static QueryType fromValue(String value) {
    for (QueryType b : QueryType.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }
}
