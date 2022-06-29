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

  public static OptionalFilters ofListOfListOfString(List<List<String>> inside) {
    return new OptionalFiltersListOfListOfString(inside);
  }

  public static OptionalFilters ofListOfString(List<String> inside) {
    return new OptionalFiltersListOfString(inside);
  }

  public static class Adapter extends TypeAdapter<OptionalFilters> {

    @Override
    public void write(final JsonWriter out, final OptionalFilters oneOf) throws IOException {
      TypeAdapter runtimeTypeAdapter = (TypeAdapter) JSON.getGson().getAdapter(TypeToken.get(oneOf.getInsideValue().getClass()));
      runtimeTypeAdapter.write(out, oneOf.getInsideValue());
    }

    @Override
    public OptionalFilters read(final JsonReader jsonReader) throws IOException {
      List<List<String>> listoflistofstring = JSON.tryDeserialize(jsonReader, new TypeToken<List<List<String>>>() {}.getType());
      if (listoflistofstring != null) {
        return OptionalFilters.ofListOfListOfString(listoflistofstring);
      }
      List<String> listofstring = JSON.tryDeserialize(jsonReader, new TypeToken<List<String>>() {}.getType());
      if (listofstring != null) {
        return OptionalFilters.ofListOfString(listofstring);
      }
      return null;
    }
  }
}

@JsonAdapter(OptionalFilters.Adapter.class)
class OptionalFiltersListOfListOfString extends OptionalFilters {

  private final List<List<String>> insideValue;

  OptionalFiltersListOfListOfString(List<List<String>> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<List<String>> getInsideValue() {
    return insideValue;
  }
}

@JsonAdapter(OptionalFilters.Adapter.class)
class OptionalFiltersListOfString extends OptionalFilters {

  private final List<String> insideValue;

  OptionalFiltersListOfString(List<String> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<String> getInsideValue() {
    return insideValue;
  }
}
