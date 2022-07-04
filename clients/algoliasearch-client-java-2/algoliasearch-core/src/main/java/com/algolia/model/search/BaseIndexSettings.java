package com.algolia.model.search;

import com.google.gson.annotations.SerializedName;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;

/** BaseIndexSettings */
public class BaseIndexSettings {

  @SerializedName("replicas")
  private List<String> replicas;

  @SerializedName("paginationLimitedTo")
  private Integer paginationLimitedTo;

  @SerializedName("disableTypoToleranceOnWords")
  private List<String> disableTypoToleranceOnWords;

  @SerializedName("attributesToTransliterate")
  private List<String> attributesToTransliterate;

  @SerializedName("camelCaseAttributes")
  private List<String> camelCaseAttributes;

  @SerializedName("decompoundedAttributes")
  private Object decompoundedAttributes;

  @SerializedName("indexLanguages")
  private List<String> indexLanguages;

  @SerializedName("disablePrefixOnAttributes")
  private List<String> disablePrefixOnAttributes;

  @SerializedName("allowCompressionOfIntegerArray")
  private Boolean allowCompressionOfIntegerArray;

  @SerializedName("numericAttributesForFiltering")
  private List<String> numericAttributesForFiltering;

  @SerializedName("separatorsToIndex")
  private String separatorsToIndex;

  @SerializedName("searchableAttributes")
  private List<String> searchableAttributes;

  @SerializedName("userData")
  private Object userData;

  @SerializedName("customNormalization")
  private Map<String, Map<String, String>> customNormalization;

  public BaseIndexSettings setReplicas(List<String> replicas) {
    this.replicas = replicas;
    return this;
  }

  public BaseIndexSettings addReplicas(String replicasItem) {
    if (this.replicas == null) {
      this.replicas = new ArrayList<>();
    }
    this.replicas.add(replicasItem);
    return this;
  }

  /**
   * Creates replicas, exact copies of an index.
   *
   * @return replicas
   */
  @javax.annotation.Nullable
  public List<String> getReplicas() {
    return replicas;
  }

  public BaseIndexSettings setPaginationLimitedTo(Integer paginationLimitedTo) {
    this.paginationLimitedTo = paginationLimitedTo;
    return this;
  }

  /**
   * Set the maximum number of hits accessible via pagination.
   *
   * @return paginationLimitedTo
   */
  @javax.annotation.Nullable
  public Integer getPaginationLimitedTo() {
    return paginationLimitedTo;
  }

  public BaseIndexSettings setDisableTypoToleranceOnWords(List<String> disableTypoToleranceOnWords) {
    this.disableTypoToleranceOnWords = disableTypoToleranceOnWords;
    return this;
  }

  public BaseIndexSettings addDisableTypoToleranceOnWords(String disableTypoToleranceOnWordsItem) {
    if (this.disableTypoToleranceOnWords == null) {
      this.disableTypoToleranceOnWords = new ArrayList<>();
    }
    this.disableTypoToleranceOnWords.add(disableTypoToleranceOnWordsItem);
    return this;
  }

  /**
   * A list of words for which you want to turn off typo tolerance.
   *
   * @return disableTypoToleranceOnWords
   */
  @javax.annotation.Nullable
  public List<String> getDisableTypoToleranceOnWords() {
    return disableTypoToleranceOnWords;
  }

  public BaseIndexSettings setAttributesToTransliterate(List<String> attributesToTransliterate) {
    this.attributesToTransliterate = attributesToTransliterate;
    return this;
  }

  public BaseIndexSettings addAttributesToTransliterate(String attributesToTransliterateItem) {
    if (this.attributesToTransliterate == null) {
      this.attributesToTransliterate = new ArrayList<>();
    }
    this.attributesToTransliterate.add(attributesToTransliterateItem);
    return this;
  }

  /**
   * Specify on which attributes in your index Algolia should apply Japanese transliteration to make
   * words indexed in Katakana or Kanji searchable in Hiragana.
   *
   * @return attributesToTransliterate
   */
  @javax.annotation.Nullable
  public List<String> getAttributesToTransliterate() {
    return attributesToTransliterate;
  }

  public BaseIndexSettings setCamelCaseAttributes(List<String> camelCaseAttributes) {
    this.camelCaseAttributes = camelCaseAttributes;
    return this;
  }

  public BaseIndexSettings addCamelCaseAttributes(String camelCaseAttributesItem) {
    if (this.camelCaseAttributes == null) {
      this.camelCaseAttributes = new ArrayList<>();
    }
    this.camelCaseAttributes.add(camelCaseAttributesItem);
    return this;
  }

