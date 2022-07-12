// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.predict;

import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonValue;

/** Gets or Sets modelsToRetrieveEnum */
public enum ModelsToRetrieveEnum {
  FUNNEL_STAGE("funnel_stage"),

  ORDER_VALUE("order_value"),

  AFFINITIES("affinities");

  private final String value;

  ModelsToRetrieveEnum(String value) {
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
  public static ModelsToRetrieveEnum fromValue(String value) {
    for (ModelsToRetrieveEnum b : ModelsToRetrieveEnum.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }
}
