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

  public static ReRankingApplyFilter of(List<MixedSearchFilters> inside) {
    return new ReRankingApplyFilterListOfMixedSearchFilters(inside);
  }

  public static ReRankingApplyFilter of(String inside) {
    return new ReRankingApplyFilterString(inside);
  }

  public static class Adapter extends TypeAdapter<ReRankingApplyFilter> {

    @Override
    public void write(final JsonWriter out, final ReRankingApplyFilter oneOf) throws IOException {
      TypeAdapter runtimeTypeAdapter = (TypeAdapter) JSON.getGson().getAdapter(TypeToken.get(oneOf.getInsideValue().getClass()));
      runtimeTypeAdapter.write(out, oneOf.getInsideValue());
    }

    @Override
    public ReRankingApplyFilter read(final JsonReader jsonReader) throws IOException {
      List<MixedSearchFilters> listofmixedsearchfilters = JSON.tryDeserialize(
        jsonReader,
        new TypeToken<List<MixedSearchFilters>>() {}.getType()
      );
      if (listofmixedsearchfilters != null) {
        return ReRankingApplyFilter.of(listofmixedsearchfilters);
      }
      String string = JSON.tryDeserialize(jsonReader, new TypeToken<String>() {}.getType());
      if (string != null) {
        return ReRankingApplyFilter.of(string);
      }
      return null;
    }
  }
}

@JsonAdapter(ReRankingApplyFilter.Adapter.class)
class ReRankingApplyFilterListOfMixedSearchFilters extends ReRankingApplyFilter {

  private final List<MixedSearchFilters> insideValue;

  ReRankingApplyFilterListOfMixedSearchFilters(List<MixedSearchFilters> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<MixedSearchFilters> getInsideValue() {
    return insideValue;
  }
}

@JsonAdapter(ReRankingApplyFilter.Adapter.class)
class ReRankingApplyFilterString extends ReRankingApplyFilter {

  private final String insideValue;

  ReRankingApplyFilterString(String insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public String getInsideValue() {
    return insideValue;
  }
}
