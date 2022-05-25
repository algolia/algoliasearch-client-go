package com.algolia.model.search;

import com.algolia.utils.CompoundType;
import com.algolia.utils.JSON;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.reflect.TypeToken;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;

@JsonAdapter(AttributeOrBuiltInOperation.Adapter.class)
public abstract class AttributeOrBuiltInOperation implements CompoundType {

  public static AttributeOrBuiltInOperation ofBuiltInOperation(BuiltInOperation inside) {
    return new AttributeOrBuiltInOperationBuiltInOperation(inside);
  }

  public static AttributeOrBuiltInOperation ofString(String inside) {
    return new AttributeOrBuiltInOperationString(inside);
  }

  public static class Adapter extends TypeAdapter<AttributeOrBuiltInOperation> {

    @Override
    public void write(final JsonWriter out, final AttributeOrBuiltInOperation oneOf) throws IOException {
      TypeAdapter runtimeTypeAdapter = (TypeAdapter) JSON.getGson().getAdapter(TypeToken.get(oneOf.getInsideValue().getClass()));
      runtimeTypeAdapter.write(out, oneOf.getInsideValue());
    }

    @Override
    public AttributeOrBuiltInOperation read(final JsonReader jsonReader) throws IOException {
      BuiltInOperation builtinoperation = JSON.tryDeserialize(jsonReader, new TypeToken<BuiltInOperation>() {}.getType());
      if (builtinoperation != null) {
        return AttributeOrBuiltInOperation.ofBuiltInOperation(builtinoperation);
      }
      String string = JSON.tryDeserialize(jsonReader, new TypeToken<String>() {}.getType());
      if (string != null) {
        return AttributeOrBuiltInOperation.ofString(string);
      }
      return null;
    }
  }
}

@JsonAdapter(AttributeOrBuiltInOperation.Adapter.class)
class AttributeOrBuiltInOperationBuiltInOperation extends AttributeOrBuiltInOperation {

  private final BuiltInOperation insideValue;

  AttributeOrBuiltInOperationBuiltInOperation(BuiltInOperation insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public BuiltInOperation getInsideValue() {
    return insideValue;
  }
}

@JsonAdapter(AttributeOrBuiltInOperation.Adapter.class)
class AttributeOrBuiltInOperationString extends AttributeOrBuiltInOperation {

  private final String insideValue;

  AttributeOrBuiltInOperationString(String insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public String getInsideValue() {
    return insideValue;
  }
}
