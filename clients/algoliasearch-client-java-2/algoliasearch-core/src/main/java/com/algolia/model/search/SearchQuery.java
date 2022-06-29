package com.algolia.model.search;

import com.algolia.utils.CompoundType;
import com.algolia.utils.JSON;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.reflect.TypeToken;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;

@JsonAdapter(SearchQuery.Adapter.class)
/** SearchQuery */
public abstract class SearchQuery implements CompoundType {

  public static SearchQuery of(SearchForFacets inside) {
    return new SearchQuerySearchForFacets(inside);
  }

  public static SearchQuery of(SearchForHits inside) {
    return new SearchQuerySearchForHits(inside);
  }

  public static class Adapter extends TypeAdapter<SearchQuery> {

    @Override
    public void write(final JsonWriter out, final SearchQuery oneOf) throws IOException {
      TypeAdapter runtimeTypeAdapter = (TypeAdapter) JSON.getGson().getAdapter(TypeToken.get(oneOf.getInsideValue().getClass()));
      runtimeTypeAdapter.write(out, oneOf.getInsideValue());
    }

    @Override
    public SearchQuery read(final JsonReader jsonReader) throws IOException {
      SearchForFacets searchforfacets = JSON.tryDeserialize(jsonReader, new TypeToken<SearchForFacets>() {}.getType());
      if (searchforfacets != null) {
        return SearchQuery.of(searchforfacets);
      }
      SearchForHits searchforhits = JSON.tryDeserialize(jsonReader, new TypeToken<SearchForHits>() {}.getType());
      if (searchforhits != null) {
        return SearchQuery.of(searchforhits);
      }
      return null;
    }
  }
}

@JsonAdapter(SearchQuery.Adapter.class)
class SearchQuerySearchForFacets extends SearchQuery {

  private final SearchForFacets insideValue;

  SearchQuerySearchForFacets(SearchForFacets insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public SearchForFacets getInsideValue() {
    return insideValue;
  }
}

@JsonAdapter(SearchQuery.Adapter.class)
class SearchQuerySearchForHits extends SearchQuery {

  private final SearchForHits insideValue;

  SearchQuerySearchForHits(SearchForHits insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public SearchForHits getInsideValue() {
    return insideValue;
  }
}
