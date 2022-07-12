// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.search;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;

/** SearchSynonymsResponse */
public class SearchSynonymsResponse {

  @JsonProperty("hits")
  private List<SynonymHit> hits = new ArrayList<>();

  @JsonProperty("nbHits")
  private Integer nbHits;

  private Map<String, Object> additionalProperties = new HashMap<>();

  @JsonAnyGetter
  public Map<String, Object> getAdditionalProperties() {
    return this.additionalProperties;
  }

  @JsonAnySetter
  public SearchSynonymsResponse setAdditionalProperty(String name, Object value) {
    this.additionalProperties.put(name, value);
    return this;
  }

  public SearchSynonymsResponse setHits(List<SynonymHit> hits) {
    this.hits = hits;
    return this;
  }

  public SearchSynonymsResponse addHits(SynonymHit hitsItem) {
    this.hits.add(hitsItem);
    return this;
  }

  /**
   * Array of synonym objects.
   *
   * @return hits
   */
  @javax.annotation.Nonnull
  public List<SynonymHit> getHits() {
    return hits;
  }

  public SearchSynonymsResponse setNbHits(Integer nbHits) {
    this.nbHits = nbHits;
    return this;
  }

  /**
   * Number of hits that the search query matched.
   *
   * @return nbHits
   */
  @javax.annotation.Nonnull
  public Integer getNbHits() {
    return nbHits;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SearchSynonymsResponse searchSynonymsResponse = (SearchSynonymsResponse) o;
    return (
      Objects.equals(this.hits, searchSynonymsResponse.hits) &&
      Objects.equals(this.nbHits, searchSynonymsResponse.nbHits) &&
      super.equals(o)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(hits, nbHits, super.hashCode());
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SearchSynonymsResponse {\n");
    sb.append("    ").append(toIndentedString(super.toString())).append("\n");
    sb.append("    hits: ").append(toIndentedString(hits)).append("\n");
    sb.append("    nbHits: ").append(toIndentedString(nbHits)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }
}
