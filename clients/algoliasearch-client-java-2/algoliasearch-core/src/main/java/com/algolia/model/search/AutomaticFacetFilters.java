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

@JsonAdapter(AutomaticFacetFilters.Adapter.class)
public abstract class AutomaticFacetFilters implements CompoundType {

  public static AutomaticFacetFilters ofListAutomaticFacetFilter(List<AutomaticFacetFilter> inside) {
    return new AutomaticFacetFiltersListAutomaticFacetFilter(inside);
  }

  public static AutomaticFacetFilters ofListString(List<String> inside) {
    return new AutomaticFacetFiltersListString(inside);
  }

  public static class Adapter extends TypeAdapter<AutomaticFacetFilters> {

    @Override
    public void write(final JsonWriter out, final AutomaticFacetFilters oneOf) throws IOException {
      TypeAdapter runtimeTypeAdapter = (TypeAdapter) JSON.getGson().getAdapter(TypeToken.get(oneOf.getInsideValue().getClass()));
      runtimeTypeAdapter.write(out, oneOf.getInsideValue());
    }

    @Override
    public AutomaticFacetFilters read(final JsonReader jsonReader) throws IOException {
      List<AutomaticFacetFilter> listautomaticfacetfilter = JSON.tryDeserialize(
        jsonReader,
        new TypeToken<List<AutomaticFacetFilter>>() {}.getType()
      );
      if (listautomaticfacetfilter != null) {
        return AutomaticFacetFilters.ofListAutomaticFacetFilter(listautomaticfacetfilter);
      }
      List<String> liststring = JSON.tryDeserialize(jsonReader, new TypeToken<List<String>>() {}.getType());
      if (liststring != null) {
        return AutomaticFacetFilters.ofListString(liststring);
      }
      return null;
    }
  }
}

@JsonAdapter(AutomaticFacetFilters.Adapter.class)
class AutomaticFacetFiltersListAutomaticFacetFilter extends AutomaticFacetFilters {

  private final List<AutomaticFacetFilter> insideValue;

  AutomaticFacetFiltersListAutomaticFacetFilter(List<AutomaticFacetFilter> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<AutomaticFacetFilter> getInsideValue() {
    return insideValue;
  }
}

@JsonAdapter(AutomaticFacetFilters.Adapter.class)
class AutomaticFacetFiltersListString extends AutomaticFacetFilters {

  private final List<String> insideValue;

  AutomaticFacetFiltersListString(List<String> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<String> getInsideValue() {
    return insideValue;
  }
}
