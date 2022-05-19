package com.algolia.model.search;

import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;

/** Gets or Sets searchStrategy */
@JsonAdapter(SearchStrategy.Adapter.class)
public enum SearchStrategy {
  NONE("none"),

  STOP_IF_ENOUGH_MATCHES("stopIfEnoughMatches");

  private final String value;

  SearchStrategy(String value) {
    this.value = value;
  }

  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  public static SearchStrategy fromValue(String value) {
    for (SearchStrategy b : SearchStrategy.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }

  public static class Adapter extends TypeAdapter<SearchStrategy> {

    @Override
    public void write(
      final JsonWriter jsonWriter,
      final SearchStrategy enumeration
    ) throws IOException {
      jsonWriter.value(enumeration.getValue());
    }

    @Override
    public SearchStrategy read(final JsonReader jsonReader) throws IOException {
      String value = jsonReader.nextString();
      return SearchStrategy.fromValue(value);
    }
  }
}
