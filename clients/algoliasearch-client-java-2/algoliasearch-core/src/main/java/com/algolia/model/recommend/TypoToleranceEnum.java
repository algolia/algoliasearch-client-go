package com.algolia.model.recommend;

import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;

/** Gets or Sets typoToleranceEnum */
@JsonAdapter(TypoToleranceEnum.Adapter.class)
public enum TypoToleranceEnum {
  MIN("min"),

  STRICT("strict");

  private final String value;

  TypoToleranceEnum(String value) {
    this.value = value;
  }

  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  public static TypoToleranceEnum fromValue(String value) {
    for (TypoToleranceEnum b : TypoToleranceEnum.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }

  public static class Adapter extends TypeAdapter<TypoToleranceEnum> {

    @Override
    public void write(final JsonWriter jsonWriter, final TypoToleranceEnum enumeration) throws IOException {
      jsonWriter.value(enumeration.getValue());
    }

    @Override
    public TypoToleranceEnum read(final JsonReader jsonReader) throws IOException {
      String value = jsonReader.nextString();
      return TypoToleranceEnum.fromValue(value);
    }
  }
}
