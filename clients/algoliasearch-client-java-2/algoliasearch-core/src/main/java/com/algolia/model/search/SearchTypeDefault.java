package com.algolia.model.search;

import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;

/** Perform a search query with `default`, will search for facet values if `facet` is given. */
@JsonAdapter(SearchTypeDefault.Adapter.class)
public enum SearchTypeDefault {
  DEFAULT("default");

  private final String value;

  SearchTypeDefault(String value) {
    this.value = value;
  }

  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  public static SearchTypeDefault fromValue(String value) {
    for (SearchTypeDefault b : SearchTypeDefault.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }

  public static class Adapter extends TypeAdapter<SearchTypeDefault> {

    @Override
    public void write(final JsonWriter jsonWriter, final SearchTypeDefault enumeration) throws IOException {
      jsonWriter.value(enumeration.getValue());
    }

    @Override
    public SearchTypeDefault read(final JsonReader jsonReader) throws IOException {
      String value = jsonReader.nextString();
      return SearchTypeDefault.fromValue(value);
    }
  }
}
