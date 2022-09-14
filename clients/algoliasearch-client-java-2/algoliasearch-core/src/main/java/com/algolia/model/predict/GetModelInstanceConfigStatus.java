// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.predict;

import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonValue;

/**
 * `pending` - model has just been created and the pipelines are being set up for the first train &
 * inference. \\ `active` - model is running and generating prediction. \\ `invalid` - model has
 * failed training (ex. can’t retrieve data from source). An additional error field will be set for
 * this status. \\ `inactive` - model has been deactivated from the dashboard. Pipelines still exist
 * but they are not currently running.
 */
public enum GetModelInstanceConfigStatus {
  PENDING("pending"),

  ACTIVE("active"),

  INVALID("invalid"),

  ERROR("error"),

  INACTIVE("inactive");

  private final String value;

  GetModelInstanceConfigStatus(String value) {
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
  public static GetModelInstanceConfigStatus fromValue(String value) {
    for (GetModelInstanceConfigStatus b : GetModelInstanceConfigStatus.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }
}
