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

/** AttributeToUpdate */
@JsonDeserialize(using = AttributeToUpdate.AttributeToUpdateDeserializer.class)
@JsonSerialize(using = AttributeToUpdate.AttributeToUpdateSerializer.class)
public abstract class AttributeToUpdate implements CompoundType {

  public static AttributeToUpdate of(BuiltInOperation inside) {
    return new AttributeToUpdateBuiltInOperation(inside);
  }

  public static AttributeToUpdate of(String inside) {
    return new AttributeToUpdateString(inside);
  }

  public static class AttributeToUpdateSerializer extends StdSerializer<AttributeToUpdate> {

    public AttributeToUpdateSerializer(Class<AttributeToUpdate> t) {
      super(t);
    }

    public AttributeToUpdateSerializer() {
      this(null);
    }

    @Override
    public void serialize(AttributeToUpdate value, JsonGenerator jgen, SerializerProvider provider)
      throws IOException, JsonProcessingException {
      jgen.writeObject(value.getInsideValue());
    }
  }

  public static class AttributeToUpdateDeserializer extends StdDeserializer<AttributeToUpdate> {

    public AttributeToUpdateDeserializer() {
      this(AttributeToUpdate.class);
    }

    public AttributeToUpdateDeserializer(Class<?> vc) {
      super(vc);
    }

    @Override
    public AttributeToUpdate deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException, JsonProcessingException {
      JsonNode tree = jp.readValueAsTree();
      AttributeToUpdate deserialized = null;

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
            AttributeToUpdate.of((BuiltInOperation) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<BuiltInOperation>() {}));
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              AttributeToUpdate.of((BuiltInOperation) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<BuiltInOperation>() {}));
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
          deserialized = AttributeToUpdate.of((String) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<String>() {}));
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized = AttributeToUpdate.of((String) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<String>() {}));
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
      throw new IOException(String.format("Failed deserialization for AttributeToUpdate: %d classes match result, expected 1", match));
    }

    /** Handle deserialization of the 'null' value. */
    @Override
    public AttributeToUpdate getNullValue(DeserializationContext ctxt) throws JsonMappingException {
      throw new JsonMappingException(ctxt.getParser(), "AttributeToUpdate cannot be null");
    }
  }
}

class AttributeToUpdateBuiltInOperation extends AttributeToUpdate {

  private final BuiltInOperation insideValue;

  AttributeToUpdateBuiltInOperation(BuiltInOperation insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public BuiltInOperation getInsideValue() {
    return insideValue;
  }
}

class AttributeToUpdateString extends AttributeToUpdate {

  private final String insideValue;

  AttributeToUpdateString(String insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public String getInsideValue() {
    return insideValue;
  }
}
