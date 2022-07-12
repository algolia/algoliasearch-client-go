// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

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

/** RecommendationsRequest */
@JsonDeserialize(using = RecommendationsRequest.RecommendationsRequestDeserializer.class)
@JsonSerialize(using = RecommendationsRequest.RecommendationsRequestSerializer.class)
public abstract class RecommendationsRequest implements CompoundType {

  public static RecommendationsRequest of(RecommendationRequest inside) {
    return new RecommendationsRequestRecommendationRequest(inside);
  }

  public static RecommendationsRequest of(TrendingRequest inside) {
    return new RecommendationsRequestTrendingRequest(inside);
  }

  public static class RecommendationsRequestSerializer extends StdSerializer<RecommendationsRequest> {

    public RecommendationsRequestSerializer(Class<RecommendationsRequest> t) {
      super(t);
    }

    public RecommendationsRequestSerializer() {
      this(null);
    }

    @Override
    public void serialize(RecommendationsRequest value, JsonGenerator jgen, SerializerProvider provider)
      throws IOException, JsonProcessingException {
      jgen.writeObject(value.getInsideValue());
    }
  }

  public static class RecommendationsRequestDeserializer extends StdDeserializer<RecommendationsRequest> {

    public RecommendationsRequestDeserializer() {
      this(RecommendationsRequest.class);
    }

    public RecommendationsRequestDeserializer(Class<?> vc) {
      super(vc);
    }

    @Override
    public RecommendationsRequest deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException, JsonProcessingException {
      JsonNode tree = jp.readValueAsTree();
      RecommendationsRequest deserialized = null;

      int match = 0;
      JsonToken token = tree.traverse(jp.getCodec()).nextToken();
      String currentType = "";
      // deserialize RecommendationRequest
      try {
        boolean attemptParsing = true;
        currentType = "RecommendationRequest";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized =
            RecommendationsRequest.of(
              (RecommendationRequest) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<RecommendationRequest>() {})
            );
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              RecommendationsRequest.of(
                (RecommendationRequest) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<RecommendationRequest>() {})
              );
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println("Failed to deserialize oneOf RecommendationRequest (error: " + e.getMessage() + ") (type: " + currentType + ")");
      }

      // deserialize TrendingRequest
      try {
        boolean attemptParsing = true;
        currentType = "TrendingRequest";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized =
            RecommendationsRequest.of((TrendingRequest) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<TrendingRequest>() {}));
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              RecommendationsRequest.of(
                (TrendingRequest) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<TrendingRequest>() {})
              );
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println("Failed to deserialize oneOf TrendingRequest (error: " + e.getMessage() + ") (type: " + currentType + ")");
      }

      if (match == 1) {
        return deserialized;
      }
      throw new IOException(
        String.format("Failed deserialization for RecommendationsRequest: %d classes match result, expected" + " 1", match)
      );
    }

    /** Handle deserialization of the 'null' value. */
    @Override
    public RecommendationsRequest getNullValue(DeserializationContext ctxt) throws JsonMappingException {
      throw new JsonMappingException(ctxt.getParser(), "RecommendationsRequest cannot be null");
    }
  }
}

class RecommendationsRequestRecommendationRequest extends RecommendationsRequest {

  private final RecommendationRequest insideValue;

  RecommendationsRequestRecommendationRequest(RecommendationRequest insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public RecommendationRequest getInsideValue() {
    return insideValue;
  }
}

class RecommendationsRequestTrendingRequest extends RecommendationsRequest {

  private final TrendingRequest insideValue;

  RecommendationsRequestTrendingRequest(TrendingRequest insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public TrendingRequest getInsideValue() {
    return insideValue;
  }
}
