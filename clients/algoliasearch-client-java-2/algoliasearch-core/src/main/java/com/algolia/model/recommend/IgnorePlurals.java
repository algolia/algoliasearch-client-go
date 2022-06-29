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

@JsonAdapter(IgnorePlurals.Adapter.class)
/**
 * Treats singular, plurals, and other forms of declensions as matching terms. ignorePlurals is used
 * in conjunction with the queryLanguages setting. list: language ISO codes for which ignoring
 * plurals should be enabled. This list will override any values that you may have set in
 * queryLanguages. true: enables the ignore plurals functionality, where singulars and plurals are
 * considered equivalent (foot = feet). The languages supported here are either every language (this
 * is the default, see list of languages below), or those set by queryLanguages. false: disables
 * ignore plurals, where singulars and plurals are not considered the same for matching purposes
 * (foot will not find feet).
 */
public abstract class IgnorePlurals implements CompoundType {

  public static IgnorePlurals of(Boolean inside) {
    return new IgnorePluralsBoolean(inside);
  }

  public static IgnorePlurals of(List<String> inside) {
    return new IgnorePluralsListOfString(inside);
  }

  public static class Adapter extends TypeAdapter<IgnorePlurals> {

    @Override
    public void write(final JsonWriter out, final IgnorePlurals oneOf) throws IOException {
      TypeAdapter runtimeTypeAdapter = (TypeAdapter) JSON.getGson().getAdapter(TypeToken.get(oneOf.getInsideValue().getClass()));
      runtimeTypeAdapter.write(out, oneOf.getInsideValue());
    }

    @Override
    public IgnorePlurals read(final JsonReader jsonReader) throws IOException {
      Boolean _boolean = JSON.tryDeserialize(jsonReader, new TypeToken<Boolean>() {}.getType());
      if (_boolean != null) {
        return IgnorePlurals.of(_boolean);
      }
      List<String> listofstring = JSON.tryDeserialize(jsonReader, new TypeToken<List<String>>() {}.getType());
      if (listofstring != null) {
        return IgnorePlurals.of(listofstring);
      }
      return null;
    }
  }
}

@JsonAdapter(IgnorePlurals.Adapter.class)
class IgnorePluralsBoolean extends IgnorePlurals {

  private final Boolean insideValue;

  IgnorePluralsBoolean(Boolean insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public Boolean getInsideValue() {
    return insideValue;
  }
}

@JsonAdapter(IgnorePlurals.Adapter.class)
class IgnorePluralsListOfString extends IgnorePlurals {

  private final List<String> insideValue;

  IgnorePluralsListOfString(List<String> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<String> getInsideValue() {
    return insideValue;
  }
}
