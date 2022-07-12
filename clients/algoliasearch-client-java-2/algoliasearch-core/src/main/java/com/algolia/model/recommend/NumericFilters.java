// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.recommend;

import com.algolia.utils.CompoundType;
import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.core.*;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.*;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.fasterxml.jackson.databind.deser.std.StdDeserializer;
import com.fasterxml.jackson.databind.ser.std.StdSerializer;
import java.io.IOException;
import java.util.List;

/** Filter on numeric attributes. */
@JsonDeserialize(using = NumericFilters.NumericFiltersDeserializer.class)
@JsonSerialize(using = NumericFilters.NumericFiltersSerializer.class)
public abstract class NumericFilters implements CompoundType {

  public static NumericFilters of(List<MixedSearchFilters> inside) {
    return new NumericFiltersListOfMixedSearchFilters(inside);
  }

  public static NumericFilters of(String inside) {
    return new NumericFiltersString(inside);
  }

  public static class NumericFiltersSerializer extends StdSerializer<NumericFilters> {

    public NumericFiltersSerializer(Class<NumericFilters> t) {
      super(t);
    }

    public NumericFiltersSerializer() {
      this(null);
    }

    @Override
    public void serialize(NumericFilters value, JsonGenerator jgen, SerializerProvider provider)
      throws IOException, JsonProcessingException {
      jgen.writeObject(value.getInsideValue());
    }
  }

  public static class NumericFiltersDeserializer extends StdDeserializer<NumericFilters> {

    public NumericFiltersDeserializer() {
      this(NumericFilters.class);
    }

    public NumericFiltersDeserializer(Class<?> vc) {
      super(vc);
    }

    @Override
    public NumericFilters deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException, JsonProcessingException {
      JsonNode tree = jp.readValueAsTree();
      NumericFilters deserialized = null;

      int match = 0;
      JsonToken token = tree.traverse(jp.getCodec()).nextToken();
      String currentType = "";
      // deserialize List<MixedSearchFilters>
      try {
        boolean attemptParsing = true;
        currentType = "List<MixedSearchFilters>";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized =
            NumericFilters.of(
              (List<MixedSearchFilters>) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<List<MixedSearchFilters>>() {})
            );
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              NumericFilters.of(
                (List<MixedSearchFilters>) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<List<MixedSearchFilters>>() {})
              );
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println(
          "Failed to deserialize oneOf List<MixedSearchFilters> (error: " + e.getMessage() + ") (type: " + currentType + ")"
        );
      }

      // deserialize String
      try {
        boolean attemptParsing = true;
        currentType = "String";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized = NumericFilters.of((String) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<String>() {}));
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized = NumericFilters.of((String) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<String>() {}));
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println("Failed to deserialize oneOf String (error: " + e.getMessage() + ") (type: " + currentType + ")");
      }

      if (match == 1) {
        return deserialized;
      }
      throw new IOException(String.format("Failed deserialization for NumericFilters: %d classes match result, expected 1", match));
    }

    /** Handle deserialization of the 'null' value. */
    @Override
    public NumericFilters getNullValue(DeserializationContext ctxt) throws JsonMappingException {
      throw new JsonMappingException(ctxt.getParser(), "NumericFilters cannot be null");
    }
  }
}

class NumericFiltersListOfMixedSearchFilters extends NumericFilters {

  private final List<MixedSearchFilters> insideValue;

  NumericFiltersListOfMixedSearchFilters(List<MixedSearchFilters> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<MixedSearchFilters> getInsideValue() {
    return insideValue;
  }
}

class NumericFiltersString extends NumericFilters {

  private final String insideValue;

  NumericFiltersString(String insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public String getInsideValue() {
    return insideValue;
  }
}
