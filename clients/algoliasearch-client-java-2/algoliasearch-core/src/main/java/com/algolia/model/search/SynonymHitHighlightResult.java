package com.algolia.model.search;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;

/** Highlighted results. */
public class SynonymHitHighlightResult {

  @JsonProperty("type")
  private Map<String, HighlightResult> type;

  @JsonProperty("synonyms")
  private List<Map<String, HighlightResult>> synonyms;

  public SynonymHitHighlightResult setType(Map<String, HighlightResult> type) {
    this.type = type;
    return this;
  }

  public SynonymHitHighlightResult putType(String key, HighlightResult typeItem) {
    if (this.type == null) {
      this.type = new HashMap<>();
    }
    this.type.put(key, typeItem);
    return this;
  }

  /**
   * Show highlighted section and words matched on a query.
   *
   * @return type
   */
  @javax.annotation.Nullable
  public Map<String, HighlightResult> getType() {
    return type;
  }

  public SynonymHitHighlightResult setSynonyms(List<Map<String, HighlightResult>> synonyms) {
    this.synonyms = synonyms;
    return this;
  }

  public SynonymHitHighlightResult addSynonyms(Map<String, HighlightResult> synonymsItem) {
    if (this.synonyms == null) {
      this.synonyms = new ArrayList<>();
    }
    this.synonyms.add(synonymsItem);
    return this;
  }

  /**
   * Get synonyms
   *
   * @return synonyms
   */
  @javax.annotation.Nullable
  public List<Map<String, HighlightResult>> getSynonyms() {
    return synonyms;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SynonymHitHighlightResult synonymHitHighlightResult = (SynonymHitHighlightResult) o;
    return Objects.equals(this.type, synonymHitHighlightResult.type) && Objects.equals(this.synonyms, synonymHitHighlightResult.synonyms);
  }

  @Override
  public int hashCode() {
    return Objects.hash(type, synonyms);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SynonymHitHighlightResult {\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    synonyms: ").append(toIndentedString(synonyms)).append("\n");
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
