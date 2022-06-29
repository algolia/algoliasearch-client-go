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
/**
 * Names of facets to which automatic filtering must be applied; they must match the facet name of a
 * facet value placeholder in the query pattern.
 */
public abstract class AutomaticFacetFilters implements CompoundType {

  public static AutomaticFacetFilters ofListOfAutomaticFacetFilter(List<AutomaticFacetFilter> inside) {
    return new AutomaticFacetFiltersListOfAutomaticFacetFilter(inside);
  }

  public static AutomaticFacetFilters ofListOfString(List<String> inside) {
    return new AutomaticFacetFiltersListOfString(inside);
  }

  public static class Adapter extends TypeAdapter<AutomaticFacetFilters> {

    @Override
    public void write(final JsonWriter out, final AutomaticFacetFilters oneOf) throws IOException {
      TypeAdapter runtimeTypeAdapter = (TypeAdapter) JSON.getGson().getAdapter(TypeToken.get(oneOf.getInsideValue().getClass()));
      runtimeTypeAdapter.write(out, oneOf.getInsideValue());
    }

    @Override
    public AutomaticFacetFilters read(final JsonReader jsonReader) throws IOException {
      List<AutomaticFacetFilter> listofautomaticfacetfilter = JSON.tryDeserialize(
        jsonReader,
        new TypeToken<List<AutomaticFacetFilter>>() {}.getType()
      );
      if (listofautomaticfacetfilter != null) {
        return AutomaticFacetFilters.ofListOfAutomaticFacetFilter(listofautomaticfacetfilter);
      }
      List<String> listofstring = JSON.tryDeserialize(jsonReader, new TypeToken<List<String>>() {}.getType());
      if (listofstring != null) {
        return AutomaticFacetFilters.ofListOfString(listofstring);
      }
      return null;
    }
  }
}

@JsonAdapter(AutomaticFacetFilters.Adapter.class)
class AutomaticFacetFiltersListOfAutomaticFacetFilter extends AutomaticFacetFilters {

  private final List<AutomaticFacetFilter> insideValue;

  AutomaticFacetFiltersListOfAutomaticFacetFilter(List<AutomaticFacetFilter> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<AutomaticFacetFilter> getInsideValue() {
    return insideValue;
  }
}

@JsonAdapter(AutomaticFacetFilters.Adapter.class)
class AutomaticFacetFiltersListOfString extends AutomaticFacetFilters {

  private final List<String> insideValue;

  AutomaticFacetFiltersListOfString(List<String> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<String> getInsideValue() {
    return insideValue;
  }
}
