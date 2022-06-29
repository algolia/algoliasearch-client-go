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

@JsonAdapter(RemoveStopWords.Adapter.class)
/**
 * Removes stop (common) words from the query before executing it. removeStopWords is used in
 * conjunction with the queryLanguages setting. list: language ISO codes for which ignoring plurals
 * should be enabled. This list will override any values that you may have set in queryLanguages.
 * true: enables the stop word functionality, ensuring that stop words are removed from
 * consideration in a search. The languages supported here are either every language, or those set
 * by queryLanguages. false: disables stop word functionality, allowing stop words to be taken into
 * account in a search.
 */
public abstract class RemoveStopWords implements CompoundType {

  public static RemoveStopWords of(Boolean inside) {
    return new RemoveStopWordsBoolean(inside);
  }

  public static RemoveStopWords of(List<String> inside) {
    return new RemoveStopWordsListOfString(inside);
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
        return RemoveStopWords.of(_boolean);
      }
      List<String> listofstring = JSON.tryDeserialize(jsonReader, new TypeToken<List<String>>() {}.getType());
      if (listofstring != null) {
        return RemoveStopWords.of(listofstring);
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
class RemoveStopWordsListOfString extends RemoveStopWords {

  private final List<String> insideValue;

  RemoveStopWordsListOfString(List<String> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<String> getInsideValue() {
    return insideValue;
  }
}
