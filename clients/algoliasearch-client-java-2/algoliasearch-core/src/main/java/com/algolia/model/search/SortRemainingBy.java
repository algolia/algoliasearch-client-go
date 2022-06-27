package com.algolia.model.search;

import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;

/**
 * How to display the remaining items. - `count`: facet count (descending). - `alpha`: alphabetical
 * (ascending). - `hidden`: show only pinned values.
 */
@JsonAdapter(SortRemainingBy.Adapter.class)
public enum SortRemainingBy {
  COUNT("count"),

  ALPHA("alpha"),

  HIDDEN("hidden");

  private final String value;

  SortRemainingBy(String value) {
    this.value = value;
  }

  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  public static SortRemainingBy fromValue(String value) {
    for (SortRemainingBy b : SortRemainingBy.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }

  public static class Adapter extends TypeAdapter<SortRemainingBy> {

    @Override
    public void write(final JsonWriter jsonWriter, final SortRemainingBy enumeration) throws IOException {
      jsonWriter.value(enumeration.getValue());
    }

    @Override
    public SortRemainingBy read(final JsonReader jsonReader) throws IOException {
      String value = jsonReader.nextString();
      return SortRemainingBy.fromValue(value);
    }
  }
}
