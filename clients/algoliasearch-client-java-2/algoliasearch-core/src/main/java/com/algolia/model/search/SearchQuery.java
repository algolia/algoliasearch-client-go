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

/** SearchQuery */
@JsonDeserialize(using = SearchQuery.SearchQueryDeserializer.class)
@JsonSerialize(using = SearchQuery.SearchQuerySerializer.class)
public abstract class SearchQuery implements CompoundType {

  public static SearchQuery of(SearchForFacets inside) {
    return new SearchQuerySearchForFacets(inside);
  }

  public static SearchQuery of(SearchForHits inside) {
    return new SearchQuerySearchForHits(inside);
  }

  public static class SearchQuerySerializer extends StdSerializer<SearchQuery> {

    public SearchQuerySerializer(Class<SearchQuery> t) {
      super(t);
    }

    public SearchQuerySerializer() {
      this(null);
    }

    @Override
    public void serialize(SearchQuery value, JsonGenerator jgen, SerializerProvider provider) throws IOException, JsonProcessingException {
      jgen.writeObject(value.getInsideValue());
    }
  }

  public static class SearchQueryDeserializer extends StdDeserializer<SearchQuery> {

    public SearchQueryDeserializer() {
      this(SearchQuery.class);
    }

    public SearchQueryDeserializer(Class<?> vc) {
      super(vc);
    }

    @Override
    public SearchQuery deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException, JsonProcessingException {
      JsonNode tree = jp.readValueAsTree();
      SearchQuery deserialized = null;

      int match = 0;
      JsonToken token = tree.traverse(jp.getCodec()).nextToken();
      String currentType = "";
      // deserialize SearchForFacets
      try {
        boolean attemptParsing = true;
        currentType = "SearchForFacets";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized =
            SearchQuery.of((SearchForFacets) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<SearchForFacets>() {}));
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              SearchQuery.of((SearchForFacets) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<SearchForFacets>() {}));
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println("Failed to deserialize oneOf SearchForFacets (error: " + e.getMessage() + ") (type: " + currentType + ")");
      }

      // deserialize SearchForHits
      try {
        boolean attemptParsing = true;
        currentType = "SearchForHits";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized = SearchQuery.of((SearchForHits) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<SearchForHits>() {}));
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized = SearchQuery.of((SearchForHits) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<SearchForHits>() {}));
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println("Failed to deserialize oneOf SearchForHits (error: " + e.getMessage() + ") (type: " + currentType + ")");
      }

      if (match == 1) {
        return deserialized;
      }
      throw new IOException(String.format("Failed deserialization for SearchQuery: %d classes match result, expected 1", match));
    }

    /** Handle deserialization of the 'null' value. */
    @Override
    public SearchQuery getNullValue(DeserializationContext ctxt) throws JsonMappingException {
      throw new JsonMappingException(ctxt.getParser(), "SearchQuery cannot be null");
    }
  }
}

class SearchQuerySearchForFacets extends SearchQuery {

  private final SearchForFacets insideValue;

  SearchQuerySearchForFacets(SearchForFacets insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public SearchForFacets getInsideValue() {
    return insideValue;
  }
}

class SearchQuerySearchForHits extends SearchQuery {

  private final SearchForHits insideValue;

  SearchQuerySearchForHits(SearchForHits insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public SearchForHits getInsideValue() {
    return insideValue;
  }
}
