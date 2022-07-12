// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.search;

import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonValue;

/** The operation to apply on the attribute. */
public enum BuiltInOperationType {
  INCREMENT("Increment"),

  DECREMENT("Decrement"),

  ADD("Add"),

  REMOVE("Remove"),

  ADD_UNIQUE("AddUnique"),

  INCREMENT_FROM("IncrementFrom"),

  INCREMENT_SET("IncrementSet");

  private final String value;

  BuiltInOperationType(String value) {
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
  public static BuiltInOperationType fromValue(String value) {
    for (BuiltInOperationType b : BuiltInOperationType.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }
}
