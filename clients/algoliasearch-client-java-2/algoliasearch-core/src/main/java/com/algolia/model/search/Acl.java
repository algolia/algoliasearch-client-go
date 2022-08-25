// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.search;

import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonValue;

/**
 * List of rights for the API key. The following rights can be used: addObject: allows to add/update
 * an object in the index (copy/move index are also allowed with this right). analytics: allows to
 * retrieve the analytics through the Analytics API. browse: allows to retrieve all index content
 * via the browse API. deleteIndex: allows to delete or clear index content. deleteObject: allows to
 * delete objects from the index. editSettings: allows to change index settings. listIndexes: allows
 * to list all accessible indices. logs: allows to get the logs. recommendation: Allows usage of the
 * Personalization dashboard and the Recommendation API. search: allows to search the index.
 * seeUnretrievableAttributes: disable unretrievableAttributes feature for all operations returning
 * records. settings: allows to get index settings.
 */
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