package com.algolia.model.search;

import com.algolia.utils.CompoundType;
import com.algolia.utils.JSON;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.reflect.TypeToken;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;

@JsonAdapter(Promote.Adapter.class)
/** Promote */
public abstract class Promote implements CompoundType {

  public static Promote of(PromoteObjectID inside) {
    return new PromotePromoteObjectID(inside);
  }

  public static Promote of(PromoteObjectIDs inside) {
    return new PromotePromoteObjectIDs(inside);
  }

  public static class Adapter extends TypeAdapter<Promote> {

    @Override
    public void write(final JsonWriter out, final Promote oneOf) throws IOException {
      TypeAdapter runtimeTypeAdapter = (TypeAdapter) JSON.getGson().getAdapter(TypeToken.get(oneOf.getInsideValue().getClass()));
      runtimeTypeAdapter.write(out, oneOf.getInsideValue());
    }

    @Override
    public Promote read(final JsonReader jsonReader) throws IOException {
      PromoteObjectID promoteobjectid = JSON.tryDeserialize(jsonReader, new TypeToken<PromoteObjectID>() {}.getType());
      if (promoteobjectid != null) {
        return Promote.of(promoteobjectid);
      }
      PromoteObjectIDs promoteobjectids = JSON.tryDeserialize(jsonReader, new TypeToken<PromoteObjectIDs>() {}.getType());
      if (promoteobjectids != null) {
        return Promote.of(promoteobjectids);
      }
      return null;
    }
  }
}

@JsonAdapter(Promote.Adapter.class)
class PromotePromoteObjectID extends Promote {

  private final PromoteObjectID insideValue;

  PromotePromoteObjectID(PromoteObjectID insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public PromoteObjectID getInsideValue() {
    return insideValue;
  }
}

@JsonAdapter(Promote.Adapter.class)
class PromotePromoteObjectIDs extends Promote {

  private final PromoteObjectIDs insideValue;

  PromotePromoteObjectIDs(PromoteObjectIDs insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public PromoteObjectIDs getInsideValue() {
    return insideValue;
  }
}
