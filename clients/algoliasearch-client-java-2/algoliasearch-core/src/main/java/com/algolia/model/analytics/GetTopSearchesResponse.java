package com.algolia.model.analytics;

import com.algolia.utils.CompoundType;
import com.algolia.utils.JSON;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.reflect.TypeToken;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;

@JsonAdapter(GetTopSearchesResponse.Adapter.class)
public abstract class GetTopSearchesResponse implements CompoundType {

  public static GetTopSearchesResponse ofTopSearchesResponse(
    TopSearchesResponse inside
  ) {
    return new GetTopSearchesResponseTopSearchesResponse(inside);
  }

  public static GetTopSearchesResponse ofTopSearchesResponseWithAnalytics(
    TopSearchesResponseWithAnalytics inside
  ) {
    return new GetTopSearchesResponseTopSearchesResponseWithAnalytics(inside);
  }

  public static class Adapter extends TypeAdapter<GetTopSearchesResponse> {

    @Override
    public void write(final JsonWriter out, final GetTopSearchesResponse oneOf)
      throws IOException {
      TypeAdapter runtimeTypeAdapter = (TypeAdapter) JSON
        .getGson()
        .getAdapter(TypeToken.get(oneOf.getInsideValue().getClass()));
      runtimeTypeAdapter.write(out, oneOf.getInsideValue());
    }

    @Override
    public GetTopSearchesResponse read(final JsonReader jsonReader)
      throws IOException {
      TopSearchesResponse topsearchesresponse = JSON.tryDeserialize(
        jsonReader,
        new TypeToken<TopSearchesResponse>() {}.getType()
      );
      if (topsearchesresponse != null) {
        return GetTopSearchesResponse.ofTopSearchesResponse(
          topsearchesresponse
        );
      }
      TopSearchesResponseWithAnalytics topsearchesresponsewithanalytics = JSON.tryDeserialize(
        jsonReader,
        new TypeToken<TopSearchesResponseWithAnalytics>() {}.getType()
      );
      if (topsearchesresponsewithanalytics != null) {
        return GetTopSearchesResponse.ofTopSearchesResponseWithAnalytics(
          topsearchesresponsewithanalytics
        );
      }
      return null;
    }
  }
}

@JsonAdapter(GetTopSearchesResponse.Adapter.class)
class GetTopSearchesResponseTopSearchesResponse extends GetTopSearchesResponse {

  private final TopSearchesResponse insideValue;

  GetTopSearchesResponseTopSearchesResponse(TopSearchesResponse insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public TopSearchesResponse getInsideValue() {
    return insideValue;
  }
}

@JsonAdapter(GetTopSearchesResponse.Adapter.class)
class GetTopSearchesResponseTopSearchesResponseWithAnalytics
  extends GetTopSearchesResponse {

  private final TopSearchesResponseWithAnalytics insideValue;

  GetTopSearchesResponseTopSearchesResponseWithAnalytics(
    TopSearchesResponseWithAnalytics insideValue
  ) {
    this.insideValue = insideValue;
  }

  @Override
  public TopSearchesResponseWithAnalytics getInsideValue() {
    return insideValue;
  }
}
