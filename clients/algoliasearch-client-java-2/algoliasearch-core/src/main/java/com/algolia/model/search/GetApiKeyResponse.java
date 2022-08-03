// This file is generated, manual changes will be lost - read more on
// https://github.com/algolia/api-clients-automation.

package com.algolia.model.search;

import com.fasterxml.jackson.annotation.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

/** GetApiKeyResponse */
public class GetApiKeyResponse {

  @JsonProperty("value")
  private String value;

  @JsonProperty("createdAt")
  private Long createdAt;

  @JsonProperty("acl")
  private List<Acl> acl = new ArrayList<>();

  @JsonProperty("description")
  private String description;

  @JsonProperty("indexes")
  private List<String> indexes;

  @JsonProperty("maxHitsPerQuery")
  private Integer maxHitsPerQuery;

  @JsonProperty("maxQueriesPerIPPerHour")
  private Integer maxQueriesPerIPPerHour;

  @JsonProperty("queryParameters")
  private String queryParameters;

  @JsonProperty("referers")
  private List<String> referers;

  @JsonProperty("validity")
  private Integer validity;

  public GetApiKeyResponse setValue(String value) {
    this.value = value;
    return this;
  }

  /**
   * The API key.
   *
   * @return value
   */
  @javax.annotation.Nullable
  public String getValue() {
    return value;
  }

  public GetApiKeyResponse setCreatedAt(Long createdAt) {
    this.createdAt = createdAt;
    return this;
  }

  /**
   * Time of the event expressed in milliseconds since the Unix epoch.
   *
   * @return createdAt
   */
  @javax.annotation.Nonnull
  public Long getCreatedAt() {
    return createdAt;
  }

  public GetApiKeyResponse setAcl(List<Acl> acl) {
    this.acl = acl;
    return this;
  }

  public GetApiKeyResponse addAcl(Acl aclItem) {
    this.acl.add(aclItem);
    return this;
  }

  /**
   * Set of permissions associated with the key.
   *
   * @return acl
   */
  @javax.annotation.Nonnull
  public List<Acl> getAcl() {
    return acl;
  }

  public GetApiKeyResponse setDescription(String description) {
    this.description = description;
    return this;
  }

  /**
   * A comment used to identify a key more easily in the dashboard. It is not interpreted by the
   * API.
   *
   * @return description
   */
  @javax.annotation.Nullable
  public String getDescription() {
    return description;
  }

  public GetApiKeyResponse setIndexes(List<String> indexes) {
    this.indexes = indexes;
    return this;
  }

  public GetApiKeyResponse addIndexes(String indexesItem) {
    if (this.indexes == null) {
      this.indexes = new ArrayList<>();
    }
    this.indexes.add(indexesItem);
    return this;
  }

  /**
   * Restrict this new API key to a list of indices or index patterns. If the list is empty, all
   * indices are allowed.
   *
   * @return indexes
   */
  @javax.annotation.Nullable
  public List<String> getIndexes() {
    return indexes;
  }

  public GetApiKeyResponse setMaxHitsPerQuery(Integer maxHitsPerQuery) {
    this.maxHitsPerQuery = maxHitsPerQuery;
    return this;
  }

  /**
   * Maximum number of hits this API key can retrieve in one query. If zero, no limit is enforced.
   *
   * @return maxHitsPerQuery
   */
  @javax.annotation.Nullable
  public Integer getMaxHitsPerQuery() {
    return maxHitsPerQuery;
  }

  public GetApiKeyResponse setMaxQueriesPerIPPerHour(Integer maxQueriesPerIPPerHour) {
    this.maxQueriesPerIPPerHour = maxQueriesPerIPPerHour;
    return this;
  }

  /**
   * Maximum number of API calls per hour allowed from a given IP address or a user token.
   *
   * @return maxQueriesPerIPPerHour
   */
  @javax.annotation.Nullable
  public Integer getMaxQueriesPerIPPerHour() {
    return maxQueriesPerIPPerHour;
  }

  public GetApiKeyResponse setQueryParameters(String queryParameters) {
    this.queryParameters = queryParameters;
    return this;
  }

  /**
   * URL-encoded query string. Force some query parameters to be applied for each query made with
   * this API key.
   *
   * @return queryParameters
   */
  @javax.annotation.Nullable
  public String getQueryParameters() {
    return queryParameters;
  }

  public GetApiKeyResponse setReferers(List<String> referers) {
    this.referers = referers;
    return this;
  }

  public GetApiKeyResponse addReferers(String referersItem) {
    if (this.referers == null) {
      this.referers = new ArrayList<>();
    }
    this.referers.add(referersItem);
    return this;
  }

  /**
   * Restrict this new API key to specific referers. If empty or blank, defaults to all referers.
   *
   * @return referers
   */
  @javax.annotation.Nullable
  public List<String> getReferers() {
    return referers;
  }

  public GetApiKeyResponse setValidity(Integer validity) {
    this.validity = validity;
    return this;
  }

  /**
   * Validity limit for this key in seconds. The key will automatically be removed after this period
   * of time.
   *
   * @return validity
   */
  @javax.annotation.Nullable
  public Integer getValidity() {
    return validity;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    GetApiKeyResponse getApiKeyResponse = (GetApiKeyResponse) o;
    return (
      Objects.equals(this.value, getApiKeyResponse.value) &&
      Objects.equals(this.createdAt, getApiKeyResponse.createdAt) &&
      Objects.equals(this.acl, getApiKeyResponse.acl) &&
      Objects.equals(this.description, getApiKeyResponse.description) &&
      Objects.equals(this.indexes, getApiKeyResponse.indexes) &&
      Objects.equals(this.maxHitsPerQuery, getApiKeyResponse.maxHitsPerQuery) &&
      Objects.equals(this.maxQueriesPerIPPerHour, getApiKeyResponse.maxQueriesPerIPPerHour) &&
      Objects.equals(this.queryParameters, getApiKeyResponse.queryParameters) &&
      Objects.equals(this.referers, getApiKeyResponse.referers) &&
      Objects.equals(this.validity, getApiKeyResponse.validity)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(
      value,
      createdAt,
      acl,
      description,
      indexes,
      maxHitsPerQuery,
      maxQueriesPerIPPerHour,
      queryParameters,
      referers,
      validity
    );
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class GetApiKeyResponse {\n");
    sb.append("    value: ").append(toIndentedString(value)).append("\n");
    sb.append("    createdAt: ").append(toIndentedString(createdAt)).append("\n");
    sb.append("    acl: ").append(toIndentedString(acl)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    indexes: ").append(toIndentedString(indexes)).append("\n");
    sb.append("    maxHitsPerQuery: ").append(toIndentedString(maxHitsPerQuery)).append("\n");
    sb.append("    maxQueriesPerIPPerHour: ").append(toIndentedString(maxQueriesPerIPPerHour)).append("\n");
    sb.append("    queryParameters: ").append(toIndentedString(queryParameters)).append("\n");
    sb.append("    referers: ").append(toIndentedString(referers)).append("\n");
    sb.append("    validity: ").append(toIndentedString(validity)).append("\n");
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
