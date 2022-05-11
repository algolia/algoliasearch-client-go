package com.algolia.model.search;

import com.algolia.utils.CompoundType;
import com.algolia.utils.JSON;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.reflect.TypeToken;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;

@JsonAdapter(AroundRadius.Adapter.class)
public abstract class AroundRadius implements CompoundType {

  public static AroundRadius ofAroundRadiusAll(AroundRadiusAll inside) {
    return new AroundRadiusAroundRadiusAll(inside);
  }

  public static AroundRadius ofInteger(Integer inside) {
    return new AroundRadiusInteger(inside);
  }

  public static class Adapter extends TypeAdapter<AroundRadius> {

    @Override
    public void write(final JsonWriter out, final AroundRadius oneOf)
      throws IOException {
      TypeAdapter runtimeTypeAdapter = (TypeAdapter) JSON
        .getGson()
        .getAdapter(TypeToken.get(oneOf.getInsideValue().getClass()));
      runtimeTypeAdapter.write(out, oneOf.getInsideValue());
    }

    @Override
    public AroundRadius read(final JsonReader jsonReader) throws IOException {
      AroundRadiusAll aroundradiusall = JSON.tryDeserialize(
        jsonReader,
        new TypeToken<AroundRadiusAll>() {}.getType()
      );
      if (aroundradiusall != null) {
        return AroundRadius.ofAroundRadiusAll(aroundradiusall);
      }
      Integer integer = JSON.tryDeserialize(
        jsonReader,
        new TypeToken<Integer>() {}.getType()
      );
      if (integer != null) {
        return AroundRadius.ofInteger(integer);
      }
      return null;
    }
  }
}

@JsonAdapter(AroundRadius.Adapter.class)
class AroundRadiusAroundRadiusAll extends AroundRadius {

  private final AroundRadiusAll insideValue;

  AroundRadiusAroundRadiusAll(AroundRadiusAll insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public AroundRadiusAll getInsideValue() {
    return insideValue;
  }
}

@JsonAdapter(AroundRadius.Adapter.class)
class AroundRadiusInteger extends AroundRadius {

  private final Integer insideValue;

  AroundRadiusInteger(Integer insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public Integer getInsideValue() {
    return insideValue;
  }
}
