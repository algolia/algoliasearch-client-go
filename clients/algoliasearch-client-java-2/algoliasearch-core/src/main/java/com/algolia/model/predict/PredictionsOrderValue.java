// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.predict;

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

/** PredictionsOrderValue */
@JsonDeserialize(using = PredictionsOrderValue.PredictionsOrderValueDeserializer.class)
@JsonSerialize(using = PredictionsOrderValue.PredictionsOrderValueSerializer.class)
public abstract class PredictionsOrderValue implements CompoundType {

  public static PredictionsOrderValue of(Error inside) {
    return new PredictionsOrderValueError(inside);
  }

  public static PredictionsOrderValue of(PredictionsOrderValueSuccess inside) {
    return new PredictionsOrderValuePredictionsOrderValueSuccess(inside);
  }

  public static class PredictionsOrderValueSerializer extends StdSerializer<PredictionsOrderValue> {

    public PredictionsOrderValueSerializer(Class<PredictionsOrderValue> t) {
      super(t);
    }

    public PredictionsOrderValueSerializer() {
      this(null);
    }

    @Override
    public void serialize(PredictionsOrderValue value, JsonGenerator jgen, SerializerProvider provider)
      throws IOException, JsonProcessingException {
      jgen.writeObject(value.getInsideValue());
    }
  }

  public static class PredictionsOrderValueDeserializer extends StdDeserializer<PredictionsOrderValue> {

    public PredictionsOrderValueDeserializer() {
      this(PredictionsOrderValue.class);
    }

    public PredictionsOrderValueDeserializer(Class<?> vc) {
      super(vc);
    }

    @Override
    public PredictionsOrderValue deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException, JsonProcessingException {
      JsonNode tree = jp.readValueAsTree();
      PredictionsOrderValue deserialized = null;

      int match = 0;
      JsonToken token = tree.traverse(jp.getCodec()).nextToken();
      String currentType = "";
      // deserialize Error
      try {
        boolean attemptParsing = true;
        currentType = "Error";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized = PredictionsOrderValue.of((Error) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<Error>() {}));
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized = PredictionsOrderValue.of((Error) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<Error>() {}));
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println("Failed to deserialize oneOf Error (error: " + e.getMessage() + ") (type: " + currentType + ")");
      }

      // deserialize PredictionsOrderValueSuccess
      try {
        boolean attemptParsing = true;
        currentType = "PredictionsOrderValueSuccess";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized =
            PredictionsOrderValue.of(
              (PredictionsOrderValueSuccess) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<PredictionsOrderValueSuccess>() {})
            );
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              PredictionsOrderValue.of(
                (PredictionsOrderValueSuccess) tree
                  .traverse(jp.getCodec())
                  .readValueAs(new TypeReference<PredictionsOrderValueSuccess>() {})
              );
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println(
          "Failed to deserialize oneOf PredictionsOrderValueSuccess (error: " + e.getMessage() + ") (type: " + currentType + ")"
        );
      }

      if (match == 1) {
        return deserialized;
      }
      throw new IOException(
        String.format("Failed deserialization for PredictionsOrderValue: %d classes match result, expected" + " 1", match)
      );
    }

    /** Handle deserialization of the 'null' value. */
    @Override
    public PredictionsOrderValue getNullValue(DeserializationContext ctxt) throws JsonMappingException {
      throw new JsonMappingException(ctxt.getParser(), "PredictionsOrderValue cannot be null");
    }
  }
}

class PredictionsOrderValueError extends PredictionsOrderValue {

  private final Error insideValue;

  PredictionsOrderValueError(Error insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public Error getInsideValue() {
    return insideValue;
  }
}

class PredictionsOrderValuePredictionsOrderValueSuccess extends PredictionsOrderValue {

  private final PredictionsOrderValueSuccess insideValue;

  PredictionsOrderValuePredictionsOrderValueSuccess(PredictionsOrderValueSuccess insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public PredictionsOrderValueSuccess getInsideValue() {
    return insideValue;
  }
}
