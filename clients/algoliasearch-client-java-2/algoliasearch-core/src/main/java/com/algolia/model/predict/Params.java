package com.algolia.model.predict;

import com.algolia.JSON;
import com.algolia.utils.CompoundType;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.reflect.TypeToken;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;

@JsonAdapter(Params.Adapter.class)
public abstract class Params implements CompoundType {

  public static Params ofAllParams(AllParams inside) {
    return new ParamsAllParams(inside);
  }

  public static Params ofModelsToRetrieve(ModelsToRetrieve inside) {
    return new ParamsModelsToRetrieve(inside);
  }

  public static Params ofTypesToRetrieve(TypesToRetrieve inside) {
    return new ParamsTypesToRetrieve(inside);
  }

  public static class Adapter extends TypeAdapter<Params> {

    @Override
    public void write(final JsonWriter out, final Params oneOf)
      throws IOException {
      TypeAdapter runtimeTypeAdapter = (TypeAdapter) JSON
        .getGson()
        .getAdapter(TypeToken.get(oneOf.getInsideValue().getClass()));
      runtimeTypeAdapter.write(out, oneOf.getInsideValue());
    }

    @Override
    public Params read(final JsonReader jsonReader) throws IOException {
      return null;
    }
  }
}

@JsonAdapter(Params.Adapter.class)
class ParamsAllParams extends Params {

  private final AllParams insideValue;

  ParamsAllParams(AllParams insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public AllParams getInsideValue() {
    return insideValue;
  }
}

@JsonAdapter(Params.Adapter.class)
class ParamsModelsToRetrieve extends Params {

  private final ModelsToRetrieve insideValue;

  ParamsModelsToRetrieve(ModelsToRetrieve insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public ModelsToRetrieve getInsideValue() {
    return insideValue;
  }
}

@JsonAdapter(Params.Adapter.class)
class ParamsTypesToRetrieve extends Params {

  private final TypesToRetrieve insideValue;

  ParamsTypesToRetrieve(TypesToRetrieve insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public TypesToRetrieve getInsideValue() {
    return insideValue;
  }
}
