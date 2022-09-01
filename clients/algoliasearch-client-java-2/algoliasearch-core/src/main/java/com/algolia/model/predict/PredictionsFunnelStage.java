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

/** PredictionsFunnelStage */
@JsonDeserialize(using = PredictionsFunnelStage.PredictionsFunnelStageDeserializer.class)
@JsonSerialize(using = PredictionsFunnelStage.PredictionsFunnelStageSerializer.class)
public abstract class PredictionsFunnelStage implements CompoundType {

  public static PredictionsFunnelStage of(Error inside) {
    return new PredictionsFunnelStageError(inside);
  }

  public static PredictionsFunnelStage of(FunnelStageSuccess inside) {
    return new PredictionsFunnelStageFunnelStageSuccess(inside);
  }

  public static class PredictionsFunnelStageSerializer extends StdSerializer<PredictionsFunnelStage> {

    public PredictionsFunnelStageSerializer(Class<PredictionsFunnelStage> t) {
      super(t);
    }

    public PredictionsFunnelStageSerializer() {
      this(null);
    }

    @Override
    public void serialize(PredictionsFunnelStage value, JsonGenerator jgen, SerializerProvider provider)
      throws IOException, JsonProcessingException {
      jgen.writeObject(value.getInsideValue());
    }
  }

  public static class PredictionsFunnelStageDeserializer extends StdDeserializer<PredictionsFunnelStage> {

    public PredictionsFunnelStageDeserializer() {
      this(PredictionsFunnelStage.class);
    }

    public PredictionsFunnelStageDeserializer(Class<?> vc) {
      super(vc);
    }

    @Override
    public PredictionsFunnelStage deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException, JsonProcessingException {
      JsonNode tree = jp.readValueAsTree();
      PredictionsFunnelStage deserialized = null;

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
          deserialized = PredictionsFunnelStage.of((Error) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<Error>() {}));
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized = PredictionsFunnelStage.of((Error) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<Error>() {}));
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
            PredictionsFunnelStage.of(
              (FunnelStageSuccess) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<FunnelStageSuccess>() {})
            );
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              PredictionsFunnelStage.of(
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
        String.format("Failed deserialization for PredictionsFunnelStage: %d classes match result, expected" + " 1", match)
      );
    }

    /** Handle deserialization of the 'null' value. */
    @Override
    public PredictionsFunnelStage getNullValue(DeserializationContext ctxt) throws JsonMappingException {
      throw new JsonMappingException(ctxt.getParser(), "PredictionsFunnelStage cannot be null");
    }
  }
}

class PredictionsFunnelStageError extends PredictionsFunnelStage {

  private final Error insideValue;

  PredictionsFunnelStageError(Error insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public Error getInsideValue() {
    return insideValue;
  }
}

class PredictionsFunnelStageFunnelStageSuccess extends PredictionsFunnelStage {

  private final FunnelStageSuccess insideValue;

  PredictionsFunnelStageFunnelStageSuccess(FunnelStageSuccess insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public FunnelStageSuccess getInsideValue() {
    return insideValue;
  }
}
