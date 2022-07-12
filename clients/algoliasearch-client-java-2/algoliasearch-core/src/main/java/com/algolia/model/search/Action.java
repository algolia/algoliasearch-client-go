// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.search;

import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonValue;

/** type of operation. */
public enum Action {
  ADD_OBJECT("addObject"),

  UPDATE_OBJECT("updateObject"),

  PARTIAL_UPDATE_OBJECT("partialUpdateObject"),

  PARTIAL_UPDATE_OBJECT_NO_CREATE("partialUpdateObjectNoCreate"),

  DELETE_OBJECT("deleteObject"),

  DELETE("delete"),

  CLEAR("clear");

  private final String value;

  Action(String value) {
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
  public static Action fromValue(String value) {
    for (Action b : Action.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }
}
