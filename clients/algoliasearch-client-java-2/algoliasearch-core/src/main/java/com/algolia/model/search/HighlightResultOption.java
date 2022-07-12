// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.search;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/** Show highlighted section and words matched on a query. */
public class HighlightResultOption {

  @JsonProperty("value")
  private String value;

  @JsonProperty("matchLevel")
  private MatchLevel matchLevel;

  @JsonProperty("matchedWords")
  private List<String> matchedWords = new ArrayList<>();

  @JsonProperty("fullyHighlighted")
  private Boolean fullyHighlighted;

  public HighlightResultOption setValue(String value) {
    this.value = value;
    return this;
  }

  /**
   * Markup text with occurrences highlighted.
   *
   * @return value
   */
  @javax.annotation.Nonnull
  public String getValue() {
    return value;
  }

  public HighlightResultOption setMatchLevel(MatchLevel matchLevel) {
    this.matchLevel = matchLevel;
    return this;
  }

  /**
   * Get matchLevel
   *
   * @return matchLevel
   */
  @javax.annotation.Nonnull
  public MatchLevel getMatchLevel() {
    return matchLevel;
  }

  public HighlightResultOption setMatchedWords(List<String> matchedWords) {
    this.matchedWords = matchedWords;
    return this;
  }

  public HighlightResultOption addMatchedWords(String matchedWordsItem) {
    this.matchedWords.add(matchedWordsItem);
    return this;
  }

  /**
   * List of words from the query that matched the object.
   *
   * @return matchedWords
   */
  @javax.annotation.Nonnull
  public List<String> getMatchedWords() {
    return matchedWords;
  }

  public HighlightResultOption setFullyHighlighted(Boolean fullyHighlighted) {
    this.fullyHighlighted = fullyHighlighted;
    return this;
  }

  /**
   * Whether the entire attribute value is highlighted.
   *
   * @return fullyHighlighted
   */
  @javax.annotation.Nullable
  public Boolean getFullyHighlighted() {
    return fullyHighlighted;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    HighlightResultOption highlightResultOption = (HighlightResultOption) o;
    return (
      Objects.equals(this.value, highlightResultOption.value) &&
      Objects.equals(this.matchLevel, highlightResultOption.matchLevel) &&
      Objects.equals(this.matchedWords, highlightResultOption.matchedWords) &&
      Objects.equals(this.fullyHighlighted, highlightResultOption.fullyHighlighted)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(value, matchLevel, matchedWords, fullyHighlighted);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class HighlightResultOption {\n");
    sb.append("    value: ").append(toIndentedString(value)).append("\n");
    sb.append("    matchLevel: ").append(toIndentedString(matchLevel)).append("\n");
    sb.append("    matchedWords: ").append(toIndentedString(matchedWords)).append("\n");
    sb.append("    fullyHighlighted: ").append(toIndentedString(fullyHighlighted)).append("\n");
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
