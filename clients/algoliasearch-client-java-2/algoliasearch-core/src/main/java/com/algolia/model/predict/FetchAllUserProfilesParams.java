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

/** FetchAllUserProfilesParams */
@JsonDeserialize(using = FetchAllUserProfilesParams.FetchAllUserProfilesParamsDeserializer.class)
@JsonSerialize(using = FetchAllUserProfilesParams.FetchAllUserProfilesParamsSerializer.class)
public abstract class FetchAllUserProfilesParams implements CompoundType {

  public static FetchAllUserProfilesParams of(LimitParam inside) {
    return new FetchAllUserProfilesParamsLimitParam(inside);
  }

  public static FetchAllUserProfilesParams of(ModelsToRetrieveParam inside) {
    return new FetchAllUserProfilesParamsModelsToRetrieveParam(inside);
  }

  public static FetchAllUserProfilesParams of(NextPageTokenParam inside) {
    return new FetchAllUserProfilesParamsNextPageTokenParam(inside);
  }

  public static FetchAllUserProfilesParams of(PreviousPageTokenParam inside) {
    return new FetchAllUserProfilesParamsPreviousPageTokenParam(inside);
  }

  public static FetchAllUserProfilesParams of(TypesToRetrieveParam inside) {
    return new FetchAllUserProfilesParamsTypesToRetrieveParam(inside);
  }

  public static class FetchAllUserProfilesParamsSerializer extends StdSerializer<FetchAllUserProfilesParams> {

    public FetchAllUserProfilesParamsSerializer(Class<FetchAllUserProfilesParams> t) {
      super(t);
    }

    public FetchAllUserProfilesParamsSerializer() {
      this(null);
    }

    @Override
    public void serialize(FetchAllUserProfilesParams value, JsonGenerator jgen, SerializerProvider provider)
      throws IOException, JsonProcessingException {
      jgen.writeObject(value.getInsideValue());
    }
  }

  public static class FetchAllUserProfilesParamsDeserializer extends StdDeserializer<FetchAllUserProfilesParams> {

    public FetchAllUserProfilesParamsDeserializer() {
      this(FetchAllUserProfilesParams.class);
    }

    public FetchAllUserProfilesParamsDeserializer(Class<?> vc) {
      super(vc);
    }

    @Override
    public FetchAllUserProfilesParams deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException, JsonProcessingException {
      JsonNode tree = jp.readValueAsTree();
      FetchAllUserProfilesParams deserialized = null;

      int match = 0;
      JsonToken token = tree.traverse(jp.getCodec()).nextToken();
      String currentType = "";
      // deserialize LimitParam
      try {
        boolean attemptParsing = true;
        currentType = "LimitParam";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized =
            FetchAllUserProfilesParams.of((LimitParam) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<LimitParam>() {}));
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              FetchAllUserProfilesParams.of((LimitParam) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<LimitParam>() {}));
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println("Failed to deserialize oneOf LimitParam (error: " + e.getMessage() + ") (type: " + currentType + ")");
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
            FetchAllUserProfilesParams.of(
              (ModelsToRetrieveParam) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<ModelsToRetrieveParam>() {})
            );
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              FetchAllUserProfilesParams.of(
                (ModelsToRetrieveParam) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<ModelsToRetrieveParam>() {})
              );
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println("Failed to deserialize oneOf ModelsToRetrieveParam (error: " + e.getMessage() + ") (type: " + currentType + ")");
      }

      // deserialize NextPageTokenParam
      try {
        boolean attemptParsing = true;
        currentType = "NextPageTokenParam";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized =
            FetchAllUserProfilesParams.of(
              (NextPageTokenParam) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<NextPageTokenParam>() {})
            );
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              FetchAllUserProfilesParams.of(
                (NextPageTokenParam) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<NextPageTokenParam>() {})
              );
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println("Failed to deserialize oneOf NextPageTokenParam (error: " + e.getMessage() + ") (type: " + currentType + ")");
      }

      // deserialize PreviousPageTokenParam
      try {
        boolean attemptParsing = true;
        currentType = "PreviousPageTokenParam";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized =
            FetchAllUserProfilesParams.of(
              (PreviousPageTokenParam) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<PreviousPageTokenParam>() {})
            );
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              FetchAllUserProfilesParams.of(
                (PreviousPageTokenParam) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<PreviousPageTokenParam>() {})
              );
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println(
          "Failed to deserialize oneOf PreviousPageTokenParam (error: " + e.getMessage() + ") (type: " + currentType + ")"
        );
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
            FetchAllUserProfilesParams.of(
              (TypesToRetrieveParam) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<TypesToRetrieveParam>() {})
            );
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              FetchAllUserProfilesParams.of(
                (TypesToRetrieveParam) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<TypesToRetrieveParam>() {})
              );
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
      throw new IOException(
        String.format("Failed deserialization for FetchAllUserProfilesParams: %d classes match result," + " expected 1", match)
      );
    }

    /** Handle deserialization of the 'null' value. */
    @Override
    public FetchAllUserProfilesParams getNullValue(DeserializationContext ctxt) throws JsonMappingException {
      throw new JsonMappingException(ctxt.getParser(), "FetchAllUserProfilesParams cannot be null");
    }
  }
}

class FetchAllUserProfilesParamsLimitParam extends FetchAllUserProfilesParams {

  private final LimitParam insideValue;

  FetchAllUserProfilesParamsLimitParam(LimitParam insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public LimitParam getInsideValue() {
    return insideValue;
  }
}

class FetchAllUserProfilesParamsModelsToRetrieveParam extends FetchAllUserProfilesParams {

  private final ModelsToRetrieveParam insideValue;

  FetchAllUserProfilesParamsModelsToRetrieveParam(ModelsToRetrieveParam insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public ModelsToRetrieveParam getInsideValue() {
    return insideValue;
  }
}

class FetchAllUserProfilesParamsNextPageTokenParam extends FetchAllUserProfilesParams {

  private final NextPageTokenParam insideValue;

  FetchAllUserProfilesParamsNextPageTokenParam(NextPageTokenParam insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public NextPageTokenParam getInsideValue() {
    return insideValue;
  }
}

class FetchAllUserProfilesParamsPreviousPageTokenParam extends FetchAllUserProfilesParams {

  private final PreviousPageTokenParam insideValue;

  FetchAllUserProfilesParamsPreviousPageTokenParam(PreviousPageTokenParam insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public PreviousPageTokenParam getInsideValue() {
    return insideValue;
  }
}

class FetchAllUserProfilesParamsTypesToRetrieveParam extends FetchAllUserProfilesParams {

  private final TypesToRetrieveParam insideValue;

  FetchAllUserProfilesParamsTypesToRetrieveParam(TypesToRetrieveParam insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public TypesToRetrieveParam getInsideValue() {
    return insideValue;
  }
}
