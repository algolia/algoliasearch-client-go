package com.algolia.model.search;

import com.algolia.utils.CompoundType;
import com.algolia.utils.JSON;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.reflect.TypeToken;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;

@JsonAdapter(TypoTolerance.Adapter.class)
public abstract class TypoTolerance implements CompoundType {

  public static TypoTolerance ofBoolean(Boolean inside) {
    return new TypoToleranceBoolean(inside);
  }

  public static TypoTolerance ofTypoToleranceEnum(TypoToleranceEnum inside) {
    return new TypoToleranceTypoToleranceEnum(inside);
  }

  public static class Adapter extends TypeAdapter<TypoTolerance> {

    @Override
    public void write(final JsonWriter out, final TypoTolerance oneOf) throws IOException {
      TypeAdapter runtimeTypeAdapter = (TypeAdapter) JSON.getGson().getAdapter(TypeToken.get(oneOf.getInsideValue().getClass()));
      runtimeTypeAdapter.write(out, oneOf.getInsideValue());
    }

    @Override
    public TypoTolerance read(final JsonReader jsonReader) throws IOException {
      Boolean _boolean = JSON.tryDeserialize(jsonReader, new TypeToken<Boolean>() {}.getType());
      if (_boolean != null) {
        return TypoTolerance.ofBoolean(_boolean);
      }
      TypoToleranceEnum typotoleranceenum = JSON.tryDeserialize(jsonReader, new TypeToken<TypoToleranceEnum>() {}.getType());
      if (typotoleranceenum != null) {
        return TypoTolerance.ofTypoToleranceEnum(typotoleranceenum);
      }
      return null;
    }
  }
}

@JsonAdapter(TypoTolerance.Adapter.class)
class TypoToleranceBoolean extends TypoTolerance {

  private final Boolean insideValue;

  TypoToleranceBoolean(Boolean insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public Boolean getInsideValue() {
    return insideValue;
  }
}

@JsonAdapter(TypoTolerance.Adapter.class)
class TypoToleranceTypoToleranceEnum extends TypoTolerance {

  private final TypoToleranceEnum insideValue;

  TypoToleranceTypoToleranceEnum(TypoToleranceEnum insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public TypoToleranceEnum getInsideValue() {
    return insideValue;
  }
}
