// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.predict;

import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonValue;

/**
 * `active` - model is running and generating predictions. The active value is allowed only if the
 * current status of the model is `inactive`. \\ `inactive` - model training and inference have been
 * paused. The inactive value is allowed only if the current status of the model is `active`.
 */
public enum ModelStatus {
  ACTIVE("active"),

  INACTIVE("inactive");

  private final String value;

  ModelStatus(String value) {
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
  public static ModelStatus fromValue(String value) {
    for (ModelStatus b : ModelStatus.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }
}
