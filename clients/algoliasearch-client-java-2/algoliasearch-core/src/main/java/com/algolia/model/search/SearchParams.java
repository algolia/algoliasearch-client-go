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

/** SearchParams */
@JsonDeserialize(using = SearchParams.SearchParamsDeserializer.class)
@JsonSerialize(using = SearchParams.SearchParamsSerializer.class)
public abstract class SearchParams implements CompoundType {

  public static SearchParams of(SearchParamsObject inside) {
    return new SearchParamsSearchParamsObject(inside);
  }

  public static SearchParams of(SearchParamsString inside) {
    return new SearchParamsSearchParamsString(inside);
  }

  public static class SearchParamsSerializer extends StdSerializer<SearchParams> {

    public SearchParamsSerializer(Class<SearchParams> t) {
      super(t);
    }

    public SearchParamsSerializer() {
      this(null);
    }

    @Override
    public void serialize(SearchParams value, JsonGenerator jgen, SerializerProvider provider) throws IOException, JsonProcessingException {
      jgen.writeObject(value.getInsideValue());
    }
  }

  public static class SearchParamsDeserializer extends StdDeserializer<SearchParams> {

    public SearchParamsDeserializer() {
      this(SearchParams.class);
    }

    public SearchParamsDeserializer(Class<?> vc) {
      super(vc);
    }

    @Override
    public SearchParams deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException, JsonProcessingException {
      JsonNode tree = jp.readValueAsTree();
      SearchParams deserialized = null;

      int match = 0;
      JsonToken token = tree.traverse(jp.getCodec()).nextToken();
      String currentType = "";
      // deserialize SearchParamsObject
      try {
        boolean attemptParsing = true;
        currentType = "SearchParamsObject";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized =
            SearchParams.of((SearchParamsObject) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<SearchParamsObject>() {}));
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              SearchParams.of((SearchParamsObject) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<SearchParamsObject>() {}));
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println("Failed to deserialize oneOf SearchParamsObject (error: " + e.getMessage() + ") (type: " + currentType + ")");
      }

      // deserialize SearchParamsString
      try {
        boolean attemptParsing = true;
        currentType = "SearchParamsString";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized =
            SearchParams.of((SearchParamsString) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<SearchParamsString>() {}));
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              SearchParams.of((SearchParamsString) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<SearchParamsString>() {}));
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println("Failed to deserialize oneOf SearchParamsString (error: " + e.getMessage() + ") (type: " + currentType + ")");
      }

      if (match == 1) {
        return deserialized;
      }
      throw new IOException(String.format("Failed deserialization for SearchParams: %d classes match result, expected 1", match));
    }

    /** Handle deserialization of the 'null' value. */
    @Override
    public SearchParams getNullValue(DeserializationContext ctxt) throws JsonMappingException {
      throw new JsonMappingException(ctxt.getParser(), "SearchParams cannot be null");
    }
  }
}

class SearchParamsSearchParamsObject extends SearchParams {

  private final SearchParamsObject insideValue;

  SearchParamsSearchParamsObject(SearchParamsObject insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public SearchParamsObject getInsideValue() {
    return insideValue;
  }
}

class SearchParamsSearchParamsString extends SearchParams {

  private final SearchParamsString insideValue;

  SearchParamsSearchParamsString(SearchParamsString insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public SearchParamsString getInsideValue() {
    return insideValue;
  }
}
