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

/** Enables de-duplication or grouping of results. */
@JsonDeserialize(using = Distinct.DistinctDeserializer.class)
@JsonSerialize(using = Distinct.DistinctSerializer.class)
public abstract class Distinct implements CompoundType {

  public static Distinct of(Boolean inside) {
    return new DistinctBoolean(inside);
  }

  public static Distinct of(Integer inside) {
    return new DistinctInteger(inside);
  }

  public static class DistinctSerializer extends StdSerializer<Distinct> {

    public DistinctSerializer(Class<Distinct> t) {
      super(t);
    }

    public DistinctSerializer() {
      this(null);
    }

    @Override
    public void serialize(Distinct value, JsonGenerator jgen, SerializerProvider provider) throws IOException, JsonProcessingException {
      jgen.writeObject(value.getInsideValue());
    }
  }

  public static class DistinctDeserializer extends StdDeserializer<Distinct> {

    public DistinctDeserializer() {
      this(Distinct.class);
    }

    public DistinctDeserializer(Class<?> vc) {
      super(vc);
    }

    @Override
    public Distinct deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException, JsonProcessingException {
      JsonNode tree = jp.readValueAsTree();
      Distinct deserialized = null;

      int match = 0;
      JsonToken token = tree.traverse(jp.getCodec()).nextToken();
      String currentType = "";
      // deserialize Boolean
      try {
        boolean attemptParsing = true;
        currentType = "Boolean";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized = Distinct.of((Boolean) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<Boolean>() {}));
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized = Distinct.of((Boolean) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<Boolean>() {}));
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println("Failed to deserialize oneOf Boolean (error: " + e.getMessage() + ") (type: " + currentType + ")");
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
          deserialized = Distinct.of((Integer) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<Integer>() {}));
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized = Distinct.of((Integer) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<Integer>() {}));
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
      throw new IOException(String.format("Failed deserialization for Distinct: %d classes match result, expected 1", match));
    }

    /** Handle deserialization of the 'null' value. */
    @Override
    public Distinct getNullValue(DeserializationContext ctxt) throws JsonMappingException {
      throw new JsonMappingException(ctxt.getParser(), "Distinct cannot be null");
    }
  }
}

class DistinctBoolean extends Distinct {

  private final Boolean insideValue;

  DistinctBoolean(Boolean insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public Boolean getInsideValue() {
    return insideValue;
  }
}

class DistinctInteger extends Distinct {

  private final Integer insideValue;

  DistinctInteger(Integer insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public Integer getInsideValue() {
    return insideValue;
  }
}
