// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

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
 * Names of facets to which automatic filtering must be applied; they must match the facet name of a
 * facet value placeholder in the query pattern.
 */
@JsonDeserialize(using = AutomaticFacetFilters.AutomaticFacetFiltersDeserializer.class)
@JsonSerialize(using = AutomaticFacetFilters.AutomaticFacetFiltersSerializer.class)
public abstract class AutomaticFacetFilters implements CompoundType {

  public static AutomaticFacetFilters ofListOfAutomaticFacetFilter(List<AutomaticFacetFilter> inside) {
    return new AutomaticFacetFiltersListOfAutomaticFacetFilter(inside);
  }

  public static AutomaticFacetFilters ofListOfString(List<String> inside) {
    return new AutomaticFacetFiltersListOfString(inside);
  }

  public static class AutomaticFacetFiltersSerializer extends StdSerializer<AutomaticFacetFilters> {

    public AutomaticFacetFiltersSerializer(Class<AutomaticFacetFilters> t) {
      super(t);
    }

    public AutomaticFacetFiltersSerializer() {
      this(null);
    }

    @Override
    public void serialize(AutomaticFacetFilters value, JsonGenerator jgen, SerializerProvider provider)
      throws IOException, JsonProcessingException {
      jgen.writeObject(value.getInsideValue());
    }
  }

  public static class AutomaticFacetFiltersDeserializer extends StdDeserializer<AutomaticFacetFilters> {

    public AutomaticFacetFiltersDeserializer() {
      this(AutomaticFacetFilters.class);
    }

    public AutomaticFacetFiltersDeserializer(Class<?> vc) {
      super(vc);
    }

    @Override
    public AutomaticFacetFilters deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException, JsonProcessingException {
      JsonNode tree = jp.readValueAsTree();
      AutomaticFacetFilters deserialized = null;

      int match = 0;
      JsonToken token = tree.traverse(jp.getCodec()).nextToken();
      String currentType = "";
      // deserialize List<AutomaticFacetFilter>
      try {
        boolean attemptParsing = true;
        currentType = "List<AutomaticFacetFilter>";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized =
            AutomaticFacetFilters.ofListOfAutomaticFacetFilter(
              (List<AutomaticFacetFilter>) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<List<AutomaticFacetFilter>>() {})
            );
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              AutomaticFacetFilters.ofListOfAutomaticFacetFilter(
                (List<AutomaticFacetFilter>) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<List<AutomaticFacetFilter>>() {})
              );
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println(
          "Failed to deserialize oneOf List<AutomaticFacetFilter> (error: " + e.getMessage() + ") (type: " + currentType + ")"
        );
      }

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
            AutomaticFacetFilters.ofListOfString(
              (List<String>) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<List<String>>() {})
            );
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              AutomaticFacetFilters.ofListOfString(
                (List<String>) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<List<String>>() {})
              );
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println("Failed to deserialize oneOf List<String> (error: " + e.getMessage() + ") (type: " + currentType + ")");
      }

      if (match == 1) {
        return deserialized;
      }
      throw new IOException(
        String.format("Failed deserialization for AutomaticFacetFilters: %d classes match result, expected" + " 1", match)
      );
    }

    /** Handle deserialization of the 'null' value. */
    @Override
    public AutomaticFacetFilters getNullValue(DeserializationContext ctxt) throws JsonMappingException {
      throw new JsonMappingException(ctxt.getParser(), "AutomaticFacetFilters cannot be null");
    }
  }
}

class AutomaticFacetFiltersListOfAutomaticFacetFilter extends AutomaticFacetFilters {

  private final List<AutomaticFacetFilter> insideValue;

  AutomaticFacetFiltersListOfAutomaticFacetFilter(List<AutomaticFacetFilter> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<AutomaticFacetFilter> getInsideValue() {
    return insideValue;
  }
}

class AutomaticFacetFiltersListOfString extends AutomaticFacetFilters {

  private final List<String> insideValue;

  AutomaticFacetFiltersListOfString(List<String> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<String> getInsideValue() {
    return insideValue;
  }
}
