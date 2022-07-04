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

@JsonAdapter(NumericFilters.Adapter.class)
/** Filter on numeric attributes. */
public abstract class NumericFilters implements CompoundType {

  public static NumericFilters of(List<MixedSearchFilters> inside) {
    return new NumericFiltersListOfMixedSearchFilters(inside);
  }

  public static NumericFilters of(String inside) {
    return new NumericFiltersString(inside);
  }

  public static class Adapter extends TypeAdapter<NumericFilters> {

    @Override
    public void write(final JsonWriter out, final NumericFilters oneOf) throws IOException {
      TypeAdapter runtimeTypeAdapter = (TypeAdapter) JSON.getGson().getAdapter(TypeToken.get(oneOf.getInsideValue().getClass()));
      runtimeTypeAdapter.write(out, oneOf.getInsideValue());
    }

    @Override
    public NumericFilters read(final JsonReader jsonReader) throws IOException {
      List<MixedSearchFilters> listofmixedsearchfilters = JSON.tryDeserialize(
        jsonReader,
        new TypeToken<List<MixedSearchFilters>>() {}.getType()
      );
      if (listofmixedsearchfilters != null) {
        return NumericFilters.of(listofmixedsearchfilters);
      }
      String string = JSON.tryDeserialize(jsonReader, new TypeToken<String>() {}.getType());
      if (string != null) {
        return NumericFilters.of(string);
      }
      return null;
    }
  }
}

@JsonAdapter(NumericFilters.Adapter.class)
class NumericFiltersListOfMixedSearchFilters extends NumericFilters {

  private final List<MixedSearchFilters> insideValue;

  NumericFiltersListOfMixedSearchFilters(List<MixedSearchFilters> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<MixedSearchFilters> getInsideValue() {
    return insideValue;
  }
}

@JsonAdapter(NumericFilters.Adapter.class)
class NumericFiltersString extends NumericFilters {

  private final String insideValue;

  NumericFiltersString(String insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public String getInsideValue() {
    return insideValue;
  }
}
