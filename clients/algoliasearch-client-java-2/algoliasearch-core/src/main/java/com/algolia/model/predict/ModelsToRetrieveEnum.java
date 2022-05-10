package com.algolia.model.predict;

import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;

/** Gets or Sets modelsToRetrieveEnum */
@JsonAdapter(ModelsToRetrieveEnum.Adapter.class)
public enum ModelsToRetrieveEnum {
  FUNNEL_STAGE("funnel_stage"),

  ORDER_VALUE("order_value"),

  AFFINITIES("affinities");

  private final String value;

  ModelsToRetrieveEnum(String value) {
    this.value = value;
  }

  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  public static ModelsToRetrieveEnum fromValue(String value) {
    for (ModelsToRetrieveEnum b : ModelsToRetrieveEnum.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }

  public static class Adapter extends TypeAdapter<ModelsToRetrieveEnum> {

    @Override
    public void write(
      final JsonWriter jsonWriter,
      final ModelsToRetrieveEnum enumeration
    ) throws IOException {
      jsonWriter.value(enumeration.getValue());
    }

    @Override
    public ModelsToRetrieveEnum read(final JsonReader jsonReader)
      throws IOException {
      String value = jsonReader.nextString();
      return ModelsToRetrieveEnum.fromValue(value);
    }
  }
}
