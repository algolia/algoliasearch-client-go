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
public abstract class TagFilters implements CompoundType {

  public static TagFilters ofListListString(List<List<String>> inside) {
    return new TagFiltersListListString(inside);
  }

  public static TagFilters ofListString(List<String> inside) {
    return new TagFiltersListString(inside);
  }

  public static class Adapter extends TypeAdapter<TagFilters> {

    @Override
    public void write(final JsonWriter out, final TagFilters oneOf)
      throws IOException {
      TypeAdapter runtimeTypeAdapter = (TypeAdapter) JSON
        .getGson()
        .getAdapter(TypeToken.get(oneOf.getInsideValue().getClass()));
      runtimeTypeAdapter.write(out, oneOf.getInsideValue());
    }

    @Override
    public TagFilters read(final JsonReader jsonReader) throws IOException {
      List<List<String>> listliststring = JSON.tryDeserialize(
        jsonReader,
        new TypeToken<List<List<String>>>() {}.getType()
      );
      if (listliststring != null) {
        return TagFilters.ofListListString(listliststring);
      }
      List<String> liststring = JSON.tryDeserialize(
        jsonReader,
        new TypeToken<List<String>>() {}.getType()
      );
      if (liststring != null) {
        return TagFilters.ofListString(liststring);
      }
      return null;
    }
  }
}

@JsonAdapter(TagFilters.Adapter.class)
class TagFiltersListListString extends TagFilters {

  private final List<List<String>> insideValue;

  TagFiltersListListString(List<List<String>> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<List<String>> getInsideValue() {
    return insideValue;
  }
}

@JsonAdapter(TagFilters.Adapter.class)
class TagFiltersListString extends TagFilters {

  private final List<String> insideValue;

  TagFiltersListString(List<String> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<String> getInsideValue() {
    return insideValue;
  }
}
