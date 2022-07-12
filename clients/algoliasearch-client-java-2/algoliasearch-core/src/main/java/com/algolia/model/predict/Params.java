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

/** Params */
@JsonDeserialize(using = Params.ParamsDeserializer.class)
@JsonSerialize(using = Params.ParamsSerializer.class)
public abstract class Params implements CompoundType {

  public static Params of(AllParams inside) {
    return new ParamsAllParams(inside);
  }

  public static Params of(ModelsToRetrieve inside) {
    return new ParamsModelsToRetrieve(inside);
  }

  public static Params of(TypesToRetrieve inside) {
    return new ParamsTypesToRetrieve(inside);
  }

  public static class ParamsSerializer extends StdSerializer<Params> {

    public ParamsSerializer(Class<Params> t) {
      super(t);
    }

    public ParamsSerializer() {
      this(null);
    }

    @Override
    public void serialize(Params value, JsonGenerator jgen, SerializerProvider provider) throws IOException, JsonProcessingException {
      jgen.writeObject(value.getInsideValue());
    }
  }

  public static class ParamsDeserializer extends StdDeserializer<Params> {

    public ParamsDeserializer() {
      this(Params.class);
    }

    public ParamsDeserializer(Class<?> vc) {
      super(vc);
    }

    @Override
    public Params deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException, JsonProcessingException {
      JsonNode tree = jp.readValueAsTree();
      Params deserialized = null;

      int match = 0;
      JsonToken token = tree.traverse(jp.getCodec()).nextToken();
      String currentType = "";
      // deserialize AllParams
      try {
        boolean attemptParsing = true;
        currentType = "AllParams";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized = Params.of((AllParams) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<AllParams>() {}));
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized = Params.of((AllParams) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<AllParams>() {}));
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println("Failed to deserialize oneOf AllParams (error: " + e.getMessage() + ") (type: " + currentType + ")");
      }

      // deserialize ModelsToRetrieve
      try {
        boolean attemptParsing = true;
        currentType = "ModelsToRetrieve";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized = Params.of((ModelsToRetrieve) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<ModelsToRetrieve>() {}));
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized = Params.of((ModelsToRetrieve) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<ModelsToRetrieve>() {}));
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println("Failed to deserialize oneOf ModelsToRetrieve (error: " + e.getMessage() + ") (type: " + currentType + ")");
      }

      // deserialize TypesToRetrieve
      try {
        boolean attemptParsing = true;
        currentType = "TypesToRetrieve";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized = Params.of((TypesToRetrieve) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<TypesToRetrieve>() {}));
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized = Params.of((TypesToRetrieve) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<TypesToRetrieve>() {}));
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println("Failed to deserialize oneOf TypesToRetrieve (error: " + e.getMessage() + ") (type: " + currentType + ")");
      }

      if (match == 1) {
        return deserialized;
      }
      throw new IOException(String.format("Failed deserialization for Params: %d classes match result, expected 1", match));
    }

    /** Handle deserialization of the 'null' value. */
    @Override
    public Params getNullValue(DeserializationContext ctxt) throws JsonMappingException {
      throw new JsonMappingException(ctxt.getParser(), "Params cannot be null");
    }
  }
}

class ParamsAllParams extends Params {

  private final AllParams insideValue;

  ParamsAllParams(AllParams insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public AllParams getInsideValue() {
    return insideValue;
  }
}

class ParamsModelsToRetrieve extends Params {

  private final ModelsToRetrieve insideValue;

  ParamsModelsToRetrieve(ModelsToRetrieve insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public ModelsToRetrieve getInsideValue() {
    return insideValue;
  }
}

class ParamsTypesToRetrieve extends Params {

  private final TypesToRetrieve insideValue;

  ParamsTypesToRetrieve(TypesToRetrieve insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public TypesToRetrieve getInsideValue() {
    return insideValue;
  }
}
