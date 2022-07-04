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

@JsonAdapter(FacetFilters.Adapter.class)
/** Filter hits by facet value. */
public abstract class FacetFilters implements CompoundType {

  public static FacetFilters of(List<MixedSearchFilters> inside) {
    return new FacetFiltersListOfMixedSearchFilters(inside);
  }

  public static FacetFilters of(String inside) {
    return new FacetFiltersString(inside);
  }

  public static class Adapter extends TypeAdapter<FacetFilters> {

    @Override
    public void write(final JsonWriter out, final FacetFilters oneOf) throws IOException {
      TypeAdapter runtimeTypeAdapter = (TypeAdapter) JSON.getGson().getAdapter(TypeToken.get(oneOf.getInsideValue().getClass()));
      runtimeTypeAdapter.write(out, oneOf.getInsideValue());
    }

    @Override
    public FacetFilters read(final JsonReader jsonReader) throws IOException {
      List<MixedSearchFilters> listofmixedsearchfilters = JSON.tryDeserialize(
        jsonReader,
        new TypeToken<List<MixedSearchFilters>>() {}.getType()
      );
      if (listofmixedsearchfilters != null) {
        return FacetFilters.of(listofmixedsearchfilters);
      }
      String string = JSON.tryDeserialize(jsonReader, new TypeToken<String>() {}.getType());
      if (string != null) {
        return FacetFilters.of(string);
      }
      return null;
    }
  }
}

@JsonAdapter(FacetFilters.Adapter.class)
class FacetFiltersListOfMixedSearchFilters extends FacetFilters {

  private final List<MixedSearchFilters> insideValue;

  FacetFiltersListOfMixedSearchFilters(List<MixedSearchFilters> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<MixedSearchFilters> getInsideValue() {
    return insideValue;
  }
}

@JsonAdapter(FacetFilters.Adapter.class)
class FacetFiltersString extends FacetFilters {

  private final String insideValue;

  FacetFiltersString(String insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public String getInsideValue() {
    return insideValue;
  }
}
