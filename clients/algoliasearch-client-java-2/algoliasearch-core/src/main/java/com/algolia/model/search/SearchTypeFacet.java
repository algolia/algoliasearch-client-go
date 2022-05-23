package com.algolia.model.search;

import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;

/** Perform a search query with `default`, will search for facet values if `facet` is given. */
@JsonAdapter(SearchTypeFacet.Adapter.class)
public enum SearchTypeFacet {
  FACET("facet");

  private final String value;

  SearchTypeFacet(String value) {
    this.value = value;
  }

  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  public static SearchTypeFacet fromValue(String value) {
    for (SearchTypeFacet b : SearchTypeFacet.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }

  public static class Adapter extends TypeAdapter<SearchTypeFacet> {

    @Override
    public void write(
      final JsonWriter jsonWriter,
      final SearchTypeFacet enumeration
    ) throws IOException {
      jsonWriter.value(enumeration.getValue());
    }

    @Override
    public SearchTypeFacet read(final JsonReader jsonReader)
      throws IOException {
      String value = jsonReader.nextString();
      return SearchTypeFacet.fromValue(value);
    }
  }
}
