package com.algolia.model.search;

import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;

/** Type of edit. */
@JsonAdapter(EditType.Adapter.class)
public enum EditType {
  REMOVE("remove"),

  REPLACE("replace");

  private final String value;

  EditType(String value) {
    this.value = value;
  }

  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  public static EditType fromValue(String value) {
    for (EditType b : EditType.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }

  public static class Adapter extends TypeAdapter<EditType> {

    @Override
    public void write(final JsonWriter jsonWriter, final EditType enumeration) throws IOException {
      jsonWriter.value(enumeration.getValue());
    }

    @Override
    public EditType read(final JsonReader jsonReader) throws IOException {
      String value = jsonReader.nextString();
      return EditType.fromValue(value);
    }
  }
}
