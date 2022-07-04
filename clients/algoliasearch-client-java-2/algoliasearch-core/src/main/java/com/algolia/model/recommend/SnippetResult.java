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

@JsonAdapter(SnippetResult.Adapter.class)
/** SnippetResult */
public abstract class SnippetResult implements CompoundType {

  public static SnippetResult of(List<SnippetResultOption> inside) {
    return new SnippetResultListOfSnippetResultOption(inside);
  }

  public static SnippetResult of(SnippetResultOption inside) {
    return new SnippetResultSnippetResultOption(inside);
  }

  public static class Adapter extends TypeAdapter<SnippetResult> {

    @Override
    public void write(final JsonWriter out, final SnippetResult oneOf) throws IOException {
      TypeAdapter runtimeTypeAdapter = (TypeAdapter) JSON.getGson().getAdapter(TypeToken.get(oneOf.getInsideValue().getClass()));
      runtimeTypeAdapter.write(out, oneOf.getInsideValue());
    }

    @Override
    public SnippetResult read(final JsonReader jsonReader) throws IOException {
      List<SnippetResultOption> listofsnippetresultoption = JSON.tryDeserialize(
        jsonReader,
        new TypeToken<List<SnippetResultOption>>() {}.getType()
      );
      if (listofsnippetresultoption != null) {
        return SnippetResult.of(listofsnippetresultoption);
      }
      SnippetResultOption snippetresultoption = JSON.tryDeserialize(jsonReader, new TypeToken<SnippetResultOption>() {}.getType());
      if (snippetresultoption != null) {
        return SnippetResult.of(snippetresultoption);
      }
      return null;
    }
  }
}

@JsonAdapter(SnippetResult.Adapter.class)
class SnippetResultListOfSnippetResultOption extends SnippetResult {

  private final List<SnippetResultOption> insideValue;

  SnippetResultListOfSnippetResultOption(List<SnippetResultOption> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<SnippetResultOption> getInsideValue() {
    return insideValue;
  }
}

@JsonAdapter(SnippetResult.Adapter.class)
class SnippetResultSnippetResultOption extends SnippetResult {

  private final SnippetResultOption insideValue;

  SnippetResultSnippetResultOption(SnippetResultOption insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public SnippetResultOption getInsideValue() {
    return insideValue;
  }
}
