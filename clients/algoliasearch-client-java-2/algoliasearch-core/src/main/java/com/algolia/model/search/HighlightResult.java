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

@JsonAdapter(HighlightResult.Adapter.class)
/** HighlightResult */
public abstract class HighlightResult implements CompoundType {

  public static HighlightResult of(HighlightResultOption inside) {
    return new HighlightResultHighlightResultOption(inside);
  }

  public static HighlightResult of(List<HighlightResultOption> inside) {
    return new HighlightResultListOfHighlightResultOption(inside);
  }

  public static class Adapter extends TypeAdapter<HighlightResult> {

    @Override
    public void write(final JsonWriter out, final HighlightResult oneOf) throws IOException {
      TypeAdapter runtimeTypeAdapter = (TypeAdapter) JSON.getGson().getAdapter(TypeToken.get(oneOf.getInsideValue().getClass()));
      runtimeTypeAdapter.write(out, oneOf.getInsideValue());
    }

    @Override
    public HighlightResult read(final JsonReader jsonReader) throws IOException {
      HighlightResultOption highlightresultoption = JSON.tryDeserialize(jsonReader, new TypeToken<HighlightResultOption>() {}.getType());
      if (highlightresultoption != null) {
        return HighlightResult.of(highlightresultoption);
      }
      List<HighlightResultOption> listofhighlightresultoption = JSON.tryDeserialize(
        jsonReader,
        new TypeToken<List<HighlightResultOption>>() {}.getType()
      );
      if (listofhighlightresultoption != null) {
        return HighlightResult.of(listofhighlightresultoption);
      }
      return null;
    }
  }
}

@JsonAdapter(HighlightResult.Adapter.class)
class HighlightResultHighlightResultOption extends HighlightResult {

  private final HighlightResultOption insideValue;

  HighlightResultHighlightResultOption(HighlightResultOption insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public HighlightResultOption getInsideValue() {
    return insideValue;
  }
}

@JsonAdapter(HighlightResult.Adapter.class)
class HighlightResultListOfHighlightResultOption extends HighlightResult {

  private final List<HighlightResultOption> insideValue;

  HighlightResultListOfHighlightResultOption(List<HighlightResultOption> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<HighlightResultOption> getInsideValue() {
    return insideValue;
  }
}
