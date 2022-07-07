package com.algolia.model.search;

import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonValue;

/** Gets or Sets acl */
public enum Acl {
  ADD_OBJECT("addObject"),

  ANALYTICS("analytics"),

  BROWSE("browse"),

  DELETE_OBJECT("deleteObject"),

  DELETE_INDEX("deleteIndex"),

  EDIT_SETTINGS("editSettings"),

  LIST_INDEXES("listIndexes"),

  LOGS("logs"),

  PERSONALIZATION("personalization"),

  RECOMMENDATION("recommendation"),

  SEARCH("search"),

  SEE_UNRETRIEVABLE_ATTRIBUTES("seeUnretrievableAttributes"),

  SETTINGS("settings"),

  USAGE("usage");

  private final String value;

  Acl(String value) {
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
  public static Acl fromValue(String value) {
    for (Acl b : Acl.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }
}