  /**
   * List of attributes on which to do a decomposition of camel case words.
   *
   * @return camelCaseAttributes
   */
  @javax.annotation.Nullable
  public List<String> getCamelCaseAttributes() {
    return camelCaseAttributes;
  }

  public BaseIndexSettings setDecompoundedAttributes(Object decompoundedAttributes) {
    this.decompoundedAttributes = decompoundedAttributes;
    return this;
  }

  /**
   * Specify on which attributes in your index Algolia should apply word segmentation, also known as
   * decompounding.
   *
   * @return decompoundedAttributes
   */
  @javax.annotation.Nullable
  public Object getDecompoundedAttributes() {
    return decompoundedAttributes;
  }

  public BaseIndexSettings setIndexLanguages(List<String> indexLanguages) {
    this.indexLanguages = indexLanguages;
    return this;
  }

  public BaseIndexSettings addIndexLanguages(String indexLanguagesItem) {
    if (this.indexLanguages == null) {
      this.indexLanguages = new ArrayList<>();
    }
    this.indexLanguages.add(indexLanguagesItem);
    return this;
  }

  /**
   * Sets the languages at the index level for language-specific processing such as tokenization and
   * normalization.
   *
   * @return indexLanguages
   */
  @javax.annotation.Nullable
  public List<String> getIndexLanguages() {
    return indexLanguages;
  }

  public BaseIndexSettings setDisablePrefixOnAttributes(List<String> disablePrefixOnAttributes) {
    this.disablePrefixOnAttributes = disablePrefixOnAttributes;
    return this;
  }

  public BaseIndexSettings addDisablePrefixOnAttributes(String disablePrefixOnAttributesItem) {
    if (this.disablePrefixOnAttributes == null) {
      this.disablePrefixOnAttributes = new ArrayList<>();
    }
    this.disablePrefixOnAttributes.add(disablePrefixOnAttributesItem);
    return this;
  }

  /**
   * List of attributes on which you want to disable prefix matching.
   *
   * @return disablePrefixOnAttributes
   */
  @javax.annotation.Nullable
  public List<String> getDisablePrefixOnAttributes() {
    return disablePrefixOnAttributes;
  }

  public BaseIndexSettings setAllowCompressionOfIntegerArray(Boolean allowCompressionOfIntegerArray) {
    this.allowCompressionOfIntegerArray = allowCompressionOfIntegerArray;
    return this;
  }

  /**
   * Enables compression of large integer arrays.
   *
   * @return allowCompressionOfIntegerArray
   */
  @javax.annotation.Nullable
  public Boolean getAllowCompressionOfIntegerArray() {
    return allowCompressionOfIntegerArray;
  }

  public BaseIndexSettings setNumericAttributesForFiltering(List<String> numericAttributesForFiltering) {
    this.numericAttributesForFiltering = numericAttributesForFiltering;
    return this;
  }

  public BaseIndexSettings addNumericAttributesForFiltering(String numericAttributesForFilteringItem) {
    if (this.numericAttributesForFiltering == null) {
      this.numericAttributesForFiltering = new ArrayList<>();
    }
    this.numericAttributesForFiltering.add(numericAttributesForFilteringItem);
    return this;
  }

  /**
   * List of numeric attributes that can be used as numerical filters.
   *
   * @return numericAttributesForFiltering
   */
  @javax.annotation.Nullable
  public List<String> getNumericAttributesForFiltering() {
    return numericAttributesForFiltering;
  }

  public BaseIndexSettings setSeparatorsToIndex(String separatorsToIndex) {
    this.separatorsToIndex = separatorsToIndex;
    return this;
  }

  /**
   * Control which separators are indexed.
   *
   * @return separatorsToIndex
   */
  @javax.annotation.Nullable
  public String getSeparatorsToIndex() {
    return separatorsToIndex;
  }

  public BaseIndexSettings setSearchableAttributes(List<String> searchableAttributes) {
    this.searchableAttributes = searchableAttributes;
    return this;
  }

  public BaseIndexSettings addSearchableAttributes(String searchableAttributesItem) {
    if (this.searchableAttributes == null) {
      this.searchableAttributes = new ArrayList<>();
    }
    this.searchableAttributes.add(searchableAttributesItem);
    return this;
  }

  /**
   * The complete list of attributes used for searching.
   *
   * @return searchableAttributes
   */
  @javax.annotation.Nullable
  public List<String> getSearchableAttributes() {
    return searchableAttributes;
  }

