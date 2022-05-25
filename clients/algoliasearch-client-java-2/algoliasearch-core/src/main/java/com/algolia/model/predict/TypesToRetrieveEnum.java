package com.algolia.model.predict;

import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;

/** Gets or Sets typesToRetrieveEnum */
@JsonAdapter(TypesToRetrieveEnum.Adapter.class)
public enum TypesToRetrieveEnum {
  PROPERTIES("properties"),

  SEGMENTS("segments");

  private final String value;

  TypesToRetrieveEnum(String value) {
    this.value = value;
  }

  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  public static TypesToRetrieveEnum fromValue(String value) {
    for (TypesToRetrieveEnum b : TypesToRetrieveEnum.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }

  public static class Adapter extends TypeAdapter<TypesToRetrieveEnum> {

    @Override
    public void write(final JsonWriter jsonWriter, final TypesToRetrieveEnum enumeration) throws IOException {
      jsonWriter.value(enumeration.getValue());
    }

    @Override
    public TypesToRetrieveEnum read(final JsonReader jsonReader) throws IOException {
      String value = jsonReader.nextString();
      return TypesToRetrieveEnum.fromValue(value);
    }
  }
}
