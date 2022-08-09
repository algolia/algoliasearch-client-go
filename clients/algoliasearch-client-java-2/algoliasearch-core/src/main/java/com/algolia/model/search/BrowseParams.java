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

/** BrowseParams */
@JsonDeserialize(using = BrowseParams.BrowseParamsDeserializer.class)
@JsonSerialize(using = BrowseParams.BrowseParamsSerializer.class)
public abstract class BrowseParams implements CompoundType {

  public static BrowseParams of(BrowseParamsObject inside) {
    return new BrowseParamsBrowseParamsObject(inside);
  }

  public static BrowseParams of(SearchParamsString inside) {
    return new BrowseParamsSearchParamsString(inside);
  }

  public static class BrowseParamsSerializer extends StdSerializer<BrowseParams> {

    public BrowseParamsSerializer(Class<BrowseParams> t) {
      super(t);
    }

    public BrowseParamsSerializer() {
      this(null);
    }

    @Override
    public void serialize(BrowseParams value, JsonGenerator jgen, SerializerProvider provider) throws IOException, JsonProcessingException {
      jgen.writeObject(value.getInsideValue());
    }
  }

  public static class BrowseParamsDeserializer extends StdDeserializer<BrowseParams> {

    public BrowseParamsDeserializer() {
      this(BrowseParams.class);
    }

    public BrowseParamsDeserializer(Class<?> vc) {
      super(vc);
    }

    @Override
    public BrowseParams deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException, JsonProcessingException {
      JsonNode tree = jp.readValueAsTree();
      BrowseParams deserialized = null;

      int match = 0;
      JsonToken token = tree.traverse(jp.getCodec()).nextToken();
      String currentType = "";
      // deserialize BrowseParamsObject
      try {
        boolean attemptParsing = true;
        currentType = "BrowseParamsObject";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized =
            BrowseParams.of((BrowseParamsObject) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<BrowseParamsObject>() {}));
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              BrowseParams.of((BrowseParamsObject) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<BrowseParamsObject>() {}));
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println("Failed to deserialize oneOf BrowseParamsObject (error: " + e.getMessage() + ") (type: " + currentType + ")");
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
            BrowseParams.of((SearchParamsString) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<SearchParamsString>() {}));
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              BrowseParams.of((SearchParamsString) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<SearchParamsString>() {}));
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
      throw new IOException(String.format("Failed deserialization for BrowseParams: %d classes match result, expected 1", match));
    }

    /** Handle deserialization of the 'null' value. */
    @Override
    public BrowseParams getNullValue(DeserializationContext ctxt) throws JsonMappingException {
      throw new JsonMappingException(ctxt.getParser(), "BrowseParams cannot be null");
    }
  }
}

class BrowseParamsBrowseParamsObject extends BrowseParams {

  private final BrowseParamsObject insideValue;

  BrowseParamsBrowseParamsObject(BrowseParamsObject insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public BrowseParamsObject getInsideValue() {
    return insideValue;
  }
}

class BrowseParamsSearchParamsString extends BrowseParams {

  private final SearchParamsString insideValue;

  BrowseParamsSearchParamsString(SearchParamsString insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public SearchParamsString getInsideValue() {
    return insideValue;
  }
}
