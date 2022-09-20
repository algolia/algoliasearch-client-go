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

/** PredictionsAffinities */
@JsonDeserialize(using = PredictionsAffinities.PredictionsAffinitiesDeserializer.class)
@JsonSerialize(using = PredictionsAffinities.PredictionsAffinitiesSerializer.class)
public abstract class PredictionsAffinities implements CompoundType {

  public static PredictionsAffinities of(Error inside) {
    return new PredictionsAffinitiesError(inside);
  }

  public static PredictionsAffinities of(PredictionsAffinitiesSuccess inside) {
    return new PredictionsAffinitiesPredictionsAffinitiesSuccess(inside);
  }

  public static class PredictionsAffinitiesSerializer extends StdSerializer<PredictionsAffinities> {

    public PredictionsAffinitiesSerializer(Class<PredictionsAffinities> t) {
      super(t);
    }

    public PredictionsAffinitiesSerializer() {
      this(null);
    }

    @Override
    public void serialize(PredictionsAffinities value, JsonGenerator jgen, SerializerProvider provider)
      throws IOException, JsonProcessingException {
      jgen.writeObject(value.getInsideValue());
    }
  }

  public static class PredictionsAffinitiesDeserializer extends StdDeserializer<PredictionsAffinities> {

    public PredictionsAffinitiesDeserializer() {
      this(PredictionsAffinities.class);
    }

    public PredictionsAffinitiesDeserializer(Class<?> vc) {
      super(vc);
    }

    @Override
    public PredictionsAffinities deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException, JsonProcessingException {
      JsonNode tree = jp.readValueAsTree();
      PredictionsAffinities deserialized = null;

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
          deserialized = PredictionsAffinities.of((Error) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<Error>() {}));
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized = PredictionsAffinities.of((Error) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<Error>() {}));
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println("Failed to deserialize oneOf Error (error: " + e.getMessage() + ") (type: " + currentType + ")");
      }

      // deserialize PredictionsAffinitiesSuccess
      try {
        boolean attemptParsing = true;
        currentType = "PredictionsAffinitiesSuccess";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized =
            PredictionsAffinities.of(
              (PredictionsAffinitiesSuccess) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<PredictionsAffinitiesSuccess>() {})
            );
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              PredictionsAffinities.of(
                (PredictionsAffinitiesSuccess) tree
                  .traverse(jp.getCodec())
                  .readValueAs(new TypeReference<PredictionsAffinitiesSuccess>() {})
              );
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println(
          "Failed to deserialize oneOf PredictionsAffinitiesSuccess (error: " + e.getMessage() + ") (type: " + currentType + ")"
        );
      }

      if (match == 1) {
        return deserialized;
      }
      throw new IOException(
        String.format("Failed deserialization for PredictionsAffinities: %d classes match result, expected" + " 1", match)
      );
    }

    /** Handle deserialization of the 'null' value. */
    @Override
    public PredictionsAffinities getNullValue(DeserializationContext ctxt) throws JsonMappingException {
      throw new JsonMappingException(ctxt.getParser(), "PredictionsAffinities cannot be null");
    }
  }
}

class PredictionsAffinitiesError extends PredictionsAffinities {

  private final Error insideValue;

  PredictionsAffinitiesError(Error insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public Error getInsideValue() {
    return insideValue;
  }
}

class PredictionsAffinitiesPredictionsAffinitiesSuccess extends PredictionsAffinities {

  private final PredictionsAffinitiesSuccess insideValue;

  PredictionsAffinitiesPredictionsAffinitiesSuccess(PredictionsAffinitiesSuccess insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public PredictionsAffinitiesSuccess getInsideValue() {
    return insideValue;
  }
}
