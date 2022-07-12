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

/** MixedSearchFilters */
@JsonDeserialize(using = MixedSearchFilters.MixedSearchFiltersDeserializer.class)
@JsonSerialize(using = MixedSearchFilters.MixedSearchFiltersSerializer.class)
public abstract class MixedSearchFilters implements CompoundType {

  public static MixedSearchFilters of(List<String> inside) {
    return new MixedSearchFiltersListOfString(inside);
  }

  public static MixedSearchFilters of(String inside) {
    return new MixedSearchFiltersString(inside);
  }

  public static class MixedSearchFiltersSerializer extends StdSerializer<MixedSearchFilters> {

    public MixedSearchFiltersSerializer(Class<MixedSearchFilters> t) {
      super(t);
    }

    public MixedSearchFiltersSerializer() {
      this(null);
    }

    @Override
    public void serialize(MixedSearchFilters value, JsonGenerator jgen, SerializerProvider provider)
      throws IOException, JsonProcessingException {
      jgen.writeObject(value.getInsideValue());
    }
  }

  public static class MixedSearchFiltersDeserializer extends StdDeserializer<MixedSearchFilters> {

    public MixedSearchFiltersDeserializer() {
      this(MixedSearchFilters.class);
    }

    public MixedSearchFiltersDeserializer(Class<?> vc) {
      super(vc);
    }

    @Override
    public MixedSearchFilters deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException, JsonProcessingException {
      JsonNode tree = jp.readValueAsTree();
      MixedSearchFilters deserialized = null;

      int match = 0;
      JsonToken token = tree.traverse(jp.getCodec()).nextToken();
      String currentType = "";
      // deserialize List<String>
      try {
        boolean attemptParsing = true;
        currentType = "List<String>";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized =
            MixedSearchFilters.of((List<String>) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<List<String>>() {}));
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              MixedSearchFilters.of((List<String>) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<List<String>>() {}));
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println("Failed to deserialize oneOf List<String> (error: " + e.getMessage() + ") (type: " + currentType + ")");
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
          deserialized = MixedSearchFilters.of((String) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<String>() {}));
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized = MixedSearchFilters.of((String) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<String>() {}));
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
      throw new IOException(String.format("Failed deserialization for MixedSearchFilters: %d classes match result, expected 1", match));
    }

    /** Handle deserialization of the 'null' value. */
    @Override
    public MixedSearchFilters getNullValue(DeserializationContext ctxt) throws JsonMappingException {
      throw new JsonMappingException(ctxt.getParser(), "MixedSearchFilters cannot be null");
    }
  }
}

class MixedSearchFiltersListOfString extends MixedSearchFilters {

  private final List<String> insideValue;

  MixedSearchFiltersListOfString(List<String> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<String> getInsideValue() {
    return insideValue;
  }
}

class MixedSearchFiltersString extends MixedSearchFilters {

  private final String insideValue;

  MixedSearchFiltersString(String insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public String getInsideValue() {
    return insideValue;
  }
}
