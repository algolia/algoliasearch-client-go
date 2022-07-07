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
import java.util.List;

/** HighlightResult */
@JsonDeserialize(using = HighlightResult.HighlightResultDeserializer.class)
@JsonSerialize(using = HighlightResult.HighlightResultSerializer.class)
public abstract class HighlightResult implements CompoundType {

  public static HighlightResult of(HighlightResultOption inside) {
    return new HighlightResultHighlightResultOption(inside);
  }

  public static HighlightResult of(List<HighlightResultOption> inside) {
    return new HighlightResultListOfHighlightResultOption(inside);
  }

  public static class HighlightResultSerializer extends StdSerializer<HighlightResult> {

    public HighlightResultSerializer(Class<HighlightResult> t) {
      super(t);
    }

    public HighlightResultSerializer() {
      this(null);
    }

    @Override
    public void serialize(HighlightResult value, JsonGenerator jgen, SerializerProvider provider)
      throws IOException, JsonProcessingException {
      jgen.writeObject(value.getInsideValue());
    }
  }

  public static class HighlightResultDeserializer extends StdDeserializer<HighlightResult> {

    public HighlightResultDeserializer() {
      this(HighlightResult.class);
    }

    public HighlightResultDeserializer(Class<?> vc) {
      super(vc);
    }

    @Override
    public HighlightResult deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException, JsonProcessingException {
      JsonNode tree = jp.readValueAsTree();
      HighlightResult deserialized = null;

      int match = 0;
      JsonToken token = tree.traverse(jp.getCodec()).nextToken();
      String currentType = "";
      // deserialize HighlightResultOption
      try {
        boolean attemptParsing = true;
        currentType = "HighlightResultOption";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized =
            HighlightResult.of(
              (HighlightResultOption) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<HighlightResultOption>() {})
            );
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              HighlightResult.of(
                (HighlightResultOption) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<HighlightResultOption>() {})
              );
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println("Failed to deserialize oneOf HighlightResultOption (error: " + e.getMessage() + ") (type: " + currentType + ")");
      }

      // deserialize List<HighlightResultOption>
      try {
        boolean attemptParsing = true;
        currentType = "List<HighlightResultOption>";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized =
            HighlightResult.of(
              (List<HighlightResultOption>) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<List<HighlightResultOption>>() {})
            );
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              HighlightResult.of(
                (List<HighlightResultOption>) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<List<HighlightResultOption>>() {})
              );
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println(
          "Failed to deserialize oneOf List<HighlightResultOption> (error: " + e.getMessage() + ") (type: " + currentType + ")"
        );
      }

      if (match == 1) {
        return deserialized;
      }
      throw new IOException(String.format("Failed deserialization for HighlightResult: %d classes match result, expected 1", match));
    }

    /** Handle deserialization of the 'null' value. */
    @Override
    public HighlightResult getNullValue(DeserializationContext ctxt) throws JsonMappingException {
      throw new JsonMappingException(ctxt.getParser(), "HighlightResult cannot be null");
    }
  }
}

class HighlightResultHighlightResultOption extends HighlightResult {

  private final HighlightResultOption insideValue;

  HighlightResultHighlightResultOption(HighlightResultOption insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public HighlightResultOption getInsideValue() {
    return insideValue;
  }
}

class HighlightResultListOfHighlightResultOption extends HighlightResult {

  private final List<HighlightResultOption> insideValue;

  HighlightResultListOfHighlightResultOption(List<HighlightResultOption> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<HighlightResultOption> getInsideValue() {
    return insideValue;
  }
}
