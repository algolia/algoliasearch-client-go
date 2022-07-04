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

@JsonAdapter(MixedSearchFilters.Adapter.class)
/** MixedSearchFilters */
public abstract class MixedSearchFilters implements CompoundType {

  public static MixedSearchFilters of(List<String> inside) {
    return new MixedSearchFiltersListOfString(inside);
  }

  public static MixedSearchFilters of(String inside) {
    return new MixedSearchFiltersString(inside);
  }

  public static class Adapter extends TypeAdapter<MixedSearchFilters> {

    @Override
    public void write(final JsonWriter out, final MixedSearchFilters oneOf) throws IOException {
      TypeAdapter runtimeTypeAdapter = (TypeAdapter) JSON.getGson().getAdapter(TypeToken.get(oneOf.getInsideValue().getClass()));
      runtimeTypeAdapter.write(out, oneOf.getInsideValue());
    }

    @Override
    public MixedSearchFilters read(final JsonReader jsonReader) throws IOException {
      List<String> listofstring = JSON.tryDeserialize(jsonReader, new TypeToken<List<String>>() {}.getType());
      if (listofstring != null) {
        return MixedSearchFilters.of(listofstring);
      }
      String string = JSON.tryDeserialize(jsonReader, new TypeToken<String>() {}.getType());
      if (string != null) {
        return MixedSearchFilters.of(string);
      }
      return null;
    }
  }
}

@JsonAdapter(MixedSearchFilters.Adapter.class)
class MixedSearchFiltersListOfString extends MixedSearchFilters {

  private final List<String> insideValue;

  MixedSearchFiltersListOfString(List<String> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<String> getInsideValue() {
    return insideValue;
  }
}

@JsonAdapter(MixedSearchFilters.Adapter.class)
class MixedSearchFiltersString extends MixedSearchFilters {

  private final String insideValue;

  MixedSearchFiltersString(String insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public String getInsideValue() {
    return insideValue;
  }
}