  public BaseIndexSettings setUserData(Object userData) {
    this.userData = userData;
    return this;
  }

  /**
   * Lets you store custom data in your indices.
   *
   * @return userData
   */
  @javax.annotation.Nullable
  public Object getUserData() {
    return userData;
  }

  public BaseIndexSettings setCustomNormalization(Map<String, Map<String, String>> customNormalization) {
    this.customNormalization = customNormalization;
    return this;
  }

  public BaseIndexSettings putCustomNormalization(String key, Map<String, String> customNormalizationItem) {
    if (this.customNormalization == null) {
      this.customNormalization = new HashMap<>();
    }
    this.customNormalization.put(key, customNormalizationItem);
    return this;
  }

  /**
   * Overrides Algolia's default normalization.
   *
   * @return customNormalization
   */
  @javax.annotation.Nullable
  public Map<String, Map<String, String>> getCustomNormalization() {
    return customNormalization;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    BaseIndexSettings baseIndexSettings = (BaseIndexSettings) o;
    return (
      Objects.equals(this.replicas, baseIndexSettings.replicas) &&
      Objects.equals(this.paginationLimitedTo, baseIndexSettings.paginationLimitedTo) &&
      Objects.equals(this.disableTypoToleranceOnWords, baseIndexSettings.disableTypoToleranceOnWords) &&
      Objects.equals(this.attributesToTransliterate, baseIndexSettings.attributesToTransliterate) &&
      Objects.equals(this.camelCaseAttributes, baseIndexSettings.camelCaseAttributes) &&
      Objects.equals(this.decompoundedAttributes, baseIndexSettings.decompoundedAttributes) &&
      Objects.equals(this.indexLanguages, baseIndexSettings.indexLanguages) &&
      Objects.equals(this.disablePrefixOnAttributes, baseIndexSettings.disablePrefixOnAttributes) &&
      Objects.equals(this.allowCompressionOfIntegerArray, baseIndexSettings.allowCompressionOfIntegerArray) &&
      Objects.equals(this.numericAttributesForFiltering, baseIndexSettings.numericAttributesForFiltering) &&
      Objects.equals(this.separatorsToIndex, baseIndexSettings.separatorsToIndex) &&
      Objects.equals(this.searchableAttributes, baseIndexSettings.searchableAttributes) &&
      Objects.equals(this.userData, baseIndexSettings.userData) &&
      Objects.equals(this.customNormalization, baseIndexSettings.customNormalization)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(
      replicas,
      paginationLimitedTo,
      disableTypoToleranceOnWords,
      attributesToTransliterate,
      camelCaseAttributes,
      decompoundedAttributes,
      indexLanguages,
      disablePrefixOnAttributes,
      allowCompressionOfIntegerArray,
      numericAttributesForFiltering,
      separatorsToIndex,
      searchableAttributes,
      userData,
      customNormalization
    );
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class BaseIndexSettings {\n");
    sb.append("    replicas: ").append(toIndentedString(replicas)).append("\n");
    sb.append("    paginationLimitedTo: ").append(toIndentedString(paginationLimitedTo)).append("\n");
    sb.append("    disableTypoToleranceOnWords: ").append(toIndentedString(disableTypoToleranceOnWords)).append("\n");
    sb.append("    attributesToTransliterate: ").append(toIndentedString(attributesToTransliterate)).append("\n");
    sb.append("    camelCaseAttributes: ").append(toIndentedString(camelCaseAttributes)).append("\n");
    sb.append("    decompoundedAttributes: ").append(toIndentedString(decompoundedAttributes)).append("\n");
    sb.append("    indexLanguages: ").append(toIndentedString(indexLanguages)).append("\n");
    sb.append("    disablePrefixOnAttributes: ").append(toIndentedString(disablePrefixOnAttributes)).append("\n");
    sb.append("    allowCompressionOfIntegerArray: ").append(toIndentedString(allowCompressionOfIntegerArray)).append("\n");
    sb.append("    numericAttributesForFiltering: ").append(toIndentedString(numericAttributesForFiltering)).append("\n");
    sb.append("    separatorsToIndex: ").append(toIndentedString(separatorsToIndex)).append("\n");
    sb.append("    searchableAttributes: ").append(toIndentedString(searchableAttributes)).append("\n");
    sb.append("    userData: ").append(toIndentedString(userData)).append("\n");
    sb.append("    customNormalization: ").append(toIndentedString(customNormalization)).append("\n");
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
