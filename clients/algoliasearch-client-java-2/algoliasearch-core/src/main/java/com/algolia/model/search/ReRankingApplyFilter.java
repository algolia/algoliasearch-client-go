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
 * When Dynamic Re-Ranking is enabled, only records that match these filters will be impacted by
 * Dynamic Re-Ranking.
 */
@JsonDeserialize(using = ReRankingApplyFilter.ReRankingApplyFilterDeserializer.class)
@JsonSerialize(using = ReRankingApplyFilter.ReRankingApplyFilterSerializer.class)
public abstract class ReRankingApplyFilter implements CompoundType {

  public static ReRankingApplyFilter of(List<MixedSearchFilters> inside) {
    return new ReRankingApplyFilterListOfMixedSearchFilters(inside);
  }

  public static ReRankingApplyFilter of(String inside) {
    return new ReRankingApplyFilterString(inside);
  }

  public static class ReRankingApplyFilterSerializer extends StdSerializer<ReRankingApplyFilter> {

    public ReRankingApplyFilterSerializer(Class<ReRankingApplyFilter> t) {
      super(t);
    }

    public ReRankingApplyFilterSerializer() {
      this(null);
    }

    @Override
    public void serialize(ReRankingApplyFilter value, JsonGenerator jgen, SerializerProvider provider)
      throws IOException, JsonProcessingException {
      jgen.writeObject(value.getInsideValue());
    }
  }

  public static class ReRankingApplyFilterDeserializer extends StdDeserializer<ReRankingApplyFilter> {

    public ReRankingApplyFilterDeserializer() {
      this(ReRankingApplyFilter.class);
    }

    public ReRankingApplyFilterDeserializer(Class<?> vc) {
      super(vc);
    }

    @Override
    public ReRankingApplyFilter deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException, JsonProcessingException {
      JsonNode tree = jp.readValueAsTree();
      ReRankingApplyFilter deserialized = null;

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
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY) |
          (token == JsonToken.VALUE_NULL)
        ) {
          deserialized =
            ReRankingApplyFilter.of(
              (List<MixedSearchFilters>) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<List<MixedSearchFilters>>() {})
            );
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              ReRankingApplyFilter.of(
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
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY) |
          (token == JsonToken.VALUE_NULL)
        ) {
          deserialized = ReRankingApplyFilter.of((String) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<String>() {}));
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized = ReRankingApplyFilter.of((String) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<String>() {}));
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
      throw new IOException(
        String.format("Failed deserialization for ReRankingApplyFilter: %d classes match result, expected" + " 1", match)
      );
    }

    /** Handle deserialization of the 'null' value. */
    @Override
    public ReRankingApplyFilter getNullValue(DeserializationContext ctxt) throws JsonMappingException {
      return null;
    }
  }
}

class ReRankingApplyFilterListOfMixedSearchFilters extends ReRankingApplyFilter {

  private final List<MixedSearchFilters> insideValue;

  ReRankingApplyFilterListOfMixedSearchFilters(List<MixedSearchFilters> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<MixedSearchFilters> getInsideValue() {
    return insideValue;
  }
}

class ReRankingApplyFilterString extends ReRankingApplyFilter {

  private final String insideValue;

  ReRankingApplyFilterString(String insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public String getInsideValue() {
    return insideValue;
  }
}
