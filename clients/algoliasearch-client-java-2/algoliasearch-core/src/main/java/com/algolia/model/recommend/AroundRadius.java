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

/** Define the maximum radius for a geo search (in meters). */
@JsonDeserialize(using = AroundRadius.AroundRadiusDeserializer.class)
@JsonSerialize(using = AroundRadius.AroundRadiusSerializer.class)
public abstract class AroundRadius implements CompoundType {

  public static AroundRadius of(AroundRadiusAll inside) {
    return new AroundRadiusAroundRadiusAll(inside);
  }

  public static AroundRadius of(Integer inside) {
    return new AroundRadiusInteger(inside);
  }

  public static class AroundRadiusSerializer extends StdSerializer<AroundRadius> {

    public AroundRadiusSerializer(Class<AroundRadius> t) {
      super(t);
    }

    public AroundRadiusSerializer() {
      this(null);
    }

    @Override
    public void serialize(AroundRadius value, JsonGenerator jgen, SerializerProvider provider) throws IOException, JsonProcessingException {
      jgen.writeObject(value.getInsideValue());
    }
  }

  public static class AroundRadiusDeserializer extends StdDeserializer<AroundRadius> {

    public AroundRadiusDeserializer() {
      this(AroundRadius.class);
    }

    public AroundRadiusDeserializer(Class<?> vc) {
      super(vc);
    }

    @Override
    public AroundRadius deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException, JsonProcessingException {
      JsonNode tree = jp.readValueAsTree();
      AroundRadius deserialized = null;

      int match = 0;
      JsonToken token = tree.traverse(jp.getCodec()).nextToken();
      String currentType = "";
      // deserialize AroundRadiusAll
      try {
        boolean attemptParsing = true;
        currentType = "AroundRadiusAll";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized =
            AroundRadius.of((AroundRadiusAll) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<AroundRadiusAll>() {}));
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              AroundRadius.of((AroundRadiusAll) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<AroundRadiusAll>() {}));
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println("Failed to deserialize oneOf AroundRadiusAll (error: " + e.getMessage() + ") (type: " + currentType + ")");
      }

      // deserialize Integer
      try {
        boolean attemptParsing = true;
        currentType = "Integer";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized = AroundRadius.of((Integer) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<Integer>() {}));
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized = AroundRadius.of((Integer) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<Integer>() {}));
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println("Failed to deserialize oneOf Integer (error: " + e.getMessage() + ") (type: " + currentType + ")");
      }

      if (match == 1) {
        return deserialized;
      }
      throw new IOException(String.format("Failed deserialization for AroundRadius: %d classes match result, expected 1", match));
    }

    /** Handle deserialization of the 'null' value. */
    @Override
    public AroundRadius getNullValue(DeserializationContext ctxt) throws JsonMappingException {
      throw new JsonMappingException(ctxt.getParser(), "AroundRadius cannot be null");
    }
  }
}

class AroundRadiusAroundRadiusAll extends AroundRadius {

  private final AroundRadiusAll insideValue;

  AroundRadiusAroundRadiusAll(AroundRadiusAll insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public AroundRadiusAll getInsideValue() {
    return insideValue;
  }
}

class AroundRadiusInteger extends AroundRadius {

  private final Integer insideValue;

  AroundRadiusInteger(Integer insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public Integer getInsideValue() {
    return insideValue;
  }
}
