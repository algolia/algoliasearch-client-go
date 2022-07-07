package com.algolia.model.querysuggestions;

import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonValue;

/** type of the record, can be one of three values (INFO, SKIP or ERROR). */
public enum LogLevel {
  I_NF_O("INFO"),

  S_KI_P("SKIP"),

  E_RR_OR("ERROR");

  private final String value;

  LogLevel(String value) {
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
  public static LogLevel fromValue(String value) {
    for (LogLevel b : LogLevel.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }
}
