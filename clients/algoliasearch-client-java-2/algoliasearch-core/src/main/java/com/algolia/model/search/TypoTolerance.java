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

/** Controls whether typo tolerance is enabled and how it is applied. */
@JsonDeserialize(using = TypoTolerance.TypoToleranceDeserializer.class)
@JsonSerialize(using = TypoTolerance.TypoToleranceSerializer.class)
public abstract class TypoTolerance implements CompoundType {

  public static TypoTolerance of(Boolean inside) {
    return new TypoToleranceBoolean(inside);
  }

  public static TypoTolerance of(TypoToleranceEnum inside) {
    return new TypoToleranceTypoToleranceEnum(inside);
  }

  public static class TypoToleranceSerializer extends StdSerializer<TypoTolerance> {

    public TypoToleranceSerializer(Class<TypoTolerance> t) {
      super(t);
    }

    public TypoToleranceSerializer() {
      this(null);
    }

    @Override
    public void serialize(TypoTolerance value, JsonGenerator jgen, SerializerProvider provider)
      throws IOException, JsonProcessingException {
      jgen.writeObject(value.getInsideValue());
    }
  }

  public static class TypoToleranceDeserializer extends StdDeserializer<TypoTolerance> {

    public TypoToleranceDeserializer() {
      this(TypoTolerance.class);
    }

    public TypoToleranceDeserializer(Class<?> vc) {
      super(vc);
    }

    @Override
    public TypoTolerance deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException, JsonProcessingException {
      JsonNode tree = jp.readValueAsTree();
      TypoTolerance deserialized = null;

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
          deserialized = TypoTolerance.of((Boolean) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<Boolean>() {}));
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized = TypoTolerance.of((Boolean) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<Boolean>() {}));
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println("Failed to deserialize oneOf Boolean (error: " + e.getMessage() + ") (type: " + currentType + ")");
      }

      // deserialize TypoToleranceEnum
      try {
        boolean attemptParsing = true;
        currentType = "TypoToleranceEnum";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized =
            TypoTolerance.of((TypoToleranceEnum) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<TypoToleranceEnum>() {}));
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              TypoTolerance.of((TypoToleranceEnum) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<TypoToleranceEnum>() {}));
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println("Failed to deserialize oneOf TypoToleranceEnum (error: " + e.getMessage() + ") (type: " + currentType + ")");
      }

      if (match == 1) {
        return deserialized;
      }
      throw new IOException(String.format("Failed deserialization for TypoTolerance: %d classes match result, expected 1", match));
    }

    /** Handle deserialization of the 'null' value. */
    @Override
    public TypoTolerance getNullValue(DeserializationContext ctxt) throws JsonMappingException {
      throw new JsonMappingException(ctxt.getParser(), "TypoTolerance cannot be null");
    }
  }
}

class TypoToleranceBoolean extends TypoTolerance {

  private final Boolean insideValue;

  TypoToleranceBoolean(Boolean insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public Boolean getInsideValue() {
    return insideValue;
  }
}

class TypoToleranceTypoToleranceEnum extends TypoTolerance {

  private final TypoToleranceEnum insideValue;

  TypoToleranceTypoToleranceEnum(TypoToleranceEnum insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public TypoToleranceEnum getInsideValue() {
    return insideValue;
  }
}
