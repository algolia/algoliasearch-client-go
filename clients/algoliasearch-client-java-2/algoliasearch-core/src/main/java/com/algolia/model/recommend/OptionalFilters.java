package com.algolia.model.recommend;

import com.algolia.utils.CompoundType;
import com.algolia.utils.JSON;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.reflect.TypeToken;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;
import java.util.List;

@JsonAdapter(OptionalFilters.Adapter.class)
/**
 * Create filters for ranking purposes, where records that match the filter are ranked higher, or
 * lower in the case of a negative optional filter.
 */
public abstract class OptionalFilters implements CompoundType {

  public static OptionalFilters of(List<MixedSearchFilters> inside) {
    return new OptionalFiltersListOfMixedSearchFilters(inside);
  }

  public static OptionalFilters of(String inside) {
    return new OptionalFiltersString(inside);
  }

  public static class Adapter extends TypeAdapter<OptionalFilters> {

    @Override
    public void write(final JsonWriter out, final OptionalFilters oneOf) throws IOException {
      TypeAdapter runtimeTypeAdapter = (TypeAdapter) JSON.getGson().getAdapter(TypeToken.get(oneOf.getInsideValue().getClass()));
      runtimeTypeAdapter.write(out, oneOf.getInsideValue());
    }

    @Override
    public OptionalFilters read(final JsonReader jsonReader) throws IOException {
      List<MixedSearchFilters> listofmixedsearchfilters = JSON.tryDeserialize(
        jsonReader,
        new TypeToken<List<MixedSearchFilters>>() {}.getType()
      );
      if (listofmixedsearchfilters != null) {
        return OptionalFilters.of(listofmixedsearchfilters);
      }
      String string = JSON.tryDeserialize(jsonReader, new TypeToken<String>() {}.getType());
      if (string != null) {
        return OptionalFilters.of(string);
      }
      return null;
    }
  }
}

@JsonAdapter(OptionalFilters.Adapter.class)
class OptionalFiltersListOfMixedSearchFilters extends OptionalFilters {

  private final List<MixedSearchFilters> insideValue;

  OptionalFiltersListOfMixedSearchFilters(List<MixedSearchFilters> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<MixedSearchFilters> getInsideValue() {
    return insideValue;
  }
}

@JsonAdapter(OptionalFilters.Adapter.class)
class OptionalFiltersString extends OptionalFilters {

  private final String insideValue;

  OptionalFiltersString(String insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public String getInsideValue() {
    return insideValue;
  }
}
