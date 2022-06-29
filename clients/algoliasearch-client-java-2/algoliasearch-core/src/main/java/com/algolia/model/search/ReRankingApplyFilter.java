package com.algolia.model.search;

import com.algolia.utils.CompoundType;
import com.algolia.utils.JSON;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.reflect.TypeToken;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;
import java.util.List;

@JsonAdapter(ReRankingApplyFilter.Adapter.class)
/**
 * When Dynamic Re-Ranking is enabled, only records that match these filters will be impacted by
 * Dynamic Re-Ranking.
 */
public abstract class ReRankingApplyFilter implements CompoundType {

  public static ReRankingApplyFilter ofListOfListOfString(List<List<String>> inside) {
    return new ReRankingApplyFilterListOfListOfString(inside);
  }

  public static ReRankingApplyFilter ofListOfString(List<String> inside) {
    return new ReRankingApplyFilterListOfString(inside);
  }

  public static class Adapter extends TypeAdapter<ReRankingApplyFilter> {

    @Override
    public void write(final JsonWriter out, final ReRankingApplyFilter oneOf) throws IOException {
      TypeAdapter runtimeTypeAdapter = (TypeAdapter) JSON.getGson().getAdapter(TypeToken.get(oneOf.getInsideValue().getClass()));
      runtimeTypeAdapter.write(out, oneOf.getInsideValue());
    }

    @Override
    public ReRankingApplyFilter read(final JsonReader jsonReader) throws IOException {
      List<List<String>> listoflistofstring = JSON.tryDeserialize(jsonReader, new TypeToken<List<List<String>>>() {}.getType());
      if (listoflistofstring != null) {
        return ReRankingApplyFilter.ofListOfListOfString(listoflistofstring);
      }
      List<String> listofstring = JSON.tryDeserialize(jsonReader, new TypeToken<List<String>>() {}.getType());
      if (listofstring != null) {
        return ReRankingApplyFilter.ofListOfString(listofstring);
      }
      return null;
    }
  }
}

@JsonAdapter(ReRankingApplyFilter.Adapter.class)
class ReRankingApplyFilterListOfListOfString extends ReRankingApplyFilter {

  private final List<List<String>> insideValue;

  ReRankingApplyFilterListOfListOfString(List<List<String>> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<List<String>> getInsideValue() {
    return insideValue;
  }
}

@JsonAdapter(ReRankingApplyFilter.Adapter.class)
class ReRankingApplyFilterListOfString extends ReRankingApplyFilter {

  private final List<String> insideValue;

  ReRankingApplyFilterListOfString(List<String> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<String> getInsideValue() {
    return insideValue;
  }
}
