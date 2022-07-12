// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.analytics;

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

/** GetTopHitsResponse */
@JsonDeserialize(using = GetTopHitsResponse.GetTopHitsResponseDeserializer.class)
@JsonSerialize(using = GetTopHitsResponse.GetTopHitsResponseSerializer.class)
public abstract class GetTopHitsResponse implements CompoundType {

  public static GetTopHitsResponse of(TopHitsResponse inside) {
    return new GetTopHitsResponseTopHitsResponse(inside);
  }

  public static GetTopHitsResponse of(TopHitsResponseWithAnalytics inside) {
    return new GetTopHitsResponseTopHitsResponseWithAnalytics(inside);
  }

  public static class GetTopHitsResponseSerializer extends StdSerializer<GetTopHitsResponse> {

    public GetTopHitsResponseSerializer(Class<GetTopHitsResponse> t) {
      super(t);
    }

    public GetTopHitsResponseSerializer() {
      this(null);
    }

    @Override
    public void serialize(GetTopHitsResponse value, JsonGenerator jgen, SerializerProvider provider)
      throws IOException, JsonProcessingException {
      jgen.writeObject(value.getInsideValue());
    }
  }

  public static class GetTopHitsResponseDeserializer extends StdDeserializer<GetTopHitsResponse> {

    public GetTopHitsResponseDeserializer() {
      this(GetTopHitsResponse.class);
    }

    public GetTopHitsResponseDeserializer(Class<?> vc) {
      super(vc);
    }

    @Override
    public GetTopHitsResponse deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException, JsonProcessingException {
      JsonNode tree = jp.readValueAsTree();
      GetTopHitsResponse deserialized = null;

      int match = 0;
      JsonToken token = tree.traverse(jp.getCodec()).nextToken();
      String currentType = "";
      // deserialize TopHitsResponse
      try {
        boolean attemptParsing = true;
        currentType = "TopHitsResponse";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized =
            GetTopHitsResponse.of((TopHitsResponse) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<TopHitsResponse>() {}));
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              GetTopHitsResponse.of((TopHitsResponse) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<TopHitsResponse>() {}));
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println("Failed to deserialize oneOf TopHitsResponse (error: " + e.getMessage() + ") (type: " + currentType + ")");
      }

      // deserialize TopHitsResponseWithAnalytics
      try {
        boolean attemptParsing = true;
        currentType = "TopHitsResponseWithAnalytics";
        if (
          ((currentType.equals("Integer") || currentType.equals("Long")) && token == JsonToken.VALUE_NUMBER_INT) |
          ((currentType.equals("Float") || currentType.equals("Double")) && token == JsonToken.VALUE_NUMBER_FLOAT) |
          (currentType.equals("Boolean") && (token == JsonToken.VALUE_FALSE || token == JsonToken.VALUE_TRUE)) |
          (currentType.equals("String") && token == JsonToken.VALUE_STRING) |
          (currentType.startsWith("List<") && token == JsonToken.START_ARRAY)
        ) {
          deserialized =
            GetTopHitsResponse.of(
              (TopHitsResponseWithAnalytics) tree.traverse(jp.getCodec()).readValueAs(new TypeReference<TopHitsResponseWithAnalytics>() {})
            );
          match++;
        } else if (token == JsonToken.START_OBJECT) {
          try {
            deserialized =
              GetTopHitsResponse.of(
                (TopHitsResponseWithAnalytics) tree
                  .traverse(jp.getCodec())
                  .readValueAs(new TypeReference<TopHitsResponseWithAnalytics>() {})
              );
            match++;
          } catch (IOException e) {
            // do nothing
          }
        }
      } catch (Exception e) {
        // deserialization failed, continue
        System.err.println(
          "Failed to deserialize oneOf TopHitsResponseWithAnalytics (error: " + e.getMessage() + ") (type: " + currentType + ")"
        );
      }

      if (match == 1) {
        return deserialized;
      }
      throw new IOException(String.format("Failed deserialization for GetTopHitsResponse: %d classes match result, expected 1", match));
    }

    /** Handle deserialization of the 'null' value. */
    @Override
    public GetTopHitsResponse getNullValue(DeserializationContext ctxt) throws JsonMappingException {
      throw new JsonMappingException(ctxt.getParser(), "GetTopHitsResponse cannot be null");
    }
  }
}

class GetTopHitsResponseTopHitsResponse extends GetTopHitsResponse {

  private final TopHitsResponse insideValue;

  GetTopHitsResponseTopHitsResponse(TopHitsResponse insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public TopHitsResponse getInsideValue() {
    return insideValue;
  }
}

class GetTopHitsResponseTopHitsResponseWithAnalytics extends GetTopHitsResponse {

  private final TopHitsResponseWithAnalytics insideValue;

  GetTopHitsResponseTopHitsResponseWithAnalytics(TopHitsResponseWithAnalytics insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public TopHitsResponseWithAnalytics getInsideValue() {
    return insideValue;
  }
}
