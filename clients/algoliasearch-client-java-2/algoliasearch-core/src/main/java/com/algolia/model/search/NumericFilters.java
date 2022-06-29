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

  public static NumericFilters ofListOfListOfString(List<List<String>> inside) {
    return new NumericFiltersListOfListOfString(inside);
  }

  public static NumericFilters ofListOfString(List<String> inside) {
    return new NumericFiltersListOfString(inside);
  }

  public static class Adapter extends TypeAdapter<NumericFilters> {

    @Override
    public void write(final JsonWriter out, final NumericFilters oneOf) throws IOException {
      TypeAdapter runtimeTypeAdapter = (TypeAdapter) JSON.getGson().getAdapter(TypeToken.get(oneOf.getInsideValue().getClass()));
      runtimeTypeAdapter.write(out, oneOf.getInsideValue());
    }

    @Override
    public NumericFilters read(final JsonReader jsonReader) throws IOException {
      List<List<String>> listoflistofstring = JSON.tryDeserialize(jsonReader, new TypeToken<List<List<String>>>() {}.getType());
      if (listoflistofstring != null) {
        return NumericFilters.ofListOfListOfString(listoflistofstring);
      }
      List<String> listofstring = JSON.tryDeserialize(jsonReader, new TypeToken<List<String>>() {}.getType());
      if (listofstring != null) {
        return NumericFilters.ofListOfString(listofstring);
      }
      return null;
    }
  }
}

@JsonAdapter(NumericFilters.Adapter.class)
class NumericFiltersListOfListOfString extends NumericFilters {

  private final List<List<String>> insideValue;

  NumericFiltersListOfListOfString(List<List<String>> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<List<String>> getInsideValue() {
    return insideValue;
  }
}

@JsonAdapter(NumericFilters.Adapter.class)
class NumericFiltersListOfString extends NumericFilters {

  private final List<String> insideValue;

  NumericFiltersListOfString(List<String> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<String> getInsideValue() {
    return insideValue;
  }
}
