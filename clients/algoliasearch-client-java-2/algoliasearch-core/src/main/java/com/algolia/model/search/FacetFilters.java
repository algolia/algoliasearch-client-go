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

@JsonAdapter(FacetFilters.Adapter.class)
/** Filter hits by facet value. */
public abstract class FacetFilters implements CompoundType {

  public static FacetFilters ofListOfListOfString(List<List<String>> inside) {
    return new FacetFiltersListOfListOfString(inside);
  }

  public static FacetFilters ofListOfString(List<String> inside) {
    return new FacetFiltersListOfString(inside);
  }

  public static class Adapter extends TypeAdapter<FacetFilters> {

    @Override
    public void write(final JsonWriter out, final FacetFilters oneOf) throws IOException {
      TypeAdapter runtimeTypeAdapter = (TypeAdapter) JSON.getGson().getAdapter(TypeToken.get(oneOf.getInsideValue().getClass()));
      runtimeTypeAdapter.write(out, oneOf.getInsideValue());
    }

    @Override
    public FacetFilters read(final JsonReader jsonReader) throws IOException {
      List<List<String>> listoflistofstring = JSON.tryDeserialize(jsonReader, new TypeToken<List<List<String>>>() {}.getType());
      if (listoflistofstring != null) {
        return FacetFilters.ofListOfListOfString(listoflistofstring);
      }
      List<String> listofstring = JSON.tryDeserialize(jsonReader, new TypeToken<List<String>>() {}.getType());
      if (listofstring != null) {
        return FacetFilters.ofListOfString(listofstring);
      }
      return null;
    }
  }
}

@JsonAdapter(FacetFilters.Adapter.class)
class FacetFiltersListOfListOfString extends FacetFilters {

  private final List<List<String>> insideValue;

  FacetFiltersListOfListOfString(List<List<String>> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<List<String>> getInsideValue() {
    return insideValue;
  }
}

@JsonAdapter(FacetFilters.Adapter.class)
class FacetFiltersListOfString extends FacetFilters {

  private final List<String> insideValue;

  FacetFiltersListOfString(List<String> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<String> getInsideValue() {
    return insideValue;
  }
}
