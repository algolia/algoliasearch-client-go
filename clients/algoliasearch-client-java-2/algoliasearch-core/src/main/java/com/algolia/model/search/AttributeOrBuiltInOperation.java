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

/** AttributeOrBuiltInOperation */
@JsonDeserialize(using = AttributeOrBuiltInOperation.AttributeOrBuiltInOperationDeserializer.class)
@JsonSerialize(using = AttributeOrBuiltInOperation.AttributeOrBuiltInOperationSerializer.class)
public abstract class AttributeOrBuiltInOperation implements CompoundType {

  public static AttributeOrBuiltInOperation of(BuiltInOperation inside) {
    return new AttributeOrBuiltInOperationBuiltInOperation(inside);
  }

  public static AttributeOrBuiltInOperation of(String inside) {
    return new AttributeOrBuiltInOperationString(inside);
  }

  public static class AttributeOrBuiltInOperationSerializer extends StdSerializer<AttributeOrBuiltInOperation> {

    public AttributeOrBuiltInOperationSerializer(Class<AttributeOrBuiltInOperation> t) {
      super(t);
    }

    public AttributeOrBuiltInOperationSerializer() {
      this(null);
    }

    @Override
    public void serialize(AttributeOrBuiltInOperation value, JsonGenerator jgen, SerializerProvider provider)
      throws IOException, JsonProcessingException {
      jgen.writeObject(value.getInsideValue());
    }
  }

  public static class AttributeOrBuiltInOperationDeserializer extends StdDeserializer<AttributeOrBuiltInOperation> {

    public AttributeOrBuiltInOperationDeserializer() {
      this(AttributeOrBuiltInOperation.class);
    }

    public AttributeOrBuiltInOperationDeserializer(Class<?> vc) {
      super(vc);
    }

    @Override
    public AttributeOrBuiltInOperation deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException, JsonProcessingException {
      JsonNode tree = jp.readValueAsTree();
      AttributeOrBuiltInOperation deserialized = null;

      int match = 0;
      JsonToken token = tree.traverse(jp.getCodec()).nextToken();
      String currentType = "";
      // deserialize BuiltInOperation
      try {
        boolean attemptParsing = true;
        currentType = "BuiltInOperation";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized =
            AttributeOrBuiltInOperation.of(
              (BuiltInOperation) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<BuiltInOperation>() {})
            );
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              AttributeOrBuiltInOperation.of(
                (BuiltInOperation) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<BuiltInOperation>() {})
              );
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println("Failed to deserialize oneOf BuiltInOperation (error: " + e.getMessage() + ") (type: " + currentType + ")");
      }

      // deserialize String
      try {
        boolean attemptParsing = true;
        currentType = "String";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized = AttributeOrBuiltInOperation.of((String) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<String>() {}));
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              AttributeOrBuiltInOperation.of((String) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<String>() {}));
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println("Failed to deserialize oneOf String (error: " + e.getMessage() + ") (type: " + currentType + ")");
      }

      if (match == 1) {
        return deserialized;
      }
      throw new IOException(
        String.format("Failed deserialization for AttributeOrBuiltInOperation: %d classes match result," + " expected 1", match)
      );
    }

    /** Handle deserialization of the 'null' value. */
    @Override
    public AttributeOrBuiltInOperation getNullValue(DeserializationContext ctxt) throws JsonMappingException {
      throw new JsonMappingException(ctxt.getParser(), "AttributeOrBuiltInOperation cannot be null");
    }
  }
}

class AttributeOrBuiltInOperationBuiltInOperation extends AttributeOrBuiltInOperation {

  private final BuiltInOperation insideValue;

  AttributeOrBuiltInOperationBuiltInOperation(BuiltInOperation insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public BuiltInOperation getInsideValue() {
    return insideValue;
  }
}

class AttributeOrBuiltInOperationString extends AttributeOrBuiltInOperation {

  private final String insideValue;

  AttributeOrBuiltInOperationString(String insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public String getInsideValue() {
    return insideValue;
  }
}
