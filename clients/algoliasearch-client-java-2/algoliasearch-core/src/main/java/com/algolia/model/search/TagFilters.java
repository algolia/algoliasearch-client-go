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

@JsonAdapter(TagFilters.Adapter.class)
/** Filter hits by tags. */
public abstract class TagFilters implements CompoundType {

  public static TagFilters ofListOfListOfString(List<List<String>> inside) {
    return new TagFiltersListOfListOfString(inside);
  }

  public static TagFilters ofListOfString(List<String> inside) {
    return new TagFiltersListOfString(inside);
  }

  public static class Adapter extends TypeAdapter<TagFilters> {

    @Override
    public void write(final JsonWriter out, final TagFilters oneOf) throws IOException {
      TypeAdapter runtimeTypeAdapter = (TypeAdapter) JSON.getGson().getAdapter(TypeToken.get(oneOf.getInsideValue().getClass()));
      runtimeTypeAdapter.write(out, oneOf.getInsideValue());
    }

    @Override
    public TagFilters read(final JsonReader jsonReader) throws IOException {
      List<List<String>> listoflistofstring = JSON.tryDeserialize(jsonReader, new TypeToken<List<List<String>>>() {}.getType());
      if (listoflistofstring != null) {
        return TagFilters.ofListOfListOfString(listoflistofstring);
      }
      List<String> listofstring = JSON.tryDeserialize(jsonReader, new TypeToken<List<String>>() {}.getType());
      if (listofstring != null) {
        return TagFilters.ofListOfString(listofstring);
      }
      return null;
    }
  }
}

@JsonAdapter(TagFilters.Adapter.class)
class TagFiltersListOfListOfString extends TagFilters {

  private final List<List<String>> insideValue;

  TagFiltersListOfListOfString(List<List<String>> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<List<String>> getInsideValue() {
    return insideValue;
  }
}

@JsonAdapter(TagFilters.Adapter.class)
class TagFiltersListOfString extends TagFilters {

  private final List<String> insideValue;

  TagFiltersListOfString(List<String> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<String> getInsideValue() {
    return insideValue;
  }
}
