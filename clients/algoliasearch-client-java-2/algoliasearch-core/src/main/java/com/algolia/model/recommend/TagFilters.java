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

@JsonAdapter(TagFilters.Adapter.class)
/** Filter hits by tags. */
public abstract class TagFilters implements CompoundType {

  public static TagFilters of(List<MixedSearchFilters> inside) {
    return new TagFiltersListOfMixedSearchFilters(inside);
  }

  public static TagFilters of(String inside) {
    return new TagFiltersString(inside);
  }

  public static class Adapter extends TypeAdapter<TagFilters> {

    @Override
    public void write(final JsonWriter out, final TagFilters oneOf) throws IOException {
      TypeAdapter runtimeTypeAdapter = (TypeAdapter) JSON.getGson().getAdapter(TypeToken.get(oneOf.getInsideValue().getClass()));
      runtimeTypeAdapter.write(out, oneOf.getInsideValue());
    }

    @Override
    public TagFilters read(final JsonReader jsonReader) throws IOException {
      List<MixedSearchFilters> listofmixedsearchfilters = JSON.tryDeserialize(
        jsonReader,
        new TypeToken<List<MixedSearchFilters>>() {}.getType()
      );
      if (listofmixedsearchfilters != null) {
        return TagFilters.of(listofmixedsearchfilters);
      }
      String string = JSON.tryDeserialize(jsonReader, new TypeToken<String>() {}.getType());
      if (string != null) {
        return TagFilters.of(string);
      }
      return null;
    }
  }
}

@JsonAdapter(TagFilters.Adapter.class)
class TagFiltersListOfMixedSearchFilters extends TagFilters {

  private final List<MixedSearchFilters> insideValue;

  TagFiltersListOfMixedSearchFilters(List<MixedSearchFilters> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<MixedSearchFilters> getInsideValue() {
    return insideValue;
  }
}

@JsonAdapter(TagFilters.Adapter.class)
class TagFiltersString extends TagFilters {

  private final String insideValue;

  TagFiltersString(String insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public String getInsideValue() {
    return insideValue;
  }
}
