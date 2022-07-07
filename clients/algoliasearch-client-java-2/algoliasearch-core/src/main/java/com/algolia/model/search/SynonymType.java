package com.algolia.model.search;

import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonValue;

/** Type of the synonym object. */
public enum SynonymType {
  SYNONYM("synonym"),

  ONEWAYSYNONYM("onewaysynonym"),

  ALTCORRECTION_1("altcorrection1"),

  ALTCORRECTION_2("altcorrection2"),

  PLACEHOLDER("placeholder");

  private final String value;

  SynonymType(String value) {
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
  public static SynonymType fromValue(String value) {
    for (SynonymType b : SynonymType.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }
}
