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

@JsonAdapter(RemoveStopWords.Adapter.class)
public abstract class RemoveStopWords implements CompoundType {

  public static RemoveStopWords ofBoolean(Boolean inside) {
    return new RemoveStopWordsBoolean(inside);
  }

  public static RemoveStopWords ofListString(List<String> inside) {
    return new RemoveStopWordsListString(inside);
  }

  public static class Adapter extends TypeAdapter<RemoveStopWords> {

    @Override
    public void write(final JsonWriter out, final RemoveStopWords oneOf) throws IOException {
      TypeAdapter runtimeTypeAdapter = (TypeAdapter) JSON.getGson().getAdapter(TypeToken.get(oneOf.getInsideValue().getClass()));
      runtimeTypeAdapter.write(out, oneOf.getInsideValue());
    }

    @Override
    public RemoveStopWords read(final JsonReader jsonReader) throws IOException {
      Boolean _boolean = JSON.tryDeserialize(jsonReader, new TypeToken<Boolean>() {}.getType());
      if (_boolean != null) {
        return RemoveStopWords.ofBoolean(_boolean);
      }
      List<String> liststring = JSON.tryDeserialize(jsonReader, new TypeToken<List<String>>() {}.getType());
      if (liststring != null) {
        return RemoveStopWords.ofListString(liststring);
      }
      return null;
    }
  }
}

@JsonAdapter(RemoveStopWords.Adapter.class)
class RemoveStopWordsBoolean extends RemoveStopWords {

  private final Boolean insideValue;

  RemoveStopWordsBoolean(Boolean insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public Boolean getInsideValue() {
    return insideValue;
  }
}

@JsonAdapter(RemoveStopWords.Adapter.class)
class RemoveStopWordsListString extends RemoveStopWords {

  private final List<String> insideValue;

  RemoveStopWordsListString(List<String> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<String> getInsideValue() {
    return insideValue;
  }
}
