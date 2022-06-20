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
public abstract class IgnorePlurals implements CompoundType {

  public static IgnorePlurals ofBoolean(Boolean inside) {
    return new IgnorePluralsBoolean(inside);
  }

  public static IgnorePlurals ofListString(List<String> inside) {
    return new IgnorePluralsListString(inside);
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
        return IgnorePlurals.ofBoolean(_boolean);
      }
      List<String> liststring = JSON.tryDeserialize(jsonReader, new TypeToken<List<String>>() {}.getType());
      if (liststring != null) {
        return IgnorePlurals.ofListString(liststring);
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
class IgnorePluralsListString extends IgnorePlurals {

  private final List<String> insideValue;

  IgnorePluralsListString(List<String> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<String> getInsideValue() {
    return insideValue;
  }
}
