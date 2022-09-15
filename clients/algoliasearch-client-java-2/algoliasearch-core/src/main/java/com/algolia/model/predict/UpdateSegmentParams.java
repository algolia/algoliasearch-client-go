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

/** UpdateSegmentParams */
@JsonDeserialize(using = UpdateSegmentParams.UpdateSegmentParamsDeserializer.class)
@JsonSerialize(using = UpdateSegmentParams.UpdateSegmentParamsSerializer.class)
public abstract class UpdateSegmentParams implements CompoundType {

  public static UpdateSegmentParams of(AllUpdateSegmentParams inside) {
    return new UpdateSegmentParamsAllUpdateSegmentParams(inside);
  }

  public static UpdateSegmentParams of(SegmentConditionsParam inside) {
    return new UpdateSegmentParamsSegmentConditionsParam(inside);
  }

  public static UpdateSegmentParams of(SegmentNameParam inside) {
    return new UpdateSegmentParamsSegmentNameParam(inside);
  }

  public static class UpdateSegmentParamsSerializer extends StdSerializer<UpdateSegmentParams> {

    public UpdateSegmentParamsSerializer(Class<UpdateSegmentParams> t) {
      super(t);
    }

    public UpdateSegmentParamsSerializer() {
      this(null);
    }

    @Override
    public void serialize(UpdateSegmentParams value, JsonGenerator jgen, SerializerProvider provider)
      throws IOException, JsonProcessingException {
      jgen.writeObject(value.getInsideValue());
    }
  }

  public static class UpdateSegmentParamsDeserializer extends StdDeserializer<UpdateSegmentParams> {

    public UpdateSegmentParamsDeserializer() {
      this(UpdateSegmentParams.class);
    }

    public UpdateSegmentParamsDeserializer(Class<?> vc) {
      super(vc);
    }

    @Override
    public UpdateSegmentParams deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException, JsonProcessingException {
      JsonNode tree = jp.readValueAsTree();
      UpdateSegmentParams deserialized = null;

      int match = 0;
      JsonToken token = tree.traverse(jp.getCodec()).nextToken();
      String currentType = "";
      // deserialize AllUpdateSegmentParams
      try {
        boolean attemptParsing = true;
        currentType = "AllUpdateSegmentParams";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized =
            UpdateSegmentParams.of(
              (AllUpdateSegmentParams) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<AllUpdateSegmentParams>() {})
            );
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              UpdateSegmentParams.of(
                (AllUpdateSegmentParams) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<AllUpdateSegmentParams>() {})
              );
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println(
          "Failed to deserialize oneOf AllUpdateSegmentParams (error: " + e.getMessage() + ") (type: " + currentType + ")"
        );
      }

      // deserialize SegmentConditionsParam
      try {
        boolean attemptParsing = true;
        currentType = "SegmentConditionsParam";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized =
            UpdateSegmentParams.of(
              (SegmentConditionsParam) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<SegmentConditionsParam>() {})
            );
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              UpdateSegmentParams.of(
                (SegmentConditionsParam) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<SegmentConditionsParam>() {})
              );
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println(
          "Failed to deserialize oneOf SegmentConditionsParam (error: " + e.getMessage() + ") (type: " + currentType + ")"
        );
      }

      // deserialize SegmentNameParam
      try {
        boolean attemptParsing = true;
        currentType = "SegmentNameParam";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized =
            UpdateSegmentParams.of((SegmentNameParam) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<SegmentNameParam>() {}));
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              UpdateSegmentParams.of((SegmentNameParam) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<SegmentNameParam>() {}));
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println("Failed to deserialize oneOf SegmentNameParam (error: " + e.getMessage() + ") (type: " + currentType + ")");
      }

      if (match == 1) {
        return deserialized;
      }
      throw new IOException(String.format("Failed deserialization for UpdateSegmentParams: %d classes match result, expected 1", match));
    }

    /** Handle deserialization of the 'null' value. */
    @Override
    public UpdateSegmentParams getNullValue(DeserializationContext ctxt) throws JsonMappingException {
      throw new JsonMappingException(ctxt.getParser(), "UpdateSegmentParams cannot be null");
    }
  }
}

class UpdateSegmentParamsAllUpdateSegmentParams extends UpdateSegmentParams {

  private final AllUpdateSegmentParams insideValue;

  UpdateSegmentParamsAllUpdateSegmentParams(AllUpdateSegmentParams insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public AllUpdateSegmentParams getInsideValue() {
    return insideValue;
  }
}

class UpdateSegmentParamsSegmentConditionsParam extends UpdateSegmentParams {

  private final SegmentConditionsParam insideValue;

  UpdateSegmentParamsSegmentConditionsParam(SegmentConditionsParam insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public SegmentConditionsParam getInsideValue() {
    return insideValue;
  }
}

class UpdateSegmentParamsSegmentNameParam extends UpdateSegmentParams {

  private final SegmentNameParam insideValue;

  UpdateSegmentParamsSegmentNameParam(SegmentNameParam insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public SegmentNameParam getInsideValue() {
    return insideValue;
  }
}
