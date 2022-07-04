package com.algolia.model.search;

import com.algolia.utils.CompoundType;
import com.algolia.utils.JSON;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.reflect.TypeToken;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;

@JsonAdapter(ConsequenceQuery.Adapter.class)
/**
 * When providing a string, it replaces the entire query string. When providing an object, it
 * describes incremental edits to be made to the query string (but you can't do both).
 */
public abstract class ConsequenceQuery implements CompoundType {

  public static ConsequenceQuery of(ConsequenceQueryObject inside) {
    return new ConsequenceQueryConsequenceQueryObject(inside);
  }

  public static ConsequenceQuery of(String inside) {
    return new ConsequenceQueryString(inside);
  }

  public static class Adapter extends TypeAdapter<ConsequenceQuery> {

    @Override
    public void write(final JsonWriter out, final ConsequenceQuery oneOf) throws IOException {
      TypeAdapter runtimeTypeAdapter = (TypeAdapter) JSON.getGson().getAdapter(TypeToken.get(oneOf.getInsideValue().getClass()));
      runtimeTypeAdapter.write(out, oneOf.getInsideValue());
    }

    @Override
    public ConsequenceQuery read(final JsonReader jsonReader) throws IOException {
      ConsequenceQueryObject consequencequeryobject = JSON.tryDeserialize(jsonReader, new TypeToken<ConsequenceQueryObject>() {}.getType());
      if (consequencequeryobject != null) {
        return ConsequenceQuery.of(consequencequeryobject);
      }
      String string = JSON.tryDeserialize(jsonReader, new TypeToken<String>() {}.getType());
      if (string != null) {
        return ConsequenceQuery.of(string);
      }
      return null;
    }
  }
}

@JsonAdapter(ConsequenceQuery.Adapter.class)
class ConsequenceQueryConsequenceQueryObject extends ConsequenceQuery {

  private final ConsequenceQueryObject insideValue;

  ConsequenceQueryConsequenceQueryObject(ConsequenceQueryObject insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public ConsequenceQueryObject getInsideValue() {
    return insideValue;
  }
}

@JsonAdapter(ConsequenceQuery.Adapter.class)
class ConsequenceQueryString extends ConsequenceQuery {

  private final String insideValue;

  ConsequenceQueryString(String insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public String getInsideValue() {
    return insideValue;
  }
}
