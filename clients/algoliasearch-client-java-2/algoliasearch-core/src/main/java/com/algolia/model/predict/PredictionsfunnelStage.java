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

/** PredictionsfunnelStage */
@JsonDeserialize(using = PredictionsfunnelStage.PredictionsfunnelStageDeserializer.class)
@JsonSerialize(using = PredictionsfunnelStage.PredictionsfunnelStageSerializer.class)
public abstract class PredictionsfunnelStage implements CompoundType {

  public static PredictionsfunnelStage of(Error inside) {
    return new PredictionsfunnelStageError(inside);
  }

  public static PredictionsfunnelStage of(FunnelStageSuccess inside) {
    return new PredictionsfunnelStageFunnelStageSuccess(inside);
  }

  public static class PredictionsfunnelStageSerializer extends StdSerializer<PredictionsfunnelStage> {

    public PredictionsfunnelStageSerializer(Class<PredictionsfunnelStage> t) {
      super(t);
    }

    public PredictionsfunnelStageSerializer() {
      this(null);
    }

    @Override
    public void serialize(PredictionsfunnelStage value, JsonGenerator jgen, SerializerProvider provider)
      throws IOException, JsonProcessingException {
      jgen.writeObject(value.getInsideValue());
    }
  }

  public static class PredictionsfunnelStageDeserializer extends StdDeserializer<PredictionsfunnelStage> {

    public PredictionsfunnelStageDeserializer() {
      this(PredictionsfunnelStage.class);
    }

    public PredictionsfunnelStageDeserializer(Class<?> vc) {
      super(vc);
    }

    @Override
    public PredictionsfunnelStage deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException, JsonProcessingException {
      JsonNode tree = jp.readValueAsTree();
      PredictionsfunnelStage deserialized = null;

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
          deserialized = PredictionsfunnelStage.of((Error) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<Error>() {}));
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized = PredictionsfunnelStage.of((Error) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<Error>() {}));
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println("Failed to deserialize oneOf Error (error: " + e.getMessage() + ") (type: " + currentType + ")");
      }

      // deserialize FunnelStageSuccess
      try {
        boolean attemptParsing = true;
        currentType = "FunnelStageSuccess";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized =
            PredictionsfunnelStage.of(
              (FunnelStageSuccess) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<FunnelStageSuccess>() {})
            );
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              PredictionsfunnelStage.of(
                (FunnelStageSuccess) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<FunnelStageSuccess>() {})
              );
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println("Failed to deserialize oneOf FunnelStageSuccess (error: " + e.getMessage() + ") (type: " + currentType + ")");
      }

      if (match == 1) {
        return deserialized;
      }
      throw new IOException(
        String.format("Failed deserialization for PredictionsfunnelStage: %d classes match result, expected" + " 1", match)
      );
    }

    /** Handle deserialization of the 'null' value. */
    @Override
    public PredictionsfunnelStage getNullValue(DeserializationContext ctxt) throws JsonMappingException {
      throw new JsonMappingException(ctxt.getParser(), "PredictionsfunnelStage cannot be null");
    }
  }
}

class PredictionsfunnelStageError extends PredictionsfunnelStage {

  private final Error insideValue;

  PredictionsfunnelStageError(Error insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public Error getInsideValue() {
    return insideValue;
  }
}

class PredictionsfunnelStageFunnelStageSuccess extends PredictionsfunnelStage {

  private final FunnelStageSuccess insideValue;

  PredictionsfunnelStageFunnelStageSuccess(FunnelStageSuccess insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public FunnelStageSuccess getInsideValue() {
    return insideValue;
  }
}
