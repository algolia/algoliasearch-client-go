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

  public static Params of(ModelsToRetrieveParam inside) {
    return new ParamsModelsToRetrieveParam(inside);
  }

  public static Params of(TypesToRetrieveParam inside) {
    return new ParamsTypesToRetrieveParam(inside);
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

      // deserialize ModelsToRetrieveParam
      try {
        boolean attemptParsing = true;
        currentType = "ModelsToRetrieveParam";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized =
            Params.of((ModelsToRetrieveParam) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<ModelsToRetrieveParam>() {}));
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              Params.of((ModelsToRetrieveParam) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<ModelsToRetrieveParam>() {}));
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println("Failed to deserialize oneOf ModelsToRetrieveParam (error: " + e.getMessage() + ") (type: " + currentType + ")");
      }

      // deserialize TypesToRetrieveParam
      try {
        boolean attemptParsing = true;
        currentType = "TypesToRetrieveParam";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized =
            Params.of((TypesToRetrieveParam) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<TypesToRetrieveParam>() {}));
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              Params.of((TypesToRetrieveParam) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<TypesToRetrieveParam>() {}));
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println("Failed to deserialize oneOf TypesToRetrieveParam (error: " + e.getMessage() + ") (type: " + currentType + ")");
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

class ParamsModelsToRetrieveParam extends Params {

  private final ModelsToRetrieveParam insideValue;

  ParamsModelsToRetrieveParam(ModelsToRetrieveParam insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public ModelsToRetrieveParam getInsideValue() {
    return insideValue;
  }
}

class ParamsTypesToRetrieveParam extends Params {

  private final TypesToRetrieveParam insideValue;

  ParamsTypesToRetrieveParam(TypesToRetrieveParam insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public TypesToRetrieveParam getInsideValue() {
    return insideValue;
  }
}
