package com.algolia.model.search;

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

/**
 * Create filters for ranking purposes, where records that match the filter are ranked higher, or
 * lower in the case of a negative optional filter.
 */
@JsonDeserialize(using = OptionalFilters.OptionalFiltersDeserializer.class)
@JsonSerialize(using = OptionalFilters.OptionalFiltersSerializer.class)
public abstract class OptionalFilters implements CompoundType {

  public static OptionalFilters of(List<MixedSearchFilters> inside) {
    return new OptionalFiltersListOfMixedSearchFilters(inside);
  }

  public static OptionalFilters of(String inside) {
    return new OptionalFiltersString(inside);
  }

  public static class OptionalFiltersSerializer extends StdSerializer<OptionalFilters> {

    public OptionalFiltersSerializer(Class<OptionalFilters> t) {
      super(t);
    }

    public OptionalFiltersSerializer() {
      this(null);
    }

    @Override
    public void serialize(OptionalFilters value, JsonGenerator jgen, SerializerProvider provider)
      throws IOException, JsonProcessingException {
      jgen.writeObject(value.getInsideValue());
    }
  }

  public static class OptionalFiltersDeserializer extends StdDeserializer<OptionalFilters> {

    public OptionalFiltersDeserializer() {
      this(OptionalFilters.class);
    }

    public OptionalFiltersDeserializer(Class<?> vc) {
      super(vc);
    }

    @Override
    public OptionalFilters deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException, JsonProcessingException {
      JsonNode tree = jp.readValueAsTree();
      OptionalFilters deserialized = null;

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
            OptionalFilters.of(
              (List<MixedSearchFilters>) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<List<MixedSearchFilters>>() {})
            );
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              OptionalFilters.of(
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
          deserialized = OptionalFilters.of((String) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<String>() {}));
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized = OptionalFilters.of((String) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<String>() {}));
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
      throw new IOException(String.format("Failed deserialization for OptionalFilters: %d classes match result, expected 1", match));
    }

    /** Handle deserialization of the 'null' value. */
    @Override
    public OptionalFilters getNullValue(DeserializationContext ctxt) throws JsonMappingException {
      throw new JsonMappingException(ctxt.getParser(), "OptionalFilters cannot be null");
    }
  }
}

class OptionalFiltersListOfMixedSearchFilters extends OptionalFilters {

  private final List<MixedSearchFilters> insideValue;

  OptionalFiltersListOfMixedSearchFilters(List<MixedSearchFilters> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<MixedSearchFilters> getInsideValue() {
    return insideValue;
  }
}

class OptionalFiltersString extends OptionalFilters {

  private final String insideValue;

  OptionalFiltersString(String insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public String getInsideValue() {
    return insideValue;
  }
}
